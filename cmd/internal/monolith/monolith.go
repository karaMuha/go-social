package monolith

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/karaMuha/go-social/internal/config"
)

type IMonolith interface {
	Config() config.Config
	DB() *sql.DB
	Mux() *http.ServeMux
}

type monolith struct {
	cfg     config.Config
	db      *sql.DB
	mux     *http.ServeMux
	context context.Context
	modules []Module
}

type Module interface {
	Startup(ctx context.Context, mono IMonolith) error
}

var _ IMonolith = (*monolith)(nil)

func NewMonolith(cfg config.Config, db *sql.DB, mux *http.ServeMux, modules []Module) monolith {
	return monolith{
		cfg:     cfg,
		db:      db,
		mux:     mux,
		context: context.Background(),
		modules: modules,
	}
}

func (m *monolith) InitModules() error {
	for _, module := range m.modules {
		err := module.Startup(m.context, m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *monolith) Config() config.Config {
	return m.cfg
}

func (m *monolith) DB() *sql.DB {
	return m.db
}

func (m *monolith) Mux() *http.ServeMux {
	return m.mux
}
