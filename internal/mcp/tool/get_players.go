package tool

import (
	"context"
	"encoding/json"

	"github.com/Sup3r-Us3r/go-soccer/internal/usecase/soccer"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

type GetPlayersArguments struct {
	TeamName string `json:"team" jsonschema:"required,description=The name of the team to get the players for"`
}

func GetPlayersTool(server *mcp_golang.Server) error {
	ctx := context.Background()
	err := server.RegisterTool("get_players", "Get the players", func(arguments GetPlayersArguments) (*mcp_golang.ToolResponse, error) {
		useCase := soccer.NewGetPlayersUseCase()

		players, err := useCase.Execute(ctx, soccer.GetPlayersUseCaseInputDTO{TeamName: arguments.TeamName})
		if err != nil {
			return nil, err
		}

		playersJson, err := json.Marshal(players)
		if err != nil {
			return nil, err
		}

		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(playersJson))), nil
	})

	return err
}
