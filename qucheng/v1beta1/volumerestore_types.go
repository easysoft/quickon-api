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
	corev1api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VolumeRestoreSpec defines the desired state of VolumeRestore
type VolumeRestoreSpec struct {
	// Pod is a reference to the pod containing the volume to be restored.
	Pod corev1api.ObjectReference `json:"pod"`

	// Volume is the name of the volume within the Pod to be restored.
	Volume string `json:"volume"`

	// StorageProfileName is the name of the backup storage location
	// where the restic repository is stored.
	StorageProfileName string `json:"storageProfileName"`

	// ResticRepositoryPath is the restic repository identifier.
	ResticRepositoryPath string `json:"resticRepositoryPath"`

	// SnapshotID is the ID of the volume snapshot to be restored.
	SnapshotID string `json:"snapshotID"`
}

// VolumeRestorePhase represents the lifecycle phase of a PodVolumeRestore.
// +kubebuilder:validation:Enum=New;InProgress;Completed;Failed
type VolumeRestorePhase string

const (
	VolumeRestorePhaseNew        VolumeRestorePhase = "New"
	VolumeRestorePhaseInProgress VolumeRestorePhase = "InProgress"
	VolumeRestorePhaseCompleted  VolumeRestorePhase = "Completed"
	VolumeRestorePhaseFailed     VolumeRestorePhase = "Failed"
)

// VolumeRestoreStatus defines the observed state of VolumeRestore
type VolumeRestoreStatus struct {
	// Phase is the current state of the PodVolumeRestore.
	// +optional
	Phase VolumeRestorePhase `json:"phase,omitempty"`

	// Message is a message about the pod volume restore's status.
	// +optional
	Message string `json:"message,omitempty"`

	// StartTimestamp records the time a restore was started.
	// The server's time is used for StartTimestamps
	// +optional
	// +nullable
	StartTimestamp *metav1.Time `json:"startTimestamp,omitempty"`

	// CompletionTimestamp records the time a restore was completed.
	// Completion time is recorded even on failed restores.
	// The server's time is used for CompletionTimestamps
	// +optional
	// +nullable
	CompletionTimestamp *metav1.Time `json:"completionTimestamp,omitempty"`

	// Progress holds the total number of bytes of the snapshot and the current
	// number of restored bytes. This can be used to display progress information
	// about the restore operation.
	// +optional
	Progress VolumeOperationProgress `json:"progress,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".spec.pod.namespace",description="Namespace of the pod containing the volume to be restored"
// +kubebuilder:printcolumn:name="Pod",type="string",JSONPath=".spec.pod.name",description="Name of the pod containing the volume to be restored"
// +kubebuilder:printcolumn:name="Volume",type="string",JSONPath=".spec.volume",description="Name of the volume to be restored"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase",description="Pod Volume Restore status such as New/InProgress"
// +kubebuilder:printcolumn:name="TotalBytes",type="integer",format="int64",JSONPath=".status.progress.totalBytes",description="Pod Volume Restore status such as New/InProgress"
// +kubebuilder:printcolumn:name="BytesDone",type="integer",format="int64",JSONPath=".status.progress.bytesDone",description="Pod Volume Restore status such as New/InProgress"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:path=volumerestores,shortName=vr

// VolumeRestore is the Schema for the volumerestores API
type VolumeRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VolumeRestoreSpec   `json:"spec,omitempty"`
	Status VolumeRestoreStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VolumeRestoreList contains a list of VolumeRestore
type VolumeRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VolumeRestore{}, &VolumeRestoreList{})
}
