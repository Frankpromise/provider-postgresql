package database

import "github.com/crossplane/upjet/pkg/config"

// Configure sets custom configuration for the postgresql_database resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("postgresql_database", func(r *config.Resource) {
		// Override the default group name (e.g., "postgresql") if needed
		r.ShortGroup = "db"

		// Optionally remove unused fields from the schema (example)
		// r.LateInitializer = config.LateInitializer{
		//     IgnoredFields: []string{"owner"},
		// }

		// Mark the `name` field as the external name (already handled via external_name.go)
		// This is only needed if not using config.NameAsIdentifier
	})
}
