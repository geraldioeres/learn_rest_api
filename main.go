package main

import (
	"learn_api/configs"
	"learn_api/routes"
)

func main() {
	configs.InitDB()
	e := routes.NewRoutes()
	e.Start(":8000")
}
