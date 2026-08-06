# Basic / Simple Plugin Template

This template provides a minimal reference implementation for Ostrich SDK Osplates. It demonstrates how parameter rendering, environment variable resolution, and different task runner kinds (`inprocess` and `container`) function within the Ostrich execution engine.

## Features

- **In-process Execution**: Runs native Python scripts (`.py.tmpl`) directly inside the engine process.
- **Container Sandboxing**: Executes multi-command container tasks using Docker (`alpine` and `alpine/curl`).
- **Jinja Parameter Interpolation**: Illustrates custom delimiters (`[[` / `]]`) and Jinja filters (`here`, `env`).
- **Argument Forwarding**: Passes CLI operation arguments down to container commands via `ARGC`, `ARGV`, and `ARGV_JSON`.

## Configuration

Sample configuration file is located at `_doc/ostrich.yaml`:

```yaml
plugin:
  name: simple-test
  version: 0.0.1
  business_name: A dummy test plugin

template:
  kind: basic
  runtime: docker
  params:
    message: "This is a message from the plugin configuration"
```

## Tasks

- **`display`**:
  - **Runner**: `inprocess` (Python)
  - **Description**: Evaluates environment variables (`HOME`) and logs plugin metadata and custom messages using the standard Python `logging` module.
  - **Usage**: `ost run display`

- **`displaydocker`**:
  - **Runner**: `container` (`alpine:latest`, `alpine/curl:latest`, `kurrier/alpine-jq:latest`)
  - **Description**: Executes sequential containerized commands, inspecting local workspace directory contents, verifying environment variables (`TEST_VAR`, `LOGLEVEL`, `DRYRUN`), reading template operation data files (`data.txt`), and parsing command line parameters using `jq`.
  - **Usage**: `ost run displaydocker [args...]`
