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

type DatabaseInitParameters struct {

	// If false then no one can connect to this
	// database. The default is true, allowing connections (except as restricted by
	// other mechanisms, such as GRANT or REVOKE CONNECT).
	// If false then no one can connect to this database
	AllowConnections *bool `json:"allowConnections,omitempty" tf:"allow_connections,omitempty"`

	// If true, the change of the database
	// owner will also include a reassignment of the ownership of preexisting
	// objects like tables or sequences from the previous owner to the new one.
	// If set to false (the default), then the previous database owner will still
	// hold the ownership of the objects in that database. To alter existing objects in
	// the database, you must be a direct or indirect member of the specified role, or
	// the username in the provider must be superuser.
	// If true, the owner of already existing objects will change if the owner changes
	AlterObjectOwnership *bool `json:"alterObjectOwnership,omitempty" tf:"alter_object_ownership,omitempty"`

	// How many concurrent connections can be
	// established to this database. -1 (the default) means no limit.
	// How many concurrent connections can be made to this database
	ConnectionLimit *float64 `json:"connectionLimit,omitempty" tf:"connection_limit,omitempty"`

	// Character set encoding to use in the database.
	// Specify a string constant (e.g. UTF8 or SQL_ASCII), or an integer encoding
	// number.  If unset or set to an empty string the default encoding is set to
	// UTF8.  Changing this value will force the creation of a new
	// resource as this value can only be changed when a database is created.
	// Character set encoding to use in the new database
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// If true, then this database can be cloned by any
	// user with CREATEDB privileges; if false (the default), then only
	// superusers or the owner of the database can clone it.
	// If true, then this database can be cloned by any user with CREATEDB privileges
	IsTemplate *bool `json:"isTemplate,omitempty" tf:"is_template,omitempty"`

	// Collation order (LC_COLLATE) to use in the
	// database.  This affects the sort order applied to strings, e.g. in queries
	// with ORDER BY, as well as the order used in indexes on text columns. If
	// unset or set to an empty string the default collation is set to C.  Changing this value will force the creation of a new
	// resource as this value can only be changed when a database is created.
	// Collation order (LC_COLLATE) to use in the new database
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// Character classification (LC_CTYPE) to use in the
	// database. This affects the categorization of characters, e.g. lower, upper and
	// digit. If unset or set to an empty string the default character classification
	// is set to C.  Changing this value will
	// force the creation of a new resource as this value can only be changed when a
	// database is created.
	// Character classification (LC_CTYPE) to use in the new database
	LcCtype *string `json:"lcCtype,omitempty" tf:"lc_ctype,omitempty"`

	// The role name of the user who will own the database, or
	// DEFAULT to use the default (namely, the user executing the command). To
	// create a database owned by another role or to change the owner of an existing
	// database, you must be a direct or indirect member of the specified role, or
	// the username in the provider is a superuser.
	// The ROLE which owns the database
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The name of the tablespace that will be
	// associated with the database, or DEFAULT to use the template database's
	// tablespace.  This tablespace will be the default tablespace used for objects
	// created in this database.
	// The name of the tablespace that will be associated with the new database
	TablespaceName *string `json:"tablespaceName,omitempty" tf:"tablespace_name,omitempty"`

	// The name of the template database from which to create
	// the database, or DEFAULT to use the default template (template0).  Changing this value
	// will force the creation of a new resource as this value can only be changed
	// when a database is created.
	// The name of the template from which to create the new database
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type DatabaseObservation struct {

	// If false then no one can connect to this
	// database. The default is true, allowing connections (except as restricted by
	// other mechanisms, such as GRANT or REVOKE CONNECT).
	// If false then no one can connect to this database
	AllowConnections *bool `json:"allowConnections,omitempty" tf:"allow_connections,omitempty"`

	// If true, the change of the database
	// owner will also include a reassignment of the ownership of preexisting
	// objects like tables or sequences from the previous owner to the new one.
	// If set to false (the default), then the previous database owner will still
	// hold the ownership of the objects in that database. To alter existing objects in
	// the database, you must be a direct or indirect member of the specified role, or
	// the username in the provider must be superuser.
	// If true, the owner of already existing objects will change if the owner changes
	AlterObjectOwnership *bool `json:"alterObjectOwnership,omitempty" tf:"alter_object_ownership,omitempty"`

	// How many concurrent connections can be
	// established to this database. -1 (the default) means no limit.
	// How many concurrent connections can be made to this database
	ConnectionLimit *float64 `json:"connectionLimit,omitempty" tf:"connection_limit,omitempty"`

	// Character set encoding to use in the database.
	// Specify a string constant (e.g. UTF8 or SQL_ASCII), or an integer encoding
	// number.  If unset or set to an empty string the default encoding is set to
	// UTF8.  Changing this value will force the creation of a new
	// resource as this value can only be changed when a database is created.
	// Character set encoding to use in the new database
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// If true, then this database can be cloned by any
	// user with CREATEDB privileges; if false (the default), then only
	// superusers or the owner of the database can clone it.
	// If true, then this database can be cloned by any user with CREATEDB privileges
	IsTemplate *bool `json:"isTemplate,omitempty" tf:"is_template,omitempty"`

	// Collation order (LC_COLLATE) to use in the
	// database.  This affects the sort order applied to strings, e.g. in queries
	// with ORDER BY, as well as the order used in indexes on text columns. If
	// unset or set to an empty string the default collation is set to C.  Changing this value will force the creation of a new
	// resource as this value can only be changed when a database is created.
	// Collation order (LC_COLLATE) to use in the new database
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// Character classification (LC_CTYPE) to use in the
	// database. This affects the categorization of characters, e.g. lower, upper and
	// digit. If unset or set to an empty string the default character classification
	// is set to C.  Changing this value will
	// force the creation of a new resource as this value can only be changed when a
	// database is created.
	// Character classification (LC_CTYPE) to use in the new database
	LcCtype *string `json:"lcCtype,omitempty" tf:"lc_ctype,omitempty"`

	// The role name of the user who will own the database, or
	// DEFAULT to use the default (namely, the user executing the command). To
	// create a database owned by another role or to change the owner of an existing
	// database, you must be a direct or indirect member of the specified role, or
	// the username in the provider is a superuser.
	// The ROLE which owns the database
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The name of the tablespace that will be
	// associated with the database, or DEFAULT to use the template database's
	// tablespace.  This tablespace will be the default tablespace used for objects
	// created in this database.
	// The name of the tablespace that will be associated with the new database
	TablespaceName *string `json:"tablespaceName,omitempty" tf:"tablespace_name,omitempty"`

	// The name of the template database from which to create
	// the database, or DEFAULT to use the default template (template0).  Changing this value
	// will force the creation of a new resource as this value can only be changed
	// when a database is created.
	// The name of the template from which to create the new database
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type DatabaseParameters struct {

	// If false then no one can connect to this
	// database. The default is true, allowing connections (except as restricted by
	// other mechanisms, such as GRANT or REVOKE CONNECT).
	// If false then no one can connect to this database
	// +kubebuilder:validation:Optional
	AllowConnections *bool `json:"allowConnections,omitempty" tf:"allow_connections,omitempty"`

	// If true, the change of the database
	// owner will also include a reassignment of the ownership of preexisting
	// objects like tables or sequences from the previous owner to the new one.
	// If set to false (the default), then the previous database owner will still
	// hold the ownership of the objects in that database. To alter existing objects in
	// the database, you must be a direct or indirect member of the specified role, or
	// the username in the provider must be superuser.
	// If true, the owner of already existing objects will change if the owner changes
	// +kubebuilder:validation:Optional
	AlterObjectOwnership *bool `json:"alterObjectOwnership,omitempty" tf:"alter_object_ownership,omitempty"`

	// How many concurrent connections can be
	// established to this database. -1 (the default) means no limit.
	// How many concurrent connections can be made to this database
	// +kubebuilder:validation:Optional
	ConnectionLimit *float64 `json:"connectionLimit,omitempty" tf:"connection_limit,omitempty"`

	// Character set encoding to use in the database.
	// Specify a string constant (e.g. UTF8 or SQL_ASCII), or an integer encoding
	// number.  If unset or set to an empty string the default encoding is set to
	// UTF8.  Changing this value will force the creation of a new
	// resource as this value can only be changed when a database is created.
	// Character set encoding to use in the new database
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// If true, then this database can be cloned by any
	// user with CREATEDB privileges; if false (the default), then only
	// superusers or the owner of the database can clone it.
	// If true, then this database can be cloned by any user with CREATEDB privileges
	// +kubebuilder:validation:Optional
	IsTemplate *bool `json:"isTemplate,omitempty" tf:"is_template,omitempty"`

	// Collation order (LC_COLLATE) to use in the
	// database.  This affects the sort order applied to strings, e.g. in queries
	// with ORDER BY, as well as the order used in indexes on text columns. If
	// unset or set to an empty string the default collation is set to C.  Changing this value will force the creation of a new
	// resource as this value can only be changed when a database is created.
	// Collation order (LC_COLLATE) to use in the new database
	// +kubebuilder:validation:Optional
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// Character classification (LC_CTYPE) to use in the
	// database. This affects the categorization of characters, e.g. lower, upper and
	// digit. If unset or set to an empty string the default character classification
	// is set to C.  Changing this value will
	// force the creation of a new resource as this value can only be changed when a
	// database is created.
	// Character classification (LC_CTYPE) to use in the new database
	// +kubebuilder:validation:Optional
	LcCtype *string `json:"lcCtype,omitempty" tf:"lc_ctype,omitempty"`

	// The role name of the user who will own the database, or
	// DEFAULT to use the default (namely, the user executing the command). To
	// create a database owned by another role or to change the owner of an existing
	// database, you must be a direct or indirect member of the specified role, or
	// the username in the provider is a superuser.
	// The ROLE which owns the database
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The name of the tablespace that will be
	// associated with the database, or DEFAULT to use the template database's
	// tablespace.  This tablespace will be the default tablespace used for objects
	// created in this database.
	// The name of the tablespace that will be associated with the new database
	// +kubebuilder:validation:Optional
	TablespaceName *string `json:"tablespaceName,omitempty" tf:"tablespace_name,omitempty"`

	// The name of the template database from which to create
	// the database, or DEFAULT to use the default template (template0).  Changing this value
	// will force the creation of a new resource as this value can only be changed
	// when a database is created.
	// The name of the template from which to create the new database
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseParameters `json:"forProvider"`
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
	InitProvider DatabaseInitParameters `json:"initProvider,omitempty"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Database is the Schema for the Databases API. Creates and manages a database on a PostgreSQL server.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,postgresql}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	Database_Kind             = "Database"
	Database_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Database_Kind}.String()
	Database_KindAPIVersion   = Database_Kind + "." + CRDGroupVersion.String()
	Database_GroupVersionKind = CRDGroupVersion.WithKind(Database_Kind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
