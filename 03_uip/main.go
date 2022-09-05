package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input \n"
	fmt.Print(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	input, _ := reader.ReadString('\n')
	fmt.Printf("thanks for your rating,%s", input)
	fmt.Printf("type of this rating is %T \n", input)
}
