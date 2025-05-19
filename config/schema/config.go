package schema

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("postgresql_schema", func(r *config.Resource) {
		r.ShortGroup = "schema"
	})
}
