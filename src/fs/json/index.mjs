import fs from "fs";
import Path from "../path.mjs";

export function read(filePath, type = "project") {
  const fullPath =
    type === "cli"
      ? Path.resolveCliPath(filePath)
      : Path.resolveProjectPath(filePath);

  const raw = fs.readFileSync(fullPath, "utf-8");
  return JSON.parse(raw);
}

export function write(filePath, data, type = "project") {
  const fullPath =
    type === "cli"
      ? Path.resolveCliPath(filePath)
      : Path.resolveProjectPath(filePath);

  fs.writeFileSync(fullPath, JSON.stringify(data, null, 2), "utf8");
}

export default { read, write };
