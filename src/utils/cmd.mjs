import chalk from "chalk";
import { select } from "@inquirer/prompts";
import checkboxPrompt from "@inquirer/checkbox";

export async function confirm(
  fn,
  message,
  success = "Success",
  error = "Error",
) {
  try {
    const res = await fn();

    console.log(`${chalk.bold(chalk.green(success))}: ${message}`);
    return res;
  } catch (e) {
    console.log(`${chalk.bold(chalk.red(error))}: ${message}`);
    console.error(e.message || e);
    return e;
  }
}

export async function radio(prompt, options = []) {
  if (!options.length) {
    throw new Error("radio: options is empty");
  }

  return await select({
    message: prompt,
    choices: options.map((o) => ({ name: o, value: o })),
  });
}

export async function checkbox(prompt, options = []) {
  if (!options.length) {
    throw new Error("checkbox: options is empty");
  }

  return await checkboxPrompt({
    message: prompt,
    choices: options.map((o) => ({ name: o, value: o })),
  });
}
