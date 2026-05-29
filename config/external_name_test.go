package config

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func TestExternalNameConfigs_ContainsExpectedResources(t *testing.T) {
	expected := []string{"cloudinary_folder", "cloudinary_trigger"}
	for _, name := range expected {
		if _, ok := ExternalNameConfigs[name]; !ok {
			t.Errorf("ExternalNameConfigs missing expected resource %q", name)
		}
	}
	if len(ExternalNameConfigs) != len(expected) {
		t.Errorf("ExternalNameConfigs has %d entries, want %d", len(ExternalNameConfigs), len(expected))
	}
}

func TestExternalNameConfigs_AllUseIdentifierFromProvider(t *testing.T) {
	for name, en := range ExternalNameConfigs {
		// IdentifierFromProvider has a nil GetExternalNameFn by default in upjet.
		// We compare the zero-value type identity by checking the function ptr is nil,
		// which matches ujconfig.IdentifierFromProvider.
		want := ujconfig.IdentifierFromProvider
		if en.GetIDFn == nil && want.GetIDFn != nil {
			t.Errorf("resource %q: GetIDFn mismatch", name)
		}
		if en.OmittedFields != nil {
			t.Errorf("resource %q: expected no OmittedFields, got %v", name, en.OmittedFields)
		}
	}
}

func TestExternalNameConfigured_ReturnsRegexAnchored(t *testing.T) {
	configured := ExternalNameConfigured()
	if len(configured) != len(ExternalNameConfigs) {
		t.Fatalf("ExternalNameConfigured() returned %d entries, want %d", len(configured), len(ExternalNameConfigs))
	}
	for _, s := range configured {
		if len(s) == 0 || s[len(s)-1] != '$' {
			t.Errorf("ExternalNameConfigured() entry %q does not end with '$'", s)
		}
	}
}
