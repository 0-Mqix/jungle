package register

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/0-Mqix/jungle/src/comment"
)

type jungleFile struct {
	Target *os.File
	Config *Config
}

func StartJungleFile(config *Config) *jungleFile {
	f, err := os.OpenFile(config.ExportTarget, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("start:", err)
	}

	return &jungleFile{Target: f, Config: config}
}

func (e *jungleFile) Clear() {
	err := e.Target.Truncate(0)

	if err != nil {
		fmt.Println("clear: ", err)
	}
}

func (e *jungleFile) Import() (pairs []comment.Method) {
	content, err := os.ReadFile(e.Target.Name())

	if err != nil {
		fmt.Println("import read:", err)
	}

	for _, line := range bytes.Split(content, []byte{'\n'}) {
		method := comment.Method{}
		err := json.Unmarshal(line, &method)

		if err != nil {
			fmt.Println("import unmarshal:", err)
			continue
		}

		pairs = append(pairs, method)
	}

	return
}

func (e *jungleFile) Add(method comment.Method) {
	content, err := json.Marshal(method)

	if err != nil {
		fmt.Println("add marshal:", err)
	}

	e.Target.Write(append(content, '\n'))

	if err != nil {
		fmt.Println("add write:", err)
	}
}
