import { radio, checkbox } from "../utils/cmd.mjs";
import { input, ask } from "../cmd.mjs";
import { read } from "../fs/json/index.mjs";
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
  let name;
  while (true) {
    name = await input(prompt);
    if (checkProjectName(name)) {
      break;
    }
    console.log(error);
  }
  return name;
}

export async function chooseNodeType(language) {
  return await radio(language.choose_node_type, ["commonjs", "module"]);
}

export async function prepareDir(path, language) {
  if (!Path.exists(path)) {
    Dir.create(path);
  }

  Path.cd(path);

  if (await Dir.hasFiles()) {
    if (
      await ask(
        language.dir_is_not_empty,
        language.question_prompt,
        language.confirm_variants,
      )
    ) {
      await Dir.clear();
    }
  }
}

export async function setupProject() {
  const language = languageLib.init(languageLib.choose());
  const name = await getName(language.enter_project_name);
  console.log();

  const config = read("config.json");
  await prepareDir(name, language);

  const allOptions = ["git", "eslint", "prettier", "README.md", "node"];
  console.log();
  const res = await radio("Выберите вариант проекта:", [
    language.template_min,
    language.template_max,
    language.template_custom,
  ]);
  switch (res) {
    case language.template_max:
      config.options = allOptions;
      break;
    case language.template_min:
      config.options = [];
      break;
    case language.template_custom:
      config.options = await checkbox(
        language.choose_project_template,
        allOptions,
      );
      break;
  }
  console.log();

  config.nodeType = await chooseNodeType(language);
  console.log();

  return { name, language, config };
}
