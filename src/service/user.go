package service

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/knwoop/lets-gnomock/src/models"
)

type UserService struct {
	conn *sqlx.DB
}

func NewUserService(conn *sqlx.DB) (*UserService, error) {
	return &UserService{
		conn: conn,
	}, nil
}

func (s *UserService) Create(ctx context.Context, username string) (string, error) {
	u := models.User{}
	u.Username = username
	u.CreatedAt = time.Now()
	if err := u.Insert(s.conn); err != nil {
		return "", fmt.Errorf("error user insert: %w", err)
	}
	return username, nil
}

func (s *UserService) Get(ctx context.Context, username string) (*models.User, error) {
	u, err := models.UserByUsername(s.conn, username)
	if err != nil {
		return nil, fmt.Errorf("error user insert: %w", err)
	}
	return u, nil
}
