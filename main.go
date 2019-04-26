package main

import (
	."fmt"
	."../basis"
)

func main() {
	var xtbg XianTianBaGua
	xtbg.Init()
	xtbg.Show()
	Println()
	var htbg HouTianBaGua
	htbg.Init()
	htbg.Show()
}
