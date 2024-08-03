package repository

import "github.com/normatov07/mini-tweet/core/model"

type PostRepo interface {
	CreatePost(model model.PostModel) error
	CreateRepost(m model.CreateRepostModel) error
	DeleteRepost(m model.PostDeleteModel) error
	DeletePost(m model.PostDeleteModel) error
	AddPostLike(m model.PostLike) error
	DelPostLike(m model.PostLike) error
	GetPosts(m model.PostPaginationModel) ([]model.PostListModel, error)
}
