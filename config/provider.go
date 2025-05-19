/*
Copyright 2021 Upbound Inc.
*/

package config


import (
	_ "embed"

	// Note(turkenh): we are importing this to embed provider schema document
	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/frankpromise/provider-postgresql/config/database"
	"github.com/frankpromise/provider-postgresql/config/schema"
	"github.com/frankpromise/provider-postgresql/config/user"
	"github.com/frankpromise/provider-postgresql/config/grant"
	"github.com/frankpromise/provider-postgresql/config/role"
	"github.com/frankpromise/provider-postgresql/config/extension"
    
)

const (
	resourcePrefix = "postgresql"
	modulePath     = "github.com/frankpromise/provider-postgresql"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	p := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("postgresql.crossplane.io"),
		ujconfig.WithIncludeList([]string{
			"postgresql_database",
			"postgresql_schema",
			"postgresql_user_mapping",
			"postgresql_grant",
			"postgresql_role",
			"postgresql_extension",
		}),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithFeaturesPackage("internal/features"),
	)

	for _, configure := range []func(*ujconfig.Provider){
		database.Configure,
		schema.Configure,
		user.Configure,
		grant.Configure,
		role.Configure,
		extension.Configure,
	} {
		configure(p)
	}

	p.ConfigureResources()
	return p
}