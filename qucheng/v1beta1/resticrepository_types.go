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

// ResticRepositorySpec defines the desired state of ResticRepository
type ResticRepositorySpec struct {
	StorageProfileName   string `json:"storageProfileName"`
	ResticRepositoryPath string `json:"resticRepositoryPath,omitempty"`
	SubPath              string `json:"subPath"`
}

type ResticRepositoryPhase string

const (
	ResticRepositoryPhaseNew      ResticRepositoryPhase = "New"
	ResticRepositoryPhaseReady    ResticRepositoryPhase = "Ready"
	ResticRepositoryPhaseNotReady ResticRepositoryPhase = "NotReady"
)

// ResticRepositoryStatus defines the observed state of ResticRepository
type ResticRepositoryStatus struct {
	Phase           ResticRepositoryPhase `json:"phase"`
	Message         string                `json:"message,omitempty"`
	Statistic       ResticRepoStatistic   `json:"statistic,omitempty"`
	PruneTimestamp  *metav1.Time          `json:"pruneTimestamp,omitempty"`
	UpdateTimestamp *metav1.Time          `json:"updateTimestamp,omitempty"`
}

type ResticRepoStatistic struct {
	TotalSize      *int64 `json:"totalSize"`
	TotalFileCount *int64 `json:"totalFileCount"`
	SnapShotCount  *int   `json:"snapshotCount"`
}

// +genclient
//+kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.phase",description="ResticRepository status such as Ready/UnReady"
//+kubebuilder:printcolumn:name="Count",type="string",JSONPath=".status.statistic.snapshotCount",description="ResticRepository snapshots counts"
//+kubebuilder:printcolumn:name="Size",type="string",JSONPath=".status.statistic.totalSize",description="ResticRepository total size"
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ResticRepository is the Schema for the resticrepositories API
type ResticRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResticRepositorySpec   `json:"spec,omitempty"`
	Status ResticRepositoryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ResticRepositoryList contains a list of ResticRepository
type ResticRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResticRepository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResticRepository{}, &ResticRepositoryList{})
}
