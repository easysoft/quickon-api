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

// StorageProfileSpec defines the desired state of StorageProfile
type StorageProfileSpec struct {
	Provider string        `json:"provider"`
	Config   StorageConfig `json:"config"`
}

type StorageConfig struct {
	S3Config *StorageS3Config `json:"s3,omitempty"`
}

type StorageS3Config struct {
	AccessKey      string `json:"accessKey"`
	SecretKey      string `json:"secretKey"`
	Url            string `json:"url"`
	Region         string `json:"region"`
	Bucket         string `json:"bucket"`
	ForcePathStyle bool   `json:"forcePathStyle"`
}

type StorageProfileStatusPhase string

const (
	StorageProfileStatusPhaseAvailable   StorageProfileStatusPhase = "Available"
	StorageProfileStatusPhaseUnAvailable StorageProfileStatusPhase = "Unavailable"
)

// StorageProfileStatus defines the observed state of StorageProfile
type StorageProfileStatus struct {
	Phase   StorageProfileStatusPhase `json:"phase,omitempty"`
	Message string                    `json:"message"`
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// StorageProfile is the Schema for the storageprofiles API
type StorageProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageProfileSpec   `json:"spec,omitempty"`
	Status StorageProfileStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StorageProfileList contains a list of StorageProfile
type StorageProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageProfile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageProfile{}, &StorageProfileList{})
}
