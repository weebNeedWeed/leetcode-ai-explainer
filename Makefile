.PHONY: cache
cache:
	go run ./cmd/cache

.PHONY: build-api
build-api:
	docker build -t go-api -f api.Dockerfile .

.PHONY: build-react
build-react:
	docker build -t react -f react.Dockerfile .