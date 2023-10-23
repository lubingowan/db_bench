package main

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

func BenchmarkPostgreSQLInsert(b *testing.B) {
	db, err := sql.Open("postgres", "user=zego dbname=mydb sslmode=disable")
	if err != nil {
		fmt.Println("open db err:", err)
		return
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		_, err = db.Exec("insert into users(name, gage) values($1, $2)", "k"+fmt.Sprint(i), "v"+fmt.Sprint(i))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkPostgreSQLInsertMultiColumn(b *testing.B) {
	db, err := sql.Open("postgres", "user=zego dbname=mydb sslmode=disable")
	if err != nil {
		fmt.Println("open db err:", err)
		return
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		_, err = db.Exec("insert into users_tbl(name, gage, gage1, gage2, gage3, gage4, gage5, gage6, gage7) values($1, $2, $3, $4, $5, $6, $7, $8, $9)", "k"+fmt.Sprint(i), "v"+fmt.Sprint(i), "v"+fmt.Sprint(i+1), "v"+fmt.Sprint(2), "v"+fmt.Sprint(3), "v"+fmt.Sprint(4), "v"+fmt.Sprint(5), "v"+fmt.Sprint(6), "v"+fmt.Sprint(7))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkPostgreSQLSelect(b *testing.B) {
	db, err := sql.Open("postgres", "user=zego dbname=mydb sslmode=disable")
	if err != nil {
		fmt.Println("open db err:", err)
		return
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		_, err = db.Exec("select * from users_tbl where name = $1", "k"+fmt.Sprint(i))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkPostgreSQLSelectByPrimary(b *testing.B) {
	db, err := sql.Open("postgres", "user=zego dbname=mydb sslmode=disable")
	if err != nil {
		fmt.Println("open db err:", err)
		return
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		//_, err = db.Exec("select * from users_tbl where name = $1", "k"+fmt.Sprint(i))
		_, err = db.Exec("select * from users_tbl where id = $1", i)
		if err != nil {
			fmt.Println(err)
		}
	}
}
