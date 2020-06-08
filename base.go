package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 函数外面是不可以用:=的
// 这里不是全局变量，没有全局变量一说
// 这里的是包内变量，只作用于这个包

var (
	aa = 5
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	// 定义一个整形和一个字符串 在没有赋初始值的情况下，int类型默认等于0，string类型默认等于""，
	// 如果不用q% 无法显示出""字符串
	var a int
	var s string
	fmt.Println("%d %q\n", a, s)
}

func variableInitiaValue() {
	// 可以像python 一样 同时定义多个
	var a, b int = 3, 4
	var s string = "fuck\n"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	// 可以不注明变量类型直接交给编译器推断类型
	var a, b, c, s = 3, 4, true, "def\n"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// 最好用这种形式定义变量
	// 只能在函数里面用这种方法，函数外面不可以
	a, b, c, s := 3, 4, true, "def\n"
	b = 6
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func traingle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitiaValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	traingle()
	consts()
	enums()
}
