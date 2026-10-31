# OAuth implementation plan (Phase A → B)

Status: design only; no OAuth implementation yet.

## Phase A: local single-user

1. Register an application in the Infomaniak developer console.
2. Verify the authorization URL, token URL, PKCE support, redirect URI rules, and required scopes against the current official documentation.
3. Use Authorization Code flow with PKCE when supported. Bind callback to loopback, validate state, and validate redirect.
4. Store tokens locally with restrictive file permissions or an OS credential store; never log credentials.
5. Implement refresh and revocation handling according to documented provider behavior.
6. Test authentication independently of MCP and begin with a least-privilege read-only capability.

## Phase B: hosted multi-user

- Remote MCP over HTTPS with appropriate MCP authorization.
- Distinguish MCP client authentication from Infomaniak OAuth authorization.
- Store per-user credentials encrypted at rest with key rotation.
- Protect against cross-tenant access; isolate token access by account.
- Implement CSRF protections, token rotation, scopes, and audit logs.
- Never expose provider client secrets to MCP clients.

## Verification checklist

- [ ] Confirm official OAuth endpoints and discovery metadata
- [ ] Confirm PKCE and refresh-token semantics
- [ ] Confirm supported scopes for kDrive file listing
- [ ] Confirm application registration and redirect settings
- [ ] Obtain successful CI build before enabling provider calls

No client ID, client secret, or access token belongs in this repository.
