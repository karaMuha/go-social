package application_test

import (
	"context"
	"database/sql"
	"testing"

	postgres_test "github.com/karaMuha/go-social/internal/database/postgres/test_container"
	"github.com/karaMuha/go-social/internal/mailer"
	"github.com/karaMuha/go-social/users/application"
	"github.com/karaMuha/go-social/users/postgres"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ApplicationTestSuite struct {
	suite.Suite
	ctx       context.Context
	app       application.Application
	dbHandler *sql.DB
}

func TestApplicationSuite(t *testing.T) {
	suite.Run(t, &ApplicationTestSuite{})
}

func (s *ApplicationTestSuite) SetupSuite() {
	s.ctx = context.Background()

	dbHandler, err := postgres_test.CreatePostgresContainer(s.ctx)
	require.NoError(s.T(), err)
	s.dbHandler = dbHandler

	usersRepository := postgres.NewUsersRepository(dbHandler)
	followersRepository := postgres.NewFollowersRepository(dbHandler)
	mailMock := mailer.MailerMock{}

	s.app = application.New(usersRepository, followersRepository, &mailMock)
}

// clear tables between tests to avoid conflicts and side effects
func (s *ApplicationTestSuite) AfterTest(suiteName, testName string) {
	queryClearUsersTable := `DELETE FROM users`
	queryClearFollowersTable := `DELETE FROM followers`

	_, err := s.dbHandler.ExecContext(s.ctx, queryClearUsersTable)
	require.NoError(s.T(), err)

	_, err = s.dbHandler.ExecContext(s.ctx, queryClearFollowersTable)
	require.NoError(s.T(), err)
}
