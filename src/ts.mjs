import { read } from "./fs/json/json5.mjs";
import { write } from "./fs/json/index.mjs";

let tsConfig = null;

export function readConfig() {
  tsConfig = read("tsconfig.json");
  return tsConfig;
}

export function updateTsConfig(projectConfig) {
  if (!tsConfig) {
    readConfig();
  }

  if (!tsConfig) {
    throw new Error("tsconfig.json not found");
  }

  const options = (tsConfig.compilerOptions ??= {});

  // shared options
  options.rootDir = "./src";
  options.outDir = "./dist";

  options.target = "ES2020";

  if (projectConfig.options.includes("node")) {
    options.types = ["node"];
  }

  options.sourceMap = true;

  options.strict = true;
  options.noUncheckedIndexedAccess = true;
  options.exactOptionalPropertyTypes = true;

  options.resolveJsonModule = true;
  options.forceConsistentCasingInFileNames = true;

  options.skipLibCheck = true;

  options.esModuleInterop = true;
  options.allowSyntheticDefaultImports = true;

  // alias @ -> src
  if (projectConfig.options.includes("src dir @ alias")) {
    options.baseUrl = ".";
    options.paths ??= {};
    options.paths["@/*"] = ["src/*"];
  }

  // cleanup
  delete options.verbatimModuleSyntax;
  delete options.moduleDetection;
  delete options.noUncheckedSideEffectImports;
  delete options.declaration;
  delete options.declarationMap;

  if (projectConfig.nodeType === "module") {
    // ESM
    options.lib = ["ESNext"];

    options.module = "esnext";
    options.moduleResolution = "bundler";
  } else {
    // CommonJS
    options.lib = ["ES2020"];

    options.module = "commonjs";
    options.moduleResolution = "node";
  }

  write("tsconfig.json", tsConfig);
}

export function addType(typeName) {
  if (!tsConfig) {
    readConfig();
  }

  if (!tsConfig) {
    throw new Error("tsconfig.json not found");
  }

  const options = (tsConfig.compilerOptions ??= {});
  options.types ??= [];

  if (!options.types.includes(typeName)) {
    options.types.push(typeName);
  }

  write("tsconfig.json", tsConfig);
}

export default {
  readConfig,
  updateTsConfig,
  addType,
};
