import json from "./fs/json/index.mjs";

const PACKAGE_PATH = "package.json";

function getPackageConfig() {
  return json.read(PACKAGE_PATH);
}

export function updatePackageJSON(nodeType) {
  const config = getPackageConfig();

  config.scripts ??= {};
  config.scripts.start = "tsc && node dist/index.js";
  config.type = nodeType;

  json.write(PACKAGE_PATH, config);
}

export function addScript(name, command) {
  const config = getPackageConfig();

  config.scripts ??= {};
  config.scripts[name] = command;

  json.write(PACKAGE_PATH, config);
}

const projectNamePattern = "^[a-z0-9]+([\\-_.][a-z0-9]+)*$";
const projectNameReg = new RegExp(projectNamePattern, "i");

export function checkProjectName(name) {
  return projectNameReg.test(name);
}
