DB_NAME=supplybooster-db

compile:
	sh scripts/compile.sh

test:
	APP_ENV=TEST go test -v ./pkg/... ./shared/... -bench=. -benchmem -race \
	-coverprofile=coverage.txt -covermode=atomic

build-docker-image:
	sh scripts/docker-build.sh

run-docker-container:
	sh scripts/docker-run.sh

dev-local-env:
	echo "Docker copmose is spawning up services"