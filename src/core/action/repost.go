package action

import (
	"github.com/google/uuid"
)

type RepostCreate struct {
	UserID      uuid.UUID `form:"user_id"`
	PostID      string    `form:"post_id" binding:"required"`
	Description string    `form:"description" binding:"required"`
}

type RepostDelete struct {
	UserID uuid.UUID `form:"user_id"`
	PostID string    `form:"post_id" binding:"required"`
}
