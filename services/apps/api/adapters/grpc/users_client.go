package grpc

import (
	"context"

	pb "github.com/jcfug8/daylier-backend/protos/backend/users"
	"github.com/jcfug8/daylier-backend/protos/commons"
	commonsGRPC "github.com/jcfug8/daylier-backend/services/adapters/grpc"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

// UsersClient is the gRPC implementation of the users client port.
type UsersClient struct {
	client pb.UsersAPIClient
	conn   *grpc.ClientConn
}

// NewUsersClient returns a new UsersClient
func NewUsersClient(addr string, dialer commonsGRPC.Dialer) (*UsersClient, error) {
	conn, err := dialer.Dial(
		addr,
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Warningf("Failed connecting to users service at %s", addr)
		return &UsersClient{}, err
	}

	client := pb.NewUsersAPIClient(conn)
	log.Infof("Successfully created new users client for address %s", addr)

	return &UsersClient{
		client: client,
		conn:   conn,
	}, nil
}

// Ping -
func (s *UsersClient) Ping(ctx context.Context, req *commons.PingReq) (*commons.PingRes, error) {
	res, err := s.client.Ping(ctx, req)
	return res, err
}

// Close -
func (s *UsersClient) Close() error {
	return s.conn.Close()
}
