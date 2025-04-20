import fs from "fs";
import ip from "ip";

const localIpAddress = ip.address();

const configFilePath = "./capacitor.config.json";
const config = JSON.parse(fs.readFileSync(configFilePath).toString());

config.server = {
  url: `http://${localIpAddress}:3000`,
  cleartext: true,
};

fs.writeFileSync(configFilePath, JSON.stringify(config, null, 2));

console.log(`Capacitor config updated with IP: ${localIpAddress}`);
