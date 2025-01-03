package main

import "fmt"


func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) {
	fmt.Println("GoodBye", name)
}

func beSocialble(name string) {
	sayHello(name)
	fmt.Println("Hows the weather?")
	sayGoodBye(name)
}

func addOne(x int) int {
	return x + 1
}

func sayHelloABunch() {
	fmt.Println("Hello")
	sayHelloABunch()
}

func main() {
	beSocialble("Bob")
	beSocialble("Alice")

	x := 5
	x = addOne(x)

	fmt.Println(x)

	x = addOne(addOne(x))
	fmt.Println(x)

	sayHelloABunch()
}