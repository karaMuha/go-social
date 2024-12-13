package application_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	postgres_test "github.com/karaMuha/go-social/internal/database/postgres/test_container"
	"github.com/karaMuha/go-social/posts/application"
	"github.com/karaMuha/go-social/posts/application/commands"
	"github.com/karaMuha/go-social/posts/postgres"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/thanhpk/randstr"
)

var userIDs []string

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

	// userIDs are needed to perform actions on posts due to foreign key policy in database
	setupUsers(s.T(), dbHandler)
	postsRepository := postgres.NewPostsRepository(dbHandler)

	s.app = application.New(postsRepository)
}

// creates a few users and saves the IDs for further usage
func setupUsers(t *testing.T, dbHandler *sql.DB) {
	query := `
		INSERT INTO users (email, username, user_password, registration_token, created_at, active)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	for i := range 10 {
		email := fmt.Sprintf("test%d@test.com", i)
		username := fmt.Sprintf("test%d", i)
		password := "test123"
		token := randstr.String(64)
		createdAt := time.Now()
		active := false
		row := dbHandler.QueryRow(query, email, username, password, token, createdAt, active)
		var id string
		err := row.Scan(&id)
		require.NoError(t, err)
		userIDs = append(userIDs, id)
	}
}

// clear tables after each test to avoid conflicts and side effects
func (s *ApplicationTestSuite) AfterTest() {
	queryClearPostsTable := `DELETE FROM posts`
	_, err := s.dbHandler.ExecContext(s.ctx, queryClearPostsTable)
	require.NoError(s.T(), err)
}

func (s *ApplicationTestSuite) TestCreatePost() {
	cmd := commands.CreatePostDto{
		Title:   "This is a title",
		UserID:  userIDs[0],
		Content: "This is the content",
	}

	postID, err := s.app.CreatePost(s.ctx, &cmd)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), postID)

	post, err := s.app.GetPost(s.ctx, postID)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), post)
	require.Equal(s.T(), cmd.UserID, post.UserID)
}
