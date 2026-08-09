const { defineConfig } = require("cypress");
const { exec, spawn } = require("child_process");
const fs = require("fs");
const path = require("path");
const cypressFailFast = require("cypress-fail-fast/plugin.js");
let firebackProcess; // Store the Fireback process reference

// Repo root is the parent of this e2e/ folder. Resolving it this way (instead
// of hardcoding a specific developer's home directory) keeps this working for
// every contributor's local checkout without edits.
const REPO_ROOT = path.resolve(__dirname, "..");

let BINARY = process.env.FIREBACK_BINARY || path.join(REPO_ROOT, "app");
let CWD = process.env.FIREBACK_CWD || REPO_ROOT;
const PORT = 7794;
let DB_VENDOR = "sqlite";
const isGitHubActions = !!process.env.GITHUB_ACTIONS;

if (isGitHubActions) {
  // CI installs Fireback from the .deb artifact instead of building it locally.
  BINARY = process.env.FIREBACK_BINARY || "/usr/local/bin/fireback";
  CWD = process.env.FIREBACK_CWD || "/home/runner/work/fireback";
}

console.log(BINARY);
console.log(CWD);

// Best-effort parse of the repo root .env (same file `fireback config ... set`
// writes to) so a locally running Postgres can be reused for e2e without
// hardcoding one developer's credentials here. Falls back to the CI service
// container's defaults (postgres/postgres@localhost:5432) when there's no
// .env yet, e.g. on a fresh CI runner.
function readDotEnv(file) {
  const out = {};
  let content;
  try {
    content = fs.readFileSync(file, "utf8");
  } catch {
    return out;
  }
  for (const line of content.split("\n")) {
    const match = line.match(/^\s*([\w.-]+)\s*=\s*(.*)\s*$/);
    if (match) out[match[1]] = match[2];
  }
  return out;
}

const dotEnv = readDotEnv(path.join(REPO_ROOT, ".env"));
const PG_HOST = process.env.POSTGRES_HOST || dotEnv.DB_HOST || "localhost";
const PG_PORT = process.env.POSTGRES_PORT || dotEnv.DB_PORT || "5432";
const PG_USER = process.env.POSTGRES_USER || dotEnv.DB_USERNAME || "postgres";
const PG_PASSWORD =
  process.env.POSTGRES_PASSWORD || dotEnv.DB_PASSWORD || "postgres";

console.log("Database", PG_HOST, PG_PORT, PG_USER, PG_PASSWORD);

const execAsync = (cmd, CWD = "") => {
  return new Promise((resolve, reject) => {
    exec(cmd, { cwd: CWD }, (error, stdout, stderr) => {
      if (error) {
        return reject(error + " --- " + cmd + "----" + CWD);
      }
      resolve(stdout || stderr);
    });
  });
};

