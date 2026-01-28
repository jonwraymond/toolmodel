# Architecture

`toolmodel` sits at the bottom of the stack. Everything else consumes its types
and uses it as the canonical source of truth.

## Component view

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

## Data model view

```mermaid
classDiagram
  class Tool {
    +string Namespace
    +string Name
    +string Version
    +[]string Tags
    +ToolID() string
  }

  class ToolBackend {
    +BackendKind Kind
  }

  class MCPBackend {
    +string ServerName
  }

  class ProviderBackend {
    +string ProviderID
    +string ToolID
  }

  class LocalBackend {
    +string Name
  }

  ToolBackend --> MCPBackend
  ToolBackend --> ProviderBackend
  ToolBackend --> LocalBackend
```

## Validation pipeline

```mermaid
flowchart LR
  A[Tool.InputSchema] --> B[DefaultValidator]
  B --> C[Resolve schema]
  C --> D[Validate instance]
  D --> E[Error or success]
```

## Design notes

- Embeds `mcp.Tool` to stay aligned with the official MCP SDK.
- Adds `Namespace`, `Version`, and `Tags` without altering MCP semantics.
- Keeps validation dependency-light and deterministic.
