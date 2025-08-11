package soccer

import (
	"encoding/json"
	"net/http"

	"github.com/Sup3r-Us3r/go-soccer/internal/apperr"
	"github.com/Sup3r-Us3r/go-soccer/internal/usecase/soccer"
)

type Player struct {
	Position string `json:"position"`
	Player   string `json:"player"`
	Age      string `json:"age"`
	Country  string `json:"country"`
}

type GetPlayersRequest struct {
	TeamName string `json:"teamName"`
}

type GetPlayersResponse struct {
	Players []Player `json:"players"`
}

type GetPlayersHandler struct {
	GetPlayersUseCase soccer.GetPlayersUseCaseInterface
}

func NewGetPlayersHandler(getPlayersUseCase soccer.GetPlayersUseCaseInterface) *GetPlayersHandler {
	return &GetPlayersHandler{
		GetPlayersUseCase: getPlayersUseCase,
	}
}

func (gph GetPlayersHandler) Handle(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get("teamName")
	if teamName == "" {
		apperr.NewHttpError(w, apperr.ErrTeamNameRequired)
		return
	}

	players, err := gph.GetPlayersUseCase.Execute(
		r.Context(),
		soccer.GetPlayersUseCaseInputDTO{TeamName: teamName},
	)
	if err != nil {
		if appErr, ok := err.(*apperr.AppErr); ok {
			apperr.NewHttpError(w, appErr)
		} else {
			apperr.NewHttpError(w, apperr.NewInternalServerError("Internal server error"))
		}
		return
	}

	response := GetPlayersResponse{
		Players: make([]Player, len(players)),
	}

	for i, p := range players {
		response.Players[i] = Player{
			Position: p.Position,
			Player:   p.Player,
			Age:      p.Age,
			Country:  p.Country,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
