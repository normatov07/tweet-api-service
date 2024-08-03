package token

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/normatov07/mini-tweet/core/app_errors"
	"github.com/normatov07/mini-tweet/core/model"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
	duration     int64
}

func NewPasetoMaker() (*PasetoMaker, error) {
	symmetricKey := os.Getenv("PASETO_SYMETRIC_KEY")
	duration, err := strconv.ParseInt(os.Getenv("TOKEN_EXPIRY"), 10, 64)
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, app_errors.NewAppErr(-500, fmt.Sprintf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize))
	}
	if err != nil {
		return nil, app_errors.NewAppErr(-500, "error on parsing token expiry")
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
		duration:     duration,
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(user model.UserModel) (string, error) {
	payload := NewPayload(user, time.Duration(maker.duration))

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, errors.New("toke is invalid")
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return nil
	}

	return errors.New("token has expired")
}
