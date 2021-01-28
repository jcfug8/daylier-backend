package api

import (
	"context"

	"github.com/jcfug8/daylier-backend/protos/commons"
)

type Service interface {
	Ping(context.Context, *commons.PingReq) (*commons.PingRes, error)
	PingUsers(context.Context, *commons.PingReq) (*commons.PingRes, error)

	Close() error
}
