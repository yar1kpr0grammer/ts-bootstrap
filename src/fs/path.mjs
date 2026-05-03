import fsSync from "fs";
import fsPromises from "fs/promises";
import JSON5 from "json5";
import path from "path";

function cd(name) {
	try {
		process.chdir(name);
		return { success: true, path: process.cwd() };
	} catch (err) {
		throw new Error(err);
	}
}

function exists(dirPath) {
	try {
		fsSync.accessSync(dirPath, fsSync.constants.F_OK | fsSync.constants.R_OK);
		return true;
	} catch (err) {
		return false;
	}
}

export default { cd, exists };
