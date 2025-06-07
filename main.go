package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"Demo 🚀",
		"1.0.0",
	)

	// Add tool handler
	s.AddTool(
		// tool
		mcp.NewTool("hello_world",
			mcp.WithDescription("Say hello to someone"),
			mcp.WithString("name",
				mcp.Required(),
				mcp.Description("Name of the person to greet"),
			),
		),

		// handler
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			name, ok := request.Params.Arguments["name"].(string)
			if !ok {
				return nil, errors.New("name must be a string")
			}

			return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
		},
	)

	// 追加してみる
	s.AddTool(
		mcp.NewTool("repeat",
			mcp.WithDescription("Repeat a message"),
			mcp.WithString("message",
				mcp.Required(),
				mcp.Description("Message to repeat"),
			),
			mcp.WithNumber("number",
				mcp.Required(),
				mcp.Description("Number of times to repeat the message"),
			),
		),

		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			message, ok := request.Params.Arguments["message"].(string)
			if !ok {
				return nil, errors.New("message must be a string")
			}

			number, ok := request.Params.Arguments["number"].(float64)
			if !ok {
				return nil, errors.New("number must be a float64")
			}

			return mcp.NewToolResultText(strings.Repeat(message, int(number))), nil
		},
	)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
