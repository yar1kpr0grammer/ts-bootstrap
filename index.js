#!/usr/bin/env node

import { init } from "./src/project/index.mjs";
import { setupProject } from "./src/project/setup.mjs";
import { closeInput } from "./src/cmd.mjs";

function logInstructions(name, dir) {
  console.log("\n--------------\n");
  if (dir !== ".") {
    console.log(`cd ${name}`);
  }
  console.log(`npm start\n`);
  console.log("--------------");
}

async function main() {
  const { name, language, config, dir } = await setupProject();
  closeInput();
  await init(language, config);
  logInstructions(name, dir);
}
main().catch(console.error);
