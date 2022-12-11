package main

import (
	"flag"

	"github.com/0-Mqix/jungle/src/register"
	"github.com/0-Mqix/jungle/src/utils/print"
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
	register.UseJungeArgs()
	flag.Parse()

	app := fiber.New()
	hey := Hey{}
	msg := msg.Msg("This is Magic")

	register.JungleRoutes(
		register.Config{
			Directories: []string{"./"},
			Debug:       true,
		},
		app, &msg, &hey)

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("test ?")
	})

	print.Jungle()

	// fmt.Println(app.Listen(":3000"))
}
