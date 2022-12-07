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
type Method struct{ Comment, Type, Name, PackageName, From string }

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

func readPackage(dir string) (methods []Method) {
	fset := token.NewFileSet()
	pkg, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)

	if err != nil {
		fmt.Println(err)
	}

	for name, f := range pkg {
		p := doc.New(f, "./", 2)

		for _, t := range p.Types {
			methods = append(methods, ToMethods(name, dir, t)...)
		}
	}

	return
}

func ToMethods(name, dir string, t *doc.Type) (methods []Method) {
	for _, m := range t.Methods {
		methods = append(methods, Method{
			PackageName: name,
			From:        dir,
			Type:        t.Name,
			Comment:     ToJungeLine(m.Doc),
			Name:        m.Name,
		})
	}

	return
}

func GetJungleMethods() (pair []Method) {
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
