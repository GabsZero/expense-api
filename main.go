package main

import (
	"time"

	"github.com/gabszero/expenses-api/pkg/application/routes"
	"github.com/gabszero/expenses-api/pkg/infrastructure/repositories"
	"github.com/gabszero/expenses-api/pkg/infrastructure/seeders"
)

func main() {
	// setando o horário local
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	repository := repositories.Repository{}
	repository.StartDabase()

	mainRouter := routes.Router{}
	seeders.Populate()
	mainRouter.StartRoutes()

}
