package main

import (
	"encoding/binary"
	"fmt"
	"github.com/bits-and-blooms/bloom"
)

func main() {
	filter := bloom.NewWithEstimates(1000000, 0.01)
	filter.Add([]byte("Love"))

	if filter.Test([]byte("Love")) {
		fmt.Println("True")
	}
	i := uint32(100)
	n1 := make([]byte, 4)
	binary.BigEndian.PutUint32(n1, i)
	filter.Add(n1)

	filter.UnmarshalJSON(n1)

}
