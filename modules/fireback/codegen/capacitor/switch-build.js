import fs from "fs";

const configFilePath = "./capacitor.config.json";
const config = JSON.parse(fs.readFileSync(configFilePath).toString());

config.server = {};

fs.writeFileSync(configFilePath, JSON.stringify(config, null, 2));

console.log(
  `Capacitor server has been deleted and '${config.webDir}' folder will be used.`
);
