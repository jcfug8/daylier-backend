package http

import (
	"net/http"
)

type Service interface {
	Register(*http.ServeMux) error
	Name() string
}
