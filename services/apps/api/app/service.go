package app

import (
	"context"

	"github.com/jcfug8/daylier-backend/protos/commons"
	"github.com/jcfug8/daylier-backend/services/apps/api/ports/users"

	log "github.com/sirupsen/logrus"
)

type Service struct {
	UsersClient users.Client
}

func NewService(usersClient users.Client) *Service {
	return &Service{
		UsersClient: usersClient,
	}
}

func (s *Service) Close() error {
	var err error
	if s.UsersClient != nil {
		err = s.UsersClient.Close()
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Ping(ctx context.Context, req *commons.PingReq) (*commons.PingRes, error) {
	log.Info("ping in domain api app")
	return &commons.PingRes{}, nil
}

func (s *Service) PingUsers(ctx context.Context, req *commons.PingReq) (*commons.PingRes, error) {
	return s.UsersClient.Ping(ctx, req)
}
