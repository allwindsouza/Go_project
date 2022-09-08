package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Welcome to time study of golang \n")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))

	createDate := time.Date(2020, time.April, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

}
