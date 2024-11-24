package domain

import "time"

type Post struct {
	ID        string
	Title     string `validate:"required"`
	UserID    string `validate:"required,uuid"`
	Content   string `validate:"required"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func Create(title string, userID string, content string) (*Post, error) {
	post := Post{
		Title:     title,
		UserID:    userID,
		Content:   content,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	err := validate.Struct(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
