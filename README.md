# transcrypto-api

### Information

- `Operating system used`: Ubuntu 20.04 LTS Debian 5.4.0-52-generic
- `GO Version`: go version go1.15.4 linux/amd64

### Abstraction

**Description:**

This repository contains all the transcrypto-api services in a monorepo manner.

List of implemented services:
    * txsigner (microservice)
        * protocol: REST over HTTP

Incoming:
    * txsigner (microservice)
        * protocol: gRPC

**Application structure:**

The application structure is split for maximum re-usability, the structure is
explained as following:

- `cmd` : Is the entrypoint package for all the applications, in a nutshell, the
main package for any service that is exposed behind an interface (API, CLI,...).

- `pkg` : The actual applications/business logic, each folder inside this folder
represent a standalone service.

- `tools` : Contains one-off processes or cli tools that the project might depend on

- `shared` : The shared components and libraries which are used by all the
systems, contains a tiny framework-like components that are being reused among all the services, e.g.
logging, http transport, grpc, etc...

The application uses a `Makefile` to define scripts that are commonly used on the application, like `make compile`  To see all the available commands, take a look at `Makefile`.

### How to run

**Prerequisite**

In order to run the system(s) on your machine natively, you have to install:
- [GoLang](https://golang.org/dl/), along with [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH), `GOROOT` and `GOBIN` configurations
(check the original documentation for installation tutorial).
- [docker](https://www.docker.com/get-started)
<!-- - [docker-compose](https://docs.docker.com/compose/install/) -->

#### Run W/O Docker

After a successful setup of the above, you should see some results when you
issue the following commands:

```bash
$ go
# some results...
$ docker
# some results...
```

After you've ensured that the dependencies are all working fine, lets start
with the installation steps.

**Installation**

Install the project dependencies, in order to do so, issue the following commands:
```bash
# install dependencies
$ go mod download
```

```bash
$ make test
```

The txsigner is following the [12factor-code](https://12factor.net/) principles
(reading further about them will ease the understanding and reasoning behind the
approaches and patterns used), and in order to inject the configurations to the
system you're running, you need to export them as an environment variables.

Below a list of environment variables that need to be set:
```bash
export PORT=8080 
export DAEMON_PRIVATE_KEY=asdasdlcx12312wqsa
```

You can also make use of a `.env` file

If you would like to generate a private key for your daemon, run the following:
```bash
$ ROOT=tools make compile
$ ./build/bin/key*
```

This command should ouput an ed25519 private/public key pair, make sure to fill the environment variable properly.

After you have finished all the previous steps successfully, you can issue the
following command to run the txsigner service under one server:
```bash
user@host:~$ go run cmd/txsigner/main.go
```

#### Run with Docker

*NOT IMPLEMENTED YET, SAME FOR DOCKER COMPOSE*

1. Generate the key pair (look in the installation section)

```bash
$ go run tools/keys/main.go 
env 8080 OmwQUAAXyKArmg9Cmt5W6u79K0SvmpAk+JXUoqr55yQufTzUhmeMoidMsg66a6PfJGrfRW85Tmab0jAGB4aJew==
Generating ed25519 Private/Public Key Pair:
Public Key: pKPMEKo7E6Ml/914t+S/3pbIow4qwoGYryEjeYA3XIk=
Private Key: iAcnxg2OfgPkuRKefX8JbjrRVUMCdx4Q4OCEE1HvKEyko8wQqjsToyX/3Xi35L/elsijDirCgZivISN5gDdciQ==
```

2. export your environment variables like the following:
```bash
$ export PRIVATE_KEY=iAcnxg2OfgPkuRKefX8JbjrRVUMCdx4Q4OCEE1HvKEyko8wQqjsToyX/3Xi35L/elsijDirCgZivISN5gDdciQ==
```

3. Compile the project
```bash
$ make compile
```

4. Copy the executable name by navigating to `./build/bin/` and copying `txsignerhttp-$-$`
```
$ ls ./build/bin
▶ txservicehttp-linux-amd64
```

5. Export the app name and the tag
```bash
$ export APP_NAME=txservicehttp-linux-amd64
$ export DOCKER_TAG=1.0
```

6. Build the docker image
```
$ make build-docker-image
```

6. Run the docker image
```
$ make run-docker-container
```

### Contribution rules

- Follow GitHub flow for branching https://guides.github.com/introduction/flow/
- Follow [12 factor code](https://12factor.net/)
- [Unix philosophy](https://en.wikipedia.org/wiki/Unix_philosophy)

### What is currently missing? (What was supposed to be in here)
    - [ ] Make use "context" for better context sharing
    - [ ] Proper responses HTTP Codes
    - [ ] e2e test for the service
    - [ ] integration tests for the processor and http layer of the service
    - [ ] Application level logging (log what method was called and how much time it took, could use open tracing for this)
    - [ ] Metrics (could use prometheus)
    - [ ] Proper error handling for a few functions here and there
    - [ ] Build in a docker image for accurate target environment target

### How to quickly test that this is working properly

1- Copy `.env.example` to `.env` file

2- Generate the key pair using:
```bash
$ go run tools/keys/main.go
```

3- Paste the private key value in the `.env` `DAEMON_PRIVATE_KEY` value

4- Run the txsigner service using:
```bash
$ go run cmd/txsigner/main.go
```

5- Open another terminal and start testing using

Retrieve the public key
```bash
$ curl -X GET http://localhost:8080/public_key -H "Content-Type: application/json" -d '{}'  
```

Create a transaction (or transactions)
```bash
$ curl -X PUT http://localhost:8080/transaction -H "Content-Type: application/json" -d '{ "tx": ["/+BBAgM="]} 
```

Copy the id (or IDs) in the response and pase them in the next request replacing the values in the ids array
```bash
curl -X POST \
    http://localhost:8080/signature \
    -H "Content-Type: application/json" \
    -d '{"ids": ["ef90d56c-ff44-4dc2-9ab4-4c52b66cf352"]}'
```
