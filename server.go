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

type SearchResult struct {
	Title       string `json:"title" jsonschema:"The title of the search result"`
	URL         string `json:"url" jsonschema:"The URL of the search result"`
	Description string `json:"description" jsonschema:"A brief description of the search result"`
}

func (s *SearchServer) Search(ctx context.Context, req *mcp.CallToolRequest, args SearchParams) (*mcp.CallToolResult, []*SearchResult, error) {
	if args.Count <= 0 {
		args.Count = 10
	}

	results, err := s.duck.Search(args.Query, args.Count)

	if err != nil {
		return nil, nil, err
	}

	response := make([]*SearchResult, len(results))
	for i, r := range results {
		response[i] = &SearchResult{
			Title:       r.Title,
			URL:         r.Link,
			Description: r.Description,
		}
	}

	return nil, response, nil
}
