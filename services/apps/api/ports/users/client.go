package users

import (
	"context"

	"github.com/jcfug8/daylier-backend/protos/commons"
)

type Client interface {
	Ping(context.Context, *commons.PingReq) (*commons.PingRes, error)

	Close() error
}
