PKG_NAME=wmservice
DUSER=aceberg

mod:
	cd src && \
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/WMService && \
	go mod tidy

run:
	cd src && \
	go run .

go-build:
	cd src && \
	go build .

fmt:
	cd src && \
	go fmt ./...

lint:
	cd src && \
	golangci-lint run
	cd src && \
	golint ./...

check: fmt lint

docker-build:
	docker build -t $(DUSER)/$(PKG_NAME) .

docker-run:
	docker rm $(PKG_NAME) || true
	docker run --name $(PKG_NAME) \
	-e "TZ=Asia/Novosibirsk" \
	-v ~/.dockerdata/$(PKG_NAME):/data/$(PKG_NAME) \
	-p 8843:8843 \
	$(DUSER)/$(PKG_NAME)