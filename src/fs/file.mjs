import fsSync from "fs";
import Path from "./path.mjs";

function create(filePath, content) {
  const fullPath = Path.resolveProjectPath(filePath);
  fsSync.writeFileSync(fullPath, content);
}

function append(filePath, content) {
  const fullPath = Path.resolveProjectPath(filePath);
  fsSync.appendFileSync(fullPath, content);
}

export default { create, append };
