package main

import (
	"fmt"
	"github.com/MauriceGit/skiplist"
)

type Element int

// Implement the interface used in skiplist
func (e Element) ExtractKey() float64 {
	return float64(e)
}
func (e Element) String() string {
	return fmt.Sprintf("%03d", e)
}

func main() {
	list := skiplist.New()

	// Insert some elements
	for i := 0; i < 100; i++ {
		list.Insert(Element(i))
	}

	// Find an element
	if e, ok := list.Find(Element(53)); ok {
		fmt.Println(e)
	}

	// Delete all elements
	for i := 0; i < 100; i++ {
		list.Delete(Element(i))
	}
}

//Function	Complexity	Description
//Find	O(log(n))	Finds an element in the skiplist
//FindGreaterOrEqual	O(log(n))	Finds the first element that is greater or equal the given value in the skiplist
//Insert	O(log(n))	Inserts an element into the skiplist
//Delete	O(log(n))	Deletes an element from the skiplist
//GetSmallestNode	O(1)	Returns the smallest element in the skiplist
//GetLargestNode	O(1)	Returns the largest element in the skiplist
//Prev	O(1)	Given a skiplist-node, it returns the previous element (Wraps around and allows to linearly iterate the skiplist)
//Next	O(1)	Given a skiplist-node, it returns the next element (Wraps around and allows to linearly iterate the skiplist)
//ChangeValue	O(1)	Given a skiplist-node, the actual value can be changed, as long as the key stays the same (Example: Change a structs data)
