package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rgynn/subscription-api/pkg/subscription"
)

// SubscriptionsListHandler for api
func (srv *Server) SubscriptionsListHandler(w http.ResponseWriter, r *http.Request) {

	result, err := srv.subscriptions.List(r.Context())
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	body, err := json.Marshal(result)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	if _, err := w.Write(body); err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}
}

// SubscriptionsGetHandler for api
func (srv *Server) SubscriptionsGetHandler(w http.ResponseWriter, r *http.Request) {

	msisdn := mux.Vars(r)["msisdn"]

	result, err := srv.subscriptions.Get(r.Context(), &msisdn)
	if err != nil {
		switch err {
		case subscription.ErrNotFound:
			NewErrorResponse(w, r, http.StatusNotFound, err)
		default:
			NewErrorResponse(w, r, http.StatusInternalServerError, err)
		}
		return
	}

	body, err := json.Marshal(result)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	if _, err := w.Write(body); err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}
}

// SubscriptionsCreateHandler for api
func (srv *Server) SubscriptionsCreateHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}
	defer r.Body.Close()

	var m *subscription.Model
	if err := json.Unmarshal(body, &m); err != nil {
		NewErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}

	result, err := srv.subscriptions.Create(r.Context(), m)
	if err != nil {
		switch err {
		case subscription.ErrAlreadyExists:
			NewErrorResponse(w, r, http.StatusConflict, err)
		case subscription.ErrNotFound:
			NewErrorResponse(w, r, http.StatusNotFound, err)
		default:
			NewErrorResponse(w, r, http.StatusInternalServerError, err)
		}
		return
	}

	resp, err := json.Marshal(result)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	if _, err := w.Write(resp); err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}
}

// SubscriptionsUpdateHandler for api
func (srv *Server) SubscriptionsUpdateHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}
	defer r.Body.Close()

	var m *subscription.Model
	if err := json.Unmarshal(body, &m); err != nil {
		NewErrorResponse(w, r, http.StatusBadRequest, err)
		return
	}

	result, err := srv.subscriptions.Update(r.Context(), m)
	if err != nil {
		switch err {
		case subscription.ErrNotFound:
			NewErrorResponse(w, r, http.StatusNotFound, err)
		default:
			NewErrorResponse(w, r, http.StatusInternalServerError, err)
		}
		return
	}

	resp, err := json.Marshal(result)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	if _, err := w.Write(resp); err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}
}

// SubscriptionsTogglePausedHandler for api
func (srv *Server) SubscriptionsTogglePausedHandler(w http.ResponseWriter, r *http.Request) {

	msisdn := mux.Vars(r)["msisdn"]

	result, err := srv.subscriptions.TogglePaused(r.Context(), &msisdn)
	if err != nil {
		switch err {
		case subscription.ErrNotFound:
			NewErrorResponse(w, r, http.StatusNotFound, err)
		default:
			NewErrorResponse(w, r, http.StatusInternalServerError, err)
		}
		return
	}

	resp, err := json.Marshal(result)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	if _, err := w.Write(resp); err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}
}

// SubscriptionsCancelHandler for api
func (srv *Server) SubscriptionsCancelHandler(w http.ResponseWriter, r *http.Request) {

	msisdn := mux.Vars(r)["msisdn"]

	result, err := srv.subscriptions.Cancel(r.Context(), &msisdn)
	if err != nil {
		switch err {
		case subscription.ErrNotFound:
			NewErrorResponse(w, r, http.StatusNotFound, err)
		default:
			NewErrorResponse(w, r, http.StatusInternalServerError, err)
		}
		return
	}

	resp, err := json.Marshal(result)
	if err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	if _, err := w.Write(resp); err != nil {
		NewErrorResponse(w, r, http.StatusInternalServerError, err)
		return
	}
}
