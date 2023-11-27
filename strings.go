package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Aditya Ricki", "Aditya"))
	fmt.Println(strings.Split("Aditya Ricki", " "))
	fmt.Println(strings.ToLower("Aditya Ricki"))
	fmt.Println(strings.ToUpper("Aditya Ricki"))
	fmt.Println(strings.Trim("      Aditya Ricki    ", " "))
	fmt.Println(strings.ReplaceAll("Aditya Ricki", "a", "A"))
}
