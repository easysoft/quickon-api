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

// DbDumpReplicaSpec defines the desired state of DbDumpReplica
type DbDumpReplicaSpec struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Path        string `json:"path"`
	StorageName string `json:"storageName"`
}

type DbDumpReplicaStatusPhase string

const (
	DbDumpReplicaStatusPhaseNew        DbDumpReplicaStatusPhase = "New"
	DbDumpReplicaStatusPhaseInProgress DbDumpReplicaStatusPhase = "InProgress"
	DbDumpReplicaStatusPhaseCompleted  DbDumpReplicaStatusPhase = "Completed"
	DbDumpReplicaStatusPhaseFailed     DbDumpReplicaStatusPhase = "Failed"
)

// DbDumpReplicaStatus defines the observed state of DbDumpReplica
type DbDumpReplicaStatus struct {
	Phase               DbDumpReplicaStatusPhase `json:"phase"`
	Size                int64                    `json:"size,omitempty"`
	Message             string                   `json:"message,omitempty"`
	StartTimestamp      *metav1.Time             `json:"startTimestamp,omitempty"`
	CompletionTimestamp *metav1.Time             `json:"completionTimestamp,omitempty"`
}

//+genclient
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase",description="Pod Volume Restore status such as New/InProgress"
// +kubebuilder:printcolumn:name="Size",type="integer",format="int64",JSONPath=".status.size",description="restic dump file size"
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:resource:path=dbdumpreplicas,shortName=ddr

// DbDumpReplica is the Schema for the dbdumpreplicas API
type DbDumpReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbDumpReplicaSpec   `json:"spec,omitempty"`
	Status DbDumpReplicaStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DbDumpReplicaList contains a list of DbDumpReplica
type DbDumpReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbDumpReplica `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DbDumpReplica{}, &DbDumpReplicaList{})
}
