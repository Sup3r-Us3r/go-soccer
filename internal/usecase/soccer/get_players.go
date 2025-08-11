package soccer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/Sup3r-Us3r/go-soccer/internal/apperr"
)

type GetPlayersUseCaseInputDTO struct {
	TeamName string
}

type GetPlayersUseCaseOutputDTO struct {
	Position string `json:"position"`
	Player   string `json:"player"`
	Age      string `json:"age"`
	Country  string `json:"country"`
}

type GetPlayersUseCaseInterface interface {
	Execute(ctx context.Context, input GetPlayersUseCaseInputDTO) ([]GetPlayersUseCaseOutputDTO, error)
}

type GetPlayersUseCase struct{}

func NewGetPlayersUseCase() *GetPlayersUseCase {
	return &GetPlayersUseCase{}
}

func (gpuc GetPlayersUseCase) Execute(ctx context.Context, input GetPlayersUseCaseInputDTO) ([]GetPlayersUseCaseOutputDTO, error) {
	res, err := http.Get(fmt.Sprintf("https://www.placardefutebol.com.br/time/%s/jogadores", input.TeamName))
	if err != nil {
		return nil, apperr.NewInternalServerError("Failed to fetch players")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, apperr.NewInternalServerError(fmt.Sprintf("Failed to fetch players: %d", res.StatusCode))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, apperr.NewInternalServerError("Failed to parse response body")
	}

	var playersList []GetPlayersUseCaseOutputDTO

	doc.Find(".table__row").Each(func(i int, s *goquery.Selection) {
		cells := s.Find(".table__row-cell.text")
		if cells.Length() >= 4 {
			position := cells.Eq(0).Text()
			player := cells.Eq(1).Text()
			age := cells.Eq(2).Text()
			country := cells.Eq(3).Text()

			playersList = append(playersList, GetPlayersUseCaseOutputDTO{
				Position: position,
				Player:   player,
				Age:      age,
				Country:  country,
			})
		}
	})

	return playersList, nil
}
