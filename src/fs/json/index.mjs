import fsSync from "fs";
import path from "path";

export function read(path) {
	return JSON.parse(fsSync.readFileSync(path, "utf-8"));
}

export function write(path, data) {
	fsSync.writeFileSync(path, JSON.stringify(data, null, 2), "utf8");
}

export default { read, write };
