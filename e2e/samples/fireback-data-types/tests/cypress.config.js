const { defineConfig } = require("cypress");
const { exec, spawn } = require("child_process");
const cypressFailFast = require("cypress-fail-fast/plugin.js");
let firebackProcess; // Store the Fireback process reference

let BINARY = "/Users/ali/work/fireback/e2e/samples/fireback-data-types/app";
let CWD = "/Users/ali/work/fireback/e2e/samples/fireback-data-types";
const PORT = 7793;
let DB_VENDOR = "sqlite";
const isGitHubActions = !!process.env.GITHUB_ACTIONS;

if (isGitHubActions) {
  // BINARY = "/usr/local/bin/fireback";
  BINARY = "/usr/local/bin/fireback/e2e/samples/fireback-data-types/app";
  CWD = "/home/runner/work/fireback/e2e/samples/fireback-data-types";
  // CWD = "/home/runner/work/fireback-private/fireback-private";
}

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
          const vendor = process.env.DB_TYPE;
          console.log(100000, vendor);
          const dbname = `test_agent_${new Date().getTime()}`;

          if (vendor === "mysql") {
            try {
              await execAsync(`${BINARY} config db-vendor set mysql`, CWD);
              await execAsync(
                `${BINARY} config db-dsn set "root:root@tcp(localhost:3306)/fireback_test?charset=utf8mb4&parseTime=True&loc=Local"`,
                CWD
              );
              await execAsync(`${BINARY} migration apply`, CWD);

              return true;
            } catch (err) {
              console.error("setup mysql failed:", err);
            }
          } else {
            await execAsync(`${BINARY} config db-vendor set sqlite`, CWD);
            await execAsync(
              `${BINARY} config db-name set /tmp/${dbname}.db`,
              CWD
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
