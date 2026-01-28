# Architecture

`toolmodel` sits at the bottom of the stack. Everything else consumes its types.

```mermaid
flowchart LR
  A[toolmodel.Tool] --> B[toolindex]
  A --> C[tooldocs]
  A --> D[toolrun]
  A --> E[metatools-mcp]

  subgraph Binding
    F[ToolBackend]
  end

  F --> D
```

## Key decisions

- Embeds the official MCP Go SDK `mcp.Tool` for 1:1 protocol alignment.
- Adds `Namespace`, `Version`, and `Tags` without altering MCP semantics.
- Keeps validation local and dependency-light (`jsonschema-go`).
