package register

import "github.com/gofiber/fiber/v2"

type Route struct {
	Path       string
	Method     string
	Middleware []func(*fiber.Ctx) error
	Handler    func(*fiber.Ctx) error
}

func Create(method, path string, handler func(*fiber.Ctx) error, middleware ...func(*fiber.Ctx) error) Route {
	return Route{Path: path, Method: method, Middleware: middleware}
}
