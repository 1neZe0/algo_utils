package main

import (
	"capnproto.org/go/capnp/v3"
	"fmt"
	books "memcap/foo"
	"os"
)

func main() {
	// Create a new decoder that reads from stdin.
	// Use capnp.NewPackedDecoder if you are expecting a
	// packed byte-stream.
	decoder := capnp.NewDecoder(os.Stdin)

	// Read the message from stdin.
	msg, err := decoder.Decode()
	if err != nil {
		panic(err)
	}

	// Extract the root struct from the message.
	book, err := books.ReadRootBook(msg)
	if err != nil {
		panic(err)
	}

	// Access fields from the struct.  Again, we're
	// ignoring errors, but you definitely shouldn't.
	title, _ := book.Title()
	pageCount := book.PageCount()
	fmt.Printf("%q has %d pages\n", title, pageCount)
}
