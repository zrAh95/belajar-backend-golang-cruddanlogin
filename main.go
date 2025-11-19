package main

import (
	"belajargolang/backend-api/config"
	"belajargolang/backend-api/database"
	"belajargolang/backend-api/routes"
)

func main() {

	//load config .env
	config.LoadEnv()

	//inisialisasi database
	database.InitDB()

	//setup router
	r := routes.SetupRouter()

	//mulai server
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
