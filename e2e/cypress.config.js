const { defineConfig } = require("cypress");
const { exec } = require("child_process");
const cypressFailFast = require("cypress-fail-fast/plugin.js");

module.exports = defineConfig({
  video: true,
  chromeWebSecurity: false,
  e2e: {
    setupNodeEvents(on, config) {
      cypressFailFast(on, config);
      on("task", {
        exec(cmd) {
          return new Promise((resolve, reject) => {
            console.log("Running:", cmd);
            exec(cmd, (error, stdout, stderr) => {
              if (error) {
                return reject(error);
              }
              resolve(stdout || stderr);
            });
          });
        },
      });
    },
  },
});
