package postgres

import (
	"github.com/google/uuid"
	"github.com/normatov07/mini-tweet/core/model"
)

type UserRepo struct{}

func (r UserRepo) GetUserByLogin(login string) (m model.UserModel, err error) {
	err = conn.QueryRow("SELECT id,login,password,first_name,last_name,address,created_at,updated_at FROM users WHERE login=$1", login).Scan(&m.ID, &m.Login, &m.Password, &m.FirstName, &m.LastName, &m.Address, &m.CreatedAt, &m.UpdatedAt)
	return
}

func (r UserRepo) GetUserByID(id uuid.UUID) (m model.UserModel, err error) {
	err = conn.QueryRow("SELECT id,login,password,first_name,last_name,address,created_at,updated_at FROM users WHERE id=$1", id).Scan(&m.ID, &m.Login, &m.Password, &m.FirstName, &m.LastName, &m.Address, &m.CreatedAt, &m.UpdatedAt)
	return
}

func (r UserRepo) CreateUser(model model.UserModel) (err error) {
	_, err = conn.Exec("INSERT INTO users (id,login,password,first_name,last_name,address,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", model.ID, model.Login, model.Password, model.FirstName, model.LastName, model.Address.String, model.CreatedAt, model.UpdatedAt)
	return
}

func (r UserRepo) CreateUserFollower(m model.StoreUserFollower) (err error) {
	_, err = conn.Exec("INSERT INTO user_followers (user_id,follower_id) VALUES ($1,$2) ON CONFLICT (user_id,follower_id) DO NOTHING", m.UserID, m.FollowerID)
	return
}

func (r UserRepo) DeleteUserFollower(m model.StoreUserFollower) (err error) {
	_, err = conn.Exec("DELETE FROM user_followers WHERE user_id=$1 AND follower_id=$2", m.UserID, m.FollowerID)
	return
}

func (r UserRepo) GetUserFolowerID(userId uuid.UUID) ([]uuid.UUID, error) {
	rows, err := conn.Query("SELECT follower_id FROM user_followers WHERE user_id=$1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]uuid.UUID, 0, 40)
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		res = append(res, id)
	}

	return res, nil
}
