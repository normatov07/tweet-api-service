package model

import (
	"time"

	"github.com/google/uuid"
)

type ResourceModel struct {
	ID           uuid.UUID
	ResourceID   uuid.UUID
	ResourceType string
	UserID       uuid.UUID
	Size         int64
	Name         string
	Path         string
	Format       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ResourceGet struct {
	ID   uuid.UUID
	Path string
}

type ResourceDelete struct {
	ResourceID   uuid.UUID
	ResourceType string
	UserID       uuid.UUID
}
