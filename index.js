#!/usr/bin/env node

import { init } from "./src/project/index.mjs";
import { setupProject } from "./src/project/setup.mjs";
import { closeInput } from "./src/cmd.mjs";

function logInstructions(name) {
  console.log("\n--------------");
  console.log(`\n\tcd ${name}`);
  console.log(`\tnpm start\n`);
  console.log("--------------");
}

async function main() {
  const { name, language, config } = await setupProject();
  closeInput();
  await init(language, config);
  logInstructions(name);
}
main().catch(console.error);
