# Local development

This is a local stdio MCP server. It is **not yet** a remotely hosted ChatGPT connector.

## Requirements

Go 1.23+ and a compatible MCP client.

```sh
go mod tidy
go test ./...
go run ./cmd/infomaniak-mcp
```

The server speaks MCP JSON-RPC on standard input/output. Do not expect a human-readable greeting.

## Available tools

- `connector_status`: local readiness check; does not require credentials.
- `infomaniak_profile`: reads `GET https://api.infomaniak.com/profile`; requires an Infomaniak bearer token.

For local development only, supply `INFOMANIAK_ACCESS_TOKEN` via your shell environment. Never put tokens in the repository or share them in issues.

The current token-based development path is temporary. OAuth application registration and per-user authorization are separate planned milestones. Do not expose this stdio server as an unauthenticated public HTTP service.

## Important limitation

ChatGPT remote app connections require a separately deployed, appropriately authenticated remote MCP endpoint. The local stdio server alone is not sufficient.
