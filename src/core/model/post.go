package model

import (
	"time"

	"github.com/google/uuid"
)

type PostModel struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Tweet     string
	ViewState int
	LikeCount int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostUpdateModel struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Tweet     string
	ViewState int
	UpdatedAt time.Time
}

type PostPaginationModel struct {
	Offset      int
	Limit       int
	Search      string
	UserFolowID string
	UserID      uuid.UUID
}

type PostLike struct {
	PostID uuid.UUID
	UserID uuid.UUID
}

type PostDeleteModel struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

type PostListModel struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID  `json:"user_id"`
	AuthorID    uuid.UUID  `json:"author_id"`
	Tweet       string     `json:"tweet"`
	ViewState   int        `json:"view_state"`
	Type        int        `json:"type"`
	LikeCount   int64      `json:"like_count"`
	Description NullString `json:"description" binding:"omitempty"`
	FileUrl     NullString `json:"file_url"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	AuthorData  UserModel  `json:"author_data,omitempty"`
}
