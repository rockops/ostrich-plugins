# AI Agent Guidelines - Ostrich Plugins

This document provides guidelines, architecture context, and technical standards for AI agents operating on the `ostrich-plugins` repository.

---

## 1. Overview & Repository Purpose

`ostrich-plugins` contains **Osplates** (Ostrich Templates) used by the `ostrich-sdk` engine. 

- **Role**: An Osplate defines defaults, task operations, and templated application configurations (Docker, Kubernetes/Helm, Buildpacks, etc.).
- **Engine Interaction**: When a user runs commands using `ost-core`, `ostd`, or `ostr` (from `ostrich-sdk`), the SDK loads the target Osplate from this repository, injects `ostrich.yaml` configuration parameters, renders templates, and executes specified tasks.

---

## 2. Directory Structure & Osplate Anatomy

```
ostrich-plugins/
├── basic/                  # Basic plugin template
│   ├── template.yaml       # Defines template metadata & default runner
│   ├── _doc/               # Plugin documentation
│   ├── display/            # 'display' task operation
│   └── displaydocker/      # 'displaydocker' task operation
└── backend/                # Cloud Native Buildpacks backend template
    ├── src/                # Root template directory
    │   ├── template.yaml   # Template metadata & version
    │   ├── default.yaml    # Default parameters injected into templates
    │   ├── _doc/           # Plugin documentation (description.md)
    │   ├── docker/         # 'docker' task operation (Cloud Native Buildpacks pack build)
    │   ├── helm/           # 'helm' task operation (Kubernetes deploy/delete/template)
    │   ├── dev/            # 'dev' task operation (local run)
    │   ├── ps/             # 'ps' task operation (container status)
    │   ├── stop/           # 'stop' task operation (container teardown)
    │   ├── log/            # 'log' task operation (container logging)
    │   └── cache/          # 'cache' task operation (pack volume management)
    └── _test/              # Template unit & integration test suites
```

---

## 3. Jinja2 Templating Rules & Custom Delimiters

> [!IMPORTANT]
> **ALWAYS use double square brackets `[[` and `]]` for Jinja2 variable interpolation in Osplates.**
> Standard Jinja `{{ }}` and `{% %}` delimiters clash with Helm charts and Kubernetes manifests and MUST NOT be used.

### Delimiter Syntax
- **Variable Interpolation**: `[[ variable ]]`
- **Control Blocks**: `[% for item in list %] ... [% endfor %]`
- **Comments**: `[# comment #]`

### Custom Jinja Filters & Injected Globals
- `noslash`: Removes leading and trailing slashes from strings.
  - *Example*: `[[ template.params.registry | noslash ]]`
- `input("key")`: Resolves input directory sources.
  - *Example*: `[[ "" | input("src") ]]`
- `here`: Formats path relative to current execution context.
  - *Example*: `[[ template.params.output.helm | here ]]`
- `fromTemplateInstance`: Resolves path within the generated template instance.
  - *Example*: `[[ "helm/kubernetes" | fromTemplateInstance ]]`
- `_dryRun`: Injected boolean flag available in Python `inprocess` scripts to log actions without mutating system state.
- `_argv`: Injected list of command-line arguments passed to the task.

---

## 4. Operation Runners

Tasks are located in subdirectories named after the operation (e.g., `docker/`, `helm/`, `dev/`). An operation specifies a runner in its configuration or inherits from `template.yaml`:

### A. `inprocess` (Python Engine Scripts)
Executes a rendered Python script natively inside `ost-core`.
- Example (`helm/helm.py.tmpl`): Handles complex subcommands (`deploy`, `delete`, `template`).
- Globals provided: `run()`, `OstrichException`, `logging`, `_dryRun`, `_argv`.

### B. `container` (Sandboxed CLI Tools)
Executes within a specified container image.
- Example (`docker/docker.yaml.tmpl`):
  ```yaml
  runner:
    kind: container
    image: buildpacksio/pack:latest
  ```

### C. `shell` (Host Commands)
Runs local shell commands directly on the host or container environment.

---

## 5. Development & Testing Workflow

When adding or editing an Osplate:
1. Update `template.yaml` metadata (name, version, description).
2. Document new tasks or parameters in `src/_doc/description.md`.
3. Provide sane defaults in `src/default.yaml`.
4. Test template rendering using `ostrich-sdk`:
   ```bash
   export PYTHONPATH=<path-to-ostrich-sdk>/ost-core:$PYTHONPATH
   python ost-core/ost template test <template_name>
   ```

---

## 6. Code Guidelines for AI Agents

- **No Hardcoded Values**: Always parameterize configurable items in `default.yaml` and reference them via `[[ template.params... ]]`.
- **Error Handling**: Use `OstrichException` in `inprocess` Python templates to report failures gracefully.
- **Structured Logging**: Use Python standard `logging.info()` / `logging.error()` instead of `print()` inside `.py.tmpl` scripts.
- **Idempotency**: Operations (like `deploy` or `stop`) should be safe to run multiple times.
