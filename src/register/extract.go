package register

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func OnlyExtractJungleFileAndExit(config Config, app *fiber.App, structs ...interface{}) {
	if !IsJungleBuild() {
		return
	}

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
	}

	fmt.Println(strings.Repeat("-", 50))
	os.Exit(0)
}
