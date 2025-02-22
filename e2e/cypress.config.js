const { defineConfig } = require("cypress");
const { exec, spawn } = require("child_process");
const cypressFailFast = require("cypress-fail-fast/plugin.js");
let firebackProcess; // Store the Fireback process reference

let CWD = "";

module.exports = defineConfig({
  video: true,
  chromeWebSecurity: false,
  e2e: {
    setupNodeEvents(on, config) {
      cypressFailFast(on, config);
      on("task", {
        exec(cmd) {
          return new Promise((resolve, reject) => {
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
            firebackProcess = spawn(`PORT=4502 ${CWD}/app`, ["start"], {
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
        execCwd(cwd) {
          return new Promise((resolve, reject) => {
            console.log("Setting cwd:", cwd);
            CWD = cwd;
            resolve(true);
          });
        },
      });
    },
  },
});
