# Infomaniak MCP Connector

An independent, open-source Model Context Protocol (MCP) connector for Infomaniak services, written in Go.

> **Status:** Early development. No Infomaniak API operations or MCP server transport are implemented yet. Independent project by Ploos AS; not an official Infomaniak integration.

## Roadmap

- **Phase A:** Single-user integration. First verify APIs and authentication, then implement a read-only MCP tool.
- **Phase B:** Optional hosted multi-user remote MCP connector with per-user OAuth authorization.
- Keep the Infomaniak API client independent of MCP transport and business logic.
- Implement only operations verified against current official API documentation.
- Default to least privilege and require confirmation for side-effecting operations.

## Development

Requires Go 1.23+.

```sh
go run ./cmd/infomaniak-mcp
```

The command currently prints a placeholder message; it **does not** start an MCP server. No API credentials are needed at this stage.

See [architecture](docs/architecture.md) and [API coverage](docs/api-coverage.md).

## Security

Never commit credentials, tokens, or `.env` files. The future hosted version will need encrypted token storage, tenant isolation, and an authorization model.

## License

MIT — see [LICENSE](LICENSE).
