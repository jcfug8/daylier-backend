package http

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jcfug8/daylier-backend/services/ports/public"

	log "github.com/sirupsen/logrus"
)

// Server - defines a dialer
type Server interface {
	Serve() (Server, error)
}

type CSServer struct {
	addr    string
	service public.HTTPService
}

func NewCSServer(addr string, service public.HTTPService) *CSServer {
	return &CSServer{
		addr:    addr,
		service: service,
	}
}

func (d *CSServer) Serve() error {
	serveMux := http.NewServeMux()
	server := &http.Server{
		Addr:    d.addr,
		Handler: serveMux,
	}
	err := d.service.Register(serveMux)
	if err != nil {
		log.Warningf("could not register %s service to http server: %v", d.service.Name(), err)
		return err
	}

	var sig os.Signal

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		sig = <-ch
		log.Infof("signal '%s' received for http server with %s service", sig, d.service.Name())
		server.Close()
	}()

	log.Infof("http server with %s service started", d.service.Name())
	if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Warnf("could not close http server with %s service: %v", d.service.Name(), err)
		return err
	}

	d.service.Close()

	return nil
}
