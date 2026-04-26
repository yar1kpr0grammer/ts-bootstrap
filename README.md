# ts-vibe

CLI tool for bootstrapping TypeScript projects

---

## 🚀 Features

- Create TypeScript project structure
- Optional Git initialization
- Automatic README generation
- Path validation (prevents invalid/non-ASCII paths)
- Run existing projects
- Execution time tracking for operations

---

## ⚙️ Usage

### Create a new project

```bash
npx ts-vibe init
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

- Go 1.20+
- Node.js (for running TypeScript projects)
- Git (optional)

---

## 🛠 Example Workflow

```bash
npx ts-vibe init --git --lang=ru
cd my-project
npm start
```
