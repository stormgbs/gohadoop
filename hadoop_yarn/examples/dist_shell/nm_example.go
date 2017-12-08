package main


import (
	"log"

	"github.com/hortonworks/gohadoop/hadoop_yarn/conf"
	"github.com/hortonworks/gohadoop/hadoop_yarn/yarn_client"
	"github.com/hortonworks/gohadoop/hadoop_yarn"
)

func main() {
	var err error
	var resp *hadoop_yarn.ResourceUtilizationResponseProto
	// Create YarnConfiguration
	conf, _ := conf.NewYarnConfiguration()

	// Create RMClient
	nmClient, _ := yarn_client.CreateNMClient(conf)

	resp, err = nmClient.GetResourceUtilization()

	if err != nil {
		log.Fatal("rmClient.UpdateNodeResource: ", err)
	}else {
		nodeStatus := resp.NodeStatus
		log.Println(nodeStatus)
	}
}
