package main

import (
	//"database/sql"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func BenchmarkMysqlInsert(b *testing.B) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lubing_wan?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < b.N; i++ {
		_, err = db.Exec("insert into users(name, gage) values(?, ?)", "k"+fmt.Sprint(i), "v"+fmt.Sprint(i))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkMysqlInsertMultiColumn(b *testing.B) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lubing_wan?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < b.N; i++ {
		_, err = db.Exec("insert into users_tbl(name, gage, gage1, gage2, gage3, gage4, gage5, gage6, gage7) values(?, ?, ?,?,?,?,?,?, ?)", "k"+fmt.Sprint(i), "v"+fmt.Sprint(i), "v"+fmt.Sprint(i+1), "v"+fmt.Sprint(2), "v"+fmt.Sprint(3), "v"+fmt.Sprint(4), "v"+fmt.Sprint(5), "v"+fmt.Sprint(6), "v"+fmt.Sprint(7))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkMysqlSelect(b *testing.B) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lubing_wan?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < b.N; i++ {
		_, err = db.Exec("select * from users_tbl where name = ?", "k"+fmt.Sprint(i))
		//_, err = db.Exec("select * from users_tbl where id = ?", i)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkMysqlSelectByPrimary(b *testing.B) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lubing_wan?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < b.N; i++ {
		//_, err = db.Exec("select * from users_tbl where name = ?", "k"+fmt.Sprint(i))
		_, err = db.Exec("select * from users_tbl where id = ?", i)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkMysqlSelectByIndex(b *testing.B) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lubing_wan?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < b.N; i++ {
		_, err = db.Exec("select * from users_tbl where gage = ? AND gage IS NOT NULL", "v"+fmt.Sprint(i))
		if err != nil {
			fmt.Println(err)
		}
	}
}
