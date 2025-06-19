package utils

import (
	"fmt"
	"github.com/rejchev/palette/v2"
	"os"
)

const vaPrefix = "{VAPKG}[vapkg]{R}:"
const vaColor = "{VAPKG}"

// Disclaimer:
//// Even if some levels of logging and print wrappers occur in stdout.
//// I prefer a semantic separation.

func VaPrefix() string {
	return vaPrefix
}

func VaSprint(args ...any) string {
	return palette.Use(fmt.Sprint(args...))
}

func VaSprintf(format string, args ...any) string {
	return palette.Use(fmt.Sprintf(format, args...))
}

func VaSprintln(args ...any) string {
	return palette.Use(fmt.Sprintln(args...))
}

func VaPrintWithPrefix(args ...any) (n int, err error) {
	return fmt.Fprint(os.Stdout, VaSprint(append([]any{VaPrefix(), " "}, args...)...))
}

func VaPrint(args ...any) (n int, err error) {
	return fmt.Fprint(os.Stdout, VaSprint(args...))
}

func VaPrintln(args ...any) (n int, err error) {
	return fmt.Fprintln(os.Stdout, VaSprint(args...))
}

func VaPrintlnWithPrefix(args ...any) (n int, err error) {
	return fmt.Fprintln(os.Stdout, VaSprint(append([]any{VaPrefix(), " "}, args...)...))
}

func VaPrintf(format string, args ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, VaSprintf(format, args...))
}

func VaPrintfWithPrefix(format string, args ...any) (n int, err error) {
	return fmt.Fprintf(os.Stdout, VaSprintf("%s "+format, append([]any{VaPrefix()}, args...)...))
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
