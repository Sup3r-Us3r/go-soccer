package soccer

import (
	"encoding/json"
	"net/http"

	"github.com/Sup3r-Us3r/go-soccer/internal/apperr"
	"github.com/Sup3r-Us3r/go-soccer/internal/usecase/soccer"
)

type Transfer struct {
	Date   string      `json:"date"`
	Type   string      `json:"type"`
	Player string      `json:"player"`
	Team   soccer.Team `json:"team"`
}

type GetTransfersRequest struct {
	TeamName string `json:"teamName"`
}

type GetTransfersResponse struct {
	Transfers []Transfer `json:"transfers"`
}

type GetTransfersHandler struct {
	GetTransfersUseCase soccer.GetTransfersUseCaseInterface
}

func NewGetTransfersHandler(getTransfersUseCase soccer.GetTransfersUseCaseInterface) *GetTransfersHandler {
	return &GetTransfersHandler{
		GetTransfersUseCase: getTransfersUseCase,
	}
}

func (gth GetTransfersHandler) Handle(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get("teamName")
	if teamName == "" {
		apperr.NewHttpError(w, apperr.ErrTeamNameRequired)
		return
	}

	input := soccer.GetTransfersUseCaseInputDTO{TeamName: teamName}
	transfers, err := gth.GetTransfersUseCase.Execute(r.Context(), input)
	if err != nil {
		if appErr, ok := err.(*apperr.AppErr); ok {
			apperr.NewHttpError(w, appErr)
		} else {
			apperr.NewHttpError(w, apperr.NewInternalServerError("Internal server error"))
		}
		return
	}

	response := GetTransfersResponse{Transfers: make([]Transfer, len(transfers))}

	for i, t := range transfers {
		response.Transfers[i] = Transfer{
			Date:   t.Date,
			Type:   t.Type,
			Player: t.Player,
			Team:   t.Team,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
