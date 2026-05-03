import fs from "fs";
import JSON5 from "json5";
import Path from "../path.mjs";

export function read(filePath) {
  const fullPath = Path.resolveProjectPath(filePath);
  const raw = fs.readFileSync(fullPath, "utf-8");

  return JSON5.parse(raw);
}

export function write(filePath, data) {
  const fullPath = Path.resolveProjectPath(filePath);

  const content = JSON5.stringify(data, null, 2);

  fs.writeFileSync(fullPath, content, "utf-8");
}

export default { read, write };
