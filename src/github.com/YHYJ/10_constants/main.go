/* 常量 */
// 常量使用关键字`const`定义，不能用`:=`定义
// 常量可以是字符、字符串、布尔或者数值类型（数值类型具有高精度）

package main

import (
	"fmt"
)

const World = "世界"
const Pi float64 = 3.1415926535897932384626

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("Hello", World)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Big))
}
