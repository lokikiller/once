package store

import (
	"github.com/coreos/go-etcd/etcd"
	"time"
)

const (
	periodicSync = 10 * time.Minute
)

type Etcd struct {
	Client       *etcd.Client
}

func InitializeEtcd(addrs []string) (*Etcd, error) {
	s := &Etcd{}

	entries := createEndpoints(addrs, "http")
	s.Client = etcd.NewClient(entries)

	go func() {
		for {
			s.Client.SyncCluster()
			time.Sleep(periodicSync)
		}
	}()
	return s, nil
}


func createEndpoints(addrs []string, scheme string) (entries []string) {
	for _, addr := range addrs {
		entries = append(entries, scheme+"://"+addr)
	}
	return entries
}