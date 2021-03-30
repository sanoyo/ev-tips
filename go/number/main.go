package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("int8   -> min:%d, max:%d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16  -> min:%d, max:%d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32  -> min:%d, max:%d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64  -> min:%d, max:%d\n", int64(math.MinInt64), int64(math.MaxInt64))

	fmt.Printf("uint8  -> min:0, max:%d\n", math.MaxUint8)          // uint8の最小値は0。最小値の定数は無い
	fmt.Printf("uint16 -> min:0, max:%d\n", math.MaxUint16)         // uint16の最小値は0。最小値の定数は無い
	fmt.Printf("uint32 -> min:0, max:%d\n", uint32(math.MaxUint32)) // uint32の最小値は0。最小値の定数は無い
	fmt.Printf("uint64 -> min:0, max:%d\n", uint64(math.MaxUint64)) // uint64の最小値は0。最小値の定数は無い
}
