package grpc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type Service interface {
	Register(server *grpc.Server) error
	Close() error
}

type BufnetServer struct {
	lis     *bufconn.Listener
	t       *testing.T
	service Service
	server  *grpc.Server
}

func (s *BufnetServer) Serve() error {
	s.server = grpc.NewServer()
	err := s.service.Register(s.server)
	require.NoError(s.t, err)
	return s.server.Serve(s.lis)
}

func (s *BufnetServer) Close() {
	err := s.service.Close()
	require.NoError(s.t, err)
	s.server.Stop()
}
