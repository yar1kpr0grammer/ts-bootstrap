# tsBootstrup

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
tsbootstrap init
```

With Git:

```bash
tsbootstrap init --git
```

Skip README creation:

```bash
tsbootstrap init --noReadme
```

---

### Run project

```bash
tsbootstrap run
```

or shorthand:

```bash
tsbootstrap r
```
---

## 🧠 Flags

| Flag         | Alias | Description               |
| ------------ | ----- | ------------------------- |
| `--init`     | `-i`  | Create a new project      |
| `--run`      | `-r`  | Run the project           |
| `--git`      | `-g`  | Initialize git repository |
| `--noReadme` | -     | Skip README creation      |

---

## ⚠️ Requirements

* Go 1.20+
* Node.js (for running TypeScript projects)
* Git (optional)

---

## 🛠 Example Workflow

```bash
tsbootstrap init --git
cd my-project
tsbootstrap run
```
