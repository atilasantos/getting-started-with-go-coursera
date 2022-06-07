package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var fname string
	var address string
	fmt.Println("Please enter name!!")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	fname = scanner.Text()
	fmt.Println("Please enter address!!")
	scanner1 := bufio.NewScanner(os.Stdin)
	scanner1.Scan() // use `for scanner.Scan()` to keep reading
	address = scanner1.Text()
	jobj := make(map[string]string)
	jobj["name"] = fname
	jobj["address"] = address
	barr, err := json.Marshal(jobj)
	_ = err
	os.Stdout.Write(barr)
}
