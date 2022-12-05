package service

import (
	"context"
	"fmt"
	"log"
)

func TesPut(key string, val string) {

	etcd, err := NewEtcd([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalf("连接etcd失败", err)
	}

	err = etcd.Put(context.Background(), key, val)
	fmt.Println("**************put**************", key, val, err)

	if err != nil {
		log.Fatalf("put value 失败", err)
	}

}

func TesGet(key string) {
	etcd, err := NewEtcd([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalf("连接etcd失败", err)
	}
	res, err := etcd.Get(context.Background(), key)
	fmt.Println("**************get**************", res, err)
}

func TesWatch(key string) {
	etcd, err := NewEtcd([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalf("连接etcd失败", err)
	}
	etcd.Watch(context.Background(), key)
}

func TesPutWithLease(key string, val string) {

	etcd, err := NewEtcd([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalf("连接etcd失败", err)
	}

	err = etcd.PutWithLease(context.Background(), key, val, 10)
	fmt.Println("**************put**************", key, val, err)

	if err != nil {
		log.Fatalf("put value 失败", err)
	}

}
