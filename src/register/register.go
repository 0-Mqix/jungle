package register

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var (
	RouteType = reflect.TypeOf(Route{})
)

type Config struct {
	Directories []string
	Debug       bool
	target      string
}

func JungleRoutes(config Config, app *fiber.App, structs ...interface{}) {

	fmt.Println("[JUNGLE]")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println(" Struct Names:")

	values := ReadStructs(structs)

	fmt.Println("\n Comment Methods:")
	methods, _ := GetMethods(&config)
	var last string

	for _, m := range methods {
		name := fmt.Sprintf("%s.%s", m.Pkg, m.Struct)
		method := values[name].MethodByName(m.Name)
		t := method.Type()

		if last != "" && m.Pkg+m.Struct != last {
			fmt.Println()
		}

		last = name

		fmt.Printf("  %s.%s.%s", m.Pkg, m.Struct, m.Name)

		if m.Annotation != "register" ||
			t.NumOut() != 1 ||
			t.NumIn() != 0 ||
			t.Out(0) != RouteType {
			fmt.Printf(" [x] \n")
			continue
		}

		fmt.Printf(" [✓] \n")

		for _, r := range method.Call([]reflect.Value{}) {
			route := r.Interface().(Route)
			FiberRegisterRoute(app, route, m)
		}
	}

	fmt.Println(strings.Repeat("-", 50))
}

func ReadStructs(structs []interface{}) map[string]reflect.Value {
	values := make(map[string]reflect.Value)

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

	return values
}
