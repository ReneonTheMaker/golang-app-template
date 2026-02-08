package main

import (
	"app/internal/web"
	"log"
)

func main() {
	app := web.NewApp()
	log.Fatal(app.FiberApp.Listen(":8080"))
}
