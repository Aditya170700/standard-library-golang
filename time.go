package main

import (
	"fmt"
	"time"
)

func main() {
	// waktu saat ini
	fmt.Println(time.Now().Local())

	// buat waktu lokal
	fmt.Println(time.Date(2023, time.November, 10, 23, 0, 0, 0, time.Local))

	// parse
	parse, _ := time.Parse(time.RFC3339, "2023-01-01T15:15:15Z")
	fmt.Println(parse)

	// "yyyy-MM-dd HH:mm:ss"
	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"

	valueTime, err := time.Parse(formatter, value)
	if err == nil {
		fmt.Println(valueTime)
		fmt.Println(valueTime.Year())
		fmt.Println(valueTime.Month())
		fmt.Println(valueTime.Day())
		fmt.Println(valueTime.Hour())
		fmt.Println(valueTime.Minute())
		fmt.Println(valueTime.Second())
	} else {
		fmt.Println("Error :", err.Error())
	}

	// duration
	duration1 := time.Second * 1000
	duration2 := time.Second * 10
	duration3 := duration1 - duration2

	fmt.Println(duration3)
}
