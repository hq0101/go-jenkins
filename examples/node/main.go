package main

import (
	"context"
	"fmt"
	"github.com/hq0101/go-jenkins/jenkins"
	v1 "github.com/hq0101/go-jenkins/jenkins/typed/node/v1"
	"github.com/hq0101/go-jenkins/rest"
	"time"
)

func main() {
	config := &rest.Config{
		Host:     "http://192.168.127.131:8080/",
		UserName: "admin",
		Password: "d706f46313924a91ac8958cfe6ab790e",
		Timeout:  10 * time.Second,
	}
	clientset, err := jenkins.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	clientset.NodeV1().Nodes().GetNodes(context.Background())

	nodeClient, err := v1.NewForConfig(config)
	node, err := nodeClient.Nodes().GetNodes(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(node.DisplayName)
}
