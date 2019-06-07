package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"src.kublr.io/kublr/common/pkg/entity"
	"time"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LoggingSpec defines the desired state of Logging
// +k8s:openapi-gen=true
type LoggingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	// Chart is a helm package
	Chart *entity.Chart `json:"chart,omitempty" yaml:"chart,omitempty"`
	// Values is a helm chart values
	//Values *entity.ChartValues `json:"values,omitempty" yaml:"values,omitempty"`

	LogCollection *entity.ClusterLogCollectionSpec `json:"logCollection" yaml:"logCollection"`
	Sinks         []*entity.ClusterLogSinksSpec    `json:"sinks,omitempty" yaml:"sinks,omitempty"`
	// StorageHostPath is host path for persistent logging data
	// This directory will be used for clusters that doesn't have dynamic volume provision.
	// Default value: /var/lib/kublr/logging
	StorageHostPath string `json:"storageHostPath,omitempty" yaml:"storageHostPath,omitempty" bson:"storageHostPath,omitempty"`
}

// LoggingStatus defines the observed state of Logging
// +k8s:openapi-gen=true
type LoggingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	FeatureName entity.FeatureName `json:"featureName" yaml:"featureName"`
	ReleaseName string             `json:"releaseName" yaml:"releaseName"`
	Created     time.Time          `json:"created" yaml:"created"`
	Modified    time.Time          `json:"modified" yaml:"modified"`
	Deleted     time.Time          `json:"deleted" yaml:"deleted"`
	FeatureOk   *entity.Condition  `json:"featureOk" yaml:"featureOk"`
	//Namespace is kubernetes namespace to which the feature is installed.
	Namespace string `json:"namespace" yaml:"namespace"`
	// Version is the version of the installed feature
	Version string `json:"version" yaml:"version"`
	// Sha256sum is the SHA256 checksum for the installed function
	Sha256sum string `json:"sha256sum" yaml:"sha256sum"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Logging is the Schema for the loggings API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Logging struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggingSpec   `json:"spec,omitempty"`
	Status LoggingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LoggingList contains a list of Logging
type LoggingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Logging `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Logging{}, &LoggingList{})
}
