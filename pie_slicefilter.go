package main

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
	"strings"
)

func main() {
	name := pie.Of([]string{"Bob", "Sally", "John", "Jane"}).
		FilterNot(func(name string) bool {
			return strings.HasPrefix(name, "J")
		}).
		Map(strings.ToUpper).
		Last()

	fmt.Println(name) // "SALLY"

	names := pie.Of([]int{1, 2, 3, 4, 5}).
		Filter(func(numb int) bool {
			return numb > 2
		})
	fmt.Println(names)
}
