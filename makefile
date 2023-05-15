run-infra:
	docker-compose -f deploy/docker-compose.yaml up -d

run:
	go run cmd/main.go

migrate-up:
	go run cmd/main.go -migrate=up
	
migrate-down:
	go run cmd/main.go -migrate=down
