# Architecture

## Initial mode (Phase A)

Go executable with a separate Infomaniak HTTP client and future local MCP stdio transport. API capabilities will be added incrementally after confirming endpoints and scopes.

## Hosted mode (Phase B)

Remote MCP over HTTPS, per-user OAuth, durable encrypted token storage, tenant isolation, rate limiting, audit controls, and explicit authorization for write actions. This mode is not yet implemented.

## Design rules

- Separate `internal/infomaniak` from `internal/mcp`.
- Never expose broad account-management permissions just to read user data.
- Do not claim mail, calendar, or document editing support until API feasibility is verified.
- Prefer read-only operations before any writes.
