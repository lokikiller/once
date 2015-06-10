package monitor

import (
	docker "iscas.cn/samalba/dockerclient"
	"fmt"
	"time"
)

func MonitorContainers(client docker.Client, id string) error {
	data, err := client.TOPContainer(id)
	if err != nil {
		return err
	}
	if !analysisData(data) {
		//do something to elastic
	}
	return nil
}

func analysisData(data []*docker.TopRes) bool {
	for _, stream := range data {
		//analysis data is satisfied with conditions
		fmt.Println(stream.MEM)
	}
	return false
}

func WatchContainers(client docker.Client) {
	tc := time.Tick(defaultDuration)
	for {
		go func() {
			containers, _ := client.ListContainers(true, false, "");
			for _, container := range containers {
				err := MonitorContainers(client, container.Id)
				if err != nil {
					//do something to restart service
				}
			}
		}()
		<-tc
	}

}