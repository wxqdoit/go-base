package operator_test

import "fmt"

func Operate() {
	var a = 20.2
	var b = 3.6
	c := "name"
	d := ":wxq"
	var f3 float32 = 444.000068889955
	var f4 float64 = 444.000068889955
	println(a + b)
	println(a - b)
	println(a * b)
	println(a / b)
	println(c + d)
	println(f3, f4)

	if a > b {
		println("a>b")
	}
	if a < b {
		println("a<b")
	}
	if a <= b {
		fmt.Printf("- a 小于等于 b\n")
	}
	if b >= a {
		fmt.Printf("- b 大于等于 a\n")
	}
	bt := true
	bf := false
	if bt && bf {
		println("bt&&bf", true)
	}
	if bt || bf {
		println("bt||bf", true)
	}
	println(!bt)

	/**
	p	q	p & q	p | q	p ^ q
	0	0	0		0		0
	0	1	0		1		1
	1	1	1		1		0
	1	0	0		1		1
	*/
	var ba uint = 8
	var bb uint = 12
	fmt.Printf("%b", bb)
	println("")
	println(ba & bb)
	println(ba | bb)
	println(ba ^ bb)

}
