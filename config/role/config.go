package role

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("postgresql_role", func(r *config.Resource) {
		r.ShortGroup = "role"
	})
}
