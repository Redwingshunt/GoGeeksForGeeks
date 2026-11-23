// Write a Go program named slice.go that performs the following steps:

// Create an empty integer slice with length 3 before entering the loop.

// Prompt the user to enter integers one at a time.

// Each time an integer is entered:

// The slice must expand dynamically to hold additional integers.

// The program should continue running until the user enters “X” instead of a number.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Empty slice: length 0, initial capacity 3.
	numbers := make([]int, 0, 3)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter an integer (or X to stop): ")

		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "X" || input == "x" {
			break
		}

		if input != "" {
			num, convErr := strconv.Atoi(input)
			if convErr != nil {
				fmt.Println("Please enter a valid integer or X.")
			} else {
				numbers = append(numbers, num)
				sort.Ints(numbers)
				fmt.Println(numbers)
			}
		}

		if err != nil { // End of input (for example, Ctrl+Z on Windows)and cmd +z on the mac book
			break
		}
	}
}
