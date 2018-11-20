package main

import (
	"fmt"
	"github.com/exlskills/demo-go-webservice/config"
	"github.com/exlskills/demo-go-webservice/models"
	"github.com/exlskills/demo-go-webservice/routes"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"time"
)

var Log = config.Cfg().GetLogger()
var CORSHandler = handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), handlers.AllowCredentials(), handlers.AllowedHeaders([]string{"x-locale", "x-api-key", "content-type", "access-control-request-headers", "access-control-request-method", "x-csrftoken"}), handlers.AllowedOrigins(config.Cfg().AllowedOrigins))

func main() {
	Log.Info("Setting up database connection ...")

	for {
		err := models.Setup()
		if err != nil {
			Log.WithError(err).Error("Error setting up database connection, retrying ...")
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}

	Log.Info("Connected to database")

	// TODO: Implement graceful stop
	Log.Info("Starting HTTP server")
	http.ListenAndServe(fmt.Sprintf("%s:%s", config.Cfg().ListenAddress, config.Cfg().ListenPort), CORSHandler(handlers.CombinedLoggingHandler(os.Stdout, routes.CreateRouter())))
	Log.Info("Stopped HTTP server")
}
