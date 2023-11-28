package main

import (
	"fmt"
	"regexp"
)

func main() {
	// `e([a-z])o` => cari e ... 1 huruf ditengah ... o
	var regex = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("aditya"))
	fmt.Println(regex.MatchString("edityo"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.FindAllString("eko elo edo adi eri ari", 10))
}
