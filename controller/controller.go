package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	SayHello(c *fiber.Ctx) error
}
