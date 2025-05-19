package user

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("postgresql_user_mapping", func(r *config.Resource) {
		r.ShortGroup = "user"
	})
}
