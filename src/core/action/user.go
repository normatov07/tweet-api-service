package action

import "github.com/normatov07/mini-tweet/core/model"

type UserStore struct {
	Login     string           `form:"login" binding:"required"`
	Password  string           `form:"password" binding:"required"`
	FirstName string           `form:"first_name" binding:"required"`
	LastName  string           `form:"last_name"  binding:"required"`
	Address   model.NullString `form:"address"`
}
