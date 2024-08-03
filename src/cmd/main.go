package main

import (
	"github.com/joho/godotenv"
	"github.com/normatov07/mini-tweet/app/api/rest"
	"github.com/normatov07/mini-tweet/common/utils"
	"github.com/normatov07/mini-tweet/db/postgres"
)

func init() {
	err := godotenv.Load()
	utils.LoadLogs()
	if err != nil {
		panic(err)
	}

	utils.SetMode()
}

func main() {
	postgres.InitConn()
	defer postgres.Close()

	app := rest.GetServer()

	app.RunHTTP()

	defer utils.LogFile.Close()
}
