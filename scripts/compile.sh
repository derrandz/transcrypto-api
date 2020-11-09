#!/usr/bin/env sh

ROOT=cmd
APP_NAMES=$(ls ${ROOT})

for APP in ${APP_NAMES}; do
#  for GOOS in darwin linux windows; do
  for GOOS in linux; do
#    for GOARCH in 386 amd64; do
    for GOARCH in amd64; do
        echo ${ROOT}/${APP}/*.go
      # build the application statically
      GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -a -tags netgo \
       -ldflags '-s -w -extldflags "-static"' \
       -o build/bin/${APP}-$GOOS-$GOARCH ${ROOT}/${APP}/*.go

       # changing the execution mode
       chmod +x build/bin/${APP}-$GOOS-$GOARCH
      done
    done
done
