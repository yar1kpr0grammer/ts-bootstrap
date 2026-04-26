#!/usr/bin/env node

const os = require("os");
const path = require("path");
const { spawnSync } = require("child_process");

const platform = os.platform();

let binary;

if (platform === "linux") {
	binary = "ts-vibe-linux";
} else if (platform === "win32") {
	binary = "ts-vibe-windows.exe";
} else if (platform === "darwin") {
	binary = "ts-vibe-mac";
} else {
	console.error("Unsupported OS");
	process.exit(1);
}

const binPath = path.join(__dirname, "bin", "release", binary);

spawnSync(binPath, process.argv.slice(2), {
	stdio: "inherit",
});
