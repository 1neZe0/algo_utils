package main

import (
	"fmt"
	"github.com/yourbasic/radix"
)

func main() {
	stringa := []string{"ciao", "come", "va", "akapulko", "delamore", "sibianto"}
	radix.Sort(stringa)
	fmt.Println(stringa)
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	radix.SortSlice(people, func(i int) string { return people[i].Name })
	fmt.Println(people)
}
