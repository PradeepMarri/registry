package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Updatedat string `json:"updatedAt,omitempty"` // Last update timestamp
	Createdat string `json:"createdAt,omitempty"` // Creation timestamp
	Customerid string `json:"customerId,omitempty"` // Customer ID associated with the artifact
	Id string `json:"id,omitempty"` // Artifact unique identifier
	Name string `json:"name,omitempty"` // Artifact name
	Size int `json:"size,omitempty"` // Artifact size in bytes
	TypeField string `json:"type,omitempty"` // Artifact type
}

// APIResponse represents the APIResponse schema from the OpenAPI specification
type APIResponse struct {
	Message string `json:"message,omitempty"` // Response message
	Success bool `json:"success,omitempty"` // Operation success status
	Data map[string]interface{} `json:"data,omitempty"` // Response data
}

// Customer represents the Customer schema from the OpenAPI specification
type Customer struct {
	Id string `json:"id,omitempty"` // Customer unique identifier
	Name string `json:"name,omitempty"` // Customer name
	Createdat string `json:"createdAt,omitempty"` // Creation timestamp
	Email string `json:"email,omitempty"` // Customer email address
}

// Success represents the Success schema from the OpenAPI specification
type Success struct {
	Success bool `json:"success,omitempty"` // Operation success status
	Data map[string]interface{} `json:"data,omitempty"` // Response data
	Message string `json:"message,omitempty"` // Success message
}

// BrokerInfo represents the BrokerInfo schema from the OpenAPI specification
type BrokerInfo struct {
	Endpoint string `json:"endpoint,omitempty"` // Broker endpoint URL
	Id string `json:"id,omitempty"` // Broker unique identifier
	Lasthealthcheck string `json:"lastHealthCheck,omitempty"` // Last health check timestamp
	Name string `json:"name,omitempty"` // Broker name
	Status string `json:"status,omitempty"` // Broker status
	TypeField string `json:"type,omitempty"` // Broker type
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	ErrorField string `json:"error,omitempty"` // Technical error details
	Message string `json:"message,omitempty"` // Error message
	Timestamp string `json:"timestamp,omitempty"` // Error timestamp
}

// Broker represents the Broker schema from the OpenAPI specification
type Broker struct {
	Endpoint string `json:"endpoint,omitempty"` // Broker endpoint URL
	Id string `json:"id,omitempty"` // Broker unique identifier
	Name string `json:"name,omitempty"` // Broker name
	Status string `json:"status,omitempty"` // Broker status
}

// CustomerResponse represents the CustomerResponse schema from the OpenAPI specification
type CustomerResponse struct {
	Success bool `json:"success,omitempty"` // Operation success status
	Customer Customer `json:"customer,omitempty"`
	Message string `json:"message,omitempty"` // Response message
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Code int `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
