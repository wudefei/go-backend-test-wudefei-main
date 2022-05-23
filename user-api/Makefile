fileList := `go list ./...`
goFiles := `find . -name '*.go' | grep -v _test.go`

build := go build -o build/user-api -i cmd/main.go
run  := build/user-api -conf config/conf.yaml

.DEFAULT_GOAL = all
.PHONY: all build cmt errcheck format gc local revive run test vet

all:vet revive test build

cmt: ## 添加缺失的评论
	@hash gocmt 2>&- || go get -u github.com/Gnouc/gocmt
	gocmt -i -t "TODO: comments" $(goFiles) 
	## 要写注释的文件
	@grep -R "TODO: comments" * |grep -v makefile

build: ##go build
	$(build)

errcheck: ## 检查是否处理error
	@hash errcheck 2>&- || go get -u github.com/kisielk/errcheck
	errcheck ./...

format : ## 代码格式化
	@hash goreturns 2>&- || go get -u github.com/sqs/goreturns
	goreturns -w $(goFiles)

gc: ## 垃圾回收
	$(build)
	GODEBUG='gctrace=1' $(run)

local: ## 本地运行，build再运行
	$(build)
	$(run)

run: ##go run 
	$(run)
main:
	go run cmd/main.go

## 代码质量，代码优化
revive: 
	@hash revive 2>&- || go get -u github.com/mgechev/revive
	revive -formatter stylish ./...

test: ## go测试
	go test $(fileList)

## 静态检查，不好的习惯或潜在bug
## 如提前return，有不可达到代码
vet: 
	go vet $(fileList)

help:
	@echo -e "make帮助： \n\
		make build: go build \n\
		make cmt: 添加注释 \n\
		make errcheck: 检查error处理 \n\
		make format: 代码goreturns格式化 \n\
		make gc: 观察垃圾回收 \n\
		make local: 本地运行 \n\
		make revive: 代码优化 \n\
		make run: go run \n\
		make test: go test \n\
		make vet: 代码检查 \n\
		\n\
		make all: vet + revive + format + test + build \n\
		make: 等价 make all\n\
	"