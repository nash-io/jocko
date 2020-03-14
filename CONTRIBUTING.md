# Contributing to Jocko

This repo is a fork of a fork of [Travis Jocko](https://github.com/nash-io/jocko), our goal is to maintain and update the project while Travis is on his hiatus.

When creating a pull-request you should:

- **Open an issue first**: 
    - Confirm that the change will be accepted
    - Describe in bullet points what your change does and why it's useful
    - Link the equivalent code from the [Kafka repo](https://github.com/apache/kafka) or [docs](https://kafka.apache.org/documentation).
- **Fork and clone the repo**:
    ```
    git clone git@github.com:your-username/jocko.git
    ```

- **Get the deps**:
    ```
    dep ensure
    ``` 
    (If you don't have dep, run: `go get -u github.com/golang/dep/cmd/dep`)

- **Check the tests pass**:
    ```
    go test ./...
    ```

- **Make your change**
- **Write tests and check they pass**
- **Lint your code**: Use `gofmt`, `golint`, and `govet` to clean up your code
- **Start your commit message with a verb**: your commit message must start a lowercase verb such as: "add", "fix", "refactor", "remove"
- **Reference the issue**: Reference your issue N by including "closes #N" in the commit message


Thanks!
