package handler

import "net/http"

type User interface {
	User(w http.ResponseWriter, r *http.Request)
}
