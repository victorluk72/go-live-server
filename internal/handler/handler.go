package handler

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

//This is for Dependency injection...
func NewHandler() *Handler {
	return &Handler{}
}

//Basic handler r - http request, w - response
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello live server")
}
