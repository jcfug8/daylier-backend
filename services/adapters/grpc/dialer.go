package grpc

import (
	"google.golang.org/grpc"
)

// Dialer - defines a dialer
type Dialer interface {
	Dial(addr string, options ...grpc.DialOption) (*grpc.ClientConn, error)
}

type CSDialer struct{}

func (d *CSDialer) Dial(addr string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(addr, options...)
}
