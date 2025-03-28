VERSION:=`cat ./VERSION`
COMMIT:=`git describe --dirty=+WiP --always`
APP:=hellosrcdist

all: version.go
	go build -v -ldflags "-X app.version=$(VERSION)-$(COMMIT)" -o ./bin/ ./cmd/...

.PHONY: test
test: version.go
	go test ./app/...

.PHONY: coverage
coverage: version.go
	go test -coverprofile=.coverage.html ./app/...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	-rm -rf ./bin/
	-rm .coverage.html
	-rm app/version.go

version.go:
	echo "package app" > ./app/version.go
	echo "const Name = \"h$(APP)\"" >> ./app/version.go
	echo "const Version = \"$(VERSION)-$(COMMIT)\"" >> ./app/version.go

srcdist:
	-mkdir dist
	git archive --format=tar.gz --prefix=$(APP)-$(VERSION) -o dist/$(APP)-$(VERSION).tar.gz HEAD
