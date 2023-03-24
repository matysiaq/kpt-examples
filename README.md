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

## Sample `Results` output

```yaml
results:
- message: index 2 not accepted
  severity: error
- file:
    path: dnn.yaml
  message: not allowed
  resourceRef:
    name: internet
    apiVersion: req.nephio.org/v1alpha1
    kind: DataNetworkName
  severity: error

```