# toolmodel

`toolmodel` is the canonical, MCP-aligned data model for tools across the stack.
It embeds the official MCP Go SDK `mcp.Tool`, adds namespace and tagging, and
provides validation helpers for JSON Schema inputs/outputs.

## What this library provides

- Canonical tool definition (`toolmodel.Tool`)
- Stable IDs (`Tool.ToolID()` + `ParseToolID`)
- Backend bindings (`ToolBackend`) for execution layers
- Tag normalization for discovery (`NormalizeTags`)
- JSON Schema validation via `SchemaValidator`

## Quickstart

```go
package main

import (
  "fmt"
  "log"

  "github.com/jonwraymond/toolmodel"
  "github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
  tool := toolmodel.Tool{
    Namespace: "github",
    Tool: mcp.Tool{
      Name:        "get_repo",
      Description: "Fetch repository metadata",
      InputSchema: map[string]any{
        "type": "object",
        "properties": map[string]any{
          "owner": {"type": "string"},
          "repo":  {"type": "string"},
        },
        "required": []string{"owner", "repo"},
      },
    },
    Tags: toolmodel.NormalizeTags([]string{"GitHub", "repos"}),
  }

  if err := tool.Validate(); err != nil {
    log.Fatal(err)
  }

  fmt.Println(tool.ToolID()) // github:get_repo
}
```

## Next

- Architecture and placement in the stack: `architecture.md`
- How to use tags, backends, and validation: `usage.md`
- Additional examples: `examples.md`
