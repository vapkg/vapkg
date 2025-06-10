package cli

import (
	"fmt"
	"vapkg/internal/core"
)

var initCommand = core.Command{
	Usage:       "vapkg init [<name>] [--<option>[, ...]]",
	Description: "",
	Handler:     initCommandHandleFn,
	Options: map[string]bool{
		"": false,
	},
}

func initCommandHandleFn(ctx *core.Context, opts map[string]string) (err error) {
	//ctx.Logger().Debugf("command '%s' called", "init")

	if ctx.Ws().Exists() {
		return fmt.Errorf("{FRD}project already exists{R}")
	}

	name, ok := opts[""]

	if !ok {
		name = ""
	}

	if err = ctx.Ws().Init(name); err != nil {
		err = fmt.Errorf("{FRD}init project fail {R}(%v)", err)
	}

	return err
}
