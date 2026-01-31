# Migration Guide: toolmodel → toolfoundation/model

## Import Path Changes

| Before | After |
|--------|-------|
| `github.com/jonwraymond/toolmodel` | `github.com/jonwraymond/toolfoundation/model` |

## Breaking Changes

- Package renamed from `toolmodel` to `model`
- Update all type references: `toolmodel.Tool` → `model.Tool`

## Migration Steps

1. Update go.mod:
   ```bash
   go get github.com/jonwraymond/toolfoundation
   ```

2. Update imports (sed example):
   ```bash
   find . -name "*.go" -exec sed -i '' 's|github.com/jonwraymond/toolmodel|github.com/jonwraymond/toolfoundation/model|g' {} +
   ```

3. Update type references:
   ```bash
   find . -name "*.go" -exec sed -i '' 's|toolmodel\.|model.|g' {} +
   ```

4. Clean up:
   ```bash
   go mod tidy
   ```

## API Compatibility

The API is fully compatible. Only the import path and package name have changed.
