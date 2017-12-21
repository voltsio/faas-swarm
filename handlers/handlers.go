package handlers

import (
	"net/http"
)

func ReplicaReader() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func ReplicaUpdater() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
