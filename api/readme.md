# What

Service that serve HTTP.

# Usage

```
go run main.go
```

or

```
./run.sh
```

# Endpoints

```sh
curl "localhost:3999/health"
```

```sh
curl "localhost:3999/login" -d '{"username":"mo","password":"mo"}'
```

```sh
curl "localhost:3999/logout" -H "Cookie: token=afb0bd5b-2f66-4b76-b376-26f2f5154385; expires=Sun, 12 Jun 2022 09:26:52 GMT" -d ''
```

```sh
curl "localhost:3999/sum?term1=1&term2=-1"
```

```sh
curl "localhost:3999/sum?term1=1&term2=-1" -H "Cookie: token=kenasd;expires=Mon, 27 Jun 2022 00:00:00 GMT"
```

# Dev note

CONTAINERIZING

```sh
docker build . -t grpc-microservice-example-api:1.0.0
```

```sh
docker run -it --rm --name=api-ex grpc-microservice-example-api:latest
```
