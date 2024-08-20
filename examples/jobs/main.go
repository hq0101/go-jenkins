package main

import (
	"context"
	"fmt"
	"github.com/hq0101/go-jenkins/jenkins"
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
	configXML, err := clientset.JobV1().Jobs().GetConfigXML(context.Background(), "pipeline-01")
	if err != nil {
		panic(err)
	}
	fmt.Println(configXML)
}
