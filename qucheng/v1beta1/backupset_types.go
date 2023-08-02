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

type BackupItem struct {
	Name        string      `json:"name"`
	Namespace   string      `json:"namespace"`
	BackupName  string      `json:"backupName,omitempty"`
	BackupPhase BackupPhase `json:"phase,omitempty"`
}

// BackupSetSpec defines the desired state of BackupSet
type BackupSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Apps []BackupItem `json:"apps"`
}

type BackupSetPhase string

const (
	BackupSetPhaseNew       BackupSetPhase = "New"
	BackupSetPhaseProcess   BackupSetPhase = "InProgress"
	BackupSetPhaseFailed    BackupSetPhase = "Failed"
	BackupSetPhaseCompleted BackupSetPhase = "Completed"
	BackupSetPhaseDeleting  BackupSetPhase = "Deleting"
)

type BackupSetProgress struct {
	Total     int `json:"total,omitempty"`
	Completed int `json:"completed,omitempty"`
}

// BackupSetStatus defines the observed state of BackupSet
type BackupSetStatus struct {
	Phase               BackupSetPhase    `json:"phase"`
	Reason              string            `json:"reason,omitempty"`
	Message             string            `json:"message,omitempty"`
	MetadataPath        string            `json:"metadataPath"`
	Progress            BackupSetProgress `json:"progress,omitempty"`
	StartTimestamp      *metav1.Time      `json:"startTimestamp,omitempty"`
	CompletionTimestamp *metav1.Time      `json:"completionTimestamp,omitempty"`
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// BackupSet is the Schema for the backupsets API
type BackupSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupSetSpec   `json:"spec,omitempty"`
	Status BackupSetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BackupSetList contains a list of BackupSet
type BackupSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupSet{}, &BackupSetList{})
}
