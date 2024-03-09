package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/myk-francis/learning-go/internal/database"
)

type User struct {
	ID		uuid.UUID `json:"id"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAtAt		time.Time `json:"updated_at"`
	Name		string `json:"name"`
	ApiKey		string `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAtAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}