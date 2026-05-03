import { read } from "../fs/json/index.mjs";

function getSystemLocale() {
	return (
		Intl.DateTimeFormat().resolvedOptions().locale ||
		process.env.LC_ALL ||
		process.env.LANG ||
		"en-US"
	);
}

export function choose() {
	const locale = getSystemLocale();
	return locale.split("-")[0];
}

export function init(language = "en") {
	return read(`./languages/${language}.json`);
}
export default { choose, init };
