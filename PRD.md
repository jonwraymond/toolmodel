# PRD: toolmodel Go Library

## Summary
The toolmodel library provides a single, canonical data model for tools in the system, based on the MCP Tool specification (2025-11-25). It is the only place that defines what a "tool" is (fields, constraints, IDs), and all downstream modules depend on it for tool definitions. Default validation behavior must match MCP's JSON Schema rules.

## Goals
- Represent MCP Tool objects accurately in Go.
- Provide a stable Tool ID and namespacing convention across backends.
- Store backend binding information (MCP, provider, local) in a small, typed structure.
- Offer JSON Schema validation helpers for inputs and outputs.
- Be safe to embed in public APIs (no framework dependencies).
- Use the official MCP Go SDK types/constants when available to avoid divergence, while keeping toolmodel as the canonical entry point.

## Non-goals
- No networking or JSON-RPC.
- No searching, registry, or execution logic.
- No MCP message types (requests/responses); those live elsewhere.

## Users and Use Cases
Primary users:
- toolindex, tooldocs, toolrun, toolcode, MCP servers, provider bridges.

Core use cases:
- A registry loads tools from multiple sources, normalizes them to toolmodel.Tool, and indexes them for search.
- An execution engine receives tool_id and args, looks up Tool + backend metadata in toolmodel, validates args, and executes.
- A docs module fetches Tool and attaches examples, then exposes it via MCP tools.

## Functional Requirements
- Represent MCP Tool with required fields:
  - name, description, inputSchema
- Include optional MCP fields:
  - title, icons, outputSchema, annotations
- Provide an extended Tool type supporting:
  - namespace (for stable IDs)
  - version (optional)
- Provide backend binding model:
  - BackendKind values: mcp, provider, local (allow future extensions)
  - MCP, Provider, and Local binding structs
  - Provider backend stores ProviderID and ToolID for external/manual tool providers
- Provide JSON helpers:
  - Serialize/deserialize Tool to JSON compatible with MCP Tool spec
- Provide validation helpers:
  - Validate input arguments against inputSchema
  - Validate result against outputSchema when present
- Enforce MCP schema rules:
  - inputSchema MUST be a valid JSON Schema object (not null)
  - If no $schema is present, treat as JSON Schema 2020-12
  - If $schema is present and unsupported by the default validator, return a clear error
- Support "no parameters" schemas per MCP guidance:
  - Recommended: { "type": "object", "additionalProperties": false }
  - Allowed: { "type": "object" }
- Provide safe defaults:
  - If inputSchema is {"type":"object"}, treat it as "any object" per MCP rules

## Non-functional Requirements
- Zero external network calls.
- Standard library plus one JSON Schema validation lib (default implementation), while still supporting swapping via interface.
- Default validator must use github.com/google/jsonschema-go/jsonschema (supports JSON Schema 2020-12 and draft-07 only).
- External $ref resolution must be disabled by default to prevent network access.
- API stability: future schema changes should be additive.
- Well-documented, small surface area.
- When the official MCP Go SDK provides compatible Tool definitions or constants, toolmodel MUST import and use them. Only mirror the spec when SDK coverage is missing or incompatible, and document the gap.

## Design Notes (Normalized)
- Tool mirrors MCP schema and adds Namespace and Version.
- ToolIcon contains Uri and optional MediaType.
- ToolBackend separates binding from Tool so a single Tool can have multiple backends.
- Provider backend represents external/manual tool providers (non-MCP) and carries ProviderID and ToolID.
- Schema validation uses a small interface (SchemaValidator) plus a built-in default validator implementation.
- Default validator uses jsonschema-go and supports only 2020-12 and draft-07; other dialects return an unsupported error.
- jsonschema-go ignores "format" and content-related keywords during validation; document this limitation.
- $ref resolution uses a loader; the default loader denies external references to enforce no-network.
- Tool identity uses canonical ID format: namespace:name (namespace optional).
- Optional helpers strip non-MCP fields for MCP-facing JSON.
- If the official MCP Go SDK exposes Tool types, prefer type aliasing; if aliasing is not viable, use a thin wrapper and document why.

## Gates
Gate: go test ./...

## Tasks
- [x] Initialize Go module and package skeleton for toolmodel (doc.go with package comment).
- [x] Define Tool and ToolIcon types matching MCP Tool fields plus Namespace and Version.
- [x] Define BackendKind, ToolBackend, MCPBackend, ProviderBackend, and LocalBackend.
- [x] Evaluate the official MCP Go SDK for Tool definitions/constants; prefer type aliasing, else thin wrapper, else mirror; document the decision and gaps.
- [x] Implement ToolID and ParseToolID helpers with unit tests.
- [ ] Implement JSON helpers (ToMCPJSON and FromMCPJSON) with MCP compatibility tests.
- [ ] Implement SchemaValidator interface and default validator using jsonschema-go, enforcing MCP schema rules (2020-12 default, explicit draft-07, unsupported dialect errors, no external refs) with tests.
- [ ] Add examples (examples_test.go) that demonstrate tool creation and validation.

## Success Criteria
- [ ] All Gates pass.
- [ ] All Functional Requirements are implemented.
- [ ] MCP JSON compatibility is verified by tests.
- [ ] Default validator uses jsonschema-go while preserving the interface to allow swapping.
- [ ] Unsupported $schema dialects return a deterministic error.

## Open Questions
- Do we want to allow opt-in external $ref loaders, or keep external refs permanently disallowed in toolmodel?
