# Repository Guidelines

## Project Structure & Module Organization

This repository is a Go MCP server for IEEE Xplore search and article retrieval.

- `main.go` defines the MCP server, HTTP transport, schemas, and response mapping.
- `ieeexplore/client.go` contains the IEEE Xplore HTTP client, headers, parsing, and HTML-to-Markdown conversion.
- `ieeexplore/schema.go` defines structs for IEEE Xplore metadata and search responses.
- `Dockerfile` and `Makefile` provide the container path.
- `tmp/` is scratch output, not source.

Keep domain logic inside `ieeexplore/` unless it is MCP wiring or startup behavior.

## Build, Test, and Development Commands

- `go run .` starts the MCP HTTP server on `:8080`, or `:$PORT` when set.
- `go test ./...` runs all Go tests.
- `go fmt ./...` formats Go files using the standard formatter.
- `go vet ./...` runs static checks.
- `make build` builds the Docker image tagged `ieeexplore-mcp`.
- `make run` builds and starts the Docker container on host port `8080`.

## Coding Style & Naming Conventions

Use standard Go formatting: tabs for indentation, `gofmt` before committing, and short lowercase package names. Export only types and functions needed across packages, such as `Client`, `NewClient`, and response schema types. Keep JSON field names stable because they define the MCP tool contract, for example `articlesPerPage`, `current_page`, and `content`.

Prefer contextual error wrapping, e.g. `fmt.Errorf("decode response: %w", err)`. Use `slog` for server logging.

## Testing Guidelines

There are currently no checked-in test files. New tests should use Go's standard `testing` package and live next to the code under test, with names like `TestClientSearchRejectsEmptyQuery`. For HTTP behavior, inject a custom `http.Client` or transport rather than calling IEEE Xplore in unit tests. Keep network-dependent checks as manual or integration tests.

Run `go test ./...` before opening a PR.

## Commit & Pull Request Guidelines

Recent history uses concise Conventional Commit prefixes such as `feat:`, `fix:`, and `chore:`. Follow that pattern, for example `feat: add article pagination` or `fix: preserve DOI in article response`.

Pull requests should include a short description, user-visible MCP tool changes, and validation performed. Link related issues when available. Include sample inputs/outputs when changing `search` or `get_article` responses because clients may rely on those schemas.

## Security & Configuration Tips

Do not commit credentials, cookies, or private IEEE access tokens. The server relies on public HTTP requests with browser-like headers. Keep request headers centralized in `setSearchHeaders`. Use `PORT` for deployment-specific binding instead of hardcoding alternate ports.
