package main

import (
 "context"
 "log"
 mcpserver "github.com/Ploos-AS/Infomaniak-MCP-Connector/internal/mcp"
)

func main() {
 if err := mcpserver.Run(context.Background()); err != nil {
  log.Fatal(err)
 }
}
