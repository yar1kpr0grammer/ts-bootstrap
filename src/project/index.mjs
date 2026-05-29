import { exec } from "../cmd.mjs";
import { updatePackageJSON, addScript } from "../npm.mjs";
import File from "../fs/file.mjs";
import Dir from "../fs/dir.mjs";
import ts from "../ts.mjs";
import { confirm } from "../utils/cmd.mjs";

async function setupGit(language, config) {
  File.create(".gitignore", config.ignores.git);
  File.create(".gitattributes", "* text=auto\n");
  const commands = [
    { command: "git init", message: "git init" },
    {
      command: "git add .",
      message: "git add .",
    },
    {
      command: 'git commit -m "init project with ts-vibe"',
      message: "git commit",
    },
  ];
  for (let cmd of commands) {
    await confirm(
      () => exec(cmd.command),
      cmd.message,
      language.success,
      language.error,
    );
  }
}

async function setupLint(language, config) {
  addScript("lint", "eslint .");
  File.create("eslint.config.js", config.default_code.eslintConfig);
  File.create(".eslintignore", config.ignores.eslint);
  await confirm(
    () =>
      exec(
        "npm i -D eslint @typescript-eslint/parser @typescript-eslint/eslint-plugin eslint-config-prettier",
      ),
    "npm i -D eslint @typescript-eslint/parser @typescript-eslint/eslint-plugin eslint-config-prettier",
    language.success,
    language.error,
  );
  if (config.options.includes("README.md")) {
    File.append("README.md", language.lint_tip);
  }
}

async function setupPrettier(language, config) {
  addScript("format", "prettier --write .");
  File.create(".prettierrc", config.default_code.prettierrc);
  File.create(".prettierignore", config.ignores.prettier);
  await confirm(
    () => exec("npm i -D prettier "),
    "npm i -D  prettier",
    language.success,
    language.error,
  );
  if (config.options.includes("README.md")) {
    File.append("README.md", language.prettier_tip);
  }
}

async function setupReadme(language) {
  await confirm(
    () => File.create("README.md", language.readme_content),
    language.create_readme,
    language.success,
    language.error,
  );
}

async function setupNode(language) {
  await confirm(
    async () => {
      await exec("npm i -D @types/node");
      ts.addType("node");
    },
    "npm i -D @types/node",
    language.success,
    language.error,
  );
}
export async function init(language, config) {
  const commands = [
    { command: "npm init -y", message: "npm init -y" },
    {
      command: "npm i -D typescript tsx @types/node",
      message: "npm i -D typescript",
    },
    { command: "npx tsc --init", message: "npx tsc --init" },
  ];

  for (const cmd of commands) {
    await confirm(
      () => exec(cmd.command),
      cmd.message,
      language.success,
      language.error,
    );
  }

  ts.readConfig();

  await confirm(
    () => ts.updateTsConfig(config),
    language.update_ts_config,
    language.success,
    language.error,
  );

  await confirm(
    () => updatePackageJSON(config.nodeType),
    language.update_package,
    language.success,
    language.error,
  );

  Dir.create("src");
  File.create("src/index.ts", config.default_code.indexTs);

  if (config.options.includes("node")) {
    await setupNode(language);
  }
  if (config.options.includes("README.md")) {
    await setupReadme(language);
  }

  if (config.options.includes("eslint")) {
    await setupLint(language, config);
  }

  if (config.options.includes("prettier")) {
    await setupPrettier(language, config);
  }

  if (config.options.includes("git")) {
    await setupGit(language, config);
  }
}
