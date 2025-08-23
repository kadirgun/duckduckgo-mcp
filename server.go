package main

import (
	"context"

	duckduckgo "github.com/kadirgun/duckduck-go"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SearchServer struct {
	duck *duckduckgo.DuckDuckGo
}

func NewSearchServer() *SearchServer {
	return &SearchServer{
		duck: duckduckgo.New(),
	}
}

type SearchParams struct {
	Query string `json:"query" jsonschema:"The search query"`
	Count int    `json:"count" jsonschema:"Number of results to return"`
}

func (s *SearchServer) Search(ctx context.Context, req *mcp.CallToolRequest, args SearchParams) (*mcp.CallToolResult, any, error) {
	if args.Count <= 0 {
		args.Count = 10
	}

	results, err := s.duck.Search(args.Query, args.Count)

	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		StructuredContent: results,
	}, nil, nil
}
