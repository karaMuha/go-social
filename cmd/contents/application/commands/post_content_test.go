package commands

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/karaMuha/go-social/contents/application/domain"
	"github.com/karaMuha/go-social/contents/application/ports/driven"
	"github.com/karaMuha/go-social/contents/postgres"
	postgres_test "github.com/karaMuha/go-social/internal/database/postgres/test_container"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/thanhpk/randstr"
)

// userIDs are needed to perform actions on posts due to foreign key policy in database
var userIDs []string

type PostContentTestSuite struct {
	suite.Suite
	ctx         context.Context
	dbHandler   *sql.DB
	cmd         PostContentCommand
	contentRepo driven.ContentsRepository
}

func TestPostContentSuite(t *testing.T) {
	suite.Run(t, &PostContentTestSuite{})
}

func (s *PostContentTestSuite) SetupSuite() {
	s.ctx = context.Background()
	initScriptPath := filepath.Join("..", "..", "..", "..", "dbscripts", "public_schema.sql")
	dbHandler, err := postgres_test.CreatePostgresContainer(s.ctx, initScriptPath)
	require.NoError(s.T(), err)
	s.dbHandler = dbHandler

	s.contentRepo = postgres.NewContentsRepository(dbHandler)
	s.cmd = NewPostContentCommand(s.contentRepo)

	domain.InitValidator()
	setupUsers(s.T(), dbHandler)
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

// clear table after each test to avoid conflicts and side effects
func (s *PostContentTestSuite) AfterTest() {
	queryClearPostsTable := `DELETE FROM posts`
	_, err := s.dbHandler.ExecContext(s.ctx, queryClearPostsTable)
	require.NoError(s.T(), err)
}

func (s *PostContentTestSuite) TestCreateContent() {
	cmd := PostContentDto{
		Title:   "This is a title",
		UserID:  userIDs[0],
		Content: "This is the content",
	}

	postID, err := s.cmd.PostContent(s.ctx, &cmd)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), postID)

	post, err := s.contentRepo.GetByID(s.ctx, postID)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), post)
	require.Equal(s.T(), cmd.UserID, post.UserID)
}
