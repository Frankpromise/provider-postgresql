/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this provider.
// ExternalNameConfigs contains all external name configurations for this
// provider.
// config/external_name.go

// ExternalNameConfigs contains all external name configurations for this provider.
var ExternalNameConfigs = map[string]config.ExternalName{
    // Simple "name"-based resources
    "postgresql_database":  config.NameAsIdentifier,
    "postgresql_role":      config.NameAsIdentifier,
    "postgresql_user":      config.NameAsIdentifier,
    "postgresql_schema":    config.NameAsIdentifier,
    "postgresql_extension": config.NameAsIdentifier,

    // Complex import: database/schema/table/role
    "postgresql_grant": config.TemplatedStringAsIdentifier("role", "{{ .parameters.database }}/{{ .parameters.schema }}/{{ .parameters.object_name }}/{{ .external_name }}"),

}


// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
