package register

import (
	"fmt"
	"os"
)

func (j *Jungle) ReadProject() {
	methods := j.GetMethods()

	var last string

	for _, m := range methods {
		name := fmt.Sprintf("%s.%s", m.Pkg, m.Struct)

		if last != "" && m.Pkg+m.Struct != last {
			fmt.Println()
		}

		last = name

		fmt.Printf("  %s.%s.%s \n", m.Pkg, m.Struct, m.Name)
	}

	if j.IsBuild() {
		os.Exit(0)
	}
}
