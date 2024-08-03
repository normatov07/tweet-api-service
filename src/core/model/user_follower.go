package model

import "github.com/google/uuid"

type StoreUserFollower struct {
	UserID     uuid.UUID
	FollowerID uuid.UUID
}
