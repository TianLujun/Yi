package gua

import (
	"fmt"
	"../../lib/biu"
)

type GUA struct {
	chu byte
	er byte
	shang byte
	sum byte
}

const (
	YIN byte = 00
	YANG byte = 01
)

func (gua *GUA) Init(chu byte, er byte, shang byte) {
	gua.chu = chu
	gua.er = er
	gua.shang = shang
	gua.sum = gua.chu + gua.er << 1 + gua.shang << 2
}

func (gua *GUA) Show() byte {
	fmt.Println(biu.ToBinaryString(gua.sum))
	if 1 == gua.shang {
		fmt.Println("---")
	} else if 0 == gua.shang {
		fmt.Println("- -")
	}
	if 1 == gua.er {
		fmt.Println("---")
	} else if 0 == gua.er {
		fmt.Println("- -")
	}
	if 1 == gua.chu {
		fmt.Println("---")
	} else if 0 == gua.chu {
		fmt.Println("- -")
	}
	return gua.sum
}