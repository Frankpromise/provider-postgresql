package extension

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("postgresql_extension", func(r *config.Resource) {
		r.ShortGroup = "extension"
	})
}