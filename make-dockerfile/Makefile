REGISTRY ?= registry.cn-hangzhou.aliyuncs.com/bocloud

DEV_REGISTRY ?= deploy.bocloud.k8s:40443/aios

IMAGE_NAME ?= aios-automl

BIN_NAME = aios-automl

VERSION ?= v1.0

RELEASE_TAG = $(VERSION)

TIME_STAMP = $(shell date +%Y%m%d%H%M)

# 版本镜像指定镜像仓库
RELEASE_IMG ?= $(REGISTRY)/$(IMAGE_NAME):$(RELEASE_TAG)

COMMIT_ID=$(shell git rev-parse HEAD)
BUILD_TIME=$(shell date +'%Y-%m-%dT%H:%M')
HOST_ARCH=$(shell go env GOHOSTARCH)
HOST_OS=$(shell go env GOHOSTOS)

.PHONY: all
all: release


EMPTY :=
SPACE := $(EMPTY) $(EMPTY)
COMMA := ,
PLATFORMS ?=amd64 arm64

#--push表示将镜像推送到远程仓库。
#--progress plain表示显示简单的进度信息。
#--platform $(subst $(SPACE),$(COMMA),$(PLATFORMS))表示构建多个平台的镜像，如arm64和amd64。

.PHONY: release  # Build a multi-arch docker image
release: bin release-binaries record-build-info  ## 同时打包arm64，amd64镜像并上传到仓库.使用前先修改 REGISTRY参数
	echo "Building and pushing image..."
	@docker buildx build \
		--push \
		--progress plain \
		--platform $(subst $(SPACE),$(COMMA),$(PLATFORMS)) \
		--build-arg COMMIT_ID=${COMMIT_ID} \
		--build-arg BUILD_TIME=${BUILD_TIME} \
		--build-arg VERSION=${VERSION} \
		-t $(RELEASE_IMG) .
	$(MAKE) migrate
	# release external image: $(RELEASE_IMG)
	echo "release image time: $(TIME_STAMP)" >> BUILD_INFO
	echo "release external image: $(RELEASE_IMG)" >> BUILD_INFO


# -trimpath - 在运行时删除二进制文件中的符号表信息，以减小文件大小。
# -ldflags="-s -w" - 这是链接器标志，用于控制生成的可执行文件的行为。-s表示去除调试信息，-w表示去除所有符号信息(如函数名、变量名等)。
# -X 这些是传递给二进制文件的自定义属性。这些属性可以在运行时通过反射进行访问和修改。
.PHONY: build
build: ## Build manager binary.
	CGO_ENABLED=0 GOOS=linux GOARCH=${HOST_ARCH} \
        go build  \
		-trimpath \
        -ldflags="-s -w \
                  -X main.GitCommitID=${COMMIT_ID}  \
                  -X main.Architecture=${HOST_OS}/${HOST_ARCH} \
                  -X main.BuildTime=${BUILD_TIME} \
                  -X main.Version=${VERSION}" \
        -o bin/$(BIN_NAME)-${HOST_ARCH} .

.PHONY: release-binaries
release-binaries: clean
	$(foreach PLATFORM,$(PLATFORMS), echo -n "$(PLATFORM)..."; make build HOST_ARCH=$(PLATFORM);)

#.PHONY: migrate
#migrate: migrate-image migrate-image-dev
#
#.PHONY: migrate-image
#migrate-image: ## Migrate image to BKE registry
#	@bke registry sync --dest-tls-verify --source $(RELEASE_IMG) --target $(BKE_REGISTRY_IMG) --multi-arch
#
#.PHONY: migrate-image-dev
#migrate-image-dev: ## Migrate image to local registry
#	@bke registry sync --dest-tls-verify --source $(RELEASE_IMG) --target $(LOCAL_DEV_IMG) --multi-arch


.PHONY: record-build-info
record-build-info:
	echo -e "\n" >> BUILD_INFO
	echo "🤯 Version=${VERSION}" >> BUILD_INFO
	echo "🤔 GitCommitId=${COMMIT_ID}" >> BUILD_INFO
	echo "👉 Architecture=${HOST_ARCH}" >> BUILD_INFO
	echo "⏲ BuildTime=${BUILD_TIME}" >> BUILD_INFO

.PHONY: bin
bin:
	if [ ! -d bin ]; then mkdir bin; fi

.PHONY: clean
clean:
	rm -f bin/$(BIN_NAME)*