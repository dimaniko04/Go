package main

import (
	"fmt"
	"log"

	"github.com/dimaniko04/Go/lb-4/server/config"
	"github.com/dimaniko04/Go/lb-4/server/controllers"
	"github.com/dimaniko04/Go/lb-4/server/database"
	"github.com/dimaniko04/Go/lb-4/server/routes"
	"github.com/dimaniko04/Go/lb-4/server/services"
	"github.com/gin-gonic/gin"
)

func main() {
	env := config.NewEnv(".env", true)
	db, err := database.Db(env)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	api := gin.Default()
	api.Static("/static", "./static")
	services := services.GetServices(db, env)
	controllers := controllers.GetControllers(services, env)
	routes.Routes(api, controllers, env)

	if err := api.Run(fmt.Sprintf(":%d", env.ServerPort)); err != nil {
		log.Fatal(err)
	}
}
