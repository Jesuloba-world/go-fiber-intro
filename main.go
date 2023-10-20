package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Item struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var items []Item = []Item{
	{1, "Item 1", 10.9},
	{2, "Item 2", 11.9},
	{3, "Item 3", 20.4},
}

// returns all items in a list
func getAllItems(c *fiber.Ctx) error {

	return c.JSON(items)
}

// returns a single item
func getItemById(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID")
	}

	for _, item := range items {
		if item.ID == id {
			return c.JSON(item)
		}
	}

	return c.Status(http.StatusNotFound).SendString("Item not found")
}

func createItem(c *fiber.Ctx) error {
	var item Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request")
	}

	item.ID = len(items) + 1
	items = append(items, item)
	return c.JSON(item)
}

func main() {
	app := fiber.New()

	app.Get("/items", getAllItems)
	app.Get("/items/:id", getItemById)
	app.Post("/items", createItem)

	log.Print("starting app on port :3000")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
