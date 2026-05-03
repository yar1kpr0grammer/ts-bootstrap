import chalk from "chalk";
import { spawnSync } from "child_process";
import readline from "readline";

export function exec(command) {
  const result = spawnSync(command, {
    shell: true,
    encoding: "utf-8",
  });

  if (result.status !== 0) {
    throw new Error(result.stderr || result.stdout);
  }

  return result.stdout;
}

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

export const input = (prompt) =>
  new Promise((resolve) => rl.question(chalk.bold(prompt), resolve));

export async function ask(
  prompt,
  questionPrompt = " [y/n] ",
  confirmVariants = ["y", "yes"],
) {
  const ans = await input(chalk.bold(prompt + questionPrompt));
  return confirmVariants.includes(ans.toLowerCase());
}

export function closeInput() {
  rl.close();
}
