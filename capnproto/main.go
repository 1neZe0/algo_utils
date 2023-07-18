package main

import (
	"fmt"
	books "memcap/foo"

	"capnproto.org/go/capnp/v3"
)

func main() {
	// Create a new Arena for a books.Book type.  The Arena wraps the underlying
	// buffer, providing a low-level access API.  You probably won't ever need to
	// interact with it directly.  We will ignore the meaning of "single segment"
	// for now.
	arena := capnp.SingleSegment(nil)

	// Make a brand new empty message.  A Message allocates Cap'n Proto structs within
	// its arena.  For convenience, NewMessage also returns the root "segment" of the
	// message, which is needed to instantiate the Book struct.  You don't need to
	// understand segments and roots yet (or maybe ever), but if you're curious, messages
	// and segments are documented here:  https://capnproto.org/encoding.html
	msg, seg, err := capnp.NewMessage(arena)
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

	// Great, we have our book!  Now let's set some fields.  Each field you declared in
	// your schema will produce two methods on the generated type.  The "getter" method
	// has the name of the field, for example:  Book.Title().  The corresponding "setter"
	// method is prefixed with "Set", for example:  Book.SetTitle().
	//
	// Some getters and setters return errors, which we are ignoring in this example for
	// the sake of clarity.  Your code SHOULD check these errors and handle them.
	//
	// To begin, we set the book's title to "War and Peace".
	_ = book.SetTitle("War and Peace")

	// Then, we set the page count.
	book.SetPageCount(1440)

	// Finally, we "get" these fields and print them.
	title, _ := book.Title()
	fmt.Printf("%s (%d pages)", title, book.PageCount())

	b, err := msg.Marshal()
	if err != nil {
		panic(err)
	}

	fmt.Println(b)

	// send b over the network, or write it to a file, or whatever...

	msg2, _ := capnp.Unmarshal(b)

	book2, _ := books.ReadRootBook(msg2)

	title2, _ := book.Title()
	fmt.Printf("%s (%d pages)", title2, book2.PageCount())
}
