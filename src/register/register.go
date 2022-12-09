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

	fmt.Println("[JUNGLE]")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println(" Struct Names:")

	for _, s := range structs {
		v := reflect.ValueOf(reflect.ValueOf(s).Interface())

		e := reflect.TypeOf(s).Elem()
		pkg := e.PkgPath()

		if pkg != "main" {
			index := strings.LastIndex(e.PkgPath(), "/") + 1
			pkg = e.PkgPath()[index:]

		}
		name := fmt.Sprintf("%s.%s", pkg, e.Name())
		fmt.Println(" ", name)
		values[name] = v
	}

	fmt.Println("\n Comment Methods:")

	var last string

	for _, m := range comment.GetJungleMethods() {
		name := fmt.Sprintf("%s.%s", m.Pkg, m.Struct)
		method := values[name].MethodByName(m.Name)
		t := method.Type()

		if last != "" && m.Pkg+m.Struct != last {
			fmt.Println()
		}

		last = m.Pkg + m.Struct

		fmt.Printf("  %s.%s.%s", m.Pkg, m.Struct, m.Name)

		if m.Annotation != "register" ||
			t.NumOut() != 1 ||
			t.NumIn() != 0 ||
			t.Out(0) != RouteType {
			fmt.Printf(" [x] \n")
			continue
		}

		fmt.Printf(" [✓] \n")

		returns := method.Call([]reflect.Value{})

		for _, r := range returns {
			route := r.Interface().(Route)
			FiberRegisterRoute(app, route, m)
		}
	}

	fmt.Println(strings.Repeat("-", 50))
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
