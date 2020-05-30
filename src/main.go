package main

import (
	"./datatype_test"
	"./operator_test"
	"./package_test"
	"fmt"
)

// 如果有init函数则会在主函数之前执行
func init() {
	println("*******func init")
	fmt.Println("3333")
}

// 主函数，程序执行入口
func main() {
	println("*******func main")
	package_test.OuterTestFunc()

	println("*******datatype_test")
	datatype_test.DataGet()

	println("*******operator_test")
	operator_test.Operate()

}
