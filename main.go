package main

import (
	"github.com/gabszero/expenses-api/pkg/application/routes"
	"github.com/gabszero/expenses-api/pkg/infrastructure/repositories"
	"github.com/gabszero/expenses-api/pkg/infrastructure/seeders"
)

func main() {

	repository := repositories.Repository{}
	repository.StartDabase()

	mainRouter := routes.Router{}
	seeders.Populate()
	mainRouter.StartRoutes()

}
