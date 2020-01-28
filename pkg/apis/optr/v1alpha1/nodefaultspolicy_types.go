package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NoDefaultsPolicySpec defines the desired state of NoDefaultsPolicy
type NoDefaultsPolicySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// NoDefaultsPolicyStatus defines the observed state of NoDefaultsPolicy
type NoDefaultsPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NoDefaultsPolicy is the Schema for the nodefaultspolicies API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=nodefaultspolicies,scope=Namespaced
type NoDefaultsPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NoDefaultsPolicySpec   `json:"spec,omitempty"`
	Status NoDefaultsPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NoDefaultsPolicyList contains a list of NoDefaultsPolicy
type NoDefaultsPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NoDefaultsPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NoDefaultsPolicy{}, &NoDefaultsPolicyList{})
}
