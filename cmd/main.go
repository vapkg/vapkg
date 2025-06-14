package main

import (
	"os"
	"vapkg/cmd/cli"
	cfg "vapkg/internal/config"
	"vapkg/internal/core"
	"vapkg/internal/logger"
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

	var ctx *core.Context
	if ctx = getContext(cfg.Get()); ctx == nil {
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

func getContext(cfg core.IConfig) (ctx *core.Context) {
	if log := getLogger(cfg); log != nil {

		if ctx = core.NewContext(log, cfg); ctx != nil {

			for k, v := range cli.Commands() {
				ctx.Commands().Register(k, v)
			}

			log.Infof("Context initialized with %d command(s)", len(cli.Commands()))
		}

	}

	return
}

func getLogger(cfg core.IConfig) core.ILogger {
	if log, err := logger.NewActualFromConfig(cfg); err == nil {
		log.Infof("Logger initialized")
		return log
	}

	return nil
}
