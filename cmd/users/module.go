package users

import (
	"context"

	"github.com/karaMuha/go-social/internal/monolith"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.IMonolith) error {
	return nil
}
