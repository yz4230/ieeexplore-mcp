# IEEE Xplore Search MCP

An MCP server that lets MCP clients search IEEE Xplore and retrieve article metadata and available article-page content as Markdown.

The server uses the MCP Streamable HTTP transport and listens on port `8080` by default. Set `PORT` to change the listening port.

## Tools

### `search`

Search IEEE Xplore for papers using short technical keyword phrases. The tool returns concise metadata and pagination information.

Input:

```json
{
  "query": "large language model retrieval",
  "page": 1,
  "articlesPerPage": 25
}
```

`page` starts at `1`; both `page` and `articlesPerPage` are optional. The default page size is `25`.

Each result includes:

- `id` — IEEE Xplore document number
- `title`
- `authors`
- `publication`
- `year`
- `doi`
- `abstract`

### `get_article`

Retrieve metadata and available article-page content for a specific IEEE Xplore document.

Input:

```json
{
  "id": "1234567"
}
```

Use the `id` returned by `search`, not a DOI. The response includes the same metadata fields plus `content`, containing available page content converted to Markdown.

## Run locally

Requirements: Go `1.26.2` or newer.

```sh
go run .
```

The server starts on `:8080`:

```sh
PORT=9000 go run .
```

## Build and test

```sh
go test ./...
go vet ./...
go fmt ./...
go build -o server .
```

## OAuth

Set both `MCP_PASSWORD` and `MCP_PUBLIC_URL` to protect `/mcp` with OAuth from [mcpauth](https://github.com/yz4230/mcpauth). Connecting from a ChatGPT or Claude connector then shows a password page. With neither set the server runs unauthenticated; setting only one is an error.

| Variable | Default | Description |
|---|---|---|
| `MCP_PASSWORD` | none | Password for the login page |
| `MCP_PUBLIC_URL` | none | Public base URL as clients see it, e.g. `https://mcp.example.com` |
| `MCP_STATE_FILE` | `oauth_state.json` (`/data/oauth_state.json` in Docker) | Where registered clients and tokens are kept |
| `MCP_REDIRECT_URIS` | none | Comma-separated redirect URIs to allow besides ChatGPT's and Claude's |

`docker:run` reads `.env` when present and mounts the `<CONTAINER_NAME>-data` volume at `/data`, so tokens survive recreating the container.

## Docker

Build and run with Docker:

```sh
docker build -t ieeexplore-mcp .
docker run --rm -p 8080:8080 ieeexplore-mcp
```

The repository also provides `mise` tasks:

```sh
mise run docker:build
mise run docker:run
```

## Notes

- Requests are made directly to public IEEE Xplore endpoints using browser-like request headers.
- Article content availability depends on what IEEE Xplore exposes on the article page; this server does not bypass access controls.
- Do not commit credentials, cookies, or private access tokens.
