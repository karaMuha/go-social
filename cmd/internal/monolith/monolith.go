package monolith

import (
	"context"
	"database/sql"
	"net/http"

	authtoken "github.com/karaMuha/go-social/internal/auth_token"
	"github.com/karaMuha/go-social/internal/config"
	"github.com/karaMuha/go-social/internal/mailer"
)

type IMonolith interface {
	Config() *config.Config
	DB() *sql.DB
	Mux() *http.ServeMux
	MailServer() mailer.IMailer
	TokenProvider() authtoken.ITokenProvider
}

type monolith struct {
	cfg           *config.Config
	db            *sql.DB
	mux           *http.ServeMux
	mailServer    mailer.IMailer
	context       context.Context
	modules       []IModule
	tokenProvider authtoken.ITokenProvider
}

type IModule interface {
	Startup(ctx context.Context, mono IMonolith) error
}

var _ IMonolith = (*monolith)(nil)

func NewMonolith(cfg *config.Config,
	db *sql.DB,
	mux *http.ServeMux,
	mailServer mailer.IMailer,
	modules []IModule,
	tokenGenerator authtoken.ITokenProvider,
) monolith {
	setProtectedEndpoints()
	return monolith{
		cfg:           cfg,
		db:            db,
		mux:           mux,
		mailServer:    mailServer,
		context:       context.Background(),
		modules:       modules,
		tokenProvider: tokenGenerator,
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

func (m *monolith) Config() *config.Config {
	return m.cfg
}

func (m *monolith) DB() *sql.DB {
	return m.db
}

func (m *monolith) Mux() *http.ServeMux {
	return m.mux
}

func (m *monolith) MailServer() mailer.IMailer {
	return m.mailServer
}

func (m *monolith) TokenProvider() authtoken.ITokenProvider {
	return m.tokenProvider
}
