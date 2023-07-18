package main

import (
	"fmt"
	cml "github.com/seiflotfy/count-min-log"
)

func main() {
	log, _ := cml.NewForCapacity16(10000000, 0.01)

	log.Update([]byte("b"))
	log.Update([]byte("c"))
	log.Update([]byte("b"))
	log.Update([]byte("d"))
	log.BulkUpdate([]byte("a"), 1000000)
	count := log.Query([]byte("a"))
	fmt.Println(uint(count))
	//fmt.Println(log.Query([]byte("a"))
}
