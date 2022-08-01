package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"sender/internal/controller"
	"sender/internal/controller/validators"
	"sender/internal/entities"
)

type Handler struct {
	controller controller.IController
}

func NewHandler(controller controller.IController) *Handler {
	handler := Handler{controller: controller}

	http.HandleFunc("/api/send", handler.send)

	return &handler
}

func (h *Handler) send(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var message entities.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "decoding error: "+err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := h.controller.SendMessage(ctx, message); err != nil {
		if errors.Is(err, validators.ValidationError) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}
}
