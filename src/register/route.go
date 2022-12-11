package register

import (
	"fmt"
	"strings"

	"github.com/0-Mqix/jungle/src/comment"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Path       string
	Method     string
	Middleware []func(*fiber.Ctx) error
	Handler    func(*fiber.Ctx) error
}

func Create(method, path string, handler func(*fiber.Ctx) error, middleware ...func(*fiber.Ctx) error) Route {
	return Route{Path: path, Method: method, Middleware: middleware}
}

func FiberRegisterRoute(app *fiber.App, route Route, source comment.Method) {
	handlers := append(route.Middleware, route.Handler)
	method := strings.ToUpper(route.Method)
	path := route.Path

	switch method {
	case "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH":
		app.Add(method, path, handlers...)
	case "ALL":
		app.All(path, handlers...)
		return
	case "":
		fmt.Printf("%s has no http method\n", source.GetPrintPrefix())
	default:
		fmt.Printf("%s %s has an invalid http method\n", source.GetPrintPrefix(), method)
	}
}
