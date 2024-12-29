package main

import "fmt"

func main() {

	/* variable assignment
	var a int = 5
	a := 5
	*/

	/* diff types of vars
	var a int8 // -127 , 127
	var b uint8 // 0, 255
	var c int16 // 
	var d uint16
	var e int  // 32 bits or 64 bits based on the system
	*/

	/* flot values
	var f1 float32
	var f2 float64
	*/

	/* can only add same types
	a := 5
	b := 3.14159 // 
	can't be added
	*/

	/* res is 3 bcz the max value is 255 beyond that it will start from 0
	var a uint8 = 255
	var b uint8 = 4
	var res uint8

	res = a + b // 
	*/

	/* Math is not real
	a := 2.0
	b := 3.0

	res := a/b
	fmt.Printf(("%.20f\n", res))
	*/

	/* booleans
	torf := true // int8
	var b bool = true
	fmt.Println(torf, b)
	*/

	/* strings
	helloworld := "Heleo"
	fmt.Println(helloworld)
	*/

	/* Infinite loop
	for {
		fmt.Println("Hlo world")
	}
	*/

	/* While loop
	i := 0
	for i < 10 {
		fmt.Println("hello world", i)
		i = i+1
	}
	*/

	/* For loop 
	for i:=0; i < 10; i++{
		fmt.Println("Hello world", i)
	}
	*/

	test := false

	if test {
		fmt.Println("Hello world")
	} else {
		fmt.Println("Good bye")
	}
	
}