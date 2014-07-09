package main

import (
    "fmt"
    "container/list"
)

func main() {
	myList := list.New()
	myList.PushBack("foobar")
	myList.PushBack(1)
	
	// print all element in myList
	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Println("val: ", element.Value)
	}

	// Failed print second element specifing index
	// firstElement := myList[0]
	// fmt.Println("firstElement: ", firstElement) //Error: "invalid operation: myList[0] (type *list.List does not support indexing)"
}
