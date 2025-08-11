package soccer

import (
	"encoding/json"
	"net/http"

	"github.com/Sup3r-Us3r/go-soccer/internal/apperr"
	"github.com/Sup3r-Us3r/go-soccer/internal/usecase/soccer"
)

type Match struct {
	Title      string `json:"title"`
	Home       string `json:"home"`
	HomeLogo   string `json:"homeLogo"`
	Away       string `json:"away"`
	AwayLogo   string `json:"awayLogo"`
	Date       string `json:"date"`
	ScoreBoard string `json:"scoreBoard"`
}

type GetLatestMatchesRequest struct {
	TeamName string `json:"teamName"`
}

type GetLatestMatchesResponse struct {
	Matches []Match `json:"matches"`
}

type GetLatestMatchesHandler struct {
	GetLatestMatchesUseCase soccer.GetLatestMatchesUseCaseInterface
}

func NewGetLatestMatchesHandler(getLatestMatchesUseCase soccer.GetLatestMatchesUseCaseInterface) *GetLatestMatchesHandler {
	return &GetLatestMatchesHandler{
		GetLatestMatchesUseCase: getLatestMatchesUseCase,
	}
}

func (glmh GetLatestMatchesHandler) Handle(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get("teamName")
	if teamName == "" {
		apperr.NewHttpError(w, apperr.ErrTeamNameRequired)
		return
	}

	matches, err := glmh.GetLatestMatchesUseCase.Execute(
		r.Context(),
		soccer.GetLatestMatchesUseCaseInputDTO{TeamName: teamName},
	)
	if err != nil {
		if appErr, ok := err.(*apperr.AppErr); ok {
			apperr.NewHttpError(w, appErr)
		} else {
			apperr.NewHttpError(w, apperr.NewInternalServerError("Internal server error"))
		}
		return
	}

	response := GetLatestMatchesResponse{
		Matches: make([]Match, len(matches)),
	}

	for i, match := range matches {
		response.Matches[i] = Match{
			Title:      match.Title,
			Home:       match.Home,
			HomeLogo:   match.HomeLogo,
			Away:       match.Away,
			AwayLogo:   match.AwayLogo,
			Date:       match.Date,
			ScoreBoard: match.ScoreBoard,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
