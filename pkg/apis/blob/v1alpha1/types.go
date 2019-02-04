package v1alpha1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Blob describes a Blob resource
type Blob struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	meta.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object, including
	// things like...
	//  - name
	//  - namespace
	//  - self link
	//  - labels
	//  - ... etc ...
	meta.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the custom resource spec
	Spec BlobSpec `json:"spec"`
}

// BlobSpec is the spec for a Blob resource
type BlobSpec struct {
	// Message and SomeValue are example custom spec fields
	//
	// this is where you would put your custom resource data
	Message   string `json:"message"`
	SomeValue *int32 `json:"someValue"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlobList is a list of Blob resources
type BlobList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata"`

	Items []Blob `json:"items"`
}
