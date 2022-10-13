package main

import (
	"restapi/database"
	"restapi/routers"
)

// @title Orders API
// @version 1.0
// @description Simple dummy order API
// @host localhost:8080
// @BasePath /

func main() {

	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}
