package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error :", err.Error())
	}

	resultInt, err := strconv.Atoi("10000")
	if err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println("Error :", err.Error())
	}

	// biner
	fmt.Println(strconv.FormatInt(999, 2))

	// int to string
	fmt.Println(strconv.Itoa(999))
}
