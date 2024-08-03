package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/normatov07/mini-tweet/core/action"
	"github.com/normatov07/mini-tweet/core/model"
	"github.com/normatov07/mini-tweet/core/repository"
)

type RepostService struct {
	repo repository.PostRepo
}

func GetRepostService(repo repository.PostRepo) *RepostService {
	return &RepostService{
		repo: repo,
	}
}

func (s RepostService) StoreRepost(acn action.RepostCreate) error {
	postUID, err := uuid.Parse(acn.PostID)
	if err != nil {
		return errors.New("post id is not valid")
	}
	return s.repo.CreateRepost(model.CreateRepostModel{
		UserID:      acn.UserID,
		PostID:      postUID,
		Description: acn.Description,
		CreatedAt:   time.Now(),
	})
}

func (s RepostService) DeleteRepost(acn action.RepostDelete) error {
	postUID, err := uuid.Parse(acn.PostID)
	if err != nil {
		return errors.New("post id is not valid")
	}
	return s.repo.DeleteRepost(model.PostDeleteModel{
		ID:     postUID,
		UserID: acn.UserID,
	})
}
