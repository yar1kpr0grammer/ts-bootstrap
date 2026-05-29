import chalk from "chalk";
import { spawn } from "child_process";
import readline from "readline";

export function exec(command) {
  return new Promise((resolve, reject) => {
    const child = spawn(command, {
      shell: true,
      stdio: "pipe",
    });

    let stdout = "";
    let stderr = "";

    child.stdout.on("data", (data) => {
      stdout += data.toString();
    });

    child.stderr.on("data", (data) => {
      stderr += data.toString();
    });

    child.on("close", (code) => {
      if (code === 0) {
        resolve(stdout);
      } else {
        reject(new Error(stderr || stdout));
      }
    });

    child.on("error", (err) => {
      reject(err);
    });
  });
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

  return confirmVariants.includes(ans.trim().toLowerCase());
}

export function closeInput() {
  rl.close();
}
