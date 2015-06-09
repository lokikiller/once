package monitor

import (
	docker "iscas.cn/samalba/dockerclient"
	"fmt"
	"time"
)

func MonitorContainers(client docker.Client, outTime, durationTime time.Duration, id string) {
	client.ChangeDefaultTime(outTime, durationTime)
	var ec chan error
	client.StartMonitorStats(id, callback, ec)
	for {
		<-ec
	}
}

func callback(id string, stats *docker.Stats, ec chan error, arg ...interface{}) {
	fmt.Println(stats.MemoryStats.Usage)
}