package persist

import (
	"context"

	"github.com/jcfug8/daylier-backend/protos/commons"
)

type Service interface {
	Ping(ctx context.Context, req *commons.PingReq) (*commons.PingRes, error)

	Close() error
}
