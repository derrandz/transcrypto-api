#!/usr/bin/env bash 

docker run \
    --name txsigner-summitto \
    -p 8080:8080 \
    -itd \
    summitto/txsigner:1.0