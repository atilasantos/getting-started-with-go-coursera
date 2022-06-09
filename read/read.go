package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	const (
		FIRSTNAME = iota
		LASTNAME
	)

	var people []Person
	var openedFile string

	fmt.Println("Enter file name: ")
	_, err := fmt.Scan(&openedFile)
	if err != nil {
		fmt.Println(err)
	}

	file, _ := os.Open(openedFile)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), " ")
		person := Person{fname: names[FIRSTNAME], lname: names[LASTNAME]}
		people = append(people, person)
	}

	for _, person := range people {
		fmt.Printf("First name: %v\tLast name: %v\n", person.fname, person.lname)
	}

}
