
# validateJenkinsfile

```bash
curl --location 'http://192.168.127.131:8080/pipeline-model-converter/validateJenkinsfile' \
--header 'Authorization: Basic YWRtaW46ZDcwNmY0NjMxMzkyNGE5MWFjODk1OGNmZTZhYjc5MGU=' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'jenkinsfile=pipeline {
    agent any

    stages {
        stage('\''Hello'\'') {
            steps {
                echo '\''Hello World'\''
            }
        }
    }
}
'
```

- response

```json
{
    "status": "ok",
    "data": {
        "result": "success"
    }
}
```

# toJson

```bash
curl --location 'http://192.168.127.131:8080/pipeline-model-converter/toJson' \
--header 'Authorization: Basic YWRtaW46ZDcwNmY0NjMxMzkyNGE5MWFjODk1OGNmZTZhYjc5MGU=' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'jenkinsfile=pipeline {
    agent any

    stages {
        stage('\''Hello'\'') {
            steps {
                echo '\''Hello World'\''
            }
        }
    }
}
'
```

- response

```json
{
    "status": "ok",
    "data": {
        "result": "success",
        "json": {
            "pipeline": {
                "stages": [
                    {
                        "name": "Hello",
                        "branches": [
                            {
                                "name": "default",
                                "steps": [
                                    {
                                        "name": "echo",
                                        "arguments": [
                                            {
                                                "key": "message",
                                                "value": {
                                                    "isLiteral": true,
                                                    "value": "Hello World"
                                                }
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "agent": {
                    "type": "any"
                }
            }
        }
    }
}
```

# toJenkinsfile

```bash
curl --location 'http://192.168.127.131:8080/pipeline-model-converter/toJenkinsfile' \
--header 'Authorization: Basic YWRtaW46ZDcwNmY0NjMxMzkyNGE5MWFjODk1OGNmZTZhYjc5MGU=' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'json={
            "pipeline": {
                "stages": [
                    {
                        "name": "Hello",
                        "branches": [
                            {
                                "name": "default",
                                "steps": [
                                    {
                                        "name": "echo",
                                        "arguments": [
                                            {
                                                "key": "message",
                                                "value": {
                                                    "isLiteral": true,
                                                    "value": "Hello World"
                                                }
                                            }
                                        ]
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "agent": {
                    "type": "any"
                }
            }
        }'
```

- response

```json
{
    "status": "ok",
    "data": {
        "result": "success",
        "jenkinsfile": "pipeline {\n  agent any\n  stages {\n    stage('Hello') {\n      steps {\n        echo 'Hello World'\n      }\n    }\n\n  }\n}"
    }
}
```


# stepsToJson

```bash
curl --location 'http://192.168.127.131:8080/pipeline-model-converter/stepsToJson' \
--header 'Authorization: Basic YWRtaW46ZDcwNmY0NjMxMzkyNGE5MWFjODk1OGNmZTZhYjc5MGU=' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'jenkinsfile=stages {
        stage('\''Hello'\'') {
            
        }
    }
'
```

- response

```json
{
    "status": "ok",
    "data": {
        "result": "success",
        "json": [
            {
                "name": "stages",
                "arguments": [],
                "children": [
                    {
                        "name": "stage",
                        "arguments": {
                            "isLiteral": true,
                            "value": "Hello"
                        },
                        "children": []
                    }
                ]
            }
        ]
    }
}
```


```bash
curl --location 'http://192.168.127.131:8080/pipeline-model-converter/stepsToJenkinsfile' \
--header 'Authorization: Basic YWRtaW46ZDcwNmY0NjMxMzkyNGE5MWFjODk1OGNmZTZhYjc5MGU=' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'json={
  "name":"steps",
  "arguments":[ ],
  "children":[
    {
      "name":"echo",
      "arguments":[
        {
          "key":"message",
          "value":{
            "isLiteral":true,
            "value":"Hello World"
          }
        }
      ]
    }
  ]
}'
```
- response

```json
{
    "status": "ok",
    "data": {
        "result": "success",
        "jenkinsfile": "steps() {\necho 'Hello World'\n}\n"
    }
}
```

# validate

```bash
curl --location 'http://192.168.127.131:8080/pipeline-model-converter/validate' \
--header 'Authorization: Basic YWRtaW46ZDcwNmY0NjMxMzkyNGE5MWFjODk1OGNmZTZhYjc5MGU=' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'jenkinsfile=pipeline {
    agent any

    stages {
        stage('\''Hello'\'') {
            steps {
                echo '\''Hello World'\''
            }
        }
    }
}
'
```

- response

```text
Jenkinsfile successfully validated.
```
