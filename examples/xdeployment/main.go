package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/Kerwood/crossplane-xrd-generator/generator"
)

func main() {
	// Define the resource metadata
	resource := generator.ResourceMeta{
		Type:  reflect.TypeOf(XDeployment{}),
		Group: "example.org",
	}

	// Build the CompositeResourceDefinition
	xrd, err := generator.BuildCompositeResourceDefinition(resource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building XRD: %v\n", err)
		os.Exit(1)
	}

	// Marshal to YAML
	out, err := generator.MarshalXRDToYAML(xrd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling XRD:  %v\n", err)
		os.Exit(1)
	}

	// Print the generated XRD
	fmt.Println(string(out))
}
