package soccer

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Sup3r-Us3r/go-soccer/internal/apperr"
	"github.com/Sup3r-Us3r/go-soccer/internal/util"
)

type Championship struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type GetTrophiesUseCaseInputDTO struct {
	TeamName string
}

type GetTrophiesUseCaseOutputDTO struct {
	Year         string       `json:"year"`
	Championship Championship `json:"championship"`
}

type GetTrophiesUseCaseInterface interface {
	Execute(ctx context.Context, input GetTrophiesUseCaseInputDTO) ([]GetTrophiesUseCaseOutputDTO, error)
}

type GetTrophiesUseCase struct{}

func NewGetTrophiesUseCase() *GetTrophiesUseCase {
	return &GetTrophiesUseCase{}
}

func (gtuc GetTrophiesUseCase) Execute(ctx context.Context, input GetTrophiesUseCaseInputDTO) ([]GetTrophiesUseCaseOutputDTO, error) {
	res, err := http.Get(fmt.Sprintf("https://www.placardefutebol.com.br/time/%s/titulos", util.Slugify(input.TeamName)))
	if err != nil {
		return nil, apperr.NewInternalServerError("Failed to fetch trophies")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, apperr.NewInternalServerError(fmt.Sprintf("Failed to fetch trophies: %d", res.StatusCode))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, apperr.NewInternalServerError("Failed to parse response body")
	}

	var trophies []GetTrophiesUseCaseOutputDTO

	doc.Find(".table__row").Each(func(i int, s *goquery.Selection) {
		yearCell := s.Find(".width_20.text")
		championshipCell := s.Find("a.width_75.link.text-left")
		championshipName := championshipCell.Find(".table__row-cell--text")

		trophy := GetTrophiesUseCaseOutputDTO{
			Year: strings.TrimSpace(yearCell.Text()),
		}

		trophy.Championship.Name = strings.TrimSpace(championshipName.Text())
		href, _ := championshipCell.Attr("href")
		trophy.Championship.URL = href

		trophies = append(trophies, trophy)
	})

	return trophies, nil
}
