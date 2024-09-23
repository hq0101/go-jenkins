package main

import (
	"context"
	"fmt"
	"github.com/hq0101/go-jenkins/jenkins"
	"github.com/hq0101/go-jenkins/rest"
	"log"
	"time"
)

func main() {
	config := &rest.Config{
		Host:     "http://192.168.127.131:8080/",
		UserName: "admin",
		Password: "11751d9d9a1f598eb6f4da20d08f432f24",
		Timeout:  10 * time.Second,
	}
	clientset, err := jenkins.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	//configXML, err := clientset.JobV1().Jobs().GetConfigXML(context.Background(), "pipeline-01")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(configXML)

	//client, err := v1.NewForConfig(config)
	//if err != nil {
	//	panic(err)
	//}
	//crumbIssuer, err := client.Crumb().GetCrumb(context.Background())
	//if err != nil {
	//	panic(err)
	//}

	//err = clientset.JobV1().Jobs().CreateWorkflowJob(context.Background(), "all", "pipeline-01", crumbIssuer.Crumb)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//err = clientset.JobV1().Jobs().CreateWorkflowMultiBranchProject(context.Background(), "all", "pipeline-06", crumbIssuer.Crumb)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//pipelineruns, err := clientset.JobV1().Jobs().GetPipelineRuns(context.Background(), "pipeline-01")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for _, run := range pipelineruns {
	//	fmt.Println(run.Name)
	//}

	pipelinerun, err := clientset.JobV1().Jobs().DescribeJobRun(context.Background(), "pipeline-01", 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pipelinerun.Name)
}
