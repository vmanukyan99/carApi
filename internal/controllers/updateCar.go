package controllers

import (
	"carApi/internal/config"
	"carApi/internal/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func UpdateCar(c *fiber.Ctx) error {
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

	objectID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Car not found",
			"error":   err,
		})
	}

	update := bson.M{
		"$set": car,
	}
	_, err = carsCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Car failed to update",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Car updated successfully",
	})
}
