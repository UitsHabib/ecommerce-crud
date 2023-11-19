serve:
	go run main.go serve-rest

seed:
	go run main.go seed

test:
	go test -gcflags=-l -cover -race ./...

perfect_go_get:
	go get package_module@commit_hash

deploy_redis:
	docker run --name redis -p 6379:6379 redis redis-server --requirepass "12345"