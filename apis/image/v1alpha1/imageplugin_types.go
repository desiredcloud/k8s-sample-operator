/*
Copyright 2023 desiredcloud.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ImagePluginSpec defines the desired state of ImagePlugin
type ImagePluginSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Size defines the number of ImagePlugin instances
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	// +kubebuilder:validation:ExclusiveMaximum=false
	Size int32 `json:"size,omitempty"`
}

// ImagePluginStatus defines the observed state of ImagePlugin
type ImagePluginStatus struct {
	// Represents the observations of a ImagePlugin's current state.
	// ImagePlugin.status.conditions.type are: "Available", "Progressing", and "Degraded"
	// ImagePlugin.status.conditions.status are one of True, False, Unknown.
	// ImagePlugin.status.conditions.reason the value should be a CamelCase string and producers of specific
	// condition types may define expected values and meanings for this field, and whether the values
	// are considered a guaranteed API.
	// ImagePlugin.status.conditions.Message is a human readable message indicating details about the transition.
	// For further information see: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties

	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ImagePlugin is the Schema for the imageplugins API
type ImagePlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImagePluginSpec   `json:"spec,omitempty"`
	Status ImagePluginStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ImagePluginList contains a list of ImagePlugin
type ImagePluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImagePlugin `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImagePlugin{}, &ImagePluginList{})
}
