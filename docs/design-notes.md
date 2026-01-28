# Design Notes

This page captures the tradeoffs and error semantics that guided `toolmodel`.

## Design tradeoffs

- **Spec alignment over custom types.** `Tool` embeds the official MCP Go SDK `mcp.Tool` to stay 1:1 with the spec and JSON tags. This minimizes drift but means `InputSchema`/`OutputSchema` are `any`, so validation must be handled explicitly.
- **Minimal extensions.** `Namespace`, `Version`, and `Tags` are the only additions to the MCP shape. These are intentionally kept small to preserve transport compatibility and keep higher layers in control of semantics.
- **Explicit tool IDs.** Canonical IDs are `namespace:name` (or just `name`), computed by `ToolID()`. This keeps IDs stable across backends while remaining human-readable.
- **Validation boundary.** `Tool.Validate()` enforces naming and required fields only. JSON Schema validation is delegated to `SchemaValidator` to keep `Tool` lightweight and reusable.
- **Safe schema validation.** The default validator blocks external `$ref` resolution to avoid network access and non-determinism. This trades off remote schema reuse for safety and predictability.

## Error semantics

`toolmodel` uses sentinel errors so callers can reliably classify failures:

- `ErrInvalidToolID` – malformed tool IDs (empty, extra `:` separators, missing parts).
- `ErrInvalidTool` – invalid tool definition (missing name, invalid characters, missing input schema).
- `ErrInvalidSchema` – schema is not valid JSON Schema or cannot be parsed.
- `ErrUnsupportedSchema` – schema dialect is not supported (only 2020-12 and draft-07 are accepted).
- `ErrExternalRef` – external `$ref` resolution attempted (blocked by default).

### Validation behavior

- `Tool.Validate()` enforces name format and `InputSchema != nil`.
- `DefaultValidator.ValidateInput` and `ValidateOutput` return `ErrInvalidSchema` or `ErrUnsupportedSchema` when schema parsing or dialect checks fail.
- Output validation is optional by design; `OutputSchema` can be absent.

## Extension points

- **Custom schema validation:** implement `SchemaValidator` if you need different dialects, format checking, or external reference resolution.
- **Tag strategies:** `NormalizeTags` can be replaced at higher layers (e.g., for hierarchical tags or full-text indexing).
- **Tool ingestion:** higher layers can deserialize MCP tool JSON via `FromMCPJSON` and then enrich with `Namespace` and `Tags`.

## Operational guidance

- Keep `Namespace` stable even if backend endpoints change.
- Use short, searchable `Tags`; `NormalizeTags` caps to 20 items and 64 chars each.
- Prefer JSON Schema 2020-12 with explicit `type` and `required` fields for best downstream schema derivation.
