package model

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type NullString sql.NullString

type UserModel struct {
	ID        uuid.UUID  `json:"id"`
	Login     string     `json:"login"`
	Password  string     `json:"-"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Address   NullString `json:"address"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (x *NullString) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.String)
}

func (ni *NullString) Scan(value interface{}) error {
	var i sql.NullString
	if err := i.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*ni = NullString{i.String, false}
	} else {
		*ni = NullString{i.String, true}
	}
	return nil
}
