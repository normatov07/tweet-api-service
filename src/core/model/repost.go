package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateRepostModel struct {
	UserID      uuid.UUID
	PostID      uuid.UUID
	Description string
	CreatedAt   time.Time
}
