package main

import (
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
)

func main() {
	test()
}

func test() {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	res := db.Update(func(tx *buntdb.Tx) error {
		value, replaced, err2 := tx.Set("test_key", "test_val", nil)
		if err2 != nil {
			return err2
		}
		fmt.Printf("old value:%q replaced:%t\n", value, replaced)

		val, err3 := tx.Get("test_key")
		if err3 != nil {
			return err3
		}
		fmt.Printf("get val=%q\n", val)
		return nil
	})
	fmt.Printf("res=%v", res)
}
