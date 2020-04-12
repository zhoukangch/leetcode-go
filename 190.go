package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i <= 31; i++ {
		var temp uint32
		temp = (1 << uint8(i)) & num
		if temp > 0 {
			result += 1 << uint8(31-i)
		}
	}
	return result
}

func main() {
	fmt.Println(reverseBits(uint32(43261596)))
}
