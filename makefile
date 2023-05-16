run-infra:
	docker-compose -f deploy/docker-compose.yaml up -d

run:
	go run cmd/main.go

migrate-up:
	go run cmd/main.go -migrate=up -migrate-search=up
	
migrate-down:
	go run cmd/main.go -migrate=down

test-benchmark:
	go test -v ./usecase/api/products -run Benchmark -bench=. -benchtime=10s -benchmem


run-load-test:
	docker run --rm --network=host -i grafana/k6 run - <deploy/loadtest.js

build:
	docker build -f deploy/dockerfile -t golang-research .

build-and-run:
	make build
	docker run -d --name golang-research \
		-p 8888:8888 \
		--add-host=host.docker.internal:host-gateway \
		--network host \
		-e DB_HOST=host.docker.internal \
		-e DB_PORT=6432 \
		-e DB_USER=user-search \
		-e DB_PASS=user-pass \
		-e DB_NAME=search \
		-e MEILI_HOST=http://host.docker.internal:7700 \
		-e MEILI_APIKEY= \
		-e TYPESENSE_HOST=http://host.docker.internal:8108 \
		-e TYPESENSE_APIKEY= \
		golang-research
