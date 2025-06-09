package cli

import (
	"fmt"
	"vapkg/internal/core"
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

func emptyCommandHandleFn(ctx *core.Context, opts map[string]string) error {

	var buf string
	for k := range opts {
		switch k {
		case "version":
			buf = fmt.Sprintf("%s version: %s", ctx.Core().Name(), ctx.Core().Version())
		case "author":
			buf = fmt.Sprintf("%s author: %s", ctx.Core().Name(), ctx.Core().Author())
		case "source":
			buf = fmt.Sprintf("%s source: %s", ctx.Core().Name(), ctx.Core().URL())
		case "license":
			buf = fmt.Sprintf("%s license: %s", ctx.Core().Name(), ctx.Core().License())
		}
		break
	}

	if buf != "" {
		fmt.Printf("%s\n", buf)
	}

	return nil
}
