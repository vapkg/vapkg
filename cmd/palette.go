package main

import "github.com/rejchev/palette/v2"

func initPalette() {
	palette.Init(getPalette())
}

func usePalette(str string) string {
	return palette.Use(str)
}

func getPalette() *palette.Config {
	return palette.NewConfig(append([]palette.Entry{}, palette.GetBasePaletteConfig().Palette()...))
}
