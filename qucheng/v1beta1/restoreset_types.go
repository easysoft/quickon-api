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
type RestoreItem struct {
	Name         string       `json:"name"`
	Namespace    string       `json:"namespace"`
	RestoreName  string       `json:"restoreName,omitempty"`
	RestorePhase RestorePhase `json:"phase,omitempty"`
}

// RestoreSetSpec defines the desired state of RestoreSet
type RestoreSetSpec struct {
	BackupSetName string        `json:"backupSetName"`
	Apps          []RestoreItem `json:"apps,omitempty"`
}

type RestoreSetPhase string

const (
	RestoreSetPhaseNew       RestoreSetPhase = "New"
	RestoreSetPhaseProcess   RestoreSetPhase = "InProgress"
	RestoreSetPhaseFailed    RestoreSetPhase = "Failed"
	RestoreSetPhaseCompleted RestoreSetPhase = "Completed"
	RestoreSetPhaseDeleting  RestoreSetPhase = "Deleting"
)

type RestoreSetProgress struct {
	Total     int `json:"total,omitempty"`
	Completed int `json:"completed,omitempty"`
}

// RestoreSetStatus defines the observed state of RestoreSet
type RestoreSetStatus struct {
	Phase               RestoreSetPhase    `json:"phase"`
	Reason              string             `json:"reason,omitempty"`
	Message             string             `json:"message,omitempty"`
	Progress            RestoreSetProgress `json:"progress,omitempty"`
	StartTimestamp      *metav1.Time       `json:"startTimestamp,omitempty"`
	CompletionTimestamp *metav1.Time       `json:"completionTimestamp,omitempty"`
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RestoreSet is the Schema for the restoresets API
type RestoreSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestoreSetSpec   `json:"spec,omitempty"`
	Status RestoreSetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RestoreSetList contains a list of RestoreSet
type RestoreSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestoreSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RestoreSet{}, &RestoreSetList{})
}
