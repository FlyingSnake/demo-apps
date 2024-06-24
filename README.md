# demo-apps

## Run Docker Compose

```bash
docker-compose up -d
```

## Request API

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
```

### Docker Hub Images

| App Name           | Image Repo                                                                                        | Tag   |
| ------------------ | ------------------------------------------------------------------------------------------------- | ----- |
| mysql              | [flyingsnake12/demo-mysql](https://hub.docker.com/r/flyingsnake12/demo-mysql)                     | 1.0.0 |
| java-sb-app        | [flyingsnake12/demo-java-springboot](https://hub.docker.com/r/flyingsnake12/demo-java-springboot) | 1.0.1 |
| nodejs-express-app | [flyingsnake12/demo-nodejs-express](https://hub.docker.com/r/flyingsnake12/demo-nodejs-express)   | 1.0.1 |
| php-flight-app     | [flyingsnake12/demo-php-flight](https://hub.docker.com/r/flyingsnake12/demo-php-flight)           | 1.0.1 |
| go-echo-app        | [flyingsnake12/demo-golang-echo](https://hub.docker.com/r/flyingsnake12/demo-golang-echo)         | 1.0.1 |