module.exports = defineConfig({
  video: true,
  chromeWebSecurity: false,
  env: {
    GITHUB_ACTIONS: process.env.GITHUB_ACTIONS,
    PORT: PORT,
  },
  e2e: {
    setupNodeEvents(on, config) {
      cypressFailFast(on, config);

      on("task", {
        log(message) {
          console.log(message);
          return null;
        },
      });

      on("task", {
        view(url) {
          return cy.visit(url);
        },
        async dbcon(cmd) {
          const vendor = process.env.DB_TYPE || "postgres";
          console.log(100000, vendor);
          const dbname = `test_agent_${new Date().getTime()}`;

          if (vendor === "mysql") {
            try {
              await execAsync(`${BINARY} config db-vendor set mysql`, CWD);
              await execAsync(
                `${BINARY} config db-dsn set "root:root@tcp(localhost:3306)/fireback_test?charset=utf8mb4&parseTime=True&loc=Local"`,
                CWD,
              );
              await execAsync(`${BINARY} migration apply`, CWD);

              return true;
            } catch (err) {
              console.error("setup mysql failed:", err);
            }
          } else if (vendor === "postgres") {
            try {
              console.log("Using postgres", PG_HOST, PG_PORT, PG_USER);
              await execAsync(`${BINARY} config db-vendor set postgres`, CWD);

              // Always a dedicated, disposable database, never whatever
              // DB_NAME happens to be in .env — that may point at a real
              // dev database on the shared instance and we don't want to
              // DROP/CREATE over someone's data.
              const dbName = "fireback_test";
              const psqlBase = `PGPASSWORD=${PG_PASSWORD} psql -U ${PG_USER} -h ${PG_HOST} -p ${PG_PORT} -d postgres`;

              // A previous spec's Fireback process may still be closing its
              // connections when the next spec's dbcon runs (kill() doesn't
              // wait for the pool to drain). DROP DATABASE fails with
              // "database is being accessed by other users" in that case, so
              // force out any lingering backends first.
              console.log(
                await execAsync(
                  `${psqlBase} -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = '${dbName}' AND pid <> pg_backend_pid();"`,
                  CWD,
                ),
              );

              // Drop and recreate database
              console.log(
                await execAsync(
                  `${psqlBase} -c "DROP DATABASE IF EXISTS ${dbName};"`,
                  CWD,
                ),
              );

              await execAsync(
                `${psqlBase} -c "CREATE DATABASE ${dbName};"`,
                CWD,
              );

              await execAsync(
                `${BINARY} config db-dsn set "host=${PG_HOST} user=${PG_USER} password=${PG_PASSWORD} dbname=${dbName} port=${PG_PORT} sslmode=disable TimeZone=UTC"`,
                CWD,
              );

              await execAsync(`${BINARY} migration apply`, CWD);

              return true;
            } catch (err) {
              // Surface the failure instead of swallowing it: silently
              // returning false here left the previous spec's stale
              // database in place with no test failure to flag it.
              console.error("setup postgres failed:", err);
              throw err;
            }
          } else {
            await execAsync(`${BINARY} config db-vendor set sqlite`, CWD);
            await execAsync(
              `${BINARY} config db-name set /tmp/${dbname}.db`,
              CWD,
            );
            await execAsync(`${BINARY} migration apply`, CWD);

            return true;
          }

          return false;
        },
      });
      on("task", {
        exec(cmd) {
          return new Promise((resolve, reject) => {
            cmd = BINARY + " " + cmd;
            console.log("Running:", cmd, " on: ", CWD);
            exec(cmd, { cwd: CWD }, (error, stdout, stderr) => {
              if (error) {
                return reject(error);
              }
              resolve(stdout || stderr);
            });
          });
        },
      });
      on("task", {
        execSupress(cmd) {
          return new Promise((resolve, reject) => {
            cmd = BINARY + " " + cmd;
            console.log("Running:", cmd, " on: ", CWD);
            exec(cmd, { cwd: CWD }, (error, stdout, stderr) => {
              if (error) {
                return resolve(`Error: ${stderr || error.message}`);
              }
              resolve(stdout || stderr);
            });
          });
        },
      });

      on("task", {
        startFireback() {
          return new Promise((resolve, reject) => {
            console.log("Starting Fireback...");
            firebackProcess = spawn(`PORT=${PORT} ${BINARY}`, ["start"], {
              stdio: "inherit",
              shell: true,
              cwd: CWD,
            });

            setTimeout(() => {
              resolve("Fireback started");
            }, 1500); // Adjust delay based on startup time
          });
        },
        stopFireback() {
          return new Promise((resolve) => {
            if (firebackProcess) {
              console.log("Stopping Fireback...");
              firebackProcess.kill();
              firebackProcess = null;
            }
            resolve("Fireback stopped");
          });
        },
      });

      process.on("exit", () => {
        if (firebackProcess) {
          console.log("Forcing Fireback shutdown...");
          firebackProcess.kill();
        }
      });

      on("task", {
        execCwd({ cwd, binary }) {
          return new Promise((resolve, reject) => {
            console.log("Setting cwd:", cwd, "binary:", binary);
            CWD = cwd;
            BINARY = binary;
            resolve(true);
          });
        },
      });
    },
  },
});

// You can change the binary and cwd using this directly in the tests
// beforeEach(() => {
//   cy.task("execCwd", { cwd, binary });
// });
