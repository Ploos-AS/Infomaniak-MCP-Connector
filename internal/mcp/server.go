package mcpserver

import (
 "context"
 "encoding/json"
 "os"

 "github.com/modelcontextprotocol/go-sdk/mcp"
 "github.com/Ploos-AS/Infomaniak-MCP-Connector/internal/infomaniak"
)

type StatusOutput struct {
 Status string `json:"status"`
 Transport string `json:"transport"`
 InfomaniakTokenConfigured bool `json:"infomaniak_token_configured"`
}

type ProfileOutput struct {
 Profile json.RawMessage `json:"profile"`
}

func NewServer(client *infomaniak.Client) *mcp.Server {
 server := mcp.NewServer(&mcp.Implementation{Name: "infomaniak-mcp-connector", Version: "0.1.0"}, nil)
 mcp.AddTool(server, &mcp.Tool{Name: "connector_status", Description: "Check connector readiness without accessing Infomaniak"}, func(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, StatusOutput, error) {
  return nil, StatusOutput{Status: "ready", Transport: "stdio", InfomaniakTokenConfigured: client.Token != ""}, nil
 })
 mcp.AddTool(server, &mcp.Tool{Name: "infomaniak_profile", Description: "Read the authenticated Infomaniak account profile (requires INFOMANIAK_ACCESS_TOKEN)"}, func(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, ProfileOutput, error) {
  profile, err := client.Profile(ctx)
  if err != nil { return nil, ProfileOutput{}, err }
  return nil, ProfileOutput{Profile: profile}, nil
 })
 return server
}

func Run(ctx context.Context) error {
 client := infomaniak.NewClient(os.Getenv("INFOMANIAK_ACCESS_TOKEN"))
 return NewServer(client).Run(ctx, &mcp.StdioTransport{})
}
