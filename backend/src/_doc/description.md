# Backend Template

This template allows you to build a Docker image for your backend application using **Cloud Native Buildpacks**.

## Features
- Automatic detection of language (Java, Node.js, etc.)
- Builds a Docker image using `pack`.

## Tasks
- **docker**: Build the Docker image using the `pack` CLI.
- **dev**: Run the application container.
- **ps**: List running containers for this project.
- **stop**: Stop any running containers for this image.
- **helm**: Manage Helm deployment.
  - `ost run helm deploy`: Deploy the application to Kubernetes using Helm.
  - `ost run helm delete`: Delete the application from Kubernetes using Helm.
  - `ost run helm template`: Perform a Helm template and write files to the output folder.
- **log**: Show logs of the running container.
- **cache**: Manage buildpack cache volumes.
