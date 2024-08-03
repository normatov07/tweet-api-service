package service

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/normatov07/mini-tweet/common/token"
	"github.com/normatov07/mini-tweet/common/utils"
	"github.com/normatov07/mini-tweet/core/action"
	"github.com/normatov07/mini-tweet/core/app_errors"
	"github.com/normatov07/mini-tweet/core/model"
	"github.com/normatov07/mini-tweet/core/repository"
)

type UserService struct {
	repo repository.UserRepo
}

func GetUserService(repo repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(atn action.UserStore) (string, error) {
	_, err := s.repo.GetUserByLogin(atn.Login)
	if err == nil {
		return "", app_errors.NewAppErr(app_errors.LOGIN_UNIQUE)
	} else if err != sql.ErrNoRows {
		log.Printf("register:check %v", err)
		return "", app_errors.NewAppErr(app_errors.SERVER_ERROR)
	}
	hash, _ := utils.HashPassword(atn.Password)

	model := model.UserModel{
		ID:        uuid.New(),
		Login:     atn.Login,
		Password:  hash,
		FirstName: atn.FirstName,
		LastName:  atn.LastName,
		Address:   atn.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = s.repo.CreateUser(model)
	if err != nil {
		log.Printf("register:create %v", err)
		return "", app_errors.NewAppErr(app_errors.SERVER_ERROR)
	}
	tMaker, err := token.NewPasetoMaker()
	if err != nil {
		return "", app_errors.NewAppErr(app_errors.USER_NOT_FOUND, atn.Login)
	}

	return tMaker.CreateToken(model)
}

func (s *UserService) CreateUserFollower(acn action.StoreUserFollower) error {
	userUID, err := uuid.Parse(acn.UserID)
	if err != nil {
		return errors.New("user id is not valid")
	}
	if userUID == acn.FollowerID {
		return errors.New("inccsesible operation")
	}

	return s.repo.CreateUserFollower(model.StoreUserFollower{
		UserID:     userUID,
		FollowerID: acn.FollowerID,
	})
}

func (s *UserService) DeleteUserFollower(acn action.StoreUserFollower) error {
	userUID, err := uuid.Parse(acn.UserID)
	if err != nil {
		return errors.New("user id is not valid")
	}
	if userUID == acn.FollowerID {
		return errors.New("inccsesible operation")
	}

	return s.repo.DeleteUserFollower(model.StoreUserFollower{
		UserID:     userUID,
		FollowerID: acn.FollowerID,
	})
}
