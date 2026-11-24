package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the name of the text file: ")
	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading filename:", err)
		return
	}
	filename = strings.TrimSpace(filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []name
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			fmt.Println("Skipping invalid line:", scanner.Text())
			continue
		}

		person := name{
			fname: fields[0],
			lname: fields[1],
		}
		names = append(names, person)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for _, person := range names {
		fmt.Println(person.fname, person.lname)
	}
}
