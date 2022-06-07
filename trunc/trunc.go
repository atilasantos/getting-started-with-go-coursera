package main

import (
	"fmt"
	"math"
)

func main() {

	var number float64

	fmt.Println("Enter a float pointing number:")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println(err)
	}

	number = math.Trunc(float64(number))

	fmt.Printf("\nFollowing the integer truncated number:\n%v", number)
}
