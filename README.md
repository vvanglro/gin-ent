# gin + ent + migrate + jwt + embed + air
## project run
1. set conf/conf.yaml
2. go generate ./models/ent
3. go run cli.go  (Generate migration file)
4. go run cli.go -w=true (Migration files to database)
5. air -c conf/.air.toml

## Default User
```shell
curl --location --request POST 'http://localhost:8080/api/login' \
--header 'Content-Type: application/json' \
--data-raw '{"username":"admin", "password":"123456"}'
```