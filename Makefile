# Binary file names.
BINARY_NAME=lucky

# Build parameters.
DIST_PATH=.dist

### help:					Makefile 帮助
.PHONY: help
help:
	@echo Makefile cmd:
	@echo
	@grep -E '^### [-A-Za-z0-9_]+:' Makefile | sed 's/###/   /'

.PHONY: build
### build:				项目打包
build: build-go

### build-go:			构建 golang 包
build-go:
	rm -rf ${DIST_PATH}
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -trimpath -mod=readonly -v -o $(DIST_PATH)/${BINARY_NAME} main.go
	cp -rf conf $(DIST_PATH)
	cp -rf static $(DIST_PATH)
	cp deployments/run.sh $(DIST_PATH)


.PHONY: migrate
### migrate:			同步数据库生成代码
migrate:
	@export USE_PRE=1 && lc migrate -c tools/gen/config.yml
	@echo "migrate done"

.PHONY: init
### init:				初始化数据
init:
	@echo "init data"
ifdef env
	export APP_CONFIG=conf/app-${env}.yml && go run tools/init/main.go
else
	go run tools/init/main.go
endif
