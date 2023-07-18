// Code generated by capnpc-go. DO NOT EDIT.

package books

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

type Book capnp.Struct

// Book_TypeID is the unique identifier for the type Book.
const Book_TypeID = 0x8100cc88d7d4d47c

func NewBook(s *capnp.Segment) (Book, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Book(st), err
}

func NewRootBook(s *capnp.Segment) (Book, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Book(st), err
}

func ReadRootBook(msg *capnp.Message) (Book, error) {
	root, err := msg.Root()
	return Book(root.Struct()), err
}

func (s Book) String() string {
	str, _ := text.Marshal(0x8100cc88d7d4d47c, capnp.Struct(s))
	return str
}

func (s Book) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Book) DecodeFromPtr(p capnp.Ptr) Book {
	return Book(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Book) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Book) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Book) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Book) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Book) Title() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Book) HasTitle() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Book) TitleBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Book) SetTitle(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Book) PageCount() int32 {
	return int32(capnp.Struct(s).Uint32(0))
}

func (s Book) SetPageCount(v int32) {
	capnp.Struct(s).SetUint32(0, uint32(v))
}

// Book_List is a list of Book.
type Book_List = capnp.StructList[Book]

// NewBook creates a new list of Book.
func NewBook_List(s *capnp.Segment, sz int32) (Book_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[Book](l), err
}

// Book_Future is a wrapper for a Book promised by a client call.
type Book_Future struct{ *capnp.Future }

func (f Book_Future) Struct() (Book, error) {
	p, err := f.Future.Ptr()
	return Book(p.Struct()), err
}

const schema_85d3acc39d94e0f8 = "x\xda\x12Ht`1\xe4\xdd\xcf\xc8\xc0\x14(\xc2\xca" +
	"\xb6\xbf\xe6\xca\x95\xeb\x1dg\x1a\x03y\x18\x19\xff\xffx" +
	"0e\xee\xe15\x97[\x19X\x19\xd9\x19\x18\x04\x8fv" +
	"\x09\x9e\x05\xd1'\xcb\x19t\xff'\xe5\xe7g\x17\xeb%" +
	"'2\x16\xe4\x15X9\xe5\xe7g30\x0402\x06" +
	"r0\xb300\xb0020\x08j\x1a10\x04\xaa" +
	"03\x06\x1a0122\x8a0\x82\xc4t\x83\x18\x18" +
	"\x02u\x98\x19\x03-\x98\x18\xe5K2KrR\x19y" +
	"\x18\x98\x18y\x18\x18\xff\x17$\xa6\xa7:\xe7\x97\xe61" +
	"0\x960\xb2001\xb200\x02\x02\x00\x00\xff\xff" +
	"N\xd3$\xbc"

func RegisterSchema(reg *schemas.Registry) {
	reg.Register(&schemas.Schema{
		String: schema_85d3acc39d94e0f8,
		Nodes: []uint64{
			0x8100cc88d7d4d47c,
		},
		Compressed: true,
	})
}