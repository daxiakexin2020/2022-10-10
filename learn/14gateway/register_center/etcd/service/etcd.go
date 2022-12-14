package service

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
)

func (e *Etcd) Name() string {
	return "etcd"
}

func (e *Etcd) Put(ctx context.Context, key string, val string, opts ...interface{}) error {
	options := make([]clientv3.OpOption, 0)
	for _, opt := range opts {
		if ropt, ok := opt.(clientv3.OpOption); ok {
			options = append(options, ropt)
		}
	}
	_, err := e.Client.Put(ctx, key, val)
	return err
}

func (e *Etcd) Delete(ctx context.Context, key string) error {
	_, err := e.Client.Delete(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}

func (e *Etcd) PutWithLease(ctx context.Context, key string, val string, deadtime int64) error {
	LeaseID, err := e.LeaseID(deadtime)
	if err != nil {
		return err
	}
	return e.Put(ctx, key, val, clientv3.WithLease(LeaseID), clientv3.WithLease(LeaseID))
}

func (e *Etcd) Get(ctx context.Context, key string) ([]string, error) {
	resp, err := e.Client.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, kv := range resp.Kvs {
		res = append(res, string(kv.Value))
	}
	return res, nil
}

func (e *Etcd) Watch(ctx context.Context, key string) {
	// watch key change
	rch := e.Client.Watch(ctx, key) // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func (e *Etcd) LeaseID(deadtime int64) (clientv3.LeaseID, error) {
	resp, err := e.Client.Grant(context.TODO(), deadtime)
	if err != nil {
		return 0, err
	}
	return resp.ID, nil
}

func (e *Etcd) KeepAlive(ID clientv3.LeaseID) {
	// the key 'foo' will be kept forever
	ch, kaerr := e.Client.KeepAlive(context.TODO(), ID)
	if kaerr != nil {
		log.Fatal(kaerr)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}
