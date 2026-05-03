import fsSync from "fs";
import fsPromises from "fs/promises";
import path from "path";

async function clear(folderPath = ".") {
  const items = await fsPromises.readdir(folderPath);

  for (const item of items) {
    const fullPath = path.join(folderPath, item);
    await fsPromises.rm(fullPath, {
      recursive: true,
      force: true,
    });
  }
}

export function create(path) {
  fsSync.mkdirSync(path, { recursive: true });
}

export async function hasFiles(dirPath = ".") {
  const items = await fsPromises.readdir(dirPath);

  for (const item of items) {
    const fullPath = path.join(dirPath, item);
    const stat = await fsPromises.stat(fullPath);
    if (stat.isFile()) {
      return true;
    }
  }
  return false;
}

export default { clear, create, hasFiles };
