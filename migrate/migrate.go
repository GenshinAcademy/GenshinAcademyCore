package main

import (
	"fmt"
	"log"

	"genshinacademycore/config"
	"genshinacademycore/models"
)

var DB config.Database

func init() {
	env, err := config.LoadENV()
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	DB = config.InitDB(&env)
}

func main() {
	DB.ORM.AutoMigrate(&models.Character{})
	fmt.Println("Migration complete")
}
