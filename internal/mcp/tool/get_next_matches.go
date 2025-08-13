package tool

import (
	"context"
	"encoding/json"

	"github.com/Sup3r-Us3r/go-soccer/internal/usecase/soccer"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

type GetNextMatchesArguments struct {
	TeamName string `json:"team" jsonschema:"required,description=The name of the team to get the next matches for"`
}

func GetNextMatchesTool(server *mcp_golang.Server) error {
	ctx := context.Background()
	err := server.RegisterTool("get_next_matches", "Get the next matches", func(arguments GetNextMatchesArguments) (*mcp_golang.ToolResponse, error) {
		useCase := soccer.NewGetNextMatchesUseCase()

		matches, err := useCase.Execute(ctx, soccer.GetNextMatchesUseCaseInputDTO{TeamName: arguments.TeamName})
		if err != nil {
			return nil, err
		}

		matchesJson, err := json.Marshal(matches)
		if err != nil {
			return nil, err
		}

		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(matchesJson))), nil
	})

	return err
}
