import chalk from "chalk";
import { select } from "@inquirer/prompts";
import checkboxPrompt from "@inquirer/checkbox";
import ora from "ora";

export async function confirm(
  fn,
  message,
  success = "Success",
  error = "Error",
) {
  const spinner = ora({
    text: chalk.cyan(message),
    color: "cyan",
  }).start();

  try {
    const res = await fn();

    spinner.succeed(`${chalk.bold.green(success)}: ${message}`);

    return res;
  } catch (e) {
    spinner.fail(`${chalk.bold.red(error)}: ${message}`);

    console.error(chalk.red(e instanceof Error ? e.message : String(e)));

    throw e;
  }
}

export async function radio(prompt, options = []) {
  if (!options.length) {
    throw new Error("radio: options is empty");
  }

  return select({
    message: chalk.bold(prompt),
    choices: options.map((option) => ({
      name: option,
      value: option,
    })),
  });
}

export async function checkbox(prompt, options = []) {
  if (!options.length) {
    throw new Error("checkbox: options is empty");
  }

  return checkboxPrompt({
    message: chalk.bold(prompt),
    choices: options.map((option) => ({
      name: option,
      value: option,
    })),
  });
}
