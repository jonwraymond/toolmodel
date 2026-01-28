# toolmodel

Canonical tool schema definitions aligned to the MCP tool format.

## What this repo provides

- The `Tool` type (name, title, description, input/output schema)
- Stable tool IDs
- JSON Schema helpers

## Example

```go
import "github.com/jonwraymond/toolmodel"

tool := toolmodel.Tool{
  Name: "get_repo",
  Title: "Get Repo",
  Description: "Fetch repository metadata",
}
```
