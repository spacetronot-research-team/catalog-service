# catalog-service

## Database with docker

1. run docker compose

```shell
docker compose up
```

2. run migration

```shell
go run database/migrate/up.go
```

3. connect from db client with credential

```
host=localhost
port=5432
db=boilerplate_catalog
user=user
password=password
```