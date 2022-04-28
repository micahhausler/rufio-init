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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BMCJobConditionType represents the condition of the BMC Job.
type BMCJobConditionType string

const (
	// JobCompleted represents successful completion of the BMC Job tasks.
	JobCompleted BMCJobConditionType = "Completed"
	// JobFailed represents failure in BMC job execution.
	JobFailed BMCJobConditionType = "Failed"
	// JobTaskCompleted represents successful completion of a BMC Job task.
	JobTaskCompleted BMCJobConditionType = "TaskCompleted"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BMCJobSpec defines the desired state of BMCJob
type BMCJobSpec struct {
	// Tasks represents a list of baseboard management actions to be executed.
	// The tasks are completed in order.
	Tasks []Task `json:"tasks"`
}

// Task represents the action to be performed.
// A single task can only perform one type of action.
// For example either PowerAction or BootAction.
// +kubebuilder:validation:MaxProperties:=1
type Task struct {
	// PowerAction represents a baseboard management power operation.
	PowerAction *PowerAction `json:"powerAction,omitempty"`

	// BootAction represents a baseboard management boot device operation.
	BootAction *BootAction `json:"bootAction,omitempty"`
}

type PowerAction struct {
	// State represents the requested power state to set for the baseboard management.
	// +kubebuilder:validation:Enum=On;Off;Status;Cycle;Reset;Soft
	State PowerState `json:"state"`
}

type BootAction struct {
	// BootDevice represents the requested boot device to set.
	// +kubebuilder:validation:Enum=PXE;Disk;BIOS;CDROM;Safe
	BootDevice BootDevice `json:"bootDevice"`

	// Persistent if True, boot device is permanently set for the baseboard management.
	// If False, boot device is set as one time boot.
	// +kubebuilder:default=false
	Persistent bool `json:"persistent,omitempty"`

	// EfiBoot specifies to EFI boot for the baseboard management.
	// +kubebuilder:default=false
	EfiBoot bool `json:"efiBoot,omitempty"`
}

// BMCJobStatus defines the observed state of BMCJob
type BMCJobStatus struct {
	// Conditions represents the latest available observations of an object's current state.
	// +optional
	Conditions []BMCJobCondition `json:"conditions,omitempty"`

	// StartTime represents time when the BMCJob controller started processing a job.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// CompletionTime represents time when the job was completed.
	// The completion time is only set when the job finishes successfully.
	// +optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`
}

type BMCJobCondition struct {
	// Type of the BMCJob condition.
	Type BMCJobConditionType `json:"type"`

	// Message represents human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=bmcjobs,scope=Namespaced,categories=tinkerbell,singular=bmcjob,shortName=bmj

// BMCJob is the Schema for the bmcjobs API
type BMCJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BMCJobSpec   `json:"spec,omitempty"`
	Status BMCJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BMCJobList contains a list of BMCJob
type BMCJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BMCJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BMCJob{}, &BMCJobList{})
}
