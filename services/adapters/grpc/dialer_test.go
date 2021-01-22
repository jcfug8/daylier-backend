package grpc_test

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type BufnetDialer struct {
	addr string
	lis  *bufconn.Listener
	t    *testing.T
}

func (d *BufnetDialer) Dial(addr string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	require.Equal(d.t, d.addr, addr)

	return grpc.Dial("bufnet",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDialer(func(string, time.Duration) (net.Conn, error) {
			return d.lis.Dial()
		}),
	)
}
