package soccer

import (
	"encoding/json"
	"net/http"

	"github.com/Sup3r-Us3r/go-soccer/internal/apperr"
	"github.com/Sup3r-Us3r/go-soccer/internal/usecase/soccer"
)

type Trophy struct {
	Year         string              `json:"year"`
	Championship soccer.Championship `json:"championship"`
}

type GetTrophiesRequest struct {
	TeamName string `json:"teamName"`
}

type GetTrophiesResponse struct {
	Trophies []Trophy `json:"trophies"`
}

type GetTrophiesHandler struct {
	GetTrophiesUseCase soccer.GetTrophiesUseCaseInterface
}

func NewGetTrophiesHandler(getTrophiesUseCase soccer.GetTrophiesUseCaseInterface) *GetTrophiesHandler {
	return &GetTrophiesHandler{
		GetTrophiesUseCase: getTrophiesUseCase,
	}
}

func (gth GetTrophiesHandler) Handle(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get("teamName")
	if teamName == "" {
		apperr.NewHttpError(w, apperr.ErrTeamNameRequired)
		return
	}

	input := soccer.GetTrophiesUseCaseInputDTO{TeamName: teamName}
	trophies, err := gth.GetTrophiesUseCase.Execute(r.Context(), input)
	if err != nil {
		if appErr, ok := err.(*apperr.AppErr); ok {
			apperr.NewHttpError(w, appErr)
		} else {
			apperr.NewHttpError(w, apperr.NewInternalServerError("Internal server error"))
		}
		return
	}

	response := GetTrophiesResponse{Trophies: make([]Trophy, len(trophies))}

	for i, t := range trophies {
		response.Trophies[i] = Trophy{
			Year:         t.Year,
			Championship: t.Championship,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
