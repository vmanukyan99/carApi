package controllers

import (
	"carApi/internal/config"
	"carApi/internal/models"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"strconv"
	"time"
)

func GetAllCars(c *fiber.Ctx) error {
	carsCollection := config.MI.DB.Collection("cars")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var cars []models.Car

	filter := bson.M{}
	findOptions := options.Find()

	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"brand": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
				{
					"model": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
				{
					"status": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limitVal, _ := strconv.Atoi(c.Query("limit", "10"))
	var limit = int64(limitVal)

	total, _ := carsCollection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page) - 1) * limit)
	findOptions.SetLimit(limit)

	cursor, err := carsCollection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Cars Not found",
			"error":   err,
		})
	}

	for cursor.Next(ctx) {
		var car models.Car
		err := cursor.Decode(&car)
		if err != nil {
			return err
		}
		cars = append(cars, car)
	}

	last := math.Ceil(float64(total / limit))
	if last < 1 && total > 0 {
		last = 1
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":      cars,
		"total":     total,
		"page":      page,
		"last_page": last,
		"limit":     limit,
	})
}
