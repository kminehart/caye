# Caye: Code-configured CI

# Disclaimer: Still a heavy work in progress. I'm a huge fan of this idea and plan to continue to expand it. Don't expect anything in here at the moment to be production ready. Ideas and contributions are always welcome!

## Configuration based CI

### Benefits
* Approachable
* Easy to read

### Problems
* Configuration is unclear (Requires browsing the docs like a chump)
* Inflexible

## Code configured CI

* Import CI libraries as a dependency
* Communicate with Caye server using Caye library (which is mostly a wrapper around the gRPC client).
* Use within any "runner" system

## Compare

Here's a basic example from a CI system that's "configuration-first", or expects a yaml configuration.

Some problems that a user and a developer may encounter:
* The schema is unclear.
  * What keys are available where?
  * What types are expected?
  * What values are acceptable?
* As a user, how do I run this without creating and pushing a new commit?
* Reusability relies on the very uncommon, often unsupported yaml pointer syntax

### Configuration-first

```yaml
defaults: &defaults
  working_directory: ~/repo
  docker:
    - image: circleci/golang:1.13

version: 2
jobs:
  test:
    <<: *defaults
    steps:
      - checkout
      - run:
          name: Run tests
          command: go test ./...
  build:
    <<: *defaults
    steps:
      - checkout
      - run:
          name: Build
          command: go build -o ./bin/app ./cmd/app 
      - persist_to_workspace: root: ~/repo
          paths: [ bin ]
workflows:
  version: 2
  test:
    jobs:
      - test
      - build:
          requires: [ test ]
```

### Code-first

This is a Go project. It's only natural that the definition for testing & building this project is another Go program.

```go
package main

import "github.com/kminehart/caye"

func main() {
  client := caye.Default()
  defer client.Done()

  golang := client.Languages().Go()

  if err := golang.Test("./..."); err != nil {
    client.Fatal(err)
  }

  err := golang.Build(&caye.BuildOptions{
    Target: "./cmd/app",
    Output: "./bin/app",
  })

  if err != nil {
    client.Fatal(err)
  }
}
```

Benefits of this approach:

* It's clear what options are available to the developer.
  * The developer likely already has tools (Autocompletion, definition jumping, native documentation from comments) for discovering available options
* Running this is completely isolated. A user simply has to run this application like they would any other one.
* This can be used with any CI system like Circle.
* Reusability is achieved through the languages standard paradigms.
