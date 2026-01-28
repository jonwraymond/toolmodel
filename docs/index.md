# toolmodel

`toolmodel` is the canonical, MCP-aligned data model for tools across the stack.
It embeds the official MCP Go SDK `mcp.Tool`, adds namespace + tags, and provides
schema validation helpers.

## Key APIs

- `Tool` (embeds `mcp.Tool`, adds `Namespace`, `Version`, `Tags`)
- `ToolBackend` (mcp/provider/local binding)
- `SchemaValidator` + `NewDefaultValidator()`
- `NormalizeTags`, `ToolID`, `ParseToolID`

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
- Usage patterns and validation: `usage.md`
- Additional examples: `examples.md`
