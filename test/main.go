package main

import (
	"github.com/0-Mqix/jungle/src/register"
	"github.com/gofiber/fiber/v2"
)

type Hey struct{}
type Msg string

// @jungle:register
func (h *Hey) Test() register.Route {
	return register.Route{
		Method:  "GET",
		Path:    "/hey/test.:name",
		Handler: h.test,
	}
}

func (h *Hey) test(c *fiber.Ctx) error {
	return c.SendString("hey thanks for testing, " + c.Params("name", "nameless person"))
}

// @jungle:register
func (h *Msg) Test() register.Route {
	return register.Route{
		Method:  "GET",
		Path:    "/msg/test",
		Handler: h.test,
	}
}

func (h *Msg) test(c *fiber.Ctx) error {
	return c.SendString(string(*h))
}

func main() {
	app := fiber.New()

	hey := Hey{}
	msg := Msg("This is Magic")

	register.JungleRoutes(app, &msg, &hey)

	app.Listen(":3000")
}
