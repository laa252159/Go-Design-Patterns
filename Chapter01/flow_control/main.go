package main

import "fmt"

func main() {
	ten := 10
	if ten == 20 {
		println("This shouldn't be printed as 10 isn't equal to 20")
	} else {
		println("Ten is not equals to 20")
	}

	// OR operator is "||" yo just one condition needs to be true
	if "a" == "b" || 10 == 10 || true == false {
		// Enter here. Although "a" isn't "b" and true isn't false, 10 is equal
		// to 10 so at least one condition is satisfied
		println("10 is equal to 10")

		// AND operator is && so BOTH conditions must be satisfied
	} else if 11 == 11 && "go" == "go" {
		println("This isn't print because previous condition was satisfied")

	} else {
		println("In case no condition is satisfied, print this")
	}

	number := 3
	switch number {
	case 1:
		println("Number is 1")
	case 2:
		println("Number is 2")
	case 3:
		println("Number is 3")
	}

	for i := 0; i <= 10; i++ {
		print(i)
		print("\t")
	}
	println()
	println("*")

	for i := 0; i <= 10; i++ {
		print(i)
		print("\t")
	}
	println()

	my_array := [3]int{33, 222, 3221}
	for index, value := range my_array {
		fmt.Printf("index -> %d", index)
		print("\t")
		fmt.Printf("value -> %d", value)
		println()
	}
}
