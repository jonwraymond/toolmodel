# Examples

## Tag normalization

```go
normalized := toolmodel.NormalizeTags([]string{
  "  GitHub Repos ",
  "GITHUB  repos",
  "  ci/cd  ",
})
// => ["github-repos", "ci-cd"]
```

## Validate output

```go
result := map[string]any{
  "full_name": "octo/hello",
}

if err := validator.ValidateOutput(&tool, result); err != nil {
  // handle schema mismatch
}
```

## Canonical ID parsing

```go
ns, name, err := toolmodel.ParseToolID("github:get_repo")
// ns = "github", name = "get_repo"
```
