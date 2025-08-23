package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "greeter",
			Version: "v1.0.0",
		},
		nil,
	)

	searchServer := NewSearchServer()

	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "search",
			Description: "Search the web using DuckDuckGo",
		},
		searchServer.Search,
	)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
