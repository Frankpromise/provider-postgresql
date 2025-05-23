// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ExtensionInitParameters struct {

	// When true, will also create any extensions that this extension depends on that are not already installed. (Default: false)
	// When true, will also create any extensions that this extension depends on that are not already installed
	CreateCascade *bool `json:"createCascade,omitempty" tf:"create_cascade,omitempty"`

	// Which database to create the extension on. Defaults to provider database.
	// Sets the database to add the extension to
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects. (Default: false)
	// When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects
	DropCascade *bool `json:"dropCascade,omitempty" tf:"drop_cascade,omitempty"`

	// Sets the schema of an extension.
	// Sets the schema of an extension
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// Sets the version number of the extension.
	// Sets the version number of the extension
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ExtensionObservation struct {

	// When true, will also create any extensions that this extension depends on that are not already installed. (Default: false)
	// When true, will also create any extensions that this extension depends on that are not already installed
	CreateCascade *bool `json:"createCascade,omitempty" tf:"create_cascade,omitempty"`

	// Which database to create the extension on. Defaults to provider database.
	// Sets the database to add the extension to
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects. (Default: false)
	// When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects
	DropCascade *bool `json:"dropCascade,omitempty" tf:"drop_cascade,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Sets the schema of an extension.
	// Sets the schema of an extension
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// Sets the version number of the extension.
	// Sets the version number of the extension
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ExtensionParameters struct {

	// When true, will also create any extensions that this extension depends on that are not already installed. (Default: false)
	// When true, will also create any extensions that this extension depends on that are not already installed
	// +kubebuilder:validation:Optional
	CreateCascade *bool `json:"createCascade,omitempty" tf:"create_cascade,omitempty"`

	// Which database to create the extension on. Defaults to provider database.
	// Sets the database to add the extension to
	// +kubebuilder:validation:Optional
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	// When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects. (Default: false)
	// When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects
	// +kubebuilder:validation:Optional
	DropCascade *bool `json:"dropCascade,omitempty" tf:"drop_cascade,omitempty"`

	// Sets the schema of an extension.
	// Sets the schema of an extension
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// Sets the version number of the extension.
	// Sets the version number of the extension
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// ExtensionSpec defines the desired state of Extension
type ExtensionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExtensionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ExtensionInitParameters `json:"initProvider,omitempty"`
}

// ExtensionStatus defines the observed state of Extension.
type ExtensionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExtensionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Extension is the Schema for the Extensions API. Creates and manages an extension on a PostgreSQL server.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,postgresql}
type Extension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExtensionSpec   `json:"spec"`
	Status            ExtensionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExtensionList contains a list of Extensions
type ExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Extension `json:"items"`
}

// Repository type metadata.
var (
	Extension_Kind             = "Extension"
	Extension_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Extension_Kind}.String()
	Extension_KindAPIVersion   = Extension_Kind + "." + CRDGroupVersion.String()
	Extension_GroupVersionKind = CRDGroupVersion.WithKind(Extension_Kind)
)

func init() {
	SchemeBuilder.Register(&Extension{}, &ExtensionList{})
}
