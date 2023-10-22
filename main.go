package main

import (
	"log"
	"log/slog"

	"github.com/gofiber/fiber/v2"

)

func main() {
	appConfig := fiber.Config{
		EnablePrintRoutes: true,
		Immutable:         true,
	}
	app := fiber.New(appConfig)

	app.Get("/books", getAllBooks).Name("Get books")

	slog.Info("starting app on port :3000")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
