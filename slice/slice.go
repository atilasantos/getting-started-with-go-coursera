package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var origValue string
	sliceOfInt := make([]int, 0, 3)
	var exitRequested bool = false

	for exitRequested == false {
		fmt.Println("Type an integer number or x to leave:")
		_, err := fmt.Scan(&origValue)
		if err != nil {
			fmt.Printf(err.Error())
		}

		inputValueConverted, err := strconv.Atoi(origValue)

		switch {
		case origValue == "x":
			exitRequested = true
			fmt.Println("Requested to exit..")
			break
		case err != nil:
			fmt.Println("Invalid character!")
			break
		default:
			sliceOfInt = append(sliceOfInt, inputValueConverted)
			sort.Ints(sliceOfInt)
			fmt.Println("Ordered slice of int:", sliceOfInt)
		}
	}
}
