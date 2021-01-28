package http

import (
	"net/http"

	"github.com/jcfug8/daylier-backend/protos/commons"
	comomonsHTTP "github.com/jcfug8/daylier-backend/services/adapters/http"
	"github.com/jcfug8/daylier-backend/services/apps/api/ports/api"

	log "github.com/sirupsen/logrus"
)

type ApiService struct {
	service api.Service
}

func NewApiService(service api.Service) *ApiService {
	return &ApiService{
		service: service,
	}
}

func (c *ApiService) Name() string {
	return "api"
}

func (c *ApiService) Register(serveMux *http.ServeMux) error {
	serveMux.HandleFunc("/api/v1/ping", c.Ping)
	return nil
}

func (c *ApiService) Close() error {
	return c.service.Close()
}

func (c *ApiService) Ping(w http.ResponseWriter, r *http.Request) {
	log.Info("ping in http api adapter")

	res, err := c.service.Ping(r.Context(), &commons.PingReq{})
	err = comomonsHTTP.SendResponse(res, w, err)
	if err != nil {
		log.Warnf("error while pinging api service: %s", err)
	}
}
