package main

import "fmt"

func main() {
	var slice []int
	array := [...]int{1, 2, 3, 4, 5}
	slice = array[0:5]
	fmt.Println(slice)
}