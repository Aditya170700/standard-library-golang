package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name, Address, Email string `required:"true" max:"10"`
}

func readField(value interface{}) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "-", valueField.Type)
		fmt.Println(valueField.Tag.Get("required"))
		fmt.Println(valueField.Tag.Get("max"))
	}
}

func IsValid(data interface{}) (result bool) {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			// ambil data, cek apakah kosong / tidak
			result = reflect.ValueOf(data).Field(i).Interface() != ""

			if !result {
				return result
			}
		}
	}

	return result
}

func main() {
	readField(Sample{"Aditya"})
	readField(Person{"Ricki", "Jogja", "aditya@gmail.com"})

	sample := Sample{"Aditya"}
	sampleType := reflect.TypeOf(sample)
	// get field name
	nameField := sampleType.Field(0)
	required := nameField.Tag.Get("required")
	max := nameField.Tag.Get("max")

	fmt.Println(required)
	fmt.Println(max)

	// contoh case validasi struct
	sampleValidation1 := Sample{"Aditya"}
	fmt.Println(IsValid(sampleValidation1)) // true
	sampleValidation2 := Sample{""}
	fmt.Println(IsValid(sampleValidation2)) // false
}
