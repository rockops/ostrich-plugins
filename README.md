# Ostrich Plugins (Osplates)

This repository contains official **Osplates** (Ostrich Templates) for use with [Ostrich SDK](file:///home/ben/src/ostrich/ostrich-sdk).

## Osplate Documentation Standard (`_doc/` folder)

Every Osplate in this repository includes a dedicated `_doc/` folder containing standardized documentation files:

1. **`ostrich.yaml`**: A comprehensive sample configuration file illustrating all supported parameters, default settings, and inline comments explaining their behavior.
2. **`description.md`**: A detailed markdown document outlining the Osplate's purpose, features, required dependencies, and available tasks/subcommands.

---

## Available Osplates

- **[basic](file:///home/ben/src/ostrich/ostrich-plugins/basic)**: Base template with simple in-process Python (`display`) and containerized Docker (`displaydocker`) tasks. See its [description.md](file:///home/ben/src/ostrich/ostrich-plugins/basic/_doc/description.md) and sample [ostrich.yaml](file:///home/ben/src/ostrich/ostrich-plugins/basic/_doc/ostrich.yaml).
- **[backend](file:///home/ben/src/ostrich/ostrich-plugins/backend)**: Production-ready template for building backend microservices using Cloud Native Buildpacks (`pack`), managing local dev containers, and deploying via Helm to Kubernetes. See its [description.md](file:///home/ben/src/ostrich/ostrich-plugins/backend/src/_doc/description.md).

---

## AI Agent & Developer Documentation

- **[AGENTS.md](file:///home/ben/src/ostrich/ostrich-plugins/AGENTS.md)**: Technical guidelines, directory architecture, Jinja delimiter rules (`[[` / `]]`), custom filters, and runner execution details for AI agents.
- **[CLAUDE.md](file:///home/ben/src/ostrich/ostrich-plugins/CLAUDE.md)**: Quick reference developer & AI rules for building and testing templates.
