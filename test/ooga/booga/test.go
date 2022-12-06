package booga

import (
	"fmt"

	"github.com/0-Mqix/jungle/src/register"
)

// @jungle:register
func ThisIsCrazy() register.Route {
	fmt.Println("hello, crazy ass monkey")

	return register.Route{}
}
