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

// VolumeBackupSpec defines the desired state of VolumeBackup
type VolumeBackupSpec struct {
	// Node is the name of the node that the Pod is running on.
	Node string `json:"node"`

	// Pod is a reference to the pod containing the volume to be backed up.
	Pod corev1api.ObjectReference `json:"pod"`

	// Volume is the name of the volume within the Pod to be backed
	// up.
	Volume string `json:"volume"`

	// StorageProfileName is the name of the backup storage location
	// where the restic repository is stored.
	StorageProfileName string `json:"storageProfileName"`

	// ResticRepositoryPath is the restic repository identifier.
	ResticRepositoryPath string `json:"resticRepositoryPath"`

	// Tags are a map of key-value pairs that should be applied to the
	// volume backup as tags.
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

// VolumeBackupPhase represents the lifecycle phase of a PodVolumeBackup.
// +kubebuilder:validation:Enum=New;InProgress;Completed;Failed
type VolumeBackupPhase string

const (
	VolumeBackupPhaseNew        VolumeBackupPhase = "New"
	VolumeBackupPhaseInProgress VolumeBackupPhase = "InProgress"
	VolumeBackupPhaseCompleted  VolumeBackupPhase = "Completed"
	VolumeBackupPhaseFailed     VolumeBackupPhase = "Failed"
)

// VolumeOperationProgress represents the progress of a
// VolumeBackup/Restore (restic) operation
type VolumeOperationProgress struct {
	// +optional
	TotalBytes int64 `json:"totalBytes,omitempty"`

	// +optional
	BytesDone int64 `json:"bytesDone,omitempty"`
}

// VolumeBackupStatus is the current status of a VolumeBackup.
type VolumeBackupStatus struct {
	// Phase is the current state of the VolumeBackup.
	// +optional
	Phase VolumeBackupPhase `json:"phase,omitempty"`

	// Path is the full path within the controller pod being backed up.
	// +optional
	Path string `json:"path,omitempty"`

	// SnapshotID is the identifier for the snapshot of the pod volume.
	// +optional
	SnapshotID string `json:"snapshotID,omitempty"`

	// Message is a message about the pod volume backup's status.
	// +optional
	Message string `json:"message,omitempty"`

	// StartTimestamp records the time a backup was started.
	// Separate from CreationTimestamp, since that value changes
	// on restores.
	// The server's time is used for StartTimestamps
	// +optional
	// +nullable
	StartTimestamp *metav1.Time `json:"startTimestamp,omitempty"`

	// CompletionTimestamp records the time a backup was completed.
	// Completion time is recorded even on failed backups.
	// Completion time is recorded before uploading the backup object.
	// The server's time is used for CompletionTimestamps
	// +optional
	// +nullable
	CompletionTimestamp *metav1.Time `json:"completionTimestamp,omitempty"`

	// Progress holds the total number of bytes of the volume and the current
	// number of backed up bytes. This can be used to display progress information
	// about the backup operation.
	// +optional
	Progress VolumeOperationProgress `json:"progress,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase",description="Pod Volume Backup status such as New/InProgress"
// +kubebuilder:printcolumn:name="Created",type="date",JSONPath=".status.startTimestamp",description="Time when this backup was started"
// +kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".spec.pod.namespace",description="Namespace of the pod containing the volume to be backed up"
// +kubebuilder:printcolumn:name="Pod",type="string",JSONPath=".spec.pod.name",description="Name of the pod containing the volume to be backed up"
// +kubebuilder:printcolumn:name="Volume",type="string",JSONPath=".spec.volume",description="Name of the volume to be backed up"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:object:root=true
// +kubebuilder:object:generate=true
// +kubebuilder:resource:path=volumebackups,shortName=vb

// VolumeBackup is the Schema for the volumebackups API
type VolumeBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VolumeBackupSpec   `json:"spec,omitempty"`
	Status VolumeBackupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VolumeBackupList contains a list of VolumeBackup
type VolumeBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VolumeBackup{}, &VolumeBackupList{})
}
