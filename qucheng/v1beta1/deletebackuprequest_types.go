/*
Copyright 2022.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DeleteBackupRequestSpec defines the desired state of DeleteBackupRequest
type DeleteBackupRequestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	BackupName string `json:"backupName"`
}

type DeleteBackupPhase string

const (
	DeleteBackupPhaseNew       DeleteBackupPhase = "New"
	DeleteBackupPhaseProcess   DeleteBackupPhase = "InProgress"
	DeleteBackupPhaseFailed    DeleteBackupPhase = "Failed"
	DeleteBackupPhaseCompleted DeleteBackupPhase = "Completed"
)

// DeleteBackupRequestStatus defines the observed state of DeleteBackupRequest
type DeleteBackupRequestStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Phase               DeleteBackupPhase `json:"phase"`
	Reason              string            `json:"reason,omitempty"`
	Message             string            `json:"message,omitempty"`
	StartTimestamp      *metav1.Time      `json:"startTimestamp,omitempty"`
	CompletionTimestamp *metav1.Time      `json:"completionTimestamp,omitempty"`
}

//+genclient
//+kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase",description="Delete backup request status such as New/InProgress"
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DeleteBackupRequest is the Schema for the deletebackuprequests API
type DeleteBackupRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeleteBackupRequestSpec   `json:"spec,omitempty"`
	Status DeleteBackupRequestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DeleteBackupRequestList contains a list of DeleteBackupRequest
type DeleteBackupRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeleteBackupRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeleteBackupRequest{}, &DeleteBackupRequestList{})
}
