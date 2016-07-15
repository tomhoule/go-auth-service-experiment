package main

import (
	"encoding/json"
	"net/http"

	"goji.io"

	"log"

	"goji.io/pat"
	"golang.org/x/net/context"
)

func getHandler() *goji.Mux {
	mux := goji.NewMux()
	mux.HandleC(pat.Post("/create-user"), handlerC(handleUserCreation))
	return mux
}

type userCreation struct {
	email    string
	password string
}

type handlerC func(ctx context.Context, w http.ResponseWriter, r *http.Request) (int, error)

func (h handlerC) ServeHTTPC(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	status, err := h(ctx, w, r)
	if err != nil {
		switch e := err.(type) {
		case (*json.InvalidUnmarshalError):
			log.Printf("Unmarshaling error (%s)", e.Error())
			w.WriteHeader(500)
		default:
			log.Printf("Error type left unhandled: %T", e)
		}
	}
	w.WriteHeader(status)
}

func handleUserCreation(ctx context.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	decoder := json.NewDecoder(r.Body)
	var payload userCreation
	if err := decoder.Decode(&payload); err != nil {
		return 400, err
	}
	return 200, nil
}
