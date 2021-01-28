package grpc

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jcfug8/daylier-backend/services/ports/public"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Server - defines a dialer
type Server interface {
	Serve() (Server, error)
}

type CSServer struct {
	addr    string
	service public.GRPCService
}

// NewCSServer -
func NewCSServer(addr string, service public.GRPCService) *CSServer {
	return &CSServer{
		addr:    addr,
		service: service,
	}
}

func (d *CSServer) Serve() error {
	lis, err := net.Listen("tcp", d.addr)
	if err != nil {
		log.Errorf("failed to listen at %s grpc: %v", d.addr, err)
	}

	// create a grpc server object
	grpcServer := grpc.NewServer()

	err = d.service.Register(grpcServer)
	if err != nil {
		log.Warningf("could not register %s grpc: %v", d.service.Name(), err)
		return err
	}

	var sig os.Signal

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		sig = <-ch
		log.Info("signal %s received in %s grpc", sig, d.service.Name())
		grpcServer.GracefulStop()
	}()

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Errorf("failed to serve: %s", err)
		return err
	}

	d.service.Close()

	return nil
}
