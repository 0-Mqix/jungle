package register

import (
	"fmt"
	"os"
	"strings"
)

func OnlyExtractJungleFileAndExit(config Config) {
	if !IsJungleBuild() {
		return
	}

	fmt.Println("[JUNGLE]")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println(" Struct Names:")

	fmt.Println("\n Comment Methods:")
	methods, _ := GetMethods(&config)
	var last string

	for _, m := range methods {
		name := fmt.Sprintf("%s.%s", m.Pkg, m.Struct)

		if last != "" && m.Pkg+m.Struct != last {
			fmt.Println()
		}

		last = name

		fmt.Printf("  %s.%s.%s [?] \n", m.Pkg, m.Struct, m.Name)
	}

	fmt.Println(strings.Repeat("-", 50))
	os.Exit(0)
}
