package datatype_test

var TBool bool = true
var TString string = "oh, go WTF!!"

var TInt int = 3252
var TFloat32 float32 = 3252.32
var TFloat64 float64 = 2.6400000000000000002330

func DataGet() {
	//变量
	intVal, intVal1 := 1, 2
	kar := "sss"
	a := 12
	a = 20

	var c, d int
	c = a
	d = 3
	c, d = d, c

	_, w, e := numbers()
	println(intVal, intVal1, kar, a, c, d, w, e)

	//常量
	const wxq string = "wxqdoit"
	const age, gender = 12, "南"
	const ( //相当于枚举
		pass = 1
		brec = 2
		dead = 3
	)
	println(wxq, age, gender, pass)

	println(TBool)
	println(TString)
	println(TInt)
	println(TFloat32)
	println(TFloat64)
}

func numbers() (int, int, int) {
	//iota
	const (
		v1 = iota
		v2 = "www"
		v3
		v4
	)
	println(v1, v2, v3, v4)
	var q, w, e = 1, 2, 3
	return q, w, e
}
