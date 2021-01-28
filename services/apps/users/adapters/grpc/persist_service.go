package grpc

import (
	"context"

	pb "github.com/jcfug8/daylier-backend/protos/backend/users"
	"github.com/jcfug8/daylier-backend/protos/commons"
	"github.com/jcfug8/daylier-backend/services/apps/users/ports/persist"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

// PersistServer - represents the proto/grpc implemention of the chat persist grpc service.
type PersistServer struct {
	service persist.Service
}

// NewPersistServer returns a new instance of the chat persist grpc service, and creates the internal proxy service.
func NewPersistServer(service persist.Service) *PersistServer {
	return &PersistServer{
		service: service,
	}
}

// Name returns the name of the chat persist service.
func (s *PersistServer) Name() string {
	return "omni-chat-persist"
}

// Register registers s to the grpc implementation of the service.
func (s *PersistServer) Register(server *grpc.Server) error {
	log.Info("registering chat persist")

	pb.RegisterUsersPersistServer(server, s)
	return nil
}

// Close call close on the service proxy
func (s *PersistServer) Close() error {
	return s.service.Close()
}

func (s *PersistServer) Ping(ctx context.Context, req *commons.PingReq) (*commons.PingRes, error) {
	return s.service.Ping(ctx, req)
}
