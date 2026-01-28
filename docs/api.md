# API Reference

This is a concise map of the public types and interfaces. See code for details.

## Tool

`toolmodel.Tool` embeds `mcp.Tool` and adds:

- `Namespace string`
- `Version string`
- `Tags []string`

Common fields from `mcp.Tool` used in this stack:

- `Name string`
- `Title string`
- `Description string`
- `InputSchema any`
- `OutputSchema any`

### IDs

- `Tool.ToolID() string`
- `ParseToolID(id string) (namespace, name string, err error)`

## Backends

```go
type BackendKind string
const (
  BackendKindMCP      BackendKind = "mcp"
  BackendKindProvider BackendKind = "provider"
  BackendKindLocal    BackendKind = "local"
)

type ToolBackend struct {
  Kind     BackendKind
  MCP      *MCPBackend
  Provider *ProviderBackend
  Local    *LocalBackend
}

type MCPBackend struct {
  ServerName string
}

type ProviderBackend struct {
  ProviderID string
  ToolID     string
}

type LocalBackend struct {
  Name string
}
```

## Validation

```go
type SchemaValidator interface {
  Validate(schema any, instance any) error
  ValidateInput(tool *Tool, args any) error
  ValidateOutput(tool *Tool, result any) error
}

func NewDefaultValidator() *DefaultValidator
```

## Utilities

- `NormalizeTags([]string) []string`
- `Tool.Validate() error`
- `ToolBackend.Validate() error`
