package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0o600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new bucket.
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("scores"))
		return err
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(db.GoString(), " ", db.Info())
}
