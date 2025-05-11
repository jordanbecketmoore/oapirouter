# oapirouter
A library for converting OpenAPI specification documents into HTTPRoutes for the
Kubernetes Gateway API

#  `oapirouter` CLI 
The `oapirouter` CLI will convert OpenAPI documents into HTTPRoute YAML manifests. 

## help
```
Usage: oapirouter [OPTIONS] COMMAND [ARGS]...

A CLI tool to convert OpenAPI specification documents into HTTPRoute YAML manifests for the Kubernetes Gateway API.

Options:
  -h, --help         Show this help message and exit.
  -v, --version      Show the version of the `oapirouter` CLI.

Commands:
  convert            Convert an OpenAPI document into HTTPRoute YAML manifests.
  validate           Validate an OpenAPI document for compatibility with HTTPRoute.

Run 'oapirouter COMMAND --help' for more information on a specific command.
```