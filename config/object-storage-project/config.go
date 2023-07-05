package object_storage_project

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("stackit_object_storage_project", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "object_storage_project"
	})
}
