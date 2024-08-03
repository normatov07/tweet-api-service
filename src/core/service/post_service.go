package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/normatov07/mini-tweet/common/enums"
	"github.com/normatov07/mini-tweet/common/utils"
	"github.com/normatov07/mini-tweet/core/action"
	"github.com/normatov07/mini-tweet/core/model"
	"github.com/normatov07/mini-tweet/core/repository"
)

type PostService struct {
	repo    repository.PostRepo
	resRepo repository.ResourceRepo
	usrRepo repository.UserRepo
}

func GetPostService(repo repository.PostRepo, resRepo repository.ResourceRepo, usrRepo repository.UserRepo) *PostService {
	return &PostService{
		repo:    repo,
		resRepo: resRepo,
		usrRepo: usrRepo,
	}
}

func (s PostService) StorePost(acn action.PostStore) error {
	m := model.PostModel{
		ID:        uuid.New(),
		UserID:    acn.UserID,
		Tweet:     acn.Tweet,
		ViewState: acn.ViewState,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.repo.CreatePost(m)
	if err != nil {
		return err
	}

	return GetResourceService(s.resRepo).ResourceStore(action.ResourceStore{
		ID:                uuid.New(),
		ResourceID:        m.ID,
		ResourceType:      enums.POST_RESOURCE_TWEET,
		UserID:            acn.UserID,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		IsRemoveDublicate: true,
		File:              acn.File,
	})
}

func (s PostService) DeletePost(acn action.RepostDelete) error {
	postUID, err := uuid.Parse(acn.PostID)
	if err != nil {
		return errors.New("post id is not valid")
	}
	fmt.Println(postUID)
	err = GetResourceService(s.resRepo).ResourceDelete(model.ResourceDelete{
		ResourceID:   postUID,
		ResourceType: enums.POST_RESOURCE_TWEET,
		UserID:       acn.UserID,
	})
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return s.repo.DeletePost(model.PostDeleteModel{
		ID:     postUID,
		UserID: acn.UserID,
	})
}

func (s PostService) AddLikePost(acn action.PostUser) error {
	postUID, err := uuid.Parse(acn.PostID)
	if err != nil {
		return errors.New("post id is not valid")
	}
	return s.repo.AddPostLike(model.PostLike{
		PostID: postUID,
		UserID: acn.UserID,
	})
}

func (s PostService) DelPostLike(acn action.PostUser) error {
	postUID, err := uuid.Parse(acn.PostID)
	if err != nil {
		return errors.New("post id is not valid")
	}
	return s.repo.DelPostLike(model.PostLike{
		PostID: postUID,
		UserID: acn.UserID,
	})
}

func (s PostService) GetPosts(m action.PostPagination) ([]model.PostListModel, error) {
	if m.Page == 0 {
		m.Page = 1
	}
	if m.Limit == 0 {
		m.Limit = 30
	}
	return s.repo.GetPosts(model.PostPaginationModel{
		Offset: (m.Page - 1) * m.Limit,
		Limit:  m.Limit,
		Search: m.Search,
		UserID: m.UserID,
	})
}

func (s PostService) GetFollowerPosts(m action.PostPagination) ([]model.PostListModel, error) {
	if m.Page == 0 {
		m.Page = 1
	}
	if m.Limit == 0 {
		m.Limit = 30
	}
	ids, err := s.usrRepo.GetUserFolowerID(m.UserID)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return []model.PostListModel{}, nil
	}
	qry := utils.GetQueryUUID(ids)

	return s.repo.GetPosts(model.PostPaginationModel{
		Offset:      (m.Page - 1) * m.Limit,
		Limit:       m.Limit,
		Search:      m.Search,
		UserFolowID: qry,
		UserID:      m.UserID,
	})
}
