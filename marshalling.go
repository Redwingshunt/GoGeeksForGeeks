package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("enter the name ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("enter the address")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)
	mp := make(map[string]string)
	mp["name"] = name
	mp["address"] = address
	data, err := json.Marshal(mp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

// structs couldve been used but i wanted to show how to use maps and json marshalling with structsd the ordering is maintained in the map as well but here its  picking randomly
