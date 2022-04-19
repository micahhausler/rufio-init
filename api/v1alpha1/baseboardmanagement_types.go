/*
Copyright 2022 Tinkerbell.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BaseboardManagementState represents the BaseboardManagement state.
type BaseboardManagementState string

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BaseboardManagementSpec defines the desired state of BaseboardManagement
type BaseboardManagementSpec struct {
	// Host is the host IP address of the BaseboardManagement.
	// +kubebuilder:validation:MinLength=1
	Host string `json:"host"`

	// AuthSecretRef is the SecretReference that contains authentication information of the BaseboardManagement.
	// The Secret must contain username and password keys.
	AuthSecretRef corev1.SecretReference `json:"authSecretRef"`

	// Vendor is the vendor name of the BaseboardManagement.
	// +kubebuilder:validation:MinLength=1
	Vendor string `json:"vendor"`

	// Power is the desired power state of the BaseboardManagement.
	// +kubebuilder:validation:MinLength=1
	Power string `json:"power"`

	// Boot is the desired boot device for the BaseboardManagement.
	Boot Boot `json:"boot"`
}

type Boot struct {
	// BootDevice is the desired boot for the BaseboardManagement.
	// +kubebuilder:validation:MinLength=1
	BootDevice string `json:"bootDevice"`

	// Persistent if True, boot device is permanently set for the BaseboardManagement.
	// If False, boot device is set as one time boot.
	// +kubebuilder:default=false
	Persistent bool `json:"persistent,omitempty"`

	// EfiBoot specifies to EFI boot for the BaseboardManagement.
	// +kubebuilder:default=false
	EfiBoot bool `json:"efiBoot,omitempty"`
}

// BaseboardManagementStatus defines the observed state of BaseboardManagement
type BaseboardManagementStatus struct {
	//+optional
	Power BaseboardManagementState `json:"powerState,omitempty"`
	//+optional
	Boot BaseboardManagementState `json:"bootState,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=baseboardmanagements,scope=Namespaced,categories=tinkerbell,singular=baseboardmanagement,shortName=bm

// BaseboardManagement is the Schema for the baseboardmanagements API
type BaseboardManagement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BaseboardManagementSpec   `json:"spec,omitempty"`
	Status BaseboardManagementStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BaseboardManagementList contains a list of BaseboardManagement
type BaseboardManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BaseboardManagement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BaseboardManagement{}, &BaseboardManagementList{})
}
