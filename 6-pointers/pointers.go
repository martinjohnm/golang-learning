package main

import "fmt"




type Position struct {
	x float32
	y float32
}

type badGuy struct {
	name string
	health int
	pos Position
}


func whereIsBadGuy(b *badGuy) {
	x := b.pos.x
	y := b.pos.y

	fmt.Println("(", x, ",",y, ")")
}

func addOne(num *int) {
	*num = *num+1
}

func main()  {
	x := 5 

	var xPtr *int = &x
	fmt.Println(xPtr)

	addOne(xPtr)
	fmt.Println(x)
	p := Position{4,2}

	badguy := badGuy{"Jabba the hut", 100, p}
	whereIsBadGuy(&badguy)
}