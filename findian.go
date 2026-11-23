package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Print("Enter a string: ")

    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')

    input = strings.ToLower(strings.TrimSpace(input))

    if len(input) >= 3 &&
        input[0] == 'i' &&
        input[len(input)-1] == 'n' &&
        strings.Contains(input[1:len(input)-1], "a") {

        fmt.Println("Found!")
    } else {
        fmt.Println("Not Found!")
    }
}