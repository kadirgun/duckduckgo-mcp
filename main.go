package main

import (
	"context"
	"fmt"
	"log"
	"os"

	duckduckgo "github.com/kadirgun/duckduck-go"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "duckduckgo-mcp",
		Usage: "DuckDuckGo MCP server",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "retry-count",
				Usage: "Number of retries for failed search requests (default: 0)",
				Value: 0,
				Validator: func(v int) error {
					if v < 0 {
						return fmt.Errorf("retry-count must be non-negative")
					}
					return nil
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			retryCount := cmd.Int("retry-count")

			var opts []duckduckgo.Option
			if retryCount > 0 {
				opts = append(opts, duckduckgo.WithRetryCount(retryCount))
			}

			searchServer := NewSearchServer(opts...)

			server := mcp.NewServer(
				&mcp.Implementation{
					Name:    "DuckDuckGo",
					Version: "v1.1.1",
				},
				nil,
			)

			mcp.AddTool(
				server,
				&mcp.Tool{
					Name:        "search",
					Description: "Search the web using DuckDuckGo",
				},
				searchServer.Search,
			)

			return server.Run(context.Background(), &mcp.StdioTransport{})
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
