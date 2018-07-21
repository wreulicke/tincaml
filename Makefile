ARCH := 386 amd64
OS := linux darwin windows

VERSION="main.version=0.1.0"
HASH=\"main.hash=$(shell git rev-parse --verify HEAD)\"
BUILDTIME=\"main.buildtime=$(shell date '+%Y/%m/%d %H:%M:%S %Z')\"
GOVERSION=\"main.goversion=$(shell go version)\"

preinstall: 
	go get golang.org/x/tools/cmd/goyacc

setup_ci: 
	go get github.com/mitchellh/gox
	go get github.com/jstemmer/go-junit-report
	go get github.com/haya14busa/goverage
	go get golang.org/x/tools/cmd/cover
	go get -u github.com/golang/lint/golint
	go get -u github.com/golang/dep/cmd/dep

status:
	dep status

install:
	dep ensure

update:
	dep ensure -update

lint: 
	golint ./...

build:
	go generate ./...

test:
	go test ./...

package: build
	gox -os="$(OS)" -arch="$(ARCH)" -ldflags="-X $(VERSION) -X $(HASH) -X $(BUILDTIME) -X $(GOVERSION)" -output "dist/{{.OS}}_{{.Arch}}/{{.Dir}}"