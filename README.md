# lets-gnomock

## set env vars
```shell script
$ direnv allow
```

## set up
```shell
$ docker-compose up
$ make migrateup
$ go run ./cmd/server/main.go
```

## curl api
```shell
$ curl -s -XPOST -d "{\"username\":\"foo\"}" http://localhost:8080/create
```

You can use this If you have `jo` cmd.
```shell
$ curl -s -XPOST -d "$(jo username=foo)" http://localhost:8080/create
```

## Play Gnomock TEST
```shell
$ go test -v -race ./... -run TestUserService_Create
```