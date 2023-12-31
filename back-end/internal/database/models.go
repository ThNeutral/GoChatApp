// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                   uuid.UUID
	CreatedAt            time.Time
	UpdatedAt            time.Time
	AccessTokenUpdatedAt time.Time
	Username             string
	Password             string
	Email                string
	AccessToken          string
}
