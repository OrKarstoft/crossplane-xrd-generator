package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// XDeployment is an example Crossplane composite resource (XR).
// Define your Spec and Status fields to match your needs.
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
