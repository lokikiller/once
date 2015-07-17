package store
import (
	"testing"
	"fmt"
//	"strings"
)

func TestStore_InitializeEtcd(t *testing.T) {
	addr := []string{"133.133.134.169:2379"}
	etcd, err := InitializeEtcd(addr)
	if err != nil {
		fmt.Println(err)
		return
	}
//	etcd.Client.CreateDir("/swarm", 0)
	data, _ := etcd.Client.Get("/swarm/docker/swarm/nodes/", false, true)
	fmt.Println(data.Node.Nodes[0].Nodes[0].Nodes[0].Nodes[0])
//	etcd.Client.Create("/192.168.1.1/abc123/{[config:123]}", "/abc456", 0)
//	etcd.Client.Create("/192.168.1.1/abc123/{[name:123]}", "/abc456", 0)
//	data, _ := etcd.Client.Get("/192.168.1.1", false, true)
//	fmt.Println(data.Node.Nodes[0].Key)
//	fmt.Println(data.Node.Key)
//	str := strings.LastIndex(data.Node.Nodes[0].Key, "/")
//	fmt.Println(data.Node.Nodes[0].Key[str + 1:])
//
//	etcd.Client.Delete("/192.168.1.1/abc123", true)
//	data, _ = etcd.Client.Get("/192.168.1.1/abc123/{[name:123]}", false, false)
//	if(data != nil) {
//		fmt.Println(data.Node.Value)
//		fmt.Println(data.Node.Key)
//	}
//
//	etcd.Client.CreateDir("/hello", 0)
//	etcd.Client.CreateDir("/hello", 0)
//	etcd.Client.SetDir("/hello", 0)
//	data, _ = etcd.Client.DeleteDir("/hello")
//	if(data != nil) {
//		fmt.Println(data.EtcdIndex)
//	}
}