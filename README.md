# cloud2butane

A small Go CLI tool that converts a limited cloud-init `write_files` configuration into a valid Flatcar Butane YAML configuration.

## Supported input

```yaml
write_files:
  - path: /etc/example.conf
    content: hello world
```

## Generated output

```yaml
variant: flatcar
version: 1.0.0
storage:
  files:
    - path: /etc/example.conf
      contents:
        inline: hello world
```

## Run

```bash
go run . input.yaml
```

## Test

```bash
go test ./...
```

## Scope

This project intentionally supports only:

- `write_files`
- `path`
- `content`

Unsupported fields such as `permissions`, `owner`, `append`, and `encoding` are ignored.

