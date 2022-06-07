package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var name string
	var address string

	fmt.Println("Enter your name: ")
	name, _ = in.ReadString('\n')
	name = strings.Trim(name, "\n")
	fmt.Println("Enter your address: ")
	address, _ = in.ReadString('\n')
	address = strings.Trim(address, "\n")

	person := make(map[string]string)
	person["name"] = name
	person["address"] = address

	p, _ := json.Marshal(person)

	fmt.Println("Json Object:")
	fmt.Println(string(p))
}
