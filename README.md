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

## Sample output of using KubeObject methods

```yaml
------------- Check KubeObject functions -------------
fn.KubeObject implements a lot of getters / setters / other functions, e.g.
        GetAPIVersion: kpt.dev/v1
        GetKind: Kptfile
        IsLocalConfig: true
-------
        GetAPIVersion: req.nephio.org/v1alpha1
        GetKind: Capacity
        IsLocalConfig: true
-------
        GetAPIVersion: infra.nephio.org/v1alpha1
        GetKind: ClusterContext
        IsLocalConfig: true
-------
        GetAPIVersion: req.nephio.org/v1alpha1
        GetKind: DataNetworkName
        IsLocalConfig: true
-------
        GetAPIVersion: req.nephio.org/v1alpha1
        GetKind: Interface
        IsLocalConfig: true
-------
        GetAPIVersion: req.nephio.org/v1alpha1
        GetKind: Interface
        IsLocalConfig: true
-------
        GetAPIVersion: req.nephio.org/v1alpha1
        GetKind: Interface
        IsLocalConfig: true
-------
        GetAPIVersion: v1
        GetKind: ConfigMap
        IsLocalConfig: true
-------
------------------------------------------------------
```

## Sample output while accessing Nested Fields in the Resource

```
Accessed pool of DataNetworkName [name=internet]: [pool1, true, <nil>]

Accessed values of Interface [name=n3]: [vpc-ran, true, <nil>]

Accessed values of Interface [name=n4]: [vpc-internal, true, <nil>]

Accessed values of Interface [name=n6]: [vpc-internet, true, <nil>]
```

## Accessing specific API Version: best practises

In fact, it's not a good idea to use below syntax

```go
o.GetAPIVersion() == "req.nephio.org/v1alpha1" && o.GetKind() == "Interface"
```

we should access the library e.g.

```go
// The same should be defined for Kind
apiVersion := corev1.SchemaGroupVersion.Identifier()
```