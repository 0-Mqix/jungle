package register

import (
	"strings"

	"github.com/0-Mqix/jungle/src/comment"
)

var ()

func (j *Jungle) IsBuild() bool {
	mode := strings.ToUpper(*j.modeArg)

	switch mode {
	case "BUILD", "EXPORT":
		return true
	default:
		return false
	}
}

func UseArgs() {
}

func (j *Jungle) GetMethods() []comment.Method {
	methods := make([]comment.Method, 0)

	for _, directory := range j.Directories {
		methods = append(methods, comment.GetJungleMethods(directory, j.Debug)...)
	}

	if j.mode != "" {
		file := StartJungleFile(j.Config)

		switch j.mode {
		case "BUILD", "EXPORT":

			for _, m := range methods {
				file.Add(m)
			}

		case "USE", "IMPORT":
			methods = file.Import()
		}

	}

	j.methods = methods
	return methods
}
