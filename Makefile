DB_NAME=supplybooster-db

compile:
	sh scripts/compile.sh

build-docker-image:
	sh scripts/docker-build.sh

dev-local-env:
	echo "Docker copmose is spawning up services"