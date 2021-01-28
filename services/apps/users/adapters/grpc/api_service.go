package grpc

import (
	"context"

	pb "github.com/jcfug8/daylier-backend/protos/backend/users"
	"github.com/jcfug8/daylier-backend/protos/commons"
	"github.com/jcfug8/daylier-backend/services/apps/users/ports/api"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

// ApiServer - represents the proto/grpc implemention of the chat api grpc service.
type ApiServer struct {
	service api.Service
}

// NewApiServer returns a new instance of the chat api grpc service, and creates the internal proxy service.
func NewApiServer(service api.Service) *ApiServer {
	return &ApiServer{
		service: service,
	}
}

// Name returns the name of the chat api service.
func (s *ApiServer) Name() string {
	return "omni-chat-api"
}

// Register registers s to the grpc implementation of the service.
func (s *ApiServer) Register(server *grpc.Server) error {
	log.Info("registering chat api")

	pb.RegisterUsersAPIServer(server, s)
	return nil
}

// Close call close on the service proxy
func (s *ApiServer) Close() error {
	return s.service.Close()
}

func (s *ApiServer) Ping(ctx context.Context, req *commons.PingReq) (*commons.PingRes, error) {
	return s.service.Ping(ctx, req)
}
