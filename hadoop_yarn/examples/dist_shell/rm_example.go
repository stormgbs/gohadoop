package main

import (
	"log"

	"github.com/hortonworks/gohadoop/hadoop_yarn/conf"
	"github.com/hortonworks/gohadoop/hadoop_yarn/yarn_client"
)

func main() {
	var err error

	// Create YarnConfiguration
	conf, _ := conf.NewYarnConfiguration()

	// Create RMClient
	rmClient, _ := yarn_client.CreateRMClient(conf)

	// Some useful information
	var (
		host     string = "node004"
		port     int32  = 8040
		memoryMB int32  = 8192
		vcores   int32  = 4
	)

	// Update node resource
	err = rmClient.UpdateNodeResource(host, port, memoryMB, vcores)
	if err != nil {
		log.Fatal("rmClient.UpdateNodeResource: ", err)
	}

}
