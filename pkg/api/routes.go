package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (srv *Server) NewRouter() (*mux.Router, error) {

	router := mux.NewRouter()
	router.HandleFunc("/api/0.1/subscriptions", srv.SubscriptionsListHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/0.1/subscriptions", srv.SubscriptionsCreateHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/0.1/subscriptions/{msisdn}", srv.SubscriptionsGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/0.1/subscriptions/{msisdn}", srv.SubscriptionsUpdateHandler).Methods(http.MethodPut)
	router.HandleFunc("/api/0.1/subscriptions/{msisdn}/toggle_paused", srv.SubscriptionsTogglePausedHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/0.1/subscriptions/{msisdn}/cancel", srv.SubscriptionsCancelHandler).Methods(http.MethodPost)

	return router, nil
}
