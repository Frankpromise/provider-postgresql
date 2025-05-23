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

type GrantInitParameters struct {

	// The specific columns to grant privileges on for this role
	// +listType=set
	Columns []*string `json:"columns,omitempty" tf:"columns,omitempty"`

	// The PostgreSQL object type to apply this security label to.
	// The PostgreSQL object type to grant the privileges on (one of: database, function, procedure, routine, schema, sequence, table, foreign_data_wrapper, foreign_server, column)
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// The specific objects to grant privileges on for this role (empty means all objects of the requested type)
	// +listType=set
	Objects []*string `json:"objects,omitempty" tf:"objects,omitempty"`

	// The list of privileges to grant
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// Permit the grant recipient to grant it to others
	WithGrantOption *bool `json:"withGrantOption,omitempty" tf:"with_grant_option,omitempty"`
}

type GrantObservation struct {

	// The specific columns to grant privileges on for this role
	// +listType=set
	Columns []*string `json:"columns,omitempty" tf:"columns,omitempty"`

	// The database to grant privileges on for this role
	Database *string `json:"database,omitempty" tf:"database,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The PostgreSQL object type to apply this security label to.
	// The PostgreSQL object type to grant the privileges on (one of: database, function, procedure, routine, schema, sequence, table, foreign_data_wrapper, foreign_server, column)
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// The specific objects to grant privileges on for this role (empty means all objects of the requested type)
	// +listType=set
	Objects []*string `json:"objects,omitempty" tf:"objects,omitempty"`

	// The list of privileges to grant
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// The database schema to grant privileges on for this role
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// Permit the grant recipient to grant it to others
	WithGrantOption *bool `json:"withGrantOption,omitempty" tf:"with_grant_option,omitempty"`
}

type GrantParameters struct {

	// The specific columns to grant privileges on for this role
	// +kubebuilder:validation:Optional
	// +listType=set
	Columns []*string `json:"columns,omitempty" tf:"columns,omitempty"`

	// The database to grant privileges on for this role
	// +kubebuilder:validation:Required
	Database *string `json:"database" tf:"database,omitempty"`

	// The PostgreSQL object type to apply this security label to.
	// The PostgreSQL object type to grant the privileges on (one of: database, function, procedure, routine, schema, sequence, table, foreign_data_wrapper, foreign_server, column)
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// The specific objects to grant privileges on for this role (empty means all objects of the requested type)
	// +kubebuilder:validation:Optional
	// +listType=set
	Objects []*string `json:"objects,omitempty" tf:"objects,omitempty"`

	// The list of privileges to grant
	// +kubebuilder:validation:Optional
	// +listType=set
	Privileges []*string `json:"privileges,omitempty" tf:"privileges,omitempty"`

	// The database schema to grant privileges on for this role
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// Permit the grant recipient to grant it to others
	// +kubebuilder:validation:Optional
	WithGrantOption *bool `json:"withGrantOption,omitempty" tf:"with_grant_option,omitempty"`
}

// GrantSpec defines the desired state of Grant
type GrantSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GrantParameters `json:"forProvider"`
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
	InitProvider GrantInitParameters `json:"initProvider,omitempty"`
}

// GrantStatus defines the observed state of Grant.
type GrantStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GrantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Grant is the Schema for the Grants API. Creates and manages privileges given to a user for a database schema.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,postgresql}
type Grant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.objectType) || (has(self.initProvider) && has(self.initProvider.objectType))",message="spec.forProvider.objectType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privileges) || (has(self.initProvider) && has(self.initProvider.privileges))",message="spec.forProvider.privileges is a required parameter"
	Spec   GrantSpec   `json:"spec"`
	Status GrantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GrantList contains a list of Grants
type GrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Grant `json:"items"`
}

// Repository type metadata.
var (
	Grant_Kind             = "Grant"
	Grant_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Grant_Kind}.String()
	Grant_KindAPIVersion   = Grant_Kind + "." + CRDGroupVersion.String()
	Grant_GroupVersionKind = CRDGroupVersion.WithKind(Grant_Kind)
)

func init() {
	SchemeBuilder.Register(&Grant{}, &GrantList{})
}
