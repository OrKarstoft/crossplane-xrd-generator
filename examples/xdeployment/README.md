# XDeployment Example

This example demonstrates how to use the `crossplane-xrd-generator` library to generate a Crossplane CompositeResourceDefinition (XRD) from Go structs.

## Files

| File       | Description                                                         |
| ---------- | ------------------------------------------------------------------- |
| `types.go` | Defines the `XDeployment` XR struct with `Spec` and `Status` fields |
| `main.go`  | Demonstrates how to use the generator library                       |

## Running the Example

```sh
cd examples/xdeployment
go run .
```

This will output the generated XRD YAML to stdout.

## How It Works

### 1. Define Your XR Struct

In `types.go`, define your composite resource using standard Go structs with Kubernetes metadata:

```go
type XDeployment struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ObjectMeta `json:"metadata,omitempty"`

  Spec   XDeploymentSpec   `json:"spec"`
  Status XDeploymentStatus `json:"status,omitempty"`
}

type XDeploymentSpec struct {
  Image    string `json:"image" required:"true"`
  Port     int    `json:"port,omitempty"`
  Hostname string `json:"hostname,omitempty"`
}

type XDeploymentStatus struct {
  Replicas int `json:"replicas,omitempty"`
}
```

### 2. Generate the XRD

In `main.go`, use the generator library to build and output the XRD:

```go
resource := generator.ResourceMeta{
  Type:  reflect.TypeOf(XDeployment{}),
  Group: "example. org",
}

xrd, err := generator.BuildCompositeResourceDefinition(resource)
if err != nil {
  // handle error
}

out, err := generator. MarshalXRDToYAML(xrd)
if err != nil {
  // handle error
}

fmt.Println(string(out))
```

