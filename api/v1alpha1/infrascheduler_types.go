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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InfraSchedulerSpec defines the desired state of InfraScheduler
type InfraSchedulerSpec struct {
	Method  string `json:"method,omitempty"`
	Managed bool   `json:"managed,omitempty"`
	// TODO either MachineSetManaged or Nodes must be populated but not both
	//    MachineSetManaged MachinSetManagedSpec `json:"machineSetManage,omitempty"`
	Nodes   string `json:"nodes,omitempty"`
	Isolate bool   `json:"isolate,omitempty"`
}

// InfraSchedulerStatus defines the observed state of InfraScheduler
type InfraSchedulerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// InfraScheduler is the Schema for the infraschedulers API
type InfraScheduler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InfraSchedulerSpec   `json:"spec,omitempty"`
	Status InfraSchedulerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// InfraSchedulerList contains a list of InfraScheduler
type InfraSchedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InfraScheduler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InfraScheduler{}, &InfraSchedulerList{})
}
