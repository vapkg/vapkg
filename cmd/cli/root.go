package cli

import (
	"fmt"
	"vapkg/internal/core"
	"vapkg/internal/utils"
)

var emptyCommand = core.Command{
	Usage:       "vapkg [--<option>[, ...]]",
	Description: "",
	Handler:     emptyCommandHandleFn,
	Options: map[string]bool{
		"author":  false,
		"version": false,
		"source":  false,
		"license": false,
	},
}

func emptyCommandHandleFn(ctx core.IContext, opts map[string]string) (err error) {
	//ctx.Logger().Debugf("root cmd called")

	for k := range opts {
		switch k {
		case "version":
			_, err = utils.VaPrintf("version: %s", ctx.Core().Version())
		case "author":
			_, err = utils.VaPrintf("author: %s", ctx.Core().Author())
		case "source":
			_, err = utils.VaPrintf("source: %s", ctx.Core().URL())
		case "license":
			_, err = utils.VaPrintf("license: %s", ctx.Core().License())
		default:
			err = fmt.Errorf("unknown option %s", k)
			continue
		}

		break
	}

	return err
}
