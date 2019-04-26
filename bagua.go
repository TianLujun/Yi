package gua

import (
	"fmt"
)

type BaGua struct {
	jiuGong[3][3] GUA
}

type XianTianBaGua struct {
	BaGua
}

type HouTianBaGua struct {
	BaGua
}

func (xtbg *XianTianBaGua) Init() {
	var (
		qian GUA
		kun GUA
		li GUA
		kan GUA
		zhen GUA
		xun GUA
		gen GUA
		dui GUA
		kong GUA
	)
	qian.Init(	YANG,	YANG, 	YANG)
	kun.Init(	YIN, 	YIN, 	YIN	)
	li.Init(	YANG, 	YIN, 	YANG)
	kan.Init(	YIN, 	YANG, 	YIN	)
	zhen.Init(	YANG, 	YIN, 	YIN	)
	xun.Init(	YIN, 	YANG, 	YANG)
	gen.Init(	YIN, 	YIN, 	YANG)
	dui.Init(	YANG, 	YANG, 	YIN	)
	kong.Init(	2,		2,		2	)
	xtbg.jiuGong = [3][3]GUA {
		{dui, qian, xun},
		{li, kong, kan},
		{zhen, kun, gen},
	}
}

func (xtbg *HouTianBaGua) Init() {
	var (
		qian GUA
		kun GUA
		li GUA
		kan GUA
		zhen GUA
		xun GUA
		gen GUA
		dui GUA
		kong GUA
	)
	
	qian.Init(	YANG,	YANG, 	YANG)
	kun.Init(	YIN, 	YIN, 	YIN	)
	li.Init(	YANG, 	YIN, 	YANG)
	kan.Init(	YIN, 	YANG, 	YIN	)
	zhen.Init(	YANG, 	YIN, 	YIN	)
	xun.Init(	YIN, 	YANG, 	YANG)
	gen.Init(	YIN, 	YIN, 	YANG)
	dui.Init(	YANG, 	YANG, 	YIN	)
	kong.Init(	2,		2,		2	)
	xtbg.jiuGong = [3][3]GUA {
		{xun, li, kun},
		{zhen, kong, dui},
		{gen, kan, qian},
	}
}

func (xtbg *BaGua) Show() {
	var i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if 1 == xtbg.jiuGong[i][j].shang {
				fmt.Print("\t---")
			} else if 0 == xtbg.jiuGong[i][j].shang {
				fmt.Print("\t- -")
			} else {
				fmt.Print("\t")
			}
		}
		fmt.Print("\n")
		for j = 0; j < 3; j++ {
			if 1 == xtbg.jiuGong[i][j].er {
				fmt.Print("\t---")
			} else if 0 == xtbg.jiuGong[i][j].er {
				fmt.Print("\t- -")
			} else {
				fmt.Print("\t")
			}
		}
		fmt.Print("\n")
		for j = 0; j < 3; j++ {
			if 1 == xtbg.jiuGong[i][j].chu {
				fmt.Print("\t---")
			} else if 0 == xtbg.jiuGong[i][j].chu {
				fmt.Print("\t- -")
			} else {
				fmt.Print("\t")
			}
		}
		fmt.Print("\n\n")
	}
}