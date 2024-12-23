// @title           IndustrialBack
// @version         1.0
// @description     This is a project for industrial technogies
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"indpr/dataBase"
	"indpr/routes"

	_ "indpr/docs"
	// Для вывода в консоль
	"log"
)

func main() {
	dataBase.InitDB()

	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("ERROR: Не удалось запустить сервер: %v", err)
	}

}
