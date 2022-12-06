package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/0-Mqix/jungle/src/register"
)

type Hey struct {
}

// @jungle:register
func (Hey) Test() register.Route {
	fmt.Println("test")

	return register.Route{}
}

func main() {
	//Create a new token set for the file
	fset := token.NewFileSet()
	//Parse the file and create an AST
	file, err := parser.ParseFile(fset, "./max/max copy.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Inspect(file, func(n ast.Node) bool {
		// Find Return Statements
		funcCall, ok := n.(*ast.CallExpr)
		if ok {
			mthd, ok := funcCall.Fun.(*ast.SelectorExpr)
			if ok {
				_, ok = mthd.X.(*ast.Ident)
				if ok {
					switch v := funcCall.Args[0].(type) {
					case *ast.BasicLit:
						fmt.Printf("Topic: %s \n", v.Value)
					case *ast.Ident:
						fmt.Printf("Topic is declared with: %s \n", v.Name)
					default:
						fmt.Println("Unrecognized type")
					}
				}
			}
		}
		return true
	})
}
