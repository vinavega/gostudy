package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/vinavega/gostudy/internal/database"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed(dbFeed)
}

func databaseUserToUser(dbUser database.User) User {
	return User(dbUser)
}
