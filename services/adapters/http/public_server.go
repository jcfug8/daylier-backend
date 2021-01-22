package http

import (
	standHTTP "net/http"
	"os"
	"os/signal"
	"syscall"

	log "git.tcncloud.net/m/neo/commons/logging"

	"git.tcncloud.net/m/neo/services/omnichannel/email/dispatcher/ports/http"
)

// Server - defines a dialer
type Server interface {
	Serve() (Server, error)
}

type CSServer struct {
	addr    string
	service http.Service
}

func NewCSServer(addr string, service http.Service) *CSServer {
	return &CSServer{
		addr:    addr,
		service: service,
	}
}

func (d *CSServer) Serve() error {
	serveMux := standHTTP.NewServeMux()
	server := &standHTTP.Server{
		Addr:    d.addr,
		Handler: serveMux,
	}
	err := d.service.Register(serveMux)
	if err != nil {
		log.Warningf("could not register %s http: %v", d.service.Name(), err)
		return err
	}

	var sig os.Signal

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		sig = <-ch
		server.Close()
	}()

	log.Infof("%s http started", d.service.Name())
	if err = server.ListenAndServe(); err != nil && err != standHTTP.ErrServerClosed {
		log.Warnf("could not close %s server: %v", d.service.Name(), err)
		return err
	}

	return nil
}
