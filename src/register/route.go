package register

import "github.com/gofiber/fiber/v2"

type Handler func(*fiber.Ctx) error

type Route struct {
	Path       string
	Method     string
	Middleware []Handler
	Handler    Handler
}
