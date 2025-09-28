package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	reverse(&arr)
	fmt.Println(arr)
}
func reverse(arr *[4]int) {
	for index, value := range *arr {
		(*arr)[len(arr)-index-1] = value
	}
}
