# Release Notes

## v0.1.0 (Initial Release)

First release of provider-cloudinary.

### New Managed Resources

- **Folder** (`folder.cloudinary.crossplane.io/v1alpha1`) — Manage Cloudinary
  folders that organise media assets. Supports create, rename (via path update),
  and delete.

- **Trigger** (`trigger.cloudinary.crossplane.io/v1alpha1`) — Manage Cloudinary
  webhook notification triggers. Supports all Cloudinary event types, JSONLogic
  filters, Mustache payload templates, and multiple authentication schemes
  (`default`, `legacy_hmac`, `eddsa_v2`).
