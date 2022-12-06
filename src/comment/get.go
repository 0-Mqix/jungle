package comment

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

type Find []string
type Pair struct{ PackageName, From, Function, Comment string }

func (f *Find) packages(start string) {
	files, err := os.ReadDir(start)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {

			path := start + file.Name()

			if start != "./" {
				path = start + "/" + file.Name()
			}

			*f = append(*f, path)
			f.packages(path)
		}
	}
}

func readPackage(dir string) (pair []Pair) {
	fset := token.NewFileSet()
	d, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)

	if err != nil {
		fmt.Println(err)
	}

	for name, f := range d {
		p := doc.New(f, "./", 2)

		for _, f := range p.Funcs {

			pair = append(pair, Pair{
				PackageName: name,
				From:        p.ImportPath,
				Comment:     ToJungeLine(f.Doc),
				Function:    f.Name,
			})
		}
	}

	return pair
}

func GetJungleFunctionComments() (pair []Pair) {
	find := Find{"./"}
	find.packages("./")

	for _, dir := range find {
		pair = append(pair, readPackage(dir)...)
	}

	return
}

func ToJungeLine(doc string) string {
	for _, line := range strings.Split(doc, "\n") {
		if strings.Contains(line, "@jungle") {
			return line
		}
	}

	return ""
}
