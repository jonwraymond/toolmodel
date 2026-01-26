# toolmodel

Canonical Go data model for MCP-style tools. This module defines tool schemas, stable IDs, backend bindings, and JSON Schema validation helpers. It is the base dependency for other tool libraries (toolindex, tooldocs, toolrun, toolcode).

## Project Layout
This module follows the official Go guidance for simple packages:

- Keep the package flat at the module root (`tool.go`, `validator.go`, `*_test.go`).
- Avoid `cmd/` and `pkg/` for library-only modules.
- Use `internal/` only when you need a strict, private package boundary.

References: the Go team’s “Organizing a Go module” guide. ([go.dev](https://go.dev/doc/modules/layout?utm_source=openai))

## Status
- MCP-compatible Tool model with provider/local backends
- JSON Schema validation via `github.com/google/jsonschema-go/jsonschema`
- External `$ref` resolution disabled by default
