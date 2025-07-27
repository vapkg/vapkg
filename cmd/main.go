package main

import (
	"fmt"
	"os"
	"vapkg/cmd/cli"
	"vapkg/config"
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

	var cliInstance *cli.Cli
	if cliInstance = cli.Parse(args); cliInstance == nil || cliInstance.Empty() {
		utils.VaPrintln("{FRD}try use {R}vapkg help {FRD}command{R}")
		return
	}

	var ctx *core.Context = nil
	if ctx, err = getContext(config.New()); err != nil {
		utils.VaPrintf("{FRD}ctx init err {R}(%v)\n", err)
		return
	}

	defer ctx.Close()

	if !ctx.Commands().Exists(cliInstance.Command()) {
		utils.VaPrintf("{FRD}command {R}'%s' {FRD}not found{R}\n", cliInstance.Command())
		return
	}

	if err = ctx.Commands().Get(cliInstance.Command()).Execute(ctx, cliInstance.Options()); err != nil {
		_, _ = utils.VaPrintfWithPrefix("%s\n", err)
		return
	}
}

func getContext(cfg core.IConfig) (*core.Context, error) {
	if log := getLogger(cfg); log != nil {

		switch ctx, err := core.NewContext(log, cfg); {
		case err != nil:
			return nil, err
		default:
			for k, v := range cli.Commands() {
				ctx.Commands().Register(k, v)
			}

			for k, v := range ProviderMap {
				ctx.Providers().Register(k, v)
			}

			log.Infof("Context initialized with %d command(s)", len(cli.Commands()))

			return ctx, nil
		}
	}

	return nil, fmt.Errorf("no context found")
}

func getLogger(cfg core.IConfig) core.ILogger {
	if log, err := logger.NewActualFromConfig(cfg); err == nil {
		log.Infof("Logger initialized")
		return log
	}

	return nil
}
