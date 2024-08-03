package action

import "github.com/google/uuid"

type StoreUserFollower struct {
	UserID     string    `form:"user_id"  binding:"required"`
	FollowerID uuid.UUID `form:"follower_id"`
}
