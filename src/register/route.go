package register

import "github.com/gofiber/fiber/v2"

type Route struct {
	Path       string
	Method     string
	Middleware []func(*fiber.Ctx) error
	Handler    func(*fiber.Ctx) error
}
