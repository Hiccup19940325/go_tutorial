package main

import "fmt"

func sum(name string, nums ...int) {
	fmt.Print(name, nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum("go", 1, 2)

	nums := []int{1, 2, 3, 4}
	sum("go", nums...)
}
