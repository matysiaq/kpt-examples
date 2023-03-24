# pkg examples

contain example of NF blueprint packages.

## Prerequisites

Install: `golang`, `kpt`

## Create package

```bash
go mod init github.com/matysiaq/kpt-examples
```

```bash
touch main.go
```

## Run kpt fn locally

```bash
kpt fn source data/pkg-upf | go run *.go
```
