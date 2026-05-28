package trigger

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the cloudinary_trigger resource for the namespaced provider.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("cloudinary_trigger", func(r *ujconfig.Resource) {
		r.Kind = "Trigger"
		r.ShortGroup = "trigger"
		r.Version = "v1alpha1"
	})
}
