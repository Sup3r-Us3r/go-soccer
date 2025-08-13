package tool

import (
	"context"
	"encoding/json"

	"github.com/Sup3r-Us3r/go-soccer/internal/usecase/soccer"
	mcp_golang "github.com/metoro-io/mcp-golang"
)

type GetTransfersArguments struct {
	TeamName string `json:"team" jsonschema:"required,description=The name of the team to get the transfers for"`
}

func GetTransfersTool(server *mcp_golang.Server) error {
	ctx := context.Background()
	err := server.RegisterTool("get_transfers", "Get the transfers", func(arguments GetTransfersArguments) (*mcp_golang.ToolResponse, error) {
		useCase := soccer.NewGetTransfersUseCase()

		transfers, err := useCase.Execute(ctx, soccer.GetTransfersUseCaseInputDTO{TeamName: arguments.TeamName})
		if err != nil {
			return nil, err
		}

		transfersJson, err := json.Marshal(transfers)
		if err != nil {
			return nil, err
		}

		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(string(transfersJson))), nil
	})

	return err
}
