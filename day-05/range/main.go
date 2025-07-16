package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for index, num := range nums {
		fmt.Println("Index: ", index, "Number: ", num)
	}

	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("Sum: ", sum)

	fruits := map[string]string{"a": "Apple", "b": "Banana", "c": "Coconut"}

	for key, value := range fruits {
		fmt.Printf("%s => %s\n", key, value)
	}

	for key := range fruits {
		fmt.Println("key:", key)
	}

	for index, char := range "proatik" {
		fmt.Printf("%d => %c\n", index, char)
	}

}
