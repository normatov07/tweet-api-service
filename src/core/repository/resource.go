package repository

import (
	"github.com/google/uuid"
	"github.com/normatov07/mini-tweet/core/model"
)

type ResourceRepo interface {
	ResourceCreate(m model.ResourceModel) error
	GetResource(m model.ResourceDelete) (*model.ResourceGet, error)
	DeleteResource(id uuid.UUID) (err error)
}
