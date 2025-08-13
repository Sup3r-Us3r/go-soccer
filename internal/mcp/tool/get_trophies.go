package tool

import (
	"context"
	"encoding/json"

	"github.com/Sup3r-Us3r/go-soccer/internal/usecase/soccer"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

type GetTrophiesArguments struct {
	TeamName string `json:"team" jsonschema:"required,description=The name of the team to get the trophies for"`
}

func GetTrophiesTool(server *mcp_golang.Server) error {
	ctx := context.Background()
	err := server.RegisterTool("get_trophies", "Get the trophies", func(arguments GetTrophiesArguments) (*mcp_golang.ToolResponse, error) {
		useCase := soccer.NewGetTrophiesUseCase()

		trophies, err := useCase.Execute(ctx, soccer.GetTrophiesUseCaseInputDTO{TeamName: arguments.TeamName})
		if err != nil {
			return nil, err
		}

		trophiesJson, err := json.Marshal(trophies)
		if err != nil {
			return nil, err
		}

		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(trophiesJson))), nil
	})

	return err
}
