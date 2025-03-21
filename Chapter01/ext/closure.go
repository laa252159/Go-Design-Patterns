package main

func main() {
	addN := func(m int) func(n int) int {
		return func(n int) int {
			return m + n
		}
	}

	addFive := addN(5)
	result := addFive(6)
	//5 + 6 must print 7
	println(result)
}
