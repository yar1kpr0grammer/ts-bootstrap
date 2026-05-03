import fsSync from "fs";
import JSON5 from "json5";
import path from "path";

export function read(path) {
	return JSON5.parse(fsSync.readFileSync(path, "utf-8"));
}

export function write(path, data) {
	fsSync.writeFileSync(path, JSON5.stringify(data, null, 2), "utf8");
}

export default { read, write };
