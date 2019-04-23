package basis

import (
	"fmt"
	"../../lib/biu"
)

const (
	YIN = 00
	YANG = 01
)

type GUA struct {
	chu byte
	er byte
	shang byte
	sum byte
}

func (gua *GUA) Init(chu byte, er byte, shang byte) {
	gua.chu = chu
	gua.er = er << 1
	gua.shang = shang << 2
	gua.sum = gua.chu + gua.er + gua.shang
}

func (gua *GUA) Show() byte {
	fmt.Println(biu.ToBinaryString(gua.sum))
	return gua.sum
}