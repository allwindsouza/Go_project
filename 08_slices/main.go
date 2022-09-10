package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is running")

	var fruitlist = []string{"Apple", "Banana", "Orange"}

	// fruitlist[0] = "Apple"
	// fruitlist[1] = "Banana"
	// fruitlist[2] = "Orange"

	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "Mango")

	fruitlist = append(fruitlist[1:])
	fmt.Println(fruitlist)

	highscores := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("This is high scores: ", highscores)

	highscores = append(highscores, 55, 66, 77)

	fmt.Println("This is high scores updated: ", highscores)

	sort.Ints(highscores)

	fmt.Println("This is high scores updated: ", highscores)

	fmt.Println(sort.IntsAreSorted(highscores))

}
