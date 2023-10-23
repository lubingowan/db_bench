package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

const (
	AUTH_USER = "fullstackcareer"
	AUTH_PWD  = "FullStackCareer"
)

func BenchmarkZKSet(b *testing.B) {
	con, _, err := zk.Connect([]string{"127.0.0.1:2181"}, 5*time.Second, zk.WithLogInfo(false))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()

	auth := AUTH_USER + ":" + AUTH_PWD
	if err := con.AddAuth("digest", []byte(auth)); err != nil {
		con.Close()
		fmt.Println("auth err", err)
		return
	}

	for i := 0; i < b.N; i++ {
		_, err = con.Create("/k"+fmt.Sprint(i), []byte("test"), 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			// fmt.Println(err)
			continue
		}
	}
}

func BenchmarkZKGet(b *testing.B) {
	con, _, err := zk.Connect([]string{"127.0.0.1:2181"}, 5*time.Second, zk.WithLogInfo(false))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()

	auth := AUTH_USER + ":" + AUTH_PWD
	if err := con.AddAuth("digest", []byte(auth)); err != nil {
		con.Close()
		fmt.Println("auth err", err)
		return
	}

	for i := 0; i < b.N; i++ {
		//_, err = con.Create("/k"+fmt.Sprint(i), []byte("test"), 0, zk.WorldACL(zk.PermAll))
		_, _, err = con.Get("/k" + fmt.Sprint(i))
		if err != nil {
			// fmt.Println(err)
			continue
		}
	}
}
