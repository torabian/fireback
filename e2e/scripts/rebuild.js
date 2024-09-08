const fs = require("fs");
const path = require("path");
const { exec } = require("child_process");

function findModule3Files(dir, results = []) {
  // Read the contents of the directory
  const files = fs.readdirSync(dir);

  for (const file of files) {
    const filePath = path.join(dir, file);

    if (!fs.existsSync(filePath)) {
      continue;
    }
    // Check if it's a directory
    if (fs.statSync(filePath).isDirectory()) {
      // Recursively search in the directory
      findModule3Files(filePath, results);
    } else if (file.endsWith("Module3.yml")) {
      // If the file ends with 'Module3.yml', add it to the results
      results.push(filePath);
    }
  }

  return results;
}

async function execa(command, file) {
  return new Promise((resolve, reject) => {
    exec(command, (error, stdout, stderr) => {
      if (error) {
        console.error(
          `Error executing command for file ${file}: ${error.message}`
        );
        reject(error);
      }
      if (stderr) {
        console.error(`stderr for file ${file}: ${stderr}`);
        reject(stderr);
      } else {
        console.log(`stdout for file ${file}: ${stdout}`);
        resolve();
      }
    });
  });
}

// Function to execute the external executable for each file and run `make` command
async function processFilesAndRunMake(files, pwd) {
  // Loop through each file and execute the external command
  for (const file of files) {
    const command = `${pwd}/artifacts/fireback/f gen gof --no-cache true --def ${pwd}/${file} --relative-to ${pwd} --gof-module github.com/torabian/fireback`;
    await execa(command, file);
  }
}

async function rebuild() {
  return new Promise((resolve, reject) => {
    exec("make", (error, stdout, stderr) => {
      if (error) {
        console.error(`Rebuild error: ${error.message}`);
        reject(error); // If error, throw it
      }
      if (stderr) {
        console.error(`stderr rebuild fireback: ${stderr}`);
        reject(stderr);
      } else {
        console.log(`stdout rebuild fireback: ${stdout}`);
        resolve();
      }
    });
  });
}

async function main() {
  const directoryToSearch = "./modules"; // Replace with your starting directory
  const module3Files = findModule3Files(directoryToSearch);

  console.log("Found Module3.yml files:", module3Files);

  await processFilesAndRunMake(module3Files, process.argv[2]);
  await rebuild();
}

main().then(() => {
  console.log(
    "This test rebuilt the gof module and made sure that even after rebuilding all definitions app still compiles"
  );
});
