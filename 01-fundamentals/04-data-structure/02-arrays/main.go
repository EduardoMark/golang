package main

import "fmt"

// arrays has a fixed size

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr[3] = 99
	fmt.Println(arr, arr[3])
}
