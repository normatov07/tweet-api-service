package app

import "github.com/normatov07/mini-tweet/common/token"

type ApplicationContext struct {
	User *token.Payload
}

func GetApplicationContext() *ApplicationContext {
	return &ApplicationContext{}
}
