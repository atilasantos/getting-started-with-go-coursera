package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Scan(&filename)

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	personSlice := make([]Person, 0, 5)

	for scanner.Scan() {

		person := strings.Split(scanner.Text(), " ")
		personSlice = append(personSlice, Person{fname: person[0], lname: person[1]})
	}

	for _, itm := range personSlice {
		fmt.Printf("%s %s\n", itm.fname, itm.lname)
	}
}
