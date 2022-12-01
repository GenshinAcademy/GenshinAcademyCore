package main

import (
	"fmt"
	"log"

	"genshinacademycore/config"
	models "genshinacademycore/models/db"
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
	DB.ORM.AutoMigrate(&models.Character{}, models.Name{}, models.StatsProfit{})
	fmt.Println("Migration complete")
}
