package register

import (
	"flag"
	"strings"

	"github.com/0-Mqix/jungle/src/comment"
)

var (
	modeArg         *string
	targetArg       *string
	exitAfterExport bool
)

func IsJungleBuild() bool {
	mode := strings.ToUpper(*modeArg)

	switch mode {
	case "BUILD", "EXPORT":
		return true
	default:
		return false
	}
}

func UseJungeArgs() {
	modeArg = flag.String("jungle-mode", "?", "")
	targetArg = flag.String("jungle-target", "?", "")
}

func GetMethods(config *Config) ([]comment.Method, string) {
	mode := strings.ToUpper(*modeArg)
	target := *targetArg

	methods := make([]comment.Method, 0)
	for _, directory := range config.Directories {
		methods = append(methods, comment.GetJungleMethods(directory, config.Debug)...)
	}

	if target != "?" {
		config.target = target
	}

	if mode != "?" {
		file := StartJungleFile(config)

		switch mode {
		case "BUILD", "EXPORT":
			exitAfterExport = true

			for _, m := range methods {
				file.Add(m)
			}
			return methods, mode

		case "USE", "IMPORT":
			methods = file.Import()
			return methods, mode
		}

	}

	return methods, mode
}
