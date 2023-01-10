.PHONY: all build lint test prepare up_local generate run lint-no-fix


OUTPUT_EXE:=./.build/be
GOSRC := ${GOPATH}/src

BUILD_TIME:=$$(TZ=UTC date +%FT%T%z)
GIT_TAG:=$$(git describe --tags --always)
GIT_COMMIT:=$$(git describe --always)


all: build

build:
	go build -installsuffix 'static' -tags dev -ldflags="-X be/pkg/version.Version=$(GIT_TAG) -X be/pkg/version.BuildTime=$(BUILD_TIME) -X be/pkg/version.GitCommit=$(GIT_COMMIT)" -o $(OUTPUT_EXE) cmd/be/main.go

run:
	source local/vars.sh && ./$(OUTPUT_EXE) run

lint-no-fix:
	golangci-lint run --fix --allow-parallel-runners --build-tags=integration --config ./build/ci/.golangci.yml ./...

lint:
	gofmt -w ./
	make lint-no-fix

test:
	source local/vars.sh && go test ./... -count=1

up_local:
	COMPOSE_PROJECT_NAME=office docker-compose -f local/docker-compose.infra.yaml up -d

clean:
	docker-compose -f local/docker-compose.infra.yaml stop
	rm tn-profile


