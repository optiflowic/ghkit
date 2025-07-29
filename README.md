# ghkit

**ghkit** is a CLI tool for installing GitHub repository templates such as issue templates, pull request templates, and meta files like CODEOWNERS and CONTRIBUTING.md.

## 📦 Installation

You can install `ghkit` via `go install` or Homebrew.

### With Go

```bash
go install github.com/optiflowic/ghkit@latest
```

### With Homebrew

```bash
brew install optiflowic/tap/ghkit
```

## 🚀 Usage

Add all templates:

```bash
ghkit add all --path ./your-repo
```

Add a specific issue template (e.g., `bug`):

```bash
ghkit add issue bug --path ./your-repo
```

Add a pull request template:

```bash
ghkit add pr --path ./your-repo
```

Add meta templates like `CODEOWNERS`, `CONTRIBUTING.md`, etc.:

```bash
ghkit add meta codeowners --path ./your-repo
```

## 🧾 Available Issue Templates

- `bug`
- `feature`
- `question`
- `task`
- `docs`
- `feedback`
- `config`

## 🛠️ Options

Most subcommands support the following flags:

- `--format`, `-f`: Specify the format of the issue template. Options: `yml`, `md`. Default: `yml`
- `--lang`, `-l`: Language for the template. Options: `en`, `ja`. Default: `en`
- `--path`: Root path to your repository. Instead, use the root path (e.g., `./your-repo`). Default: `.`
- `--force`: Overwrite existing files.
- `--verbose`: Outputs the log.
- `--debug`: Outputs the debug log.

## 💡 Examples

Add a Japanese markdown template for feature requests:

```bash
ghkit add issue feature --format md --lang ja --path ./your-repo
```

Add all templates to a repository root and overwrite if needed:

```bash
ghkit add all --format yml --lang en --path ./your-repo --force
```

## 📁 Output Structure

Given `--path ./your-repo`, files will be created in:

```bash
your-repo/
└── .github/
    ├── ISSUE_TEMPLATE/
    │   └── bug.yml, feature.yml, ...
    ├── PULL_REQUEST_TEMPLATE.md
    ├── CODEOWNERS
    ├── CONTRIBUTING.md
    └── ...
```

## ⚖️ License

This project is licensed under the [MIT License](./LICENSE).
