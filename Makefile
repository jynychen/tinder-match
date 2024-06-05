.PHONY: clean
clean:
	rm -rf bin/


.PHONY: test
test:
	go test -v ./...


.PHONY: lint
lint:
	golangci-lint run


.PHONY: docs
docs:
	swag init -g cmd/server/main.go --pd -o docs/api

.PHONY: build-bin
build-bin:
	go build -o bin/server cmd/server/main.go
	@echo "Binary file is created at ./bin/server"

.PHONY: run-bin
run-bin: build-bin
	./bin/server


.PHONY: build-image
build-image:
	docker build . -t tinder-match
	@echo "Docker image is created with name tinder-match"

.PHONY: run-docker
run-docker: build-image
	docker run -p 8080:8080 tinder-match


.PHONY: k6-test
k6-test:
	k6 run k6/match.js