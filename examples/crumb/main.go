package main

import (
	"context"
	"fmt"
	v1 "github.com/hq0101/go-jenkins/jenkins/typed/crumb/v1"
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

	client, err := v1.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	crumbIssuer, err := client.Crumb().GetCrumb(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(crumbIssuer)
}
