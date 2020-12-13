test:
	go test -v ./...

start: 
	docker-compose up -d
	go run main.go

stop:
	docker-compose down