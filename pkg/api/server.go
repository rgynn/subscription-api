package api

import (
	"fmt"
	"net/http"

	"github.com/rgynn/subscription-api/pkg/config"
	"github.com/rgynn/subscription-api/pkg/subscription"
	subs "github.com/rgynn/subscription-api/pkg/subscription/service"
)

type Server struct {
	*http.Server
	subscriptions subscription.Repository
}

func NewServerFromConfig(cfg *config.Config) (*Server, error) {

	srv := &Server{}

	subscriptions, err := subs.NewServiceFromConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize subscription service for server: %w", err)
	}

	srv.subscriptions = subscriptions

	router, err := srv.NewRouter()
	if err != nil {
		return nil, fmt.Errorf("failed to get routes for server: %w", err)
	}

	srv.Server = &http.Server{
		Addr:         cfg.Port,
		IdleTimeout:  cfg.IdleTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		Handler:      router,
	}

	return srv, nil
}
