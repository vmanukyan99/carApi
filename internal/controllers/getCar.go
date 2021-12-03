package controllers

import (
	"carApi/internal/config"
	"carApi/internal/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func GetCar(c *fiber.Ctx) error {
	carsCollection := config.MI.DB.Collection("cars")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var car models.Car
	objectID, err := primitive.ObjectIDFromHex(c.Params("id"))
	findResult := carsCollection.FindOne(ctx, bson.M{"_id": objectID})
	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Car Not found",
			"error":   err,
		})
	}

	err = findResult.Decode(&car)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Car Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    car,
		"success": true,
	})
}
