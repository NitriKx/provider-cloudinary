package upload_preset

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the cloudinary_upload_preset resource for the namespaced provider.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("cloudinary_upload_preset", func(r *ujconfig.Resource) {
		r.Kind = "UploadPreset"
		r.ShortGroup = "uploadpreset"
		r.Version = "v1alpha1"
	})
}
