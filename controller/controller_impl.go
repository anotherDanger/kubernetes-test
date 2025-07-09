package controller

import "github.com/gofiber/fiber/v2"

type ControllerImpl struct{}

func NewControllerImpl() Controller {
	return &ControllerImpl{}
}

func (ctrl *ControllerImpl) SayHello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
