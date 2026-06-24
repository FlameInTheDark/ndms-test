NDM Systems test task

## Run

Simple run (default url `localhost:8080`):
```shell
go run ./cmd/mq
```

Run with custom URL and Port:
```shell
go run ./cmd/mq -url localhost:6000
```

## Usage

Consume from a topic:
```shell
curl --request GET \
  --url 'http://localhost:8080/pet'
```

With timeout seconds:
```shell
curl --request GET \
  --url 'http://localhost:8080/pet?timeout=20'
```

Produce message:
```shell
curl --request POST \
  --url 'http://localhost:8080/pet?v=dog'
```