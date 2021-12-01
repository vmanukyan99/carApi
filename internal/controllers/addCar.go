package controllers

import (
	"carApi/internal/config"
	"carApi/internal/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func AddCar(c *fiber.Ctx) error {
	carsCollection := config.MI.DB.Collection("cars")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	car := new(models.Car)

	if err := c.BodyParser(car); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	result, err := carsCollection.InsertOne(ctx, car)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Car failed to insert",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "Car inserted successfully",
	})
}
