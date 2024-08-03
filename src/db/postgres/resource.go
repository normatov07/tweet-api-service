package postgres

import (
	"github.com/google/uuid"
	"github.com/normatov07/mini-tweet/core/model"
)

type ResourceRepo struct{}

func (r ResourceRepo) ResourceCreate(m model.ResourceModel) (err error) {
	_, err = conn.Exec("INSERT INTO resources (id,resource_id,resource_type,user_id,size,name,path,format,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", m.ID, m.ResourceID, m.ResourceType, m.UserID, m.Size, m.Name, m.Path, m.Format, m.CreatedAt, m.UpdatedAt)
	return
}

func (r ResourceRepo) DeleteResource(id uuid.UUID) (err error) {
	_, err = conn.Exec("DELETE FROM resources WHERE id=$1", id)
	return
}

func (r ResourceRepo) GetResource(m model.ResourceDelete) (*model.ResourceGet, error) {
	var res model.ResourceGet
	err := conn.QueryRow("SELECT id,path FROM resources WHERE resource_id=$1 AND resource_type=$2 AND user_id=$3", m.ResourceID, m.ResourceType, m.UserID).Scan(&res.ID, &res.Path)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
