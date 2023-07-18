package main

import (
	"fmt"
	"github.com/Henry-Sarabia/blank"
)

func main() {
	phrase := "this is a phrase"

	str := blank.Remove(phrase)

	fmt.Println(str)

	// output: "thisisaphrase"

	//func search(qrs []string) error {
	//	if blank.Has(qrs) {
	//	// return error
	//}
	//
	//	// rest of code
	//}
}
