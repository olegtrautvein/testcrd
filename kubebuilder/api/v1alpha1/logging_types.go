/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"src.kublr.io/kublr/common/pkg/entity"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LoggingSpec defines the desired state of Logging
type LoggingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
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
type LoggingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Logging is the Schema for the loggings API
type Logging struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggingSpec   `json:"spec,omitempty"`
	Status LoggingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoggingList contains a list of Logging
type LoggingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Logging `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Logging{}, &LoggingList{})
}
