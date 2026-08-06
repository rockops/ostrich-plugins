# CLAUDE.md - Ostrich Plugins Guidelines

## Repository Overview
`ostrich-plugins` contains Osplates (templates) executed by the `ostrich-sdk` engine.

## Key Rules for AI Agents

### 1. Jinja Templating Delimiters
- **ALWAYS** use `[[` and `]]` for Jinja variable interpolation.
- **ALWAYS** use `[%` and `%]` for loops and conditional control blocks.
- **NEVER** use standard `{{ }}` or `{% %}` as they conflict with Helm and Kubernetes syntax.

### 2. Osplate Structure
- `template.yaml`: Template metadata (name, version, description, default runner).
- `default.yaml`: Default parameters injected into `template.params`.
- `_doc/description.md`: Documentation for tasks supported by the Osplate.
- Operation Folders (`docker/`, `helm/`, `dev/`, `ps/`, `stop/`, `log/`, `cache/`): Task definitions.

### 3. Task Runners
- `inprocess`: Native Python script (`.py.tmpl`) running inside `ost-core`.
- `container`: Containerized tool execution (e.g. `buildpacksio/pack:latest`).
- `shell`: Host shell commands.

### 4. Custom Jinja Filters & Globals
- `noslash`: Strip leading/trailing slashes.
- `input("key")`: Resolve input path.
- `here`: Resolve output path.
- `fromTemplateInstance`: Resolve path relative to template instance.
- `_dryRun` & `_argv`: Injected globals for `inprocess` Python scripts.

## Testing Templates
Run integration test suite via `ostrich-sdk`:
```bash
export PYTHONPATH=$(pwd)/../ostrich-sdk/ost-core:$PYTHONPATH
python ../ostrich-sdk/ost-core/ost template test <template-name>
```
