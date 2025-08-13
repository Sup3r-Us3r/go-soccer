package main

import (
	"fmt"

	"github.com/Sup3r-Us3r/go-soccer/internal/mcp/tool"
	"github.com/invopop/jsonschema"
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

func main() {
	jsonschema.Version = "https://json-schema.org/draft-07/schema#"

	done := make(chan struct{})

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())
	if err := tool.GetLatestMatchesTool(server); err != nil {
		panic(err)
	}
	if err := tool.GetNextMatchesTool(server); err != nil {
		panic(err)
	}
	if err := tool.GetPlayersTool(server); err != nil {
		panic(err)
	}
	if err := tool.GetTransfersTool(server); err != nil {
		panic(err)
	}
	if err := tool.GetTrophiesTool(server); err != nil {
		panic(err)
	}

	if err := server.Serve(); err != nil {
		panic(err)
	}

	fmt.Println("MCP Server started")

	<-done
}
