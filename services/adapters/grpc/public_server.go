package grpc

import (
	"log"
	"google.golang.org/grpc"
)

// Server - defines a dialer
type Server interface {
	Serve(settings ...service.Setting) (Server, error)
}

type CSServer struct {
	addr    string
	service service.Service
}

func NewCSServer(addr string, service service.Service) *CSServer {
	return &CSServer{
		addr:    addr,
		service: service,
	}
}

func (d *CSServer) Serve() error {
	server, err := grpc.Serve(
		service.WithService(d.service),
		service.WithAddr(d.addr),
		service.WithReflection(true),
	)

	if err != nil {
		log.Warningf("could not serve %s grpc: %v", d.service.Name(), err)
		return err
	}

	log.Infof("%s grpc started", d.service.Name())
	if err = server.Block(); err != nil {
		log.Warnf("could not close %s server: %v", d.service.Name(), err)
		return err
	}

	return nil
}
