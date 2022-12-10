package comment

import "fmt"

type Config struct {
	Directory    string
	Debug        bool
	Export       bool
	ExportTarget string
}

type Method struct {
	Pkg        string `json:"package"`
	Annotation string `json:"annotation"`
	Struct     string `json:"struct"`
	Name       string `json:"name"`
}

func (m Method) GetPrintPrefix() string {
	return fmt.Sprintf("@jungle:%s\n%s.%s()", m.Annotation, m.Struct, m.Name)
}
