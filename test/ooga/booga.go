package ooga

import (
	"fmt"

	"github.com/0-Mqix/jungle/src/register"
)

// @jungle:register
func Ooga() register.Route {
	fmt.Println("ooga booga")

	return register.Route{}
}
