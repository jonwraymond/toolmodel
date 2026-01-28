# Usage

## Define a tool

```go
import (
  "github.com/jonwraymond/toolmodel"
  "github.com/modelcontextprotocol/go-sdk/mcp"
)

tool := toolmodel.Tool{
  Namespace: "tickets",
  Tool: mcp.Tool{
    Name:        "create",
    Description: "Create a support ticket",
    InputSchema: map[string]any{
      "type": "object",
      "properties": map[string]any{
        "title":    {"type": "string"},
        "priority": {"type": "string", "enum": []string{"low", "high"}},
      },
      "required": []string{"title"},
    },
  },
  Tags: toolmodel.NormalizeTags([]string{"Support", "Tickets"}),
}
```

## Tool IDs

```go
id := tool.ToolID()                 // tickets:create
ns, name, _ := toolmodel.ParseToolID(id)
```

## Backends

`toolmodel` does not execute anything. It only describes how a tool can be
executed. These backends are consumed by `toolrun`.

```go
mcpBackend := toolmodel.ToolBackend{
  Kind: toolmodel.BackendKindMCP,
  MCP:  &toolmodel.MCPBackend{ServerName: "github"},
}

providerBackend := toolmodel.ToolBackend{
  Kind: toolmodel.BackendKindProvider,
  Provider: &toolmodel.ProviderBackend{
    ProviderID: "internal-api",
    ToolID:     "tickets.create",
  },
}

localBackend := toolmodel.ToolBackend{
  Kind: toolmodel.BackendKindLocal,
  Local: &toolmodel.LocalBackend{Name: "create_ticket"},
}
```

## Validation

```go
validator := toolmodel.NewDefaultValidator()

if err := validator.ValidateInput(&tool, map[string]any{"title": "Help"}); err != nil {
  // handle validation error
}
```

### Dialects and safety

- JSON Schema 2020-12 is assumed when `$schema` is missing.
- draft-07 is accepted (normalized internally).
- External `$ref` is blocked (no network resolution).
