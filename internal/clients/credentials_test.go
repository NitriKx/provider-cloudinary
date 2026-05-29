package clients

import (
	"testing"
)

// applyCredentials replicates the credential-mapping logic from TerraformSetupBuilder
// so it can be unit-tested without a Kubernetes client or provider scheduler.
func applyCredentials(creds map[string]string) map[string]any {
	cfg := map[string]any{}
	if v := creds["cloudinary_url"]; v != "" {
		cfg["cloudinary_url"] = v
	} else {
		if v := creds["cloud_name"]; v != "" {
			cfg["cloud_name"] = v
		}
		if v := creds["api_key"]; v != "" {
			cfg["api_key"] = v
		}
		if v := creds["api_secret"]; v != "" {
			cfg["api_secret"] = v
		}
	}
	return cfg
}

func TestCredentials_CloudinaryURLTakesPrecedence(t *testing.T) {
	creds := map[string]string{
		"cloudinary_url": "cloudinary://key:secret@mycloud",
		"cloud_name":     "ignored",
		"api_key":        "ignored",
		"api_secret":     "ignored",
	}
	cfg := applyCredentials(creds)

	if v, ok := cfg["cloudinary_url"]; !ok || v != "cloudinary://key:secret@mycloud" {
		t.Errorf("cloudinary_url = %v, want %q", v, "cloudinary://key:secret@mycloud")
	}
	if _, ok := cfg["cloud_name"]; ok {
		t.Error("cloud_name should be absent when cloudinary_url is set")
	}
}

func TestCredentials_TriplePath(t *testing.T) {
	creds := map[string]string{
		"cloud_name": "mycloud",
		"api_key":    "mykey",
		"api_secret": "mysecret",
	}
	cfg := applyCredentials(creds)

	if _, ok := cfg["cloudinary_url"]; ok {
		t.Error("cloudinary_url should be absent when triple path is used")
	}
	if v, ok := cfg["cloud_name"]; !ok || v != "mycloud" {
		t.Errorf("cloud_name = %v, want %q", v, "mycloud")
	}
	if v, ok := cfg["api_key"]; !ok || v != "mykey" {
		t.Errorf("api_key = %v, want %q", v, "mykey")
	}
	if v, ok := cfg["api_secret"]; !ok || v != "mysecret" {
		t.Errorf("api_secret = %v, want %q", v, "mysecret")
	}
}

func TestCredentials_EmptyCredsProducesEmptyConfig(t *testing.T) {
	cfg := applyCredentials(map[string]string{})
	if len(cfg) != 0 {
		t.Errorf("expected empty config for empty credentials, got %v", cfg)
	}
}

func TestCredentials_PartialTripleOmitsMissingKeys(t *testing.T) {
	creds := map[string]string{"api_key": "k"}
	cfg := applyCredentials(creds)
	if _, ok := cfg["cloud_name"]; ok {
		t.Error("cloud_name should be absent when not provided")
	}
	if v, ok := cfg["api_key"]; !ok || v != "k" {
		t.Errorf("api_key = %v, want %q", v, "k")
	}
}
