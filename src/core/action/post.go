package action

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type PostStore struct {
	Tweet     string                `form:"tweet" binding:"required"`
	ViewState int                   `form:"view_state" binding:"required"`
	File      *multipart.FileHeader `form:"file"`
	UserID    uuid.UUID             `form:"user_id"`
}

type PostUser struct {
	UserID uuid.UUID `form:"user_id"`
	PostID string    `form:"post_id" binding:"required"`
}

type PostPagination struct {
	Page        int       `form:"page"`
	Limit       int       `form:"limit"`
	Search      string    `form:"search"`
	UserFolowID string    `form:"user_flow_id"`
	UserID      uuid.UUID `form:"user_id"`
}
