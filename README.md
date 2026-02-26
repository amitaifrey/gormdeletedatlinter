# gormdeletedat

A Go linter that reports GORM model structs whose `DeletedAt` field is not typed as `*gorm.DeletedAt`.

## Why

GORM's soft-delete functionality requires `DeletedAt` to be `*gorm.DeletedAt` (from `gorm.io/gorm`). Using a different type (e.g. `time.Time`, `sql.NullTime`, or a non-pointer `gorm.DeletedAt`) silently breaks soft-delete filtering.

## Example

```go
// Bad - will be flagged
type User struct {
    ID        uint
    DeletedAt time.Time
}

// Good
type User struct {
    ID        uint
    DeletedAt *gorm.DeletedAt
}
```

## Usage

### golangci-lint plugin

Add to `.custom-gcl.yml`:

```yaml
plugins:
  - module: "github.com/amitaifrey/gormdeletedat"
    import: "github.com/amitaifrey/gormdeletedat/plugin"
    version: <version>
```

Add to `.golangci.yml`:

```yaml
linters:
  enable:
    - gormdeletedat
  settings:
    custom:
      gormdeletedat:
        type: module
        description: "checks that GORM DeletedAt fields use *gorm.DeletedAt"
```

### Standalone

```sh
go install github.com/amitaifrey/gormdeletedat/cmd/gormdeletedat@latest
gormdeletedat ./...
```
