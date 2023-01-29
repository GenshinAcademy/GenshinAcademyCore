package main

import (
	"genshinacademycore/config"
	"genshinacademycore/controllers"
	"genshinacademycore/logger"
	"genshinacademycore/repository"
	"genshinacademycore/router"
	"genshinacademycore/service"
	"log"
)

func main() {
	env, err := config.LoadENV()
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	db := config.InitDB(&env)

	// init repository
	repoFerret := repository.NewFerretRepository(db)

	// init service
	servicePokemon := service.NewFerretService(repoFerret)

	// init controller
	controllerFerret := controllers.NewFerretController(servicePokemon)

	// setup router
	controller := router.RouterController{
		Ferret: controllerFerret,
	}

	r := router.NewRouter(controller)

	err = r.Run(":" + env.ServerPort)
	if err != nil {
		logger.Log.Error(err)
	}
}
