package main

import (
	"github.com/deenikarim/gudu"
	"myapp/controllers"
	"myapp/data"
	"myapp/middleware"
)

type application struct {
	App        *gudu.Gudu
	Controller *controllers.Handler
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	i := initApp()
	i.App.ListenAndServe()
}
