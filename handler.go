package main

import (
	"github.com/gofiber/fiber/v2"
)

func getAllBooks(c *fiber.Ctx) error {
	person := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		ID:   1,
		Name: "John",
		Age:  30,
	}

	return c.JSON(person)
}
