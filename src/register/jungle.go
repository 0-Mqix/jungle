package register

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/0-Mqix/jungle/src/comment"
)

type Jungle struct {
	*Config

	exitAfterRead bool
	mode          string

	methods []comment.Method

	modeArg            *string
	routeTargetArg     *string
	interfaceTargetArg *string
}

func (j *Jungle) afterParse() {
	mode := strings.ToUpper(*j.modeArg)
	target := *j.routeTargetArg

	if target != "?" {
		j.routeTarget = target
	}

	switch mode {
	case "USE", "IMPORT", "?", "":
		break

	case "BUILD", "EXPORT":
		j.exitAfterRead = true

	default:
		fmt.Println("[JUNGLE] INVALID MODE")
		os.Exit(1)
	}

	j.mode = mode

	j.ReadProject()
}

func Init(config Config, parseFlags bool, beforeParseFlags ...func()) *Jungle {
	j := &Jungle{Config: &config}

	j.modeArg = flag.String("jungle-mode", "", "")
	j.routeTargetArg = flag.String("jungle-target", "", "")
	j.interfaceTargetArg = flag.String("jungle-interface-target", "", "")

	for _, function := range beforeParseFlags {
		function()
	}

	if parseFlags {
		flag.Parse()
		j.afterParse()
	}

	return j
}
