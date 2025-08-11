package webserver

import (
	"context"
	"fmt"
	"net/http"

	handler "gitgub.com/Sup3r-Us3r/go-soccer/internal/infra/web/handler/soccer"
	usecase "gitgub.com/Sup3r-Us3r/go-soccer/internal/usecase/soccer"
)

type WebServer struct {
	mux    *http.ServeMux
	server *http.Server
}

func NewWebServer(addr string) *WebServer {
	mux := http.NewServeMux()

	latestMatchesUseCase := usecase.NewGetLatestMatchesUseCase()
	latestMatchesHandler := handler.NewGetLatestMatchesHandler(latestMatchesUseCase)

	nextMatchesUseCase := usecase.NewGetNextMatchesUseCase()
	nextMatchesHandler := handler.NewGetNextMatchesHandler(nextMatchesUseCase)

	playersUseCase := usecase.NewGetPlayersUseCase()
	playersHandler := handler.NewGetPlayersHandler(playersUseCase)

	transfersUseCase := usecase.NewGetTransfersUseCase()
	transfersHandler := handler.NewGetTransfersHandler(transfersUseCase)

	trophiesUseCase := usecase.NewGetTrophiesUseCase()
	trophiesHandler := handler.NewGetTrophiesHandler(trophiesUseCase)

	mux.HandleFunc("GET /api/soccer/matches/latest", latestMatchesHandler.Handle)
	mux.HandleFunc("GET /api/soccer/matches/next", nextMatchesHandler.Handle)
	mux.HandleFunc("GET /api/soccer/players", playersHandler.Handle)
	mux.HandleFunc("GET /api/soccer/transfers", transfersHandler.Handle)
	mux.HandleFunc("GET /api/soccer/trophies", trophiesHandler.Handle)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &WebServer{
		mux:    mux,
		server: srv,
	}
}

func (ws *WebServer) Start() error {
	fmt.Printf("Server started at %s\n", ws.server.Addr)
	return ws.server.ListenAndServe()
}

func (ws *WebServer) Stop(ctx context.Context) error {
	return ws.server.Shutdown(ctx)
}
