MODULE_NAME = rate-limitter

BUILD_PACKAGE = github.com/iwasherd/go-rate-limiter
BUILD_DIR = $(PWD)/bin
BUILD_SHA = $$(git rev-parse --short HEAD)
BUILD_TIME = $$(date -u +"%Y/%m/%d %T")
BUILD_VERSION = $$(cat version)


build: vet
	CGO_ENABLED=0 go build -ldflags="-w -s -X '$(BUILD_PACKAGE).VersionNumber=$(BUILD_VERSION)' -X '$(BUILD_PACKAGE).GitSHA=$(BUILD_SHA)'  -X '$(BUILD_PACKAGE).Time=$(BUILD_TIME)'" -o $(BUILD_DIR)/$(MODULE_NAME) .

vet:
	go vet .
	gosec -quiet .
	govulncheck -show verbose .
	staticcheck .

test:
	go test -v -timeout 30s -count=1 -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

install-tools:
	# Profiling
	go get -u github.com/google/pprof@latest
	go install github.com/google/pprof@latest

	# WebUI for Code Coverage
	go get -u github.com/smartystreets/goconvey@latest
	go install github.com/smartystreets/goconvey@latest

	# Security scanning tools
	go get -u github.com/securego/gosec/v2/cmd/gosec@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest

	go get -u golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest

	# Linting and Formatting
	go get -u golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/cmd/goimports@latest

	go get -u honnef.co/go/tools/cmd/staticcheck@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

	go get -u mvdan.cc/gofumpt@latest
	go install mvdan.cc/gofumpt@latest