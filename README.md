# go-jenkins

[![Go Reference](https://pkg.go.dev/badge/github.com/hq0101/go-jenkins.svg)](https://pkg.go.dev/github.com/hq0101/go-jenkins)

Jenkins API Client

- Jobs
    - 创建Job（构建一个自由风格的软件项目、流水线、多分支流水线、文件夹）
    - 复制Job
    - 删除Job
    - Get Pipelineruns、Describe PipelineRun
    

- View

    - 创建视图（列表视图、我的视图）

- Queue

    - GetQueue

- Nodes

    - GetNodes


# Example

- Get Categories

```go
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

	itemCategories, err := clientset.ViewV1().View().GetCategories(context.Background(), "all", 3)
	fmt.Println(itemCategories)
}

```

- 创建流水线、多分支流水线

```go
package main

import (
	"context"
	"github.com/hq0101/go-jenkins/jenkins"
	v1 "github.com/hq0101/go-jenkins/jenkins/typed/crumb/v1"
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
	
	client, err := v1.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	crumbIssuer, err := client.Crumb().GetCrumb(context.Background())
	if err != nil {
		panic(err)
	}

	// 流水线
	err = clientset.JobV1().Jobs().CreateWorkflowJob(context.Background(), "all", "pipeline-01", crumbIssuer.Crumb)
	if err != nil {
		log.Fatalln(err)
	}

	// 多分支流水线
	err = clientset.JobV1().Jobs().CreateWorkflowMultiBranchProject(context.Background(), "all", "pipeline-02", crumbIssuer.Crumb)
	if err != nil {
		log.Fatalln(err)
	}

}

```

- Get Pipelineruns

```go
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


	pipelineruns, err := clientset.JobV1().Jobs().GetPipelineRuns(context.Background(), "pipeline-01")
	if err != nil {
		log.Fatalln(err)
	}
	for _, run := range pipelineruns {
		fmt.Println(run.Name)
	}
}

```

- Describe PipelineRun

```go
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
	
	pipelinerun, err := clientset.JobV1().Jobs().DescribeJobRun(context.Background(), "pipeline-01", 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pipelinerun.Name)
}

```

- Get Console Output

```go
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
	
	output, err := clientset.JobV1().Jobs().GetConsoleText(context.Background(), "pipeline-01", 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(output)
}

```


- Run Jenkins

```shell
docker run --name jenkins --restart=on-failure -d -p 8080:8080 -p 50000:50000 --volume jenkins-data:/var/jenkins_home registry.cn-shanghai.aliyuncs.com/kubesec/jenkins:lts-jdk17
```
- Get Jenkins Initial Password

```shell
docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```



