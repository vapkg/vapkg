package cli

import (
	"fmt"
	"vapkg/internal/core"
	"vapkg/internal/utils"
)

var spinner = utils.NewSpinnerPrinter([]string{"| ", "/ ", "- ", "\\ "})

var downloadCommand = core.Command{
	Usage:       "vapkg download [--<option>[, ...]]",
	Description: "",
	Handler:     downloadCommandHandleFn,
	Options: map[string]bool{
		"":           true,
		"provider":   true,
		"attachment": false,
	},
}

func downloadCommandHandleFn(ctx core.IContext, opts map[string]string) error {
	return fmt.Errorf("uniml")
}
