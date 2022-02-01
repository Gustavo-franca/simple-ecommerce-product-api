package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
	"os"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Allow-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			_, err := w.Write([]byte(nil))
			if err != nil {
				_, _ = w.Write([]byte(nil))
				return
			}
			return

		}
		next.ServeHTTP(w, r)
	})
}

func NewRouter() chi.Router {
	handler := chi.NewRouter()
	handler.Use(Cors)
	return handler
}

func New(handler chi.Router) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
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
