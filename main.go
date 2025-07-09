package main

import (
	"kubernetes-test/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctrl := controller.NewControllerImpl()

	app := fiber.New()
	app.Get("/v1/hello", ctrl.SayHello)

	app.Listen(":8080")
}
