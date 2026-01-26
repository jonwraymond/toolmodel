package toolmodel

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ErrInvalidToolID is returned when a tool ID string is malformed.
var ErrInvalidToolID = errors.New("invalid tool ID format")

// Decision Log:
// We evaluate the official MCP Go SDK (github.com/modelcontextprotocol/go-sdk/mcp)
// and choose to embed mcp.Tool in our Tool struct.
// - Usage: Embedding allows us to inherit all standard fields and JSON tags from the SDK,
//   ensuring 1:1 compatibility with the spec as interpreted by the official SDK.
// - Gaps: mcp.Tool uses `any` for InputSchema and OutputSchema, which is correct for
//   flexibility but requires us to handle validation carefully (which is a separate requirement).
//   mcp.Tool does not support Namespace or Version, so we add them.
// - Type Aliasing: We use type aliasing for ToolIcon (mcp.Icon) as it matches our needs.

// Tool mirrors the MCP Tool definition and adds Namespace and Version.
// It embeds mcp.Tool to ensure compatibility with the official SDK.
type Tool struct {
	mcp.Tool
	// Namespace provides a way to namespace tools, e.g. for stable IDs.
	Namespace string `json:"namespace,omitempty"`
	// Version is an optional version string for the tool.
	Version string `json:"version,omitempty"`
}

// ToolIcon is an alias for mcp.Icon from the official SDK.
type ToolIcon = mcp.Icon

// BackendKind defines the type of backend backing a tool.
type BackendKind string

const (
	BackendKindMCP      BackendKind = "mcp"
	BackendKindProvider BackendKind = "provider"
	BackendKindLocal    BackendKind = "local"
)

// ToolBackend defines the binding information for a tool's execution.
// A tool can have multiple backends, but typically one active one.
type ToolBackend struct {
	Kind     BackendKind      `json:"kind"`
	MCP      *MCPBackend      `json:"mcp,omitempty"`
	Provider *ProviderBackend `json:"provider,omitempty"`
	Local    *LocalBackend    `json:"local,omitempty"`
}

// MCPBackend defines metadata for an MCP server backend.
type MCPBackend struct {
	// ServerName identifies the MCP server (e.g. in a registry or config).
	ServerName string `json:"serverName,omitempty"`
}

// ProviderBackend defines metadata for an external/manual tool provider.
type ProviderBackend struct {
	ProviderID string `json:"providerId"`
	ToolID     string `json:"toolId"`
}

// LocalBackend defines metadata for a locally executed tool.
type LocalBackend struct {
	// Name identifies the local function or handler.
	Name string `json:"name"`
}

// ToolID returns the canonical identifier for a tool.
// Format: "namespace:name" when namespace is present, otherwise just "name".
func (t *Tool) ToolID() string {
	if t.Namespace == "" {
		return t.Name
	}
	return t.Namespace + ":" + t.Name
}

// ParseToolID parses a tool ID string into namespace and name components.
// The format is "namespace:name" or just "name" (empty namespace).
// Returns an error if the ID is empty or contains multiple colons.
func ParseToolID(id string) (namespace, name string, err error) {
	if id == "" {
		return "", "", ErrInvalidToolID
	}

	// Count colons - we only allow at most one
	colonCount := strings.Count(id, ":")
	if colonCount > 1 {
		return "", "", ErrInvalidToolID
	}

	if colonCount == 0 {
		// No namespace, just the name
		return "", id, nil
	}

	// Split on the single colon
	parts := strings.SplitN(id, ":", 2)
	namespace = parts[0]
	name = parts[1]

	// Both namespace and name must be non-empty when colon is present
	if namespace == "" || name == "" {
		return "", "", ErrInvalidToolID
	}

	return namespace, name, nil
}

// ToMCPJSON serializes the Tool to JSON that is compatible with the MCP Tool spec.
// This strips toolmodel-specific fields (Namespace, Version) and returns only
// the standard MCP Tool fields.
func (t *Tool) ToMCPJSON() ([]byte, error) {
	return json.Marshal(t.Tool)
}

// ToJSON serializes the full Tool (including toolmodel extensions) to JSON.
func (t *Tool) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}

// FromMCPJSON deserializes an MCP Tool JSON into a Tool struct.
// The Namespace and Version fields will be empty after this call.
func FromMCPJSON(data []byte) (*Tool, error) {
	var mcpTool mcp.Tool
	if err := json.Unmarshal(data, &mcpTool); err != nil {
		return nil, err
	}
	return &Tool{Tool: mcpTool}, nil
}

// FromJSON deserializes a full Tool JSON (including toolmodel extensions) into a Tool struct.
func FromJSON(data []byte) (*Tool, error) {
	var tool Tool
	if err := json.Unmarshal(data, &tool); err != nil {
		return nil, err
	}
	return &tool, nil
}
