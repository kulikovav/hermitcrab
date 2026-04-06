package server

import (
	"context"

	"github.com/seal-io/walrus/utils/gopool"

	"github.com/kulikovav/hermitcrab/pkg/database"
	"github.com/kulikovav/hermitcrab/pkg/health"
)

// registerHealthCheckers registers the health checkers into the global health registry.
func (r *Server) registerHealthCheckers(ctx context.Context, opts initOptions) error {
	cs := health.Checkers{
		health.CheckerFunc("database", getDatabaseHealthChecker(opts.BoltDriver)),
		health.CheckerFunc("gopool", getGoPoolHealthChecker()),
	}

	return health.Register(ctx, cs)
}

func getDatabaseHealthChecker(db database.BoltDriver) health.Check {
	return func(ctx context.Context) error {
		return database.IsConnected(ctx, db)
	}
}

func getGoPoolHealthChecker() health.Check {
	return func(_ context.Context) error {
		return gopool.IsHealthy()
	}
}
