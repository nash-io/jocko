# Jocko

[![Build Status](https://travis-ci.org/nash-io/jocko.svg?branch=master)](https://travis-ci.org/nash-io/jocko) [![codecov](https://codecov.io/gh/nash-io/jocko/branch/master/graph/badge.svg)](https://codecov.io/gh/nash-io/jocko) [![Go Report Card](https://goreportcard.com/badge/github.com/nash-io/jocko)](https://goreportcard.com/report/github.com/nash-io/jocko)

Distributed commit log service in Go that is wire compatible with Kafka.

Created by [@travisjeffery](https://github.com/travisjeffery), continued by [nash](https://nash.io).

## Goals:

- Protocol compatible with Kafka so Kafka clients and services work with Jocko
- Distribute a single binary
- Use Serf for discovery, Raft for consensus (and remove the need to run ZooKeeper)
- Simpler configuration settings

## TODO

- [ ] Map features missing
- [ ] Extensive protocol compliance test suit
- [ ] Update all dependencies and code to up-to-date Go (as of MAR2020)

## Reading

- [How Jocko's built-in service discovery and consensus works](https://medium.com/the-hoard/building-a-kafka-that-doesnt-depend-on-zookeeper-2c4701b6e961#.uamxtq1yz)
- [How Jocko's (and Kafka's) storage internals work](https://medium.com/the-hoard/how-kafkas-storage-internals-work-3a29b02e026#.qfbssm978)

## Project Layout

```
├── broker        broker subsystem
├── cmd           commands
│   └── jocko     command to run a Jocko broker and manage topics
├── commitlog     low-level commit log implementation
├── examples      examples running/using Jocko
│   ├── cluster   example booting up a 3-broker Jocko cluster
│   └── sarama    example producing/consuming with Sarama
├── protocol      golang implementation of Kafka's protocol
├── prometheus    wrapper around Prometheus' client lib to handle metrics
├── server        API subsystem
└── testutil      test utils
    └── mock      mocks of the various subsystems
```

## Building

### Local

1. Clone Jocko

    ```
    $ go get github.com/nash-io/jocko
    ```

1. Build Jocko

    ```
    $ cd $GOPATH/src/github.com/nash-io/jocko
    $ make build
    ```

    (If you see an error about `dep` not being found, ensure that
    `$GOPATH/bin` is in your `PATH`)

### Docker

`docker build -t nash-io/jocko:latest .`

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches and the contribution workflow.

## License

Jocko is under the MIT license, see the [LICENSE](LICENSE) file for details.
