package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func BenchmarkEtcdPut(b *testing.B) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cli.Close()

	for i := 0; i < b.N; i++ {
		cli.Put(context.TODO(), "k"+fmt.Sprint(i), "v"+fmt.Sprint(i))
	}
}

func BenchmarkEtcdGet(b *testing.B) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cli.Close()

	for i := 0; i < b.N; i++ {
		_, err := cli.Get(context.TODO(), "k"+fmt.Sprint(i))
		// cli.Put(context.TODO(), "k"+fmt.Sprint(i), "v"+fmt.Sprint(i))
		if err != nil {
			fmt.Println(err)
		}
	}
}
