// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
//собственное множество.

package main

import "fmt"

func createSet(sequence []string) map[string]bool {
	set := make(map[string]bool)
	for _, s := range sequence {
		set[s] = true
	}
	return set
}
func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := createSet(sequence)
	for item := range set {
		fmt.Println(item)
	}
}
