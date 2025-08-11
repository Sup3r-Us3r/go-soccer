package soccer

import (
	"encoding/json"
	"net/http"

	"github.com/Sup3r-Us3r/go-soccer/internal/apperr"
	"github.com/Sup3r-Us3r/go-soccer/internal/usecase/soccer"
)

type NextMatch struct {
	Title    string `json:"title"`
	Home     string `json:"home"`
	HomeLogo string `json:"homeLogo"`
	Away     string `json:"away"`
	AwayLogo string `json:"awayLogo"`
	Date     string `json:"date"`
}

type GetNextMatchesRequest struct {
	TeamName string `json:"teamName"`
}

type GetNextMatchesResponse struct {
	Matches []NextMatch `json:"matches"`
}

type GetNextMatchesHandler struct {
	GetNextMatchesUseCase soccer.GetNextMatchesUseCaseInterface
}

func NewGetNextMatchesHandler(getNextMatchesUseCase soccer.GetNextMatchesUseCaseInterface) *GetNextMatchesHandler {
	return &GetNextMatchesHandler{
		GetNextMatchesUseCase: getNextMatchesUseCase,
	}
}

func (gnmh GetNextMatchesHandler) Handle(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get("teamName")
	if teamName == "" {
		apperr.NewHttpError(w, apperr.ErrTeamNameRequired)
		return
	}

	matches, err := gnmh.GetNextMatchesUseCase.Execute(
		r.Context(),
		soccer.GetNextMatchesUseCaseInputDTO{TeamName: teamName},
	)
	if err != nil {
		if appErr, ok := err.(*apperr.AppErr); ok {
			apperr.NewHttpError(w, appErr)
		} else {
			apperr.NewHttpError(w, apperr.NewInternalServerError("Internal server error"))
		}
		return
	}

	response := GetNextMatchesResponse{
		Matches: make([]NextMatch, len(matches)),
	}

	for i, match := range matches {
		response.Matches[i] = NextMatch{
			Title:    match.Title,
			Home:     match.Home,
			HomeLogo: match.HomeLogo,
			Away:     match.Away,
			AwayLogo: match.AwayLogo,
			Date:     match.Date,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
