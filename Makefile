
APPNAME=raid-champ-api
GORELEASER_VERSION=v0.129.0


default:
	$(MAKE) build

test:
	go test ./... -v

run:
	$(MAKE) build
	./${APPNAME}

build:
	go build
	$(MAKE) swagger

swagger:
	@swagger version > /dev/null 2>&1 || go get -u github.com/go-swagger/go-swagger/cmd/swagger
	swagger generate spec -o ./swagger.json

clean:
	rm ${APPNAME}


setup:
	rm -rf goreleaser
	mkdir goreleaser
	cd goreleaser && wget https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/goreleaser_Linux_x86_64.tar.gz
	cd goreleaser && tar xvzf goreleaser_Linux_x86_64.tar.gz

release:
	git push --tags
	goreleaser/goreleaser --rm-dist

release-build:
	goreleaser/goreleaser --rm-dist --snapshot


