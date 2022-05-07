package main

import (
	"fmt"
	"golang.org/x/gmsm/sm4"
	"io/ioutil"
	"os"
)

func main() {

	key := []byte("douknowwhoami?dr") // 自定义16位密钥
	if len(os.Args) < 2 {
		fmt.Println("Please select a shellcode file. (E.g.: encrypt payload.bin)")
		return
	}
	filename := os.Args[1]
	payload, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Println("[+]成功加密，文件已生成...")
	}

	ecbDec, err := sm4.Sm4Ecb(key, payload, true) //sm4Ecb模式pksc7填充加密
	ioutil.WriteFile("zha0gongz1", ecbDec, 0644)
}

