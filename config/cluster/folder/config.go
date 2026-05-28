package folder

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the cloudinary_folder resource for the cluster-scoped provider.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("cloudinary_folder", func(r *ujconfig.Resource) {
		r.Kind = "Folder"
		r.ShortGroup = "folder"
		r.Version = "v1alpha1"
	})
}
