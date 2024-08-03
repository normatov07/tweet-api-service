package token

import (
	"time"

	"github.com/google/uuid"
	"github.com/normatov07/mini-tweet/core/model"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Login     string    `json:"username"`
	Address   string    `json:"address"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(user model.UserModel, duration time.Duration) *Payload {
	return &Payload{
		ID:        user.ID,
		Login:     user.Login,
		Address:   user.Address.String,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

}
