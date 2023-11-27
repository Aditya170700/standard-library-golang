package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Aditya")
	data.PushBack("Ricki")
	data.PushBack("Julianto")

	// queue
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// stack
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
