package main

import (
	"books_online_api/configs"
	"books_online_api/routes"
)

func main() {
	configs.InitDB()
	e := routes.NewRoute()
	e.Logger.Fatal(e.Start(":8000"))
}
