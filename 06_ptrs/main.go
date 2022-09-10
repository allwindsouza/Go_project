package main

import "fmt"

func main() {

	fmt.Println("This is running")

	mynum := 22

	var ptr = &mynum

	print("This is the value of ptr: ", ptr, "\n")
	print("This is the value of *ptr: ", *ptr)

}
