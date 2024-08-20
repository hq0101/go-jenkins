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

	ret, err := clientset.PipelineModeV1().PipelineMode().ValidateJenkinsFile(context.Background(), "pipeline {\n    agent any\n\n    stages {\n        stage('Hello') {\n            steps {\n                echo 'Hello World'\n            }\n        }\n    }\n}")
	fmt.Println("ValidateJenkinsFile", ret)
	ret, err = clientset.PipelineModeV1().PipelineMode().ToJson(context.Background(), "pipeline {\n    agent any\n\n    stages {\n        stage('Hello') {\n            steps {\n                echo 'Hello World'\n            }\n        }\n    }\n}")
	fmt.Println("ToJson", ret)

	ret, err = clientset.PipelineModeV1().PipelineMode().StepsToJson(context.Background(), "pipeline {\n    agent any\n\n    stages {\n        stage('Hello') {\n            steps {\n                echo 'Hello World'\n            }\n        }\n    }\n}")
	fmt.Println("StepsToJson", ret)
	r, err := clientset.PipelineModeV1().PipelineMode().Validate(context.Background(), "pipeline {\n    agent any\n\n    stages {\n        stage('Hello') {\n            steps {\n                echo 'Hello World'\n            }\n        }\n    }\n}")
	fmt.Println("Validate", r)

	ret, err = clientset.PipelineModeV1().PipelineMode().ValidateJson(context.Background(), "{\n  \"pipeline\":{\n    \"stages\":[\n      {\n        \"name\":\"Hello\",\n        \"branches\":[\n          {\n            \"name\":\"default\",\n            \"steps\":[\n              {\n                \"name\":\"echo\",\n                \"arguments\":[\n                  {\n                    \"key\":\"message\",\n                    \"value\":{\n                      \"isLiteral\":true,\n                      \"value\":\"Hello World\"\n                    }\n                  }\n                ]\n              }\n            ]\n          }\n        ]\n      }\n    ],\n    \"agent\":{\n      \"type\":\"any\"\n    }\n  }\n}")
	fmt.Println("ValidateJson", ret)
	ret, err = clientset.PipelineModeV1().PipelineMode().StepsToJenkinsFile(context.Background(), "{\n  \"name\":\"steps\",\n  \"arguments\":[ ],\n  \"children\":[\n    {\n      \"name\":\"echo\",\n      \"arguments\":[\n        {\n          \"key\":\"message\",\n          \"value\":{\n            \"isLiteral\":true,\n            \"value\":\"Hello World\"\n          }\n        }\n      ]\n    }\n  ]\n}")
	fmt.Println("StepsToJenkinsFile", ret)
	ret, err = clientset.PipelineModeV1().PipelineMode().ToJenkinsFile(context.Background(), "{\n  \"pipeline\":{\n    \"stages\":[\n      {\n        \"name\":\"Hello\",\n        \"branches\":[\n          {\n            \"name\":\"default\",\n            \"steps\":[\n              {\n                \"name\":\"echo\",\n                \"arguments\":[\n                  {\n                    \"key\":\"message\",\n                    \"value\":{\n                      \"isLiteral\":true,\n                      \"value\":\"Hello World\"\n                    }\n                  }\n                ]\n              }\n            ]\n          }\n        ]\n      }\n    ],\n    \"agent\":{\n      \"type\":\"any\"\n    }\n  }\n}")
	fmt.Println("ToJenkinsFile", ret)
}
