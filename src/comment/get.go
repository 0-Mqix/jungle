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

func ToMethods(pkg, dir string, t *doc.Type) (methods []Method) {

	for _, m := range t.Methods {
		name := fmt.Sprintf("%s.%s.%s", pkg, t.Name, m.Name)
		annotation := ToJungeAnnotation(name, m.Doc, ":")

		if annotation == "" {
			continue
		}

		methods = append(methods, Method{
			Pkg:        pkg,
			Struct:     t.Name,
			Annotation: annotation,
			Name:       m.Name,
		})
	}

	return
}

func GetJungleMethods(directory string, print bool) (pair []Method) {
	find := Find{"./"}
	find.packages(directory)

	if print {
		fmt.Println("  Directories:")
		for _, dir := range find {
			fmt.Println("  ", dir)
		}
		fmt.Println()
	}

	for _, dir := range find {
		pair = append(pair, readPackage(dir)...)
	}

	return
}

func ToJungeAnnotation(name, doc, spliter string) string {
	for _, line := range strings.Split(doc, "\n") {
		index := strings.Index(line, "@")

		if index != 0 {
			return ""
		}

		index = strings.Index(line, "@jungle"+spliter)

		if index == 0 {
			return strings.Replace(line, "@jungle"+spliter, "", 1)
		}

		fmt.Printf("  %s [x]\n   ANNOTATION UNKNOWN\n   %s\n\n", name, line[1:])
	}

	return ""
}
