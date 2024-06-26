package main

import "fmt"

func main() {
	fmt.Print(sort([]int{6, 4, 7, 8, 9, 0}))
}

func sort(items []int) []int {
	for i := 1; i < len(items); i++ {
		curr := i // replicate i to maintain its integrity
		prev := i - 1 // previous element
		for { // while-statement equivalent
			if !(items[prev] > items[curr]) { break }
			items[prev], items[curr] = items[curr], items[prev] // in-line swap
			curr-- // backward index shift due to swap
			if curr <= 0 { break } // avoid out of range index with prev
			prev--
		}
	}
	return items
}