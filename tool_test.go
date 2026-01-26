package toolmodel

import (
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestTool_ToolID(t *testing.T) {
	tests := []struct {
		name      string
		tool      Tool
		wantID    string
	}{
		{
			name: "with namespace",
			tool: Tool{
				Tool:      mcp.Tool{Name: "read"},
				Namespace: "filesystem",
			},
			wantID: "filesystem:read",
		},
		{
			name: "without namespace",
			tool: Tool{
				Tool: mcp.Tool{Name: "read"},
			},
			wantID: "read",
		},
		{
			name: "empty namespace explicitly",
			tool: Tool{
				Tool:      mcp.Tool{Name: "write"},
				Namespace: "",
			},
			wantID: "write",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tool.ToolID()
			if got != tt.wantID {
				t.Errorf("ToolID() = %q, want %q", got, tt.wantID)
			}
		})
	}
}

func TestParseToolID(t *testing.T) {
	tests := []struct {
		name          string
		id            string
		wantNamespace string
		wantName      string
		wantErr       bool
	}{
		{
			name:          "with namespace",
			id:            "filesystem:read",
			wantNamespace: "filesystem",
			wantName:      "read",
			wantErr:       false,
		},
		{
			name:          "without namespace",
			id:            "read",
			wantNamespace: "",
			wantName:      "read",
			wantErr:       false,
		},
		{
			name:    "empty string",
			id:      "",
			wantErr: true,
		},
		{
			name:    "multiple colons",
			id:      "a:b:c",
			wantErr: true,
		},
		{
			name:    "leading colon",
			id:      ":name",
			wantErr: true,
		},
		{
			name:    "trailing colon",
			id:      "namespace:",
			wantErr: true,
		},
		{
			name:    "just a colon",
			id:      ":",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNamespace, gotName, err := ParseToolID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToolID(%q) error = %v, wantErr %v", tt.id, err, tt.wantErr)
				return
			}
			if err != nil {
				if err != ErrInvalidToolID {
					t.Errorf("ParseToolID(%q) error = %v, want ErrInvalidToolID", tt.id, err)
				}
				return
			}
			if gotNamespace != tt.wantNamespace {
				t.Errorf("ParseToolID(%q) namespace = %q, want %q", tt.id, gotNamespace, tt.wantNamespace)
			}
			if gotName != tt.wantName {
				t.Errorf("ParseToolID(%q) name = %q, want %q", tt.id, gotName, tt.wantName)
			}
		})
	}
}

func TestParseToolID_RoundTrip(t *testing.T) {
	// Test that ToolID() output can be parsed back correctly
	tests := []struct {
		namespace string
		name      string
	}{
		{"filesystem", "read"},
		{"", "read"},
		{"my-namespace", "my-tool"},
	}

	for _, tt := range tests {
		tool := Tool{
			Tool:      mcp.Tool{Name: tt.name},
			Namespace: tt.namespace,
		}
		id := tool.ToolID()
		gotNamespace, gotName, err := ParseToolID(id)
		if err != nil {
			t.Errorf("ParseToolID(ToolID()) failed for namespace=%q, name=%q: %v", tt.namespace, tt.name, err)
			continue
		}
		if gotNamespace != tt.namespace || gotName != tt.name {
			t.Errorf("Round-trip failed: got (%q, %q), want (%q, %q)", gotNamespace, gotName, tt.namespace, tt.name)
		}
	}
}

func TestBackendKind_Constants(t *testing.T) {
	// Verify the constants have expected string values
	if BackendKindMCP != "mcp" {
		t.Errorf("BackendKindMCP = %q, want %q", BackendKindMCP, "mcp")
	}
	if BackendKindProvider != "provider" {
		t.Errorf("BackendKindProvider = %q, want %q", BackendKindProvider, "provider")
	}
	if BackendKindLocal != "local" {
		t.Errorf("BackendKindLocal = %q, want %q", BackendKindLocal, "local")
	}
}

func TestToolBackend_Structures(t *testing.T) {
	// Test that backend structures can be instantiated correctly
	t.Run("MCP backend", func(t *testing.T) {
		backend := ToolBackend{
			Kind: BackendKindMCP,
			MCP: &MCPBackend{
				ServerName: "my-server",
			},
		}
		if backend.Kind != BackendKindMCP {
			t.Errorf("Kind = %q, want %q", backend.Kind, BackendKindMCP)
		}
		if backend.MCP.ServerName != "my-server" {
			t.Errorf("MCP.ServerName = %q, want %q", backend.MCP.ServerName, "my-server")
		}
	})

	t.Run("Provider backend", func(t *testing.T) {
		backend := ToolBackend{
			Kind: BackendKindProvider,
			Provider: &ProviderBackend{
				ProviderID: "openai",
				ToolID:     "gpt-4-tool",
			},
		}
		if backend.Kind != BackendKindProvider {
			t.Errorf("Kind = %q, want %q", backend.Kind, BackendKindProvider)
		}
		if backend.Provider.ProviderID != "openai" {
			t.Errorf("Provider.ProviderID = %q, want %q", backend.Provider.ProviderID, "openai")
		}
		if backend.Provider.ToolID != "gpt-4-tool" {
			t.Errorf("Provider.ToolID = %q, want %q", backend.Provider.ToolID, "gpt-4-tool")
		}
	})

	t.Run("Local backend", func(t *testing.T) {
		backend := ToolBackend{
			Kind: BackendKindLocal,
			Local: &LocalBackend{
				Name: "my-handler",
			},
		}
		if backend.Kind != BackendKindLocal {
			t.Errorf("Kind = %q, want %q", backend.Kind, BackendKindLocal)
		}
		if backend.Local.Name != "my-handler" {
			t.Errorf("Local.Name = %q, want %q", backend.Local.Name, "my-handler")
		}
	})
}

func TestTool_EmbedsMCPTool(t *testing.T) {
	// Verify Tool correctly embeds mcp.Tool and can access its fields
	tool := Tool{
		Tool: mcp.Tool{
			Name:        "test-tool",
			Description: "A test tool",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"input": map[string]any{"type": "string"},
				},
			},
		},
		Namespace: "test",
		Version:   "1.0.0",
	}

	if tool.Name != "test-tool" {
		t.Errorf("Name = %q, want %q", tool.Name, "test-tool")
	}
	if tool.Description != "A test tool" {
		t.Errorf("Description = %q, want %q", tool.Description, "A test tool")
	}
	if tool.Namespace != "test" {
		t.Errorf("Namespace = %q, want %q", tool.Namespace, "test")
	}
	if tool.Version != "1.0.0" {
		t.Errorf("Version = %q, want %q", tool.Version, "1.0.0")
	}
}

func TestToolIcon_Alias(t *testing.T) {
	// Verify ToolIcon is a proper alias for mcp.Icon
	icon := ToolIcon{
		Source:   "https://example.com/icon.png",
		MIMEType: "image/png",
	}

	// ToolIcon should be usable as mcp.Icon
	var mcpIcon mcp.Icon = icon
	if mcpIcon.Source != "https://example.com/icon.png" {
		t.Errorf("Icon Source = %q, want %q", mcpIcon.Source, "https://example.com/icon.png")
	}
}
