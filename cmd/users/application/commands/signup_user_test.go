package commands

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"

	postgres_test "github.com/karaMuha/go-social/internal/database/postgres/test_container"
	"github.com/karaMuha/go-social/internal/mailer"
	"github.com/karaMuha/go-social/users/application/domain"
	"github.com/karaMuha/go-social/users/application/ports/driven"
	"github.com/karaMuha/go-social/users/postgres"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SignupUserTestSuite struct {
	suite.Suite
	ctx       context.Context
	cmd       SignupUserCommand
	dbHandler *sql.DB
	userRepo  driven.IUsersRepsitory
}

func TestSignupUserSuite(t *testing.T) {
	suite.Run(t, &SignupUserTestSuite{})
}

func (s *SignupUserTestSuite) SetupSuite() {
	s.ctx = context.Background()
	initScriptPath := filepath.Join("..", "..", "..", "..", "dbscripts", "public_schema.sql")
	dbHandler, err := postgres_test.CreatePostgresContainer(s.ctx, initScriptPath)
	require.NoError(s.T(), err)
	s.dbHandler = dbHandler

	usersRepository := postgres.NewUsersRepository(dbHandler)
	mailMock := mailer.MailerMock{}

	s.userRepo = usersRepository
	s.cmd = NewSignupUserCommand(usersRepository, &mailMock)

	domain.InitValidator()
}

// clear tables between tests to avoid conflicts and side effects
func (s *SignupUserTestSuite) AfterTest(suiteName, testName string) {
	queryClearUsersTable := `DELETE FROM users`

	_, err := s.dbHandler.ExecContext(s.ctx, queryClearUsersTable)
	require.NoError(s.T(), err)
}

func (s *SignupUserTestSuite) TestSignupSuccess() {
	params := SignupUserDto{
		Email:    "test@test.com",
		Username: "Testian",
		Password: "test123",
	}

	err := s.cmd.SignupUser(s.ctx, &params)
	require.NoError(s.T(), err)

	user, err := s.userRepo.GetByEmail(s.ctx, params.Email)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), user.ID)
	require.False(s.T(), user.Active)
}

func (s *SignupUserTestSuite) TestSignupFailToSendMail() {
	params := SignupUserDto{
		Email:    "error@error.com",
		Username: "Testian",
		Password: "test123",
	}

	err := s.cmd.SignupUser(s.ctx, &params)
	require.Error(s.T(), err)

	user, err := s.userRepo.GetByEmail(s.ctx, params.Email)
	require.Error(s.T(), err)
	require.Equal(s.T(), "user not found", err.Error())
	require.Nil(s.T(), user)
}
