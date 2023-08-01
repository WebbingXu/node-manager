/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Node struct {
	Ip string `json:"ip"`
	Role [] string `json:"role"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodeScaleSpec defines the desired state of NodeScale
type NodeScaleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of NodeScale. Edit nodescale_types.go to remove/update
	ClusterName string `json:"cluster_name"`
	Nodes [] Node `json:"nodes"`
}

// NodeScaleStatus defines the observed state of NodeScale
type NodeScaleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// NodeScale is the Schema for the nodescales API
type NodeScale struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeScaleSpec   `json:"spec,omitempty"`
	Status NodeScaleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NodeScaleList contains a list of NodeScale
type NodeScaleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeScale `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeScale{}, &NodeScaleList{})
}
