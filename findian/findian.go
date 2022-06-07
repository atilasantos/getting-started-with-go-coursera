package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string
	fmt.Println("Enter a string so we can find the ['i','a','n'] strings: ")
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Printf(err.Error())
	}

	str = strings.ToLower(str)

	switch {
	case strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a"):
		fmt.Println("Found!")
	default:
		fmt.Println("Not Found!")
	}

}
