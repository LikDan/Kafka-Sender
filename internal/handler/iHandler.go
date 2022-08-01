package handler

import "net/http"

type IHandler interface {
	send(w http.ResponseWriter, r *http.Request)
}
