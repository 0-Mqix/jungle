package main

import (
	"github.com/0-Mqix/jungle/src/register"
	"github.com/0-Mqix/jungle/test/msg"
	"github.com/gofiber/fiber/v2"
)

type Hey struct{}

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

func main() {
	app := fiber.New()
	hey := Hey{}
	msg := msg.Msg("This is Magic")

	register.JungleRoutes(register.Config{
		Directories:  []string{"./"},
		Export:       true,
		Debug:        true,
		ExportTarget: "./jungle",
	}, app, &msg, &hey)

	app.Listen(":3000")
}
