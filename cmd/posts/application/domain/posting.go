package domain

import (
	"errors"
	"time"
)

type Post struct {
	ID        string
	Title     string `validate:"required"`
	UserID    string `validate:"required,uuid"`
	Content   string `validate:"required"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func CreatePost(title string, userID string, content string) (*Post, error) {
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

func (p *Post) Update(title string, content string, userID string) error {
	if p.UserID != userID {
		return errors.New("access denied")
	}

	if title != "" {
		p.Title = title
	}

	if content != "" {
		p.Content = content
	}

	p.UpdatedAt = time.Now()

	return nil
}

func (p *Post) Delete(userID string) error {
	if p.UserID != userID {
		return errors.New("access denied")
	}

	return nil
}
