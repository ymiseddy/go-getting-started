# go-getting-started

This is the companion code for my [Go Project Setup](https://seddy.com/story/go-project-01)
article.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.24 or later
- Make
- Docker and Docker BuildKit (optional, for building docker images)

## Building

To build the application, run:

```bash
make
```

Or

```bash
make build
```

To build a docker image, run:

```bash
make docker-build
```

To build an exportable tarball of the docker image, run:

```bash
make image
    ```
