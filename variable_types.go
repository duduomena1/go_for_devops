package main

import "fmt"

func main() {
	fmt.Println("Constants")
	const PI = 3.14
	fmt.Println("Pi=", PI)

	fmt.Println("Integers Variables")

	fmt.Println("Signed integers:")
	var a int8 = 127
	var b int16 = -32768
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	fmt.Println("Unsigned integers:")
	var c uint8 = 255
	var d uint16 = 65535
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	fmt.Println("Machine Dependent Types:")
	var e int = 123456789
	var f uint = 123456789
	var g uintptr = 0xdeadbeef
	fmt.Println("e=", e)
	fmt.Println("f=", f)
	fmt.Println("g=", g)

	fmt.Println("Floats")
	var h float32 = 1.230970970970
	var i float64 = 3.14159456456456425464646
	fmt.Println("h=", h)
	fmt.Println("i=", i)

	fmt.Println("Complex Numbers")

	var j complex64 = 1.23456789 + 2.3456789i
	var k complex128 = 1.23456789 + 2.3456789i
	fmt.Println("j=", j)
	fmt.Println("k=", k)

	fmt.Println("Booleans")
	var l bool = true
	var m bool = false
	fmt.Println("l=", l)
	fmt.Println("m=", m)

	fmt.Println("Strings")
	var n string = "Hello, World!"
	fmt.Println("n=", n)
}
