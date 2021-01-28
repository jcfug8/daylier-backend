package grpc

import (
	"context"

	pb "github.com/jcfug8/daylier-backend/protos/backend/users"
	"github.com/jcfug8/daylier-backend/protos/commons"
	commonsGRPC "github.com/jcfug8/daylier-backend/services/adapters/grpc"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

// PersistClient is the gRPC implementation of the users client port.
type PersistClient struct {
	client pb.UsersPersistClient
	conn   *grpc.ClientConn
}

// NewPersistClient returns a new PersistClient
func NewPersistClient(addr string, dialer commonsGRPC.Dialer) (*PersistClient, error) {
	conn, err := dialer.Dial(
		addr,
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Warningf("Failed connecting to users persist service at %s", addr)
		return &PersistClient{}, err
	}

	client := pb.NewUsersPersistClient(conn)
	log.Infof("Successfully created new users persist client for address %s", addr)

	return &PersistClient{
		client: client,
		conn:   conn,
	}, nil
}

// Ping -
func (s *PersistClient) Ping(ctx context.Context, req *commons.PingReq) (*commons.PingRes, error) {
	res, err := s.client.Ping(ctx, req)
	return res, err
}

// Close -
func (s *PersistClient) Close() error {
	return s.conn.Close()
}
