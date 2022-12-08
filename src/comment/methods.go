package comment

import "fmt"

type Method struct{ Dir, Pkg, Annotation, Type, Name string }

func (m Method) GetPrintPrefix() string {
	return fmt.Sprintf("@jungle:%s\n%s.%s()", m.Annotation, m.Type, m.Name)
}
