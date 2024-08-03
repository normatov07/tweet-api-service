package action

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type ResourceStore struct {
	ID                uuid.UUID
	ResourceID        uuid.UUID
	ResourceType      string
	UserID            uuid.UUID
	Size              int64
	Name              string
	Path              string
	Format            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	IsRemoveDublicate bool
	File              *multipart.FileHeader
}
