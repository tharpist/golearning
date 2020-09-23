package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 //vars can be declared on multple lines interesting here that the var are declared and then the type is given to them so bc are type int.
	fmt.Println(b, c)

	var d = true //boooooooooolllllllllll
	fmt.Println(d)

	var e int //variables that are undefined are 0
	fmt.Println(e)

	var g string //undeclared strings still come up as nothing.

	fmt.Println("g is:", g, " YOLO")

	f := "apple" //shorthand
	fmt.Println(f)

}
