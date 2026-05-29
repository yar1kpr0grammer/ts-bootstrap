import { read } from "./fs/json/json5.mjs";
import { write } from "./fs/json/index.mjs";

let config = null;

export function readConfig() {
  config = read("tsconfig.json");
  return config;
}

export function updateTsConfig() {
  if (!config) {
    readConfig();
  }

  if (!config) {
    throw new Error("tsconfig.json not found");
  }

  const options = (config.compilerOptions ??= {});

  options.rootDir = "./src"; // source code in src dir
  options.outDir = "./dist"; // code output in dist dir
  options.lib = ["ESNext"]; // adds modern JS API
  options.esModuleInterop = true; // more comfortable imports
  options.allowSyntheticDefaultImports = true;
  options.resolveJsonModule = true; // allows to import JSON files
  options.forceConsistentCasingInFileNames = true; // strict file names in imports
  options.moduleResolution = "bundler"; // allows to use bundlers
  options.module = "esnext"; // ts generates modern ESM

  // @ in imports as src dir
  options.baseUrl = ".";
  options.paths ??= {};
  options.paths["@/*"] = ["src/*"];

  delete options.verbatimModuleSyntax;
  delete options.moduleDetection;
  delete options.noUncheckedSideEffectImports;
  delete options.declaration;
  delete options.declarationMap;

  write("tsconfig.json", config);
}

export function addType(typeName) {
  if (!config) {
    readConfig();
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
