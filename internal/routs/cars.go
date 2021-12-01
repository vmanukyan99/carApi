package routes

import (
	"carApi/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func CarsRoute(route fiber.Router) {
	route.Get("/", controllers.GetAllCars)
	route.Get("/:id", controllers.GetCar)
	route.Post("/", controllers.AddCar)
	route.Put("/:id", controllers.UpdateCar)
	route.Delete("/:id", controllers.DeleteCar)
}
