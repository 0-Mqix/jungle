package msg

import (
	"github.com/0-Mqix/jungle/src/register"
	"github.com/gofiber/fiber/v2"
)

type Msg string

// @junge:register
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
