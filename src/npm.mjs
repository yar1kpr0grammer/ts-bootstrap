import json from "./fs/json/index.mjs";

const path = "package.json";
const config = json.read(path);

export function updatePackageJSON(nodeType) {
  config.scripts.start = "tsc && node dist/index.js";
  config.type = nodeType;
  json.write(path, config);
}
export function addScript(name, command) {
  config.scripts[name] = command;
  json.write(path, config);
}

const projectNamePattern = "^[a-z0-9]+([\-_.][a-z0-9]+)*$";
const projectNameReg = new RegExp(projectNamePattern, "i");
export function checkProjectName(name) {
  return projectNameReg.test(name);
}
