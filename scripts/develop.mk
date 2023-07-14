# generate 缩写
.PHONY: gen
gen:
	go mod tidy
	GOFLAGS=-mod=mod go generate ./...
	make fmt

# fmt 格式化代码
fmt:
	gofumpt -w .

# lint 代码检查
lint:
	make fmt
	golangci-lint run  ./...

# 运行程序http
.PHONY: http
http:
	go run main.go server http
