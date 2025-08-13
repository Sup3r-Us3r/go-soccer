package tool

import (
	"context"
	"encoding/json"

	"github.com/Sup3r-Us3r/go-soccer/internal/usecase/soccer"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

type GetLatestMatchesArguments struct {
	TeamName string `json:"team" jsonschema:"required,description=The name of the team to get the latest matches for"`
}

func GetLatestMatchesTool(server *mcp_golang.Server) error {
	ctx := context.Background()
	err := server.RegisterTool("get_latest_matches", "Get the latest matches", func(arguments GetLatestMatchesArguments) (*mcp_golang.ToolResponse, error) {
		useCase := soccer.NewGetLatestMatchesUseCase()

		matches, err := useCase.Execute(ctx, soccer.GetLatestMatchesUseCaseInputDTO{TeamName: arguments.TeamName})
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
