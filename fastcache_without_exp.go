package main

import (
	"fmt"
	"github.com/VictoriaMetrics/fastcache"
	"log"
)

func main() {
	c := fastcache.New(1)
	defer c.Reset()

	c.SetBig([]byte("key"), []byte("value"))
	var v []byte
	v = c.GetBig(nil, []byte("key"))
	fmt.Println(string(v))

	if v := c.Get(nil, []byte("aaa")); len(v) != 0 {
		log.Fatalf("unexpected non-empty value obtained from small cache: %q", v)
	}
	if v, exist := c.HasGet(nil, []byte("aaa")); exist || len(v) != 0 {
		log.Fatalf("unexpected non-empty value obtained from small cache: %q", v)
	}

	c.Set([]byte("key"), []byte("value"))
	if v := c.Get(nil, []byte("key")); string(v) != "value" {
		log.Fatalf("unexpected value obtained; got %q; want %q", v, "value")
	}
	if v := c.Get(nil, nil); len(v) != 0 {
		log.Fatalf("unexpected non-empty value obtained from small cache: %q", v)
	}
	if v, exist := c.HasGet(nil, nil); exist {
		log.Fatalf("unexpected nil-keyed value obtained in small cache: %q", v)
	}
	if v := c.Get(nil, []byte("aaa")); len(v) != 0 {
		log.Fatalf("unexpected non-empty value obtained from small cache: %q", v)
	}

	c.Set([]byte("aaa"), []byte("bbb"))
	if v := c.Get(nil, []byte("aaa")); string(v) != "bbb" {
		log.Fatalf("unexpected value obtained; got %q; want %q", v, "bbb")
	}
	if v, exist := c.HasGet(nil, []byte("aaa")); !exist || string(v) != "bbb" {
		log.Fatalf("unexpected value obtained; got %q; want %q", v, "bbb")
	}

	c.Reset()
	if v := c.Get(nil, []byte("aaa")); len(v) != 0 {
		log.Fatalf("unexpected non-empty value obtained from empty cache: %q", v)
	}
	if v, exist := c.HasGet(nil, []byte("aaa")); exist || len(v) != 0 {
		log.Fatalf("unexpected non-empty value obtained from small cache: %q", v)
	}

	// Test empty value
	k := []byte("empty")
	c.Set(k, nil)
	if v := c.Get(nil, k); len(v) != 0 {
		log.Fatalf("unexpected non-empty value obtained from empty entry: %q", v)
	}
	if v, exist := c.HasGet(nil, k); !exist {
		log.Fatalf("cannot find empty entry for key %q", k)
	} else if len(v) != 0 {
		log.Fatalf("unexpected non-empty value obtained from empty entry: %q", v)
	}
	if !c.Has(k) {
		log.Fatalf("cannot find empty entry for key %q", k)
	}
	if c.Has([]byte("foobar")) {
		log.Fatalf("non-existing entry found in the cache")
	}
}
