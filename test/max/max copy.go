package max

import (
	"fmt"

	"github.com/0-Mqix/jungle/src/register"
)

// @jungle:register
func Monkey() register.Route {
	fmt.Println("hello, Monke")

	return register.Route{}
}
