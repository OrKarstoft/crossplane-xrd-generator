# `xrdgen-cli` Example

This example demonstrates how to use the `crossplane-xrd-generator` library to generate a Crossplane CompositeResourceDefinition (XRD) from Go structs in a CLI maner.

## Files

| File       | Description                                                                   |
| ---------- | ----------------------------------------------------------------------------- |
| `types.go` | Defines the `XDeployment` XR struct with `Spec` and `Status` fields           |
| `main.go`  | Demonstrates how to use the generator library when creating a simple CLI tool |

## Running the Example

```sh
cd examples/xrdgen-cli
go run . <all | xdeployment>
```

This CLI tool takes in a single argument. Either `all` or the name of the resource you want to output. The CLI tool will output the generated XRD YAML to stdout.

Detailed instructions on how to use the library can be found in the [XDeployment Example](../xdeployment/README.md).
