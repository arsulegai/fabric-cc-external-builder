# Fabric Chaincode External Builder

The external builder is used to run the chaincode as a service. This
chaincode builder is written in Go. The builder supports reading the
chaincode TLS parameters at runtime. Consumes them through `peer`'s
chaincode builder parameters configured in the `core.yaml` file.

The example here illustrates how to make use of external builder pattern
to run chaincode as a service. The same concept can be used to run chaincode
brought up by the peer node. A custom base image can be used for chaincode
building.

## Build & Run

A [Makefile](./Makefile) is provided to ease up the build option. In
order to run these external builder binaries, detailed instructions are
given later in this document.

Run the following command to generate the required artefact that can be run
with a Hyperledger Fabric peer node alpine image. 

```shell
make build
```

In case you would like to copy these artefacts to the peer container.
Build these for alpine environment. You can do that by

```shell
make server-build-alpine
```

## Hyperledger Fabric Documentation

Refer to the section
[External Builder and Launcher](https://hyperledger-fabric.readthedocs.io/en/release-2.2/cc_launcher.html#external-builder-and-launcher-api)
in the Hyperledger Fabric documentation.

Read the reference documentation in
[chaincode builder and launcher details](./docs/builder.md).

## Environment Variables

### Chaincode as a Server

These environment variables are to be set in the peer's process environment.
Pass them through to the builder in the `core.yaml` configuration.

| ENV Variable | Purpose |
|:--- |:--- |
| CORE_PEER_CC_BUILDER_LOG_LEVEL | Log level for the builder |
| CORE_CHAINCODE_ADDRESS | Chaincode server's address |
| CORE_CHAINCODE_TIMEOUT | Timeout parameter |
| CORE_CHAINCODE_TLS_REQUIRED | If TLS is required for chaincode connection (`true` or `false` ) |
| CORE_CHAINCODE_ROOT_CERT | Which Chaincode server to trust |
| CORE_CHAINCODE_CLIENT_AUTH_REQUIRED | Is mTLS enabled (`true` or `false`) |
| CORE_CHAINCODE_TLS_CLIENT_KEY | Peer's client key if mTLS is enabled |
| CORE_CHAINCODE_TLS_CLIENT_CERT | Peer's Client certificate if mTLS is enabled |

## Logging

Chaincode server builder logs are captured by the peer process. Remember to
print any of your log to the `STDERR` pipe.
