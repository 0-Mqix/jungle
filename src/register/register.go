package register

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/0-Mqix/jungle/src/comment"
	"github.com/gofiber/fiber/v2"
)

var (
	RouteType = reflect.TypeOf(Route{})
)

func JungleRoutes(app *fiber.App, structs ...interface{}) {
	values := make(map[string]reflect.Value)

	for _, s := range structs {
		v := reflect.ValueOf(reflect.ValueOf(s).Interface())
		values[reflect.TypeOf(s).Elem().Name()] = v
	}

	for _, m := range comment.GetJungleMethods() {
		method := values[m.Type].MethodByName(m.Name)
		t := method.Type()

		if m.Annotation != "register" ||
			t.NumOut() != 1 ||
			t.NumIn() != 0 ||
			t.Out(0) != RouteType {
			continue
		}

		returns := method.Call([]reflect.Value{})

		for _, r := range returns {
			route := r.Interface().(Route)
			FiberRegisterRoute(app, route, m)
		}
	}
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
