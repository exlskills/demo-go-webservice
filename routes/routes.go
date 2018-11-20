package routes

import (
	"github.com/exlinc/golang-utils/httpmiddleware"
	"github.com/exlskills/demo-go-webservice/config"
	"github.com/exlskills/demo-go-webservice/controllers/v1"
	"github.com/exlskills/demo-go-webservice/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

var Log = config.Cfg().GetLogger()

func CreateRouter() http.Handler {
	router := mux.NewRouter()
	router.StrictSlash(true)

	// V1 Routes
	v1Router := router.PathPrefix("/v1").Subrouter()
	v1Router.HandleFunc("/", v1.API).Methods("GET")
	v1Router.HandleFunc("/gophers", httpmiddleware.Use(v1.GetGophers, middleware.RequireAPIKey)).Methods("GET")

	return httpmiddleware.Use(router.ServeHTTP, middleware.GetContext, httpmiddleware.RecoverInternalServerError)
}
