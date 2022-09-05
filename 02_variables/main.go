package main

import "fmt"

func main() {
	fmt.Println("variables")
	var username string = "Allwin"
	fmt.Print(username)

	var isLoggedin bool = true
	fmt.Printf("variable is of type %T \n", isLoggedin)

	var isnum int = 66
	fmt.Printf("This is a %T type of variable whose value is %d \n", isnum, isnum)

	newNum := 66
	fmt.Println(newNum)

	const LoginTacken string = "This is hpw u login"
	fmt.Printf("this is the way to login %s \n", LoginTacken)
}
