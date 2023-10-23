package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/boltdb/bolt"
)

func BenchmarkBoltPut(b *testing.B) {
	db, err := bolt.Open("my.db", 0o600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		err = db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("scores"))
			err := bucket.Put([]byte("Tom"+fmt.Sprint(i)), []byte("100"))
			return err
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkBoltPut2(b *testing.B) {
	db, err := bolt.Open("my.db", 0o600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		tx, err := db.Begin(true)
		if err != nil {
			fmt.Println(err)
			continue
		}

		bucket, err := tx.CreateBucketIfNotExists([]byte("scores"))
		if err != nil {
			fmt.Println(err)
			continue
		}

		if err = bucket.Put([]byte("Tom"+fmt.Sprint(i)), []byte("100")); err != nil {
			fmt.Println(err)
			continue
		}

		if err := tx.Commit(); err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkBoltView(b *testing.B) {
	db, err := bolt.Open("my.db", 0o600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		err = db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("scores"))
			_ = bucket.Get([]byte("Tom" + fmt.Sprint(i)))
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}
