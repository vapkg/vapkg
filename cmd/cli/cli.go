package cli

import (
	"strings"
	"vapkg/internal/core"
)

var commandsMap = map[string]*core.Command{
	"":         &emptyCommand,
	"init":     &initCommand,
	"download": &downloadCommand,
}

func Commands() map[string]*core.Command {
	return commandsMap
}

type Cli struct {
	command string
	options map[string]string
}

func Parse(args []string) Cli {

	var (
		val string
		key string
		idx = 0
	)

	cli := Cli{"", make(map[string]string)}

	for i, arg := range args {

		key = arg
		idx = strings.Index(key, core.OptionPrefix)

		if i == 0 && idx != 0 {
			cli.setCommand(key)
			continue
		}

		if idx != 0 && len(cli.Command()) != 0 {
			val = ""

			if cli.Exists(key) {
				val = cli.GetOption("")
			}

			if len(val) != 0 {
				val += " "
			}

			cli.setOption("", val+key)
		}

		if idx != 0 {
			continue
		}

		val = "1"
		if idx = strings.Index(key, "="); idx != -1 {
			val = key[idx+1:]
		}

		if idx == -1 {
			idx = len(key)
		}

		cli.setOption(key[len(core.OptionPrefix):idx], val)
	}

	return cli
}

func (cli *Cli) setCommand(val string) {
	cli.command = val
}

func (cli *Cli) Command() string {
	return cli.command
}

func (cli *Cli) Options() map[string]string {
	return cli.options
}

func (cli *Cli) Exists(opt string) bool {
	_, ok := cli.options[opt]
	return ok
}

func (cli *Cli) setOption(opt string, val string) {
	cli.options[opt] = val
}

func (cli *Cli) GetOption(opt string) string {
	return cli.options[opt]
}

func (cli *Cli) Empty() bool {
	return len(cli.options) == 0 && len(cli.command) == 0
}
