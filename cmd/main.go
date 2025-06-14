package main

import (
	"os"
	"vapkg/cmd/cli"
	"vapkg/internal/core"
	"vapkg/internal/utils"
)

func main() {

	var err error
	var args []string

	utils.InitVaPalette()

	if args = os.Args[1:]; len(args) < 1 {
		utils.VaPrintln("{FRD}no args found{R}")
		return
	}

	var cliInstance cli.Cli
	if cliInstance = cli.Parse(args); cliInstance.Empty() {
		utils.VaPrintln("{FRD}try use {R}vapkg help {FRD}command{R}")
		return
	}

	var config *core.Config
	if config, err = core.GetConfig(); err != nil {
		utils.VaPrintf("{FRD}config initialization failed (%s){R}\n", err)
		return
	}

	var ctx *core.Context
	if ctx = getContext(config); ctx == nil {
		utils.VaPrintf("{FRD}ctx init err {R}(%s)\n", err)
		return
	}
	defer ctx.Close()

	if !ctx.Commands().Exists(cliInstance.Command()) {
		utils.VaPrintf("{FRD}command {R}'%s' {FRD}not found{R}\n", cliInstance.Command())
		return
	}

	if err = ctx.Commands().Get(cliInstance.Command()).Execute(ctx, cliInstance.Options()); err != nil {
		utils.VaPrintf("%s\n", err)
		return
	}
}

func getContext(cfg *core.Config) *core.Context {
	if pwd, err := os.Getwd(); err == nil {

		ctx := core.NewContext(pwd, cfg)

		for k, v := range cli.Commands() {
			ctx.Commands().Register(k, v)
		}
	}

	return nil
}
