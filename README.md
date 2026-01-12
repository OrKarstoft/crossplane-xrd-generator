# Crossplane XRD Generator
Instead of handcrafting the OpenAPI schema in your CompositeResourceDefinition like a caveman,
you define your composite resource (XR) as Go structs, and this tool generates the XRD for you automatically.

The real advantage of defining your XR as Go structs is type reuse.
If you’re writing a Go Function for your composition, you can deserialize the observed XR resource directly into the same
Go structs that were used to generate the XRD.

In short: define once, generate everywhere. Your XRs become type-safe and maintainable, with zero hand-crafted OpenAPI YAML to maintain.

## Example

Create a folder for your resource inside `src/pkg/` and add a `types.go` file:

```sh
mkdir src/pkg/xdeployment
touch src/pkg/xdeployment/types.go
```

In `types.go`, define your resource struct. You can use `src/pkg/example/types.go` as a reference.

Customize the `Spec` and `Status` fields to match your needs:

```go
package xdeployment

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

Next, open `main.go` and import your new package.

Add your resource to the `xResources` map as a `ResourceMeta` struct, specifying your Kubernetes API group:

```go
var xResources = map[string]generator.ResourceMeta{
	"xdeployment": {
		Type:  reflect.TypeOf(xdeployment.XDeployment{}),
		Group: "example.org",
	},
}
```

Finally, run the application to generate the XRD:

```go
go run main.go -resource xdeployment
```

After running, your Crossplane XRD will be generated automatically, ready to be applied to your cluster.

```yaml
apiVersion: apiextensions.crossplane.io/v2
kind: CompositeResourceDefinition
metadata:
  name: xdeployments.example.org
spec:
  group: example.org
  names:
    kind: XDeployment
    plural: xdeployments
  scope: Namespaced
  versions:
  - name: v1
    referenceable: true
    schema:
      openAPIV3Schema:
        properties:
          spec:
            properties:
              hostname:
                type: string
              image:
                type: string
              port:
                type: integer
            required:
            - image
            type: object
          status:
            properties:
              replicas:
                type: integer
            type: object
        type: object
    served: true
```

## Crossplane Function
If you are writing a [Crossplane Composite Function in Go](https://docs.crossplane.io/latest/guides/write-a-composition-function-in-go/),
you can now import your XR Go struct and deserialize the observed composite resource directly into a strongly typed Go struct.

This allows you to reuse the same structs you used to generate your XRD and work with type-safe fields instead of unstructured maps.
