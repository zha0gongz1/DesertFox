package main

import (
	"fmt"
	"strconv"
)

var XorKey []byte = []byte{0x32, 0x34, 0x85, 0x6A, 0xA3, 0xFF, 0xF4, 0x7B}
var result string

func main() {
	j := 0
	s := ""
	src := "http://10.10.99.190:8080/payload.bin"
	bt := []rune(src)
	for i := 0; i < len(bt); i++ {
		s = strconv.FormatInt(int64(byte(bt[i])^XorKey[j]), 16)
		if len(s) == 1 {
			s = "0" + s
		}
		result = result + (s)
		j = (j + 1) % 8
	}
	fmt.Println(result)

}
