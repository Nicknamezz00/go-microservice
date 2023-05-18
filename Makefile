apps = 'products' 'details' 'ratings' 'reviews'
.PHONY: build
build:
	for app in $(apps) ;\
	do \
		GOOS=linux GOARCH="amd64" go build -o dist/$$app-linux-amd64 ./cmd/$$app/; \
		GOOS=darwin GOARCH="amd64" go build -o dist/$$app-darwin-amd64 ./cmd/$$app/; \
	done

.PHONY: docker
docker-compose: build
	docker-compose -f deployments/docker-compose.yml up --build -d