package main

import (
	"fmt"
	"reflect"
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
)

func main() {
	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(reflect.TypeOf(GB))
}
