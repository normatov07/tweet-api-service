package repository

import (
	"github.com/google/uuid"
	"github.com/normatov07/mini-tweet/core/model"
)

type UserRepo interface {
	GetUserByLogin(login string) (model.UserModel, error)
	CreateUser(model model.UserModel) error
	GetUserFolowerID(userId uuid.UUID) ([]uuid.UUID, error)
	CreateUserFollower(m model.StoreUserFollower) error
	DeleteUserFollower(m model.StoreUserFollower) error
}
