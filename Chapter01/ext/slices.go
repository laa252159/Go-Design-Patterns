package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50, 60, 70, 80}
	fmt.Println(mySlice)
	mySlice = append(mySlice, 5)
	fmt.Println(mySlice)
	mySlice2 := append(mySlice[:3], mySlice[4:]...)
	fmt.Println(mySlice2)

}
