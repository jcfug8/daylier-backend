package grpc_test

import (
	"testing"

	"github.com/jcfug8/daylier-backend/services/ports/public"

	"github.com/stretchr/testify/require"
	gRPC "google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type BufnetServer struct {
	lis     *bufconn.Listener
	t       *testing.T
	service public.GRPCService
	server  *gRPC.Server
}

func (s *BufnetServer) Serve() error {
	s.server = gRPC.NewServer()
	err := s.service.Register(s.server)
	require.NoError(s.t, err)
	return s.server.Serve(s.lis)
}

func (s *BufnetServer) Close() {
	err := s.service.Close()
	require.NoError(s.t, err)
	s.server.Stop()
}
