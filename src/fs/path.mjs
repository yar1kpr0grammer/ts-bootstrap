import fsSync from "fs";
import path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

function cd(name) {
  try {
    process.chdir(name);
    return {
      success: true,
      path: process.cwd(),
    };
  } catch (err) {
    throw new Error(err.message);
  }
}

// путь до самого cli пакета
export function resolveCliPath(p) {
  if (path.isAbsolute(p)) return p;
  return path.join(__dirname, "../../", p);
}

// путь до текущего проекта
export function resolveProjectPath(p) {
  if (path.isAbsolute(p)) return p;
  return path.join(process.cwd(), p);
}

function exists(dirPath) {
  try {
    fsSync.accessSync(dirPath, fsSync.constants.F_OK);
    return true;
  } catch {
    return false;
  }
}

export default {
  cd,
  exists,
  resolveCliPath,
  resolveProjectPath,
};
