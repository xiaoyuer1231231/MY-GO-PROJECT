package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串转换为int: %d \n", num)
	str1 := strconv.Itoa(num)
	fmt.Printf("int转换为字符串: %s \n", str1)

	ui64, err := strconv.ParseUint(str, 10, 32)
	fmt.Printf("字符串转换为uint64: %d \n", ui64)

	str2 := strconv.FormatUint(ui64, 2)
	fmt.Printf("uint64转换为字符串: %s \n", str2)

	str16 := strconv.FormatUint(str, 32) // 十六进制}
	fmt.Printf("uint64转换为字符串: %s \n", str16)

}
