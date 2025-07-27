package cli

import (
	"fmt"
	"vapkg/internal/core"
	"vapkg/internal/core/vapkg"
)

var initCommand = core.Command{
	Usage:       "vapkg init [<name>] [--<option>[, ...]]",
	Description: "",
	Handler:     initCommandHandleFn,
	Options: map[string]bool{
		"":     false,
		"type": false,
	},
}

func initCommandHandleFn(ctx core.IContext, opts map[string]string) error {
	ctx.Logger().Debugf("command '%s' called", "init")

	if ctx.Workspace().IsExist() {
		ctx.Logger().Errorf("init call err then workspace already exists")
		return fmt.Errorf("{FRD}workspace already exists")
	}

	if err := ctx.Workspace().CreateWorkspace(getPackageName(opts), getPackageType(opts)); err != nil {
		ctx.Logger().Errorf("init call err then create workspace: %v", err)
		return fmt.Errorf("{FRD}%s", err.Error())
	}

	return nil
}

func getPackageName(opts map[string]string) string {
	if v, ok := opts[""]; ok {
		return v
	}
	return ""
}

func getPackageType(opts map[string]string) vapkg.VaPackageType {
	if v, ok := opts["type"]; ok {
		return vapkg.ParseVapkgType(v)
	}

	return vapkg.GetType()
}
