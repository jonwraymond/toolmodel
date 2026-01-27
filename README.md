# toolmodel

`toolmodel` is the canonical, MCP-aligned data model for tools across this
stack. It defines:
- what a tool is,
- how it is identified, and
- how it is bound to execution backends.

Target MCP protocol version: `2025-11-25` (via `toolmodel.MCPVersion`).

This module is intentionally small and dependency-light. All other libraries
(`toolindex`, `tooldocs`, `toolrun`, `toolcode`) build on it.

## Install

```bash
go get github.com/jonwraymond/toolmodel
```

## Core concepts

### Tool (canonical tool definition)

`toolmodel.Tool` embeds the official MCP Go SDK `mcp.Tool` and adds:
- `Namespace` for stable, canonical IDs,
- `Version` for optional tool versioning, and
- `Tags` for discovery/search layers.

Canonical IDs use:
- `name` when namespace is empty, or
- `namespace:name` when namespace is set.

Use `Tool.ToolID()` to compute the canonical ID.

### ToolBackend (execution binding)

Backends describe how a tool is executed. `toolmodel` defines three kinds:
- `mcp`
- `provider`
- `local`

These are execution metadata only. Transport/execution is handled by `toolrun`
via injected executors/registries.

## Quick start

Define a tool, validate it, and compute its canonical ID:

```go
import (
  "fmt"
  "log"

  "github.com/jonwraymond/toolmodel"
  "github.com/modelcontextprotocol/go-sdk/mcp"
)

t := toolmodel.Tool{
  Namespace: "tickets",
  Tool: mcp.Tool{
    Name:        "create",
    Description: "Create a support ticket",
    InputSchema: map[string]any{
      "type": "object",
      "properties": map[string]any{
        "title": {"type": "string"},
      },
      "required": []string{"title"},
    },
  },
  Tags: toolmodel.NormalizeTags([]string{"Tickets", "Support Ops"}),
}

if err := t.Validate(); err != nil {
  log.Fatal(err)
}

fmt.Println(t.ToolID()) // "tickets:create"
```

## Validation

`toolmodel` provides a default JSON Schema validator:
- `toolmodel.NewDefaultValidator()`

Key behaviors:
- JSON Schema 2020-12 is assumed when `$schema` is omitted
- draft-07 schemas are accepted
- external `$ref` resolution is disabled (no network access)

Example:

```go
import "log"

v := toolmodel.NewDefaultValidator()
if err := v.ValidateInput(&t, map[string]any{"title": "Help"}); err != nil {
  log.Fatal(err)
}
```

## Backends (examples)

```go
// MCP backend binding
mcpBackend := toolmodel.ToolBackend{
  Kind: toolmodel.BackendKindMCP,
  MCP:  &toolmodel.MCPBackend{ServerName: "github"},
}

// Provider backend binding
providerBackend := toolmodel.ToolBackend{
  Kind: toolmodel.BackendKindProvider,
  Provider: &toolmodel.ProviderBackend{
    ProviderID: "internal-api",
    ToolID:     "tickets.create",
  },
}

// Local backend binding
localBackend := toolmodel.ToolBackend{
  Kind:  toolmodel.BackendKindLocal,
  Local: &toolmodel.LocalBackend{Name: "tickets.create"},
}
```

## Version compatibility (current tags)

- `toolmodel`: `v0.1.0`
- `toolindex`: `v0.1.1`
- `tooldocs`: `v0.1.1`
- `toolrun`: `v0.1.0`
- `toolcode`: `v0.1.0`
- `toolruntime`: `v0.1.0`
- `metatools-mcp`: `v0.1.2`

Downstream modules should import tagged versions when wiring the stack
together.
