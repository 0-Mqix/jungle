package main

import (
	"fmt"
	"reflect"

	"github.com/0-Mqix/jungle/src/comment"
)

type Hey struct{}

type Hello string

// @jungle:register
func (h *Hey) Monkey(name string) {
	fmt.Println("hey, Monkey", name)
}

// @jungle:register
func (h *Hey) CallMe(name string) {
	fmt.Println("hey, i got called " + name)
}

// @jungle:register
func (h *Hello) Monkey(name string) {
	fmt.Println(string(*h) + name)
}

// todo make it safe
func InvokeJungleMethods(structs ...interface{}) {
	values := make(map[string]reflect.Value)

	for _, s := range structs {
		v := reflect.ValueOf(reflect.ValueOf(s).Interface())
		values[reflect.TypeOf(s).Elem().Name()] = v
	}

	for _, m := range comment.GetJungleMethods() {
		values[m.Type].MethodByName(m.Name).Call([]reflect.Value{reflect.ValueOf("you ;)")})
	}
}

func main() {
	var hey Hey
	hello := Hello("this hello is for -> ")

	InvokeJungleMethods(&hey, &hello)
}
