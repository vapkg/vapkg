package utils

import (
	"fmt"
	"github.com/rejchev/palette/v2"
	"os"
)

const vaPrefix = "{VAPKG}[vapkg]{R}"
const vaColor = "{VAPKG}"

// Disclaimer:
//// Even if some levels of logging and print wrappers occur in stdout.
//// I prefer a semantic separation.

func VaPrefix() string {
	return vaPrefix
}

func VaPrint(args ...any) (n int, err error) {
	return fmt.Fprint(os.Stdout, palette.Use(VaPrefix()+" "+fmt.Sprint(args...)))
}

func VaPrintln(args ...any) (n int, err error) {
	return fmt.Fprintln(os.Stdout, palette.Use(VaPrefix()+" "+fmt.Sprintln(args...)))
}

func VaPrintf(format string, args ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, palette.Use(VaPrefix()+" "+fmt.Sprintf(format, args...)))
}

func InitVaPalette() {
	if !palette.IsInit() {
		palette.Init(getPaletteConfig())
	}
}

func getPaletteConfig() *palette.Config {
	return palette.NewConfig(append([]palette.Entry{
		palette.CreateEntry(vaColor, "FHC:#007d9c"),
	}, palette.GetBasePaletteConfig().Palette()...))
}
