package service

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/normatov07/mini-tweet/common/utils"
	"github.com/normatov07/mini-tweet/core/action"
	"github.com/normatov07/mini-tweet/core/model"
	"github.com/normatov07/mini-tweet/core/repository"
)

type ResourceService struct {
	repo repository.ResourceRepo
}

func GetResourceService(repo repository.ResourceRepo) *ResourceService {
	return &ResourceService{
		repo: repo,
	}
}

func (s ResourceService) ResourceStore(acn action.ResourceStore) error {
	path := os.Getenv("RESOURCE_PATH")
	if acn.IsRemoveDublicate {
		err := s.ResourceDelete(model.ResourceDelete{
			ResourceID:   acn.ResourceID,
			ResourceType: acn.ResourceType,
			UserID:       acn.UserID,
		})
		if err != nil && err != sql.ErrNoRows {
			return err
		}
	}
	if acn.File == nil {
		return nil
	}
	ext, err := utils.GetFileExtension(acn.File.Filename)
	if err != nil {
		return err
	}
	if !utils.PostFileExtensionValidate(ext) {
		return errors.New("unsupproted file format")
	}

	params := model.ResourceModel{
		ID:           acn.ID,
		ResourceID:   acn.ResourceID,
		ResourceType: acn.ResourceType,
		UserID:       acn.UserID,
		Size:         acn.File.Size,
		Name:         acn.File.Filename,
		Path:         fmt.Sprintf("%s/%s.%s", path, acn.ID, ext),
		Format:       ext,
		CreatedAt:    acn.CreatedAt,
		UpdatedAt:    acn.UpdatedAt,
	}

	err = s.SaveDisc(acn.File, fmt.Sprintf("%s/%s.%s", path, acn.ID, ext))
	if err != nil {
		return err
	}

	return s.repo.ResourceCreate(params)
}

func (s ResourceService) ResourceDelete(m model.ResourceDelete) error {
	fInfo, err := s.repo.GetResource(m)
	if err != nil {
		return err
	}
	if err = os.Remove(fInfo.Path); err != nil {
		return err
	}

	return s.repo.DeleteResource(fInfo.ID)
}

func (s ResourceService) SaveDisc(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	return err
}
