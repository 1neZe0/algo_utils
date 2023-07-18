package main

import (
	"encoding/json"
	"fmt"
	"github.com/kelindar/column"
	"log"
	"time"
)

type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (l Location) MarshalBinary() ([]byte, error) {
	return json.Marshal(l)
}

func (l *Location) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, l)
}

func main() {
	// Create a new collection with some columns
	players := column.NewCollection()
	players.CreateColumn("name", column.ForString())
	players.CreateColumn("class", column.ForString())
	players.CreateColumn("balance", column.ForFloat64())
	players.CreateColumn("age", column.ForInt16())

	index, err := players.Insert(func(r column.Row) error {
		r.SetString("name", "merlin")
		r.SetString("class", "rogue")
		r.SetFloat64("balance", 99.95)
		r.SetInt16("age", 107)
		return nil
	})
	fmt.Println(index, err)

	_, err = players.Insert(func(r column.Row) error {
		r.SetString("name", "merlin")
		r.SetString("class", "mage")
		r.SetFloat64("balance", 99.95)
		r.SetInt16("age", 107)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// This query performs a full scan of "class" column
	players.Query(func(txn *column.Txn) error {
		count := txn.WithValue("class", func(v interface{}) bool {
			return v == "mage"
		}).Count()
		fmt.Println(count)
		return nil
	})

	//out.CreateIndex("rogue", "class", func(v interface{}) bool {
	//	return v == "rogue"
	//})
	// How many rogues and mages?
	players.Query(func(txn *column.Txn) error {
		count := txn.With("rogue").Union("mage").Count()
		fmt.Println(count)
		return nil
	})

	// How many rogues that are over 30 years old?
	players.Query(func(txn *column.Txn) error {
		txn.With("rogue").WithFloat("age", func(v float64) bool {
			return v >= 30
		}).Count()
		return nil
	})

	// Expiring values
	players.Insert(func(r column.Row) error {
		r.SetString("name", "Merlin")
		r.SetString("class", "mage")
		r.SetTTL(5 * time.Second) // time-to-live of 5 seconds
		return nil
	})

	// Range over all of the players and update (successfully their balance)
	players.Query(func(txn *column.Txn) error {
		balance := txn.Float64("balance")
		//txn.Range(func(i uint32) {
		//	v.Set(10.0) // Update the "balance" to 10.0
		//})
		fmt.Println(balance)
		// Returns an error, transaction will be rolled back
		return fmt.Errorf("bug")
	})

	players = column.NewCollection()
	players.CreateColumn("name", column.ForKey())     // Create a "name" as a primary-key
	players.CreateColumn("class", column.ForString()) // .. and some other columns

	// Insert a player with "merlin" as its primary key
	players.InsertKey("merlin", func(r column.Row) error {
		r.SetString("class", "mage")
		return nil
	})
	fmt.Println(players.Count())

	//players.CreateColumn("location", ForRecord(func() *Location {
	//	return new(Location)
	//}))
}
