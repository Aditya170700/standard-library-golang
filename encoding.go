package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// base64
	encoded := base64.StdEncoding.EncodeToString([]byte("Aditya Ricki Julianto"))
	fmt.Println(encoded)

	var decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err == nil {
		fmt.Println(string(decoded))
	} else {
		fmt.Println("Error :", err.Error())
	}

	// csv
	csvStr := "aditya,ricki,julianto\n" +
		"ricki,aditya,julianto\n" +
		"julianto,ricki,aditya\n" +
		"senin,selasa,rabu\n"

	reader := csv.NewReader(strings.NewReader(csvStr))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"aditya", "ricki", "julianto"})
	_ = writer.Write([]string{"ricki", "aditya", "julianto"})
	_ = writer.Write([]string{"julianto", "ricki", "aditya"})
	_ = writer.Write([]string{"senin", "selasa", "rabu"})
	writer.Flush()
}
