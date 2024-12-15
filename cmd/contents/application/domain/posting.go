package domain

import (
	"errors"
	"time"
)

type Content struct {
	ID        string
	Title     string `validate:"required"`
	UserID    string `validate:"required,uuid"`
	Infill    string `validate:"required"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func PostContent(title string, userID string, infill string) (*Content, error) {
	content := Content{
		Title:     title,
		UserID:    userID,
		Infill:    infill,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	err := validate.Struct(&content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (c *Content) Update(title string, infill string, userID string) error {
	if c.UserID != userID {
		return errors.New("access denied")
	}

	if title == "" {
		return errors.New("title cannot be empty")
	}
	if infill == "" {
		return errors.New("content cannot be empty")
	}

	c.Title = title
	c.Infill = infill
	c.UpdatedAt = time.Now()

	return nil
}

func (c *Content) Delete(userID string) error {
	if c.UserID != userID {
		return errors.New("access denied")
	}

	return nil
}
