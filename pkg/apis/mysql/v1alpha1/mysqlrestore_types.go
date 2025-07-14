/*
Copyright 2019 Pressinfra SRL

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
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MysqlRestoreSpec defines the desired state of MysqlRestore.
type MysqlRestoreSpec struct {

	// ClustterName represents the cluster for which to take restore
	ClusterName string `json:"clusterName"`

	// BackupRef is the backup where the database will be restore.
	BackupRef string `json:"backupRef,omitempty"`
}

// RestoreCondition defines condition struct for backup resource
type RestoreCondition struct {
	// type of cluster condition, values in (\"Ready\")
	Type BackupConditionType `json:"type"`
	// Status of the condition, one of (\"True\", \"False\", \"Unknown\")
	Status core.ConditionStatus `json:"status"`

	// LastTransitionTime
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// Reason
	Reason string `json:"reason"`
	// Message
	Message string `json:"message"`
}

// MysqlRestoreStatus defines the observed state of MysqlRestore.
type MysqlRestoreStatus struct {
	// Completed indicates whether the restore is in a final state,
	// no matter whether its' corresponding job failed or succeeded
	Completed bool `json:"completed,omitempty"`
	// Conditions represents the restore resource conditions list.
	Conditions []RestoreCondition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MysqlRestore is the Schema for the mysqlrestores API.
type MysqlRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MysqlRestoreSpec   `json:"spec,omitempty"`
	Status MysqlRestoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MysqlRestoreList contains a list of MysqlRestore.
type MysqlRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MysqlRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MysqlRestore{}, &MysqlRestoreList{})
}
