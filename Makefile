DB_NAME=supplybooster-db

compile:
	sh scripts/compile.sh

test:
	go test -v ./pkg/... ./shared/... -bench=. -benchmem -race \
	-coverprofile=coverage.txt -covermode=atomic

build-docker-image:
	sh scripts/docker-build.sh

dev-local-env:
	echo "Docker copmose is spawning up services"