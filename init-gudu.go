package main

import (
	"github.com/deenikarim/gudu"
	"log"
	"myapp/controllers"
	"myapp/data"
	"myapp/middleware"
	"os"
)

// InitApp initializes the project for users using Gudu package
func initApp() *application {
	// get the current user working directory
	currentWorkingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// initialize the Gudu package
	gud := &gudu.Gudu{}
	err = gud.New(currentWorkingDir)
	if err != nil {
		log.Fatal(err)
	}

	// populate the handler struct  type
	myController := controllers.Handler{App: gud}
	myMiddleware := &middleware.Middleware{App: gud}

	// populate application struct with values
	apps := &application{
		App:        gud,
		Controller: &myController,
		Middleware: myMiddleware,
	}
	// adding user defined routes to the default routes
	apps.App.Router = apps.userRouters()

	// initialize the models
	apps.Models = data.New(apps.App.DBConnection.SqlConnPool)
	// populate the models in the handlers type
	myController.Models = apps.Models
	apps.Middleware.Models = apps.Models

	return apps
}
