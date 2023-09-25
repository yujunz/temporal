[![Build status](https://badge.buildkite.com/98c157ed502d55722ef7f28e6470aa20702c76d6989a0a5a89.svg?branch=main)](https://buildkite.com/temporal/temporal-public)
[![Coverage Status](https://coveralls.io/repos/github/temporalio/temporal/badge.svg?branch=main)](https://coveralls.io/github/temporalio/temporal?branch=main)
[![Discourse](https://img.shields.io/static/v1?label=Discourse&message=Get%20Help&color=informational)](https://community.temporal.io)
[![Go Report Card][go-report-image]][go-report-url]

[go-report-image]: https://goreportcard.com/badge/github.com/temporalio/temporal
[go-report-url]: https://goreportcard.com/report/github.com/temporalio/temporal

# Temporal

Temporal is a durable execution and service orchestration platform that enables developers to build scalable applications without sacrificing productivity or reliability.
The Temporal server executes units of application logic called Workflows in a resilient manner that automatically handles intermittent failures, and retries failed operations.

Temporal is a mature technology that originated as a fork of Uber's Cadence.
It is developed by [Temporal Technologies](https://temporal.io/), a startup by the creators of Cadence.

[![image](https://user-images.githubusercontent.com/12602502/136433917-98abe0d7-4f81-4f97-9b11-62b331c76608.png)](http://www.youtube.com/watch?v=f-18XztyN6c 'Temporal')

# High-level view

At the highest level, the components of a system using Temporal fall into two categories:

- **User-hosted processes**

  - The user's application communicates with the Temporal server using one of the [Temporal SDKs](https://docs.temporal.io/dev-guide).
  - The user runs Temporal [Worker](https://docs.temporal.io/workers) processes. These also use one of the Temporal SDKs and host the user's [Workflow](https://docs.temporal.io/workflows) and [Activity](https://docs.temporal.io/activities) code.

- **Temporal Server**<br>
  Users can host and operate the Temporal server and its database themselves, or use [Temporal Cloud](https://temporal.io/cloud).

Learn more about Temporal at [docs.temporal.io](https://docs.temporal.io).

For description of Temporal internals see [ARCHITECTURE.md](./ARCHITECTURE.md)

## Getting Started

### Download and Start Temporal Server Locally

Execute the following commands to start a pre-built image along with all the dependencies.

```bash
brew install temporal
temporal server start-dev
```

Refer to [Temporal CLI](https://docs.temporal.io/cli/#installation) documentation for more installation options.

### Run the Samples

Clone or download samples for [Go](https://github.com/temporalio/samples-go) or [Java](https://github.com/temporalio/samples-java) and run them with the local Temporal server.
We have a number of [HelloWorld type scenarios](https://github.com/temporalio/samples-java#helloworld) available, as well as more advanced ones. Note that the sets of samples are currently different between Go and Java.

### Use CLI

Use [Temporal CLI](https://docs.temporal.io/cli/) to interact with the running Temporal server.

```bash
temporal operator namespace list
temporal workflow list
```

### Use Temporal Web UI

Try [Temporal Web UI](https://docs.temporal.io/web-ui) by opening [http://localhost:8233](http://localhost:8233) for viewing your sample workflows executing on Temporal.

## Repository

This repository contains the source code of the Temporal server. To implement Workflows, Activities and Workers, use one of the [supported languages](https://docs.temporal.io/dev-guide/).

## Contributing

We'd love your help in making Temporal great. Please review our [contribution guide](CONTRIBUTING.md).

If you'd like to work on or propose a new feature, first peruse [feature requests](https://community.temporal.io/c/feature-requests/6) and our [proposals repo](https://github.com/temporalio/proposals) to discover existing active and accepted proposals.

Feel free to join the Temporal communty [forum](https://community.temporal.io) or [Slack](https://t.mp/slack) to start a discussion or check if a feature has already been discussed.
Once you're sure the proposal is not covered elsewhere, please follow our [proposal instructions](https://github.com/temporalio/proposals#creating-a-new-proposal) or submit a [feature request](https://community.temporal.io/c/feature-requests/6).

## License

[MIT License](https://github.com/temporalio/temporal/blob/main/LICENSE)
