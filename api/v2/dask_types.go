/*
Copyright 2021.

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

package v2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DaskSpec defines the desired state of Dask
type DaskSpec struct {
	SchedulerTemplate PodTemplate `json:",omitempty"`
	WorkerTemplate    PodTemplate `json:",omitempty"`
	WorkerReplicas    int32       `json:",omitempty"`
}

type PodTemplate struct {
	Spec corev1.PodSpec `json:".omitempty"`
}

// DaskStatus defines the observed state of Dask
type DaskStatus struct {
	Scheduler int32 `json:",omitempty"`
	Worker    int32 `json:",omitempty"`
	// State is the state of the cluster - available waiting|running|terminating
	State string `json:",omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=dasks,singular=dask,scope=Namespaced
//+kubebuilder:printcolumn:name="Scheduler",type="integer",JSONPath=".status.scheduler",description="The number of Scheduler Requested in the Dask",priority=0
//+kubebuilder:printcolumn:name="Worker",type="integer",JSONPath=".status.worker",description="The number of ready workers in the Dask",priority=0
//+kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state",description="Status of the Dask",priority=0
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The age of dask cluster",priority=0

// Dask is the Schema for the dasks API
type Dask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DaskSpec   `json:"spec,omitempty"`
	Status DaskStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DaskList contains a list of Dask
type DaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dask `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dask{}, &DaskList{})
}
