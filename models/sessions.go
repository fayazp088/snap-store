package models

import (
	"database/sql"
	"fmt"

	"github.com/fayazp088/snap-store/rand"
)

const (
	MinBytesPerToken = 32
)

type Session struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Token     string `json:"token"`
	TokenHash string `json:"token_hash"`
}

type SessionService struct {
	DB            *sql.DB
	BytesPerToken int
}

func (ss SessionService) Create(userID int) (*Session, error) {

	bytePerToken := ss.BytesPerToken

	if bytePerToken < MinBytesPerToken {
		bytePerToken = MinBytesPerToken
	}

	token, err := rand.String(bytePerToken)

	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}

	session := Session{
		UserID: userID,
		Token:  token,
	}
	// TODO: Create the session token
	// TODO: Implement SessionService.Create
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
