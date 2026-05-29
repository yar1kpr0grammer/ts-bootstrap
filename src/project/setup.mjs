import { radio, checkbox } from "../utils/cmd.mjs";
import { input, ask } from "../cmd.mjs";
import fs from "fs";
import path from "path";

import languageLib from "../utils/language.mjs";
import Path from "../fs/path.mjs";
import { checkProjectName } from "../npm.mjs";
import Dir from "../fs/dir.mjs";

export async function getName(
  prompt = "Enter project name: ",
  error = "Not valid",
) {
  process.stdin.setRawMode(false);
  process.stdin.resume();

  while (true) {
    const value = await input(prompt);

    // "." означает текущую папку
    if (value === ".") {
      const currentDirName = path.basename(process.cwd());

      if (checkProjectName(currentDirName)) {
        return {
          name: currentDirName,
          dir: ".",
        };
      }

      console.log(error);
      continue;
    }

    if (checkProjectName(value)) {
      return {
        name: value,
        dir: value,
      };
    }

    console.log(error);
  }
}

export async function chooseNodeType(language) {
  return await radio(language.choose_node_type, ["module", "commonjs"]);
}

export async function prepareDir(dirPath, language) {
  if (!Path.exists(dirPath)) {
    Dir.create(dirPath);
  }

  Path.cd(dirPath);

  if (await Dir.hasFiles()) {
    const shouldClear = await ask(
      language.dir_is_not_empty,
      language.question_prompt,
      language.confirm_variants,
    );

    if (shouldClear) {
      await Dir.clear();
    }
  }
}

export async function setupProject() {
  const language = languageLib.init(languageLib.choose());

  const { name, dir } = await getName(language.enter_project_name);

  console.log();

  const configPath = Path.resolveCliPath("config.json");

  const config = JSON.parse(fs.readFileSync(configPath, "utf-8"));

  await prepareDir(dir, language);
  const recommendedOptions = ["git", "README.md", "node", "src dir @ alias"];
  const allOptions = [...recommendedOptions, "eslint", "prettier"];

  console.log();

  const template = await radio(language.choose_project_template, [
    language.template_recommended,
    language.template_custom,
  ]);

  switch (template) {
    case language.template_recommended:
      config.options = recommendedOptions;
      config.nodeType = "module";
      break;

    case language.template_custom:
      config.options = await checkbox(
        language.choose_project_template,
        allOptions,
      );
      console.log();
      config.nodeType = await chooseNodeType(language);
      break;
  }

  console.log();

  return {
    name,
    language,
    config,
    dir,
  };
}
