import { read } from "./fs/json/json5.mjs";
import { write } from "./fs/json/index.mjs";

let config = null;

export function readConfig() {
  config = read("tsconfig.json");
  return config;
}

export function updateTsConfig() {
  if (!config) {
    readConfig(); // автоматически читаем, если еще не прочитан
  }

  if (!config) {
    throw new Error("tsconfig.json not found");
  }

  config.compilerOptions.rootDir = "./src";
  config.compilerOptions.outDir = "./dist";
  write("tsconfig.json", config);
}

export function addType(typeName) {
  if (!config) {
    readConfig(); // автоматически читаем, если еще не прочитан
  }

  if (!config) {
    throw new Error("tsconfig.json not found");
  }

  if (!config.compilerOptions.types) {
    config.compilerOptions.types = [];
  }

  config.compilerOptions.types.push(typeName);
  write("tsconfig.json", config);
}

export default { readConfig, updateTsConfig, addType };
