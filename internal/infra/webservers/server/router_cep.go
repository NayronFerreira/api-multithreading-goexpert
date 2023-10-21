package server

import (
	"net/http"

	"github.com/NayronFerreira/api-multithreading-goexpert/config"
	"github.com/NayronFerreira/api-multithreading-goexpert/internal/infra/webservers/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitAddressServer(appConfig config.Config, addrHandle handlers.AddressHandle) {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/myaddress", func(router chi.Router) {
		router.Get("/{cep}", addrHandle.GetAddressHandle)
	})

	http.ListenAndServe(appConfig.WebServerPort, router)
}
