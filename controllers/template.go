package controllers

import "net/http"

type Template interface {
	Execute(w http.ResponseWriter, h *http.Request, data interface{})
}
