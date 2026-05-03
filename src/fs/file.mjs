import fsSync from "fs";

function create(path, content) {
  fsSync.writeFileSync(path, content);
}

function append(path, content) {
  fsSync.appendFileSync(path, content);
}

export default { create, append };
