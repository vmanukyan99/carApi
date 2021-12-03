package controllers

import (
	"carApi/internal/config"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func DeleteCar(c *fiber.Ctx) error {
	carsCollection := config.MI.DB.Collection("cars")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	objectID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Car not found",
			"error":   err,
		})
	}
	_, err = carsCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Car failed to delete",
			"error":   err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Car deleted successfully",
	})
}
