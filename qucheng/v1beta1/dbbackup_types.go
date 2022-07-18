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
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DbBackupSpec defines the desired state of DbBackup
type DbBackupSpec struct {
	Db v1.ObjectReference `json:"db"`
}

type DbBackupPhase string

const (
	DbBackupPhasePhaseNew        DbBackupPhase = "New"
	DbBackupPhasePhaseInProgress DbBackupPhase = "InProgress"
	DbBackupPhasePhaseCompleted  DbBackupPhase = "Completed"
	DbBackupPhasePhaseFailed     DbBackupPhase = "Failed"
)

// DbBackupStatus defines the observed state of DbBackup
type DbBackupStatus struct {
	Phase               DbBackupPhase      `json:"phase"`
	Path                string             `json:"path"`
	Size                *resource.Quantity `json:"size,omitempty"`
	Message             string             `json:"message,omitempty"`
	StartTimestamp      *metav1.Time       `json:"startTimestamp,omitempty"`
	CompletionTimestamp *metav1.Time       `json:"completionTimestamp,omitempty"`
}

//+genclient
//+kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase",description="DbBackup status such as New/InProgress"
//+kubebuilder:printcolumn:name="Size",type="string",JSONPath=".status.size",description="backup file size"
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DbBackup is the Schema for the dbbackups API
type DbBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbBackupSpec   `json:"spec,omitempty"`
	Status DbBackupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DbBackupList contains a list of DbBackup
type DbBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DbBackup{}, &DbBackupList{})
}
