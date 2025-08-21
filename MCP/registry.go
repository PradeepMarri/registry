package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_headers "github.com/input-api/mcp-server/tools/headers"
	tools_api "github.com/input-api/mcp-server/tools/api"
	tools_v0 "github.com/input-api/mcp-server/tools/v0"
	tools_authorization "github.com/input-api/mcp-server/tools/authorization"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_headers.CreateHead_content_typeTool(cfg),
		tools_api.CreateGetTool(cfg),
		tools_v0.CreateGet_v0_serversTool(cfg),
		tools_v0.CreateGet_v0_servers_server_1Tool(cfg),
		tools_v0.CreateGet_v0_servers_server_2Tool(cfg),
		tools_authorization.CreateHead_authorizationTool(cfg),
	}
}
