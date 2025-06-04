export PATH := $(PATH):`go env GOPATH`/bin
export GO111MODULE=on
export BASE_DIR := $(shell pwd)
LDFLAGS := -s -w
os-archs=darwin:amd64 darwin:arm64 freebsd:amd64 linux:amd64 linux:arm:7 linux:arm:5 linux:arm64 windows:amd64 windows:arm64 linux:mips64 linux:mips64le linux:mips:softfloat linux:mipsle:softfloat linux:riscv64


# Binary file names.
BINARY_NAME=lc-kit-web

# Build parameters.
DIST_PATH=${BASE_DIR}/.dist

.PHONY: build
build: build-web build-go


build-web:
	pnpm i
	pnpm run build
	rm -rf server/ui
	mv dist server/ui

build-go:
	@rm -rf ${DIST_PATH}
	@mkdir -p ${DIST_PATH}
	@cp -r server/ui ${DIST_PATH}/
	@cd server && \
	go mod tidy && \
	$(foreach n, $(os-archs), \
		os=$(shell echo "$(n)" | cut -d : -f 1); \
		arch=$(shell echo "$(n)" | cut -d : -f 2); \
		extra=$(shell echo "$(n)" | cut -d : -f 3); \
		flags=''; \
		target_suffix=$${os}_$${arch}; \
		if [ "$${os}" = "linux" ] && [ "$${arch}" = "arm" ] && [ "$${extra}" != "" ] ; then \
			if [ "$${extra}" = "7" ]; then \
				flags=GOARM=7; \
				target_suffix=$${os}_arm_hf; \
			elif [ "$${extra}" = "5" ]; then \
				flags=GOARM=5; \
				target_suffix=$${os}_arm; \
			fi; \
		elif [ "$${os}" = "linux" ] && ([ "$${arch}" = "mips" ] || [ "$${arch}" = "mipsle" ]) && [ "$${extra}" != "" ] ; then \
		    flags=GOMIPS=$${extra}; \
		fi; \
		echo "Build $${os}-$${arch}$${extra:+ ($${extra})}..."; \
		env CGO_ENABLED=0 GOOS=$${os} GOARCH=$${arch} $${flags} go build -trimpath -ldflags "$(LDFLAGS)" -o ${DIST_PATH}/${BINARY_NAME}_$${target_suffix} server.go; \
		echo "Build $${os}-$${arch}$${extra:+ ($${extra})} done"; \
	)
	@mv ${DIST_PATH}/${BINARY_NAME}_windows_amd64 ${DIST_PATH}/${BINARY_NAME}_windows_amd64.exe
	@mv ${DIST_PATH}/${BINARY_NAME}_windows_arm64 ${DIST_PATH}/${BINARY_NAME}_windows_arm64.exe

