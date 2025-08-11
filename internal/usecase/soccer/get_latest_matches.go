package soccer

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Sup3r-Us3r/go-soccer/internal/apperr"
)

type GetLatestMatchesUseCaseInputDTO struct {
	TeamName string
}

type GetLatestMatchesUseCaseOutputDTO struct {
	Title      string `json:"title"`
	Home       string `json:"home"`
	HomeLogo   string `json:"homeLogo"`
	Away       string `json:"away"`
	AwayLogo   string `json:"awayLogo"`
	Date       string `json:"date"`
	ScoreBoard string `json:"scoreBoard"`
}

type GetLatestMatchesUseCaseInterface interface {
	Execute(ctx context.Context, input GetLatestMatchesUseCaseInputDTO) ([]GetLatestMatchesUseCaseOutputDTO, error)
}

type GetLatestMatchesUseCase struct{}

func NewGetLatestMatchesUseCase() *GetLatestMatchesUseCase {
	return &GetLatestMatchesUseCase{}
}

func (glmuc GetLatestMatchesUseCase) Execute(ctx context.Context, input GetLatestMatchesUseCaseInputDTO) ([]GetLatestMatchesUseCaseOutputDTO, error) {
	res, err := http.Get(fmt.Sprintf("https://www.placardefutebol.com.br/time/%s/ultimos-jogos", input.TeamName))
	if err != nil {
		return nil, apperr.NewInternalServerError("Failed to fetch latest matches")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, apperr.NewInternalServerError(fmt.Sprintf("Failed to fetch latest matches: %s", res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, apperr.NewInternalServerError("Failed to parse response body")
	}

	var matches []GetLatestMatchesUseCaseOutputDTO

	doc.Find("a.match__lg div.match__lg_card").Each(func(i int, s *goquery.Selection) {
		title := s.Children().Eq(0).Text()
		home := s.Children().Eq(1).Text()
		away := s.Children().Eq(2).Text()

		homeLogo, _ := s.Children().Eq(3).Find("img").Attr("src")
		awayLogo, _ := s.Children().Eq(4).Find("img").Attr("src")

		date := s.Children().Eq(5).Find("div.match__lg_card--date").Text()
		if date == "" {
			date = "Encerrado"
		}

		scoreBoard := strings.TrimSpace(s.Children().Eq(5).Find("div.match__lg_card--scoreboard").Text())

		matches = append(matches, GetLatestMatchesUseCaseOutputDTO{
			Title:      title,
			Home:       home,
			HomeLogo:   homeLogo,
			Away:       away,
			AwayLogo:   awayLogo,
			Date:       date,
			ScoreBoard: scoreBoard,
		})
	})

	return matches, nil
}
