package soccer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/Sup3r-Us3r/go-soccer/internal/apperr"
	"github.com/Sup3r-Us3r/go-soccer/internal/util"
)

type GetNextMatchesUseCaseInputDTO struct {
	TeamName string
}

type GetNextMatchesUseCaseOutputDTO struct {
	Title    string `json:"title"`
	Home     string `json:"home"`
	HomeLogo string `json:"homeLogo"`
	Away     string `json:"away"`
	AwayLogo string `json:"awayLogo"`
	Date     string `json:"date"`
}

type GetNextMatchesUseCaseInterface interface {
	Execute(ctx context.Context, input GetNextMatchesUseCaseInputDTO) ([]GetNextMatchesUseCaseOutputDTO, error)
}

type GetNextMatchesUseCase struct{}

func NewGetNextMatchesUseCase() *GetNextMatchesUseCase {
	return &GetNextMatchesUseCase{}
}

func (gnmuc GetNextMatchesUseCase) Execute(ctx context.Context, input GetNextMatchesUseCaseInputDTO) ([]GetNextMatchesUseCaseOutputDTO, error) {
	res, err := http.Get(fmt.Sprintf("https://www.placardefutebol.com.br/time/%s/proximos-jogos", util.Slugify(input.TeamName)))
	if err != nil {
		return nil, apperr.NewInternalServerError("Failed to fetch next matches")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, apperr.NewInternalServerError(fmt.Sprintf("Failed to fetch next matches: %s", res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, apperr.NewInternalServerError("Failed to parse response body")
	}

	var matches []GetNextMatchesUseCaseOutputDTO

	doc.Find("a.match__lg div.match__lg_card").Each(func(i int, s *goquery.Selection) {
		title := s.Children().Eq(0).Text()
		home := s.Children().Eq(1).Text()
		away := s.Children().Eq(2).Text()

		homeLogo, _ := s.Children().Eq(3).Find("img").Attr("src")
		awayLogo, _ := s.Children().Eq(4).Find("img").Attr("src")

		date := s.Children().Eq(5).Find("div.match__lg_card--datetime").Text()

		matches = append(matches, GetNextMatchesUseCaseOutputDTO{
			Title:    title,
			Home:     home,
			HomeLogo: homeLogo,
			Away:     away,
			AwayLogo: awayLogo,
			Date:     date,
		})
	})

	return matches, nil
}
