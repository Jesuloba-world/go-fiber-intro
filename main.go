package main

import (
	"log"
	"log/slog"

	"github.com/gofiber/fiber/v2"

)

// returns all items in a list
func getHandler(c *fiber.Ctx) error {
	n := c.Params("name")
	slog.Info("request received", "name", n)
	return nil
}

func main() {
	appConfig := fiber.Config{
		AppName:           "My Awesome App v0.0.0-beta",
		EnablePrintRoutes: true,
		ServerHeader:      "Awesome Server",
		Immutable: true,
	}
	app := fiber.New(appConfig)

	app.Get("/:name", getHandler).Name("Test")

	slog.Info("starting app on port :3000")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
