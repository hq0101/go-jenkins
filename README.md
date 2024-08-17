# go-jenkins

[![Go Reference](https://pkg.go.dev/badge/github.com/hq0101/go-jenkins.svg)](https://pkg.go.dev/github.com/hq0101/go-jenkins)

Jenkins API Client

- Run Jenkins

```shell
docker run --name jenkins --restart=on-failure -d -p 8080:8080 -p 50000:50000 --volume jenkins-data:/var/jenkins_home registry.cn-shanghai.aliyuncs.com/kubesec/jenkins:lts-jdk17
```
- Get Jenkins Initial Password

```shell
docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```

