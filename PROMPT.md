# Ralph Loop: toolmodel Go Library Implementation

## Task
Implement the toolmodel Go library according to PRD.md specifications. Work through tasks sequentially, completing each one fully before moving to the next.

## Context
- This is a Go library for MCP Tool definitions
- The official MCP Go SDK is already imported (github.com/mark3labs/mcp-go)
- Tool and ToolID/ParseToolID are partially implemented
- Remaining work: backend types, JSON helpers, schema validation

## Requirements
For each task in PRD.md:
- [ ] Implement the functionality as specified
- [ ] Write comprehensive unit tests (>80% coverage for new code)
- [ ] Ensure `go test ./...` passes
- [ ] Ensure `go vet ./...` passes
- [ ] Update PRD.md to mark task as complete `[x]`

## Work Process

### Each Iteration:
1. Read PRD.md to find the highest-priority unchecked task
2. Implement that single task completely
3. Write tests for the implementation
4. Run gates: `go test ./... && go vet ./...`
5. If gates pass, mark task complete in PRD.md
6. Commit changes with descriptive message
7. Continue to next task

### Gate Commands
```bash
go test ./...
go vet ./...
```

## Self-Correction

### If tests fail:
1. Read the error message carefully
2. Identify the root cause (not symptoms)
3. Fix the implementation or test
4. Re-run tests before proceeding

### If stuck on same issue for 3+ iterations:
1. Document what's blocking in a code comment
2. Try an alternative approach
3. If still stuck after 5 iterations:
   - Document the blocker clearly
   - Output: <promise>BLOCKED</promise>

### If unsure about MCP SDK usage:
1. Check the mcp-go SDK source for existing types
2. Prefer embedding/aliasing SDK types over reimplementing
3. Document any gaps between SDK and spec in code comments

## Implementation Guidelines

### Type Design
- Embed mcp.Tool in toolmodel.Tool (already done)
- Use typed constants for BackendKind
- Keep structs small and focused

### Testing
- Table-driven tests preferred
- Test edge cases: empty strings, nil values, invalid input
- Test JSON round-tripping for serialization

### Documentation
- Package doc in doc.go (already exists)
- Doc comments on all exported types and functions
- Examples in examples_test.go for key usage patterns

## Success Criteria
ALL must be true:
- All PRD.md tasks marked `[x]`
- `go test ./...` passes
- `go vet ./...` passes
- All Success Criteria in PRD.md satisfied

## Completion

When ALL tasks in PRD.md are complete AND all gates pass:
Output: <promise>ALL_PRD_TASKS_COMPLETE</promise>

When a single task is complete and committed (for progress tracking):
Continue to next task (do not output promise yet)
