package monitor

import (
	"testing"
	docker "iscas.cn/samalba/dockerclient"
//	"time"
)

func TestMonitor_MonitorContainers(t *testing.T) {
	//	url := "133.133.134.168:2375"
	url := "133.133.134.169:2375"
	client, _ := docker.NewDockerClient(url, nil)
	//	MonitorContainers(client, 25 * time.Second, 5 * time.Second, "d69d56")
	WatchContainers(client)
	//	err := MonitorContainers(client, "ec8")
	//	fmt.Println(err)
}