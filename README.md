# demo-apps

## Run Docker Compose

```bash
docker-compose up -d
```

## Check API

```bash
# java springboot
curl localhost:8080
curl localhost:8080/sleep/3
curl localhost:8080/status/random
curl localhost:8080/exception
curl localhost:8080/users

# nodejs express
curl localhost:8081
curl localhost:8081/sleep/3
curl localhost:8081/status/random
curl localhost:8081/exception
curl localhost:8081/users

# php flight
curl localhost:8082
curl localhost:8082/sleep/3
curl localhost:8082/status/random
curl localhost:8082/exception
curl localhost:8082/users

# go echo
curl localhost:8083
curl localhost:8083/sleep/3
curl localhost:8083/status/random
curl localhost:8083/exception
curl localhost:8083/users

# python flask
curl localhost:8084
curl localhost:8084/sleep/3
curl localhost:8084/status/random
curl localhost:8084/exception
curl localhost:8084/users

# dotnet carter
curl localhost:8085
curl localhost:8085/sleep/3
curl localhost:8085/status/random
curl localhost:8085/exception
curl localhost:8085/users

# via React
curl localhost:8000/java/[PATH]
curl localhost:8000/nodejs/[PATH]
curl localhost:8000/php/[PATH]
curl localhost:8000/golang/[PATH]
curl localhost:8000/python/[PATH]
curl localhost:8000/dotnet/[PATH]
```

### Docker Hub Images

| App Name           | Image Repo                                                                                        | Tag          |
| ------------------ | ------------------------------------------------------------------------------------------------- | ------------ |
| mysql              | [flyingsnake12/demo-mysql](https://hub.docker.com/r/flyingsnake12/demo-mysql)                     | 1.0.0        |
| java-sb-app        | [flyingsnake12/demo-java-springboot](https://hub.docker.com/r/flyingsnake12/demo-java-springboot) | 1.0.2        |
| nodejs-express-app | [flyingsnake12/demo-nodejs-express](https://hub.docker.com/r/flyingsnake12/demo-nodejs-express)   | 1.0.2        |
| php-flight-app     | [flyingsnake12/demo-php-flight](https://hub.docker.com/r/flyingsnake12/demo-php-flight)           | 1.0.2        |
| go-echo-app        | [flyingsnake12/demo-golang-echo](https://hub.docker.com/r/flyingsnake12/demo-golang-echo)         | 1.0.2        |
| python-flask-app   | [flyingsnake12/demo-python-flask](https://hub.docker.com/r/flyingsnake12/demo-python-flask)       | 1.0.0        |
| dotnet-carter-app  | [flyingsnake12/demo-dotnet-carter](https://hub.docker.com/r/flyingsnake12/demo-dotnet-carter)     | 1.0.0        |
| react-app          | [flyingsnake12/demo-react](https://hub.docker.com/r/flyingsnake12/demo-react)                     | nginx-stable |
