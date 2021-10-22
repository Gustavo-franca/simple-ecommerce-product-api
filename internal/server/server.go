package server

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
)

func NewRouter() chi.Router {
	return chi.NewRouter()
}

func New(handler chi.Router) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
}

func StartServer(l fx.Lifecycle, server *http.Server) {
	l.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					err := server.ListenAndServe()
					if err != nil && !errors.Is(err, http.ErrServerClosed) {
						panic(err)
					}

				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				_ = server.Close()
				return nil
			},
		},
	)
}
