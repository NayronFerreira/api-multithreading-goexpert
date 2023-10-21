package main

import (
	"github.com/NayronFerreira/api-multithreading-goexpert/config"
	"github.com/NayronFerreira/api-multithreading-goexpert/internal/infra/database"
	"github.com/NayronFerreira/api-multithreading-goexpert/internal/infra/webservers/handlers"
	"github.com/NayronFerreira/api-multithreading-goexpert/internal/infra/webservers/server"
)

func main() {
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dbBrasil, dbVia := database.InitDBConnections(appConfig)
	addrHandle := handlers.NewAddressHandle(appConfig, dbVia, dbBrasil)
	server.InitAddressServer(appConfig, addrHandle)
}
