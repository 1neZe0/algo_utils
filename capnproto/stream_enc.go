package main

import (
	"capnproto.org/go/capnp/v3"
	books "memcap/foo"
	"os"
)

func main() {
	arena := capnp.SingleSegment(nil)

	_, seg, err := capnp.NewMessage(arena)
	if err != nil {
		panic(err)
	}

	// Create a new Book struct.  Every message must have a root struct.  Again, it is
	// not important to understand "root structs" at this point.  For now, just understand
	// that every type you instantiate needs to be a "root", unless you plan on assigning
	// it to another object.  When in doubt, use NewRootXXX.
	//
	// If you're insatiably curious, see:  https://capnproto.org/encoding.html#messages
	book, err := books.NewRootBook(seg)
	if err != nil {
		panic(err)
	}

	// To begin, we set the book's title to "War and Peace".
	_ = book.SetTitle("War and Peace")

	// Then, we set the page count.
	book.SetPageCount(1440)

	// Create a new encoder that streams messages to stdout.
	// You can also use NewPackedEncoder if you want to compress
	// the data.
	encoder := capnp.NewEncoder(os.Stdout)

	// Send the book's underlying *capnp.Message.  Note that we
	// could have also passed the 'msg' variable we obtained from
	// our previous call to capnp.Unmarshal or capnp.NewMessage.
	// In most cases, however, it is more convenient to use the
	// generated type's Message() method.
	err = encoder.Encode(book.Message())
	if err != nil {
		panic(err)
	}

}
