package main

import "fmt"

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func (a SortBy) Slice() {
	for i := 1; i < a.Len(); i++ {
		for j := i; j > 0 && a.Less(j, j-1); j-- {
			a.Swap(j, j-1)
		}
	}
}

func main() {
	a := SortBy([]int{9, 0, 3, 2, 10, 11, 34})

	a.Slice()
	fmt.Println(a)
}
