package soccer

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Sup3r-Us3r/go-soccer/internal/apperr"
)

type Team struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type GetTransfersUseCaseInputDTO struct {
	TeamName string
}

type GetTransfersUseCaseOutputDTO struct {
	Date   string `json:"date"`
	Type   string `json:"type"`
	Player string `json:"player"`
	Team   Team   `json:"team"`
}

type GetTransfersUseCaseInterface interface {
	Execute(ctx context.Context, input GetTransfersUseCaseInputDTO) ([]GetTransfersUseCaseOutputDTO, error)
}

type GetTransfersUseCase struct{}

func NewGetTransfersUseCase() *GetTransfersUseCase {
	return &GetTransfersUseCase{}
}

func (gtuc GetTransfersUseCase) Execute(ctx context.Context, input GetTransfersUseCaseInputDTO) ([]GetTransfersUseCaseOutputDTO, error) {
	res, err := http.Get(fmt.Sprintf("https://www.placardefutebol.com.br/time/%s/transferencias", input.TeamName))
	if err != nil {
		return nil, apperr.NewInternalServerError("Failed to fetch transfers")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, apperr.NewInternalServerError(fmt.Sprintf("Failed to fetch transfers: %d", res.StatusCode))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, apperr.NewInternalServerError("Failed to parse response body")
	}

	var transfers []GetTransfersUseCaseOutputDTO

	doc.Find(".table__row").Each(func(i int, s *goquery.Selection) {
		dateCell := s.Find(".table__row-cell--transfers-text")
		playerCell := s.Find(".table__row-cell.text").Eq(1)
		teamCell := s.Find(".table__row-cell.text").Eq(2)
		teamLink := teamCell.Find("a")

		htmlContent, _ := dateCell.Html()
		dateParts := strings.Split(htmlContent, "<br/>")

		var date, transferType string
		if len(dateParts) >= 2 {
			date = strings.TrimSpace(stripHTML(dateParts[0]))
			transferType = strings.TrimSpace(stripHTML(dateParts[1]))
		}

		transfer := GetTransfersUseCaseOutputDTO{
			Date:   date,
			Type:   transferType,
			Player: strings.TrimSpace(playerCell.Text()),
		}

		transfer.Team.Name = strings.TrimSpace(teamLink.Text())
		href, _ := teamLink.Attr("href")
		transfer.Team.URL = href

		transfers = append(transfers, transfer)
	})

	return transfers, nil
}

func stripHTML(input string) string {
	re := regexp.MustCompile("<[^>]*>")

	return re.ReplaceAllString(input, "")
}
