package register

import (
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
	f, err := os.Open(config.ExportTarget)

	if err != nil {
		fmt.Println(err)
	}

	return &jungleFile{Target: f, Config: config}
}

func (e *jungleFile) Clear() {

}

func (e *jungleFile) Import() (pairs []comment.Method) {
	return
}

func (e *jungleFile) Add(method comment.Method) {
	json.Marshal(method)

	// e.Target.Write()
}
