# ts-vibe

CLI tool for bootstrapping TypeScript projects written on golang

---

## 🚀 Features

- Create TypeScript project structure
- Optional Git initialization
- Automatic README generation
- Path validation (prevents invalid/non-ASCII paths)
- Execution time tracking for operations
- 🇷🇺 Ru / 🇺🇲 En languages support

---

## ⚙️ Usage

### Create a new project

```bash
npx ts-vibe
```

With Git:

```bash
npx ts-vibe init --git
```

Skip README creation:

```bash
npx ts-vibe init --noReadme
```

---

### Run project

```bash
npm start
```

---

## 🧠 Flags

| Flag         | Alias | Description               |
| ------------ | ----- | ------------------------- |
| `--init`     | `-i`  | Create a new project      |
| `--clear`    | `-c`  | Clear current directory   |
| `--git`      | `-g`  | Initialize git repository |
| `--lang`     | `-l`  | Choose cli lang en/ru     |
| `--noReadme` | -     | Skip README creation      |

---

## ⚠️ Requirements

- Go 1.20+ (for build from source)
- Node.js (for running TypeScript projects)
- Git (optional)

---

## 🛠 Example Workflow

```bash
npx ts-vibe init --git --lang=ru
cd my-project
npm start
```

---

## Buid from source

Clone the source:
[source](https://github.com/yar1kpr0grammer/ts-vibe)

Run `Makefile` in the root dir:
```bash
make releace
```
it will compile go source to linux / mac / windows
