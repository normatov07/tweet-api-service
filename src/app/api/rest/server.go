package rest

import (
	"log"
	"net/http"
	"os"

	"github.com/normatov07/mini-tweet/app"
)

type Server struct {
	app *app.ApplicationContext
}

func GetServer() *Server {
	return &Server{
		app: app.GetApplicationContext(),
	}
}

func (s *Server) RunHTTP() {
	addr := os.Getenv("HTTP_ADDRESS")
	srv := &http.Server{
		Addr:    addr,
		Handler: s.Routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
