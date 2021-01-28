package public

import (
	"net/http"

	"google.golang.org/grpc"
)

type GRPCService interface {
	Register(*grpc.Server) error
	Name() string
	Close() error
}

type HTTPService interface {
	Register(*http.ServeMux) error
	Name() string
	Close() error
}
