package main

import (
	"github.com/gabszero/expenses-api/pkg/application/routes"
	"github.com/gabszero/expenses-api/pkg/infrastructure/repositories"
)

func main() {

	repository := repositories.Repository{}
	repository.StartDabase()

	mainRouter := routes.Router{}
	mainRouter.StartRoutes()
}
