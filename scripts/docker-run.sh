#!/usr/bin/env bash 

docker run \
    --name txsigner-transcrypto \
    -p 8080:8080 \
    -itd \
    transcrypto/txsigner:1.0