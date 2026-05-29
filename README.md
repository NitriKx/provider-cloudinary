# provider-cloudinary

`provider-cloudinary` is a [Crossplane](https://crossplane.io/) provider that
manages [Cloudinary](https://cloudinary.com/) resources using the Cloudinary
**Admin API**. It is built with [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources.

## Supported Resources

| Kind | API Group | Cloudinary Object |
|---|---|---|
| `Folder` | `folder.cloudinary.crossplane.io` | Cloudinary folder (organises media assets) |
| `Trigger` | `trigger.cloudinary.crossplane.io` | Webhook notification trigger |

Both resources are also available in a namespaced variant under the
`*.cloudinary.m.crossplane.io` API groups.

## Installation

Requires [Crossplane](https://docs.crossplane.io/latest/software/install/) ≥ v1.14.

```bash
kubectl crossplane install provider ghcr.io/nitrikx/provider-cloudinary:latest
```

Or apply directly:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-cloudinary
spec:
  package: ghcr.io/nitrikx/provider-cloudinary:latest
```

## Credentials Setup

Create a Kubernetes Secret with your Cloudinary credentials. You can use either a
full `cloudinary_url` or the individual `cloud_name`, `api_key`, and `api_secret`
fields:

```bash
# Option A — Cloudinary URL form
kubectl create secret generic cloudinary-creds -n crossplane-system \
  --from-literal=credentials='{"cloudinary_url":"cloudinary://API_KEY:API_SECRET@CLOUD_NAME"}'

# Option B — individual fields
kubectl create secret generic cloudinary-creds -n crossplane-system \
  --from-literal=credentials='{"cloud_name":"mycloud","api_key":"123456789012345","api_secret":"abcdefghijklmnop"}'
```

Then create a `ProviderConfig` that references the Secret:

```yaml
apiVersion: cloudinary.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: cloudinary-creds
      key: credentials
```

## Quick Start

```bash
kubectl apply -f examples/cluster/providerconfig/providerconfig.yaml

# Create a Cloudinary folder
kubectl apply -f examples/cluster/folder/v1alpha1/folder.yaml
kubectl get folder -o wide

# Create a webhook trigger
kubectl apply -f examples/cluster/trigger/v1alpha1/trigger.yaml
kubectl get trigger -o wide
```

## Developing

Regenerate from the local sibling Terraform provider:

```bash
make generate-local
```

Regenerate from a published GitHub release of the Terraform provider:

```bash
make generate
```

Build the provider binary:

```bash
make build
```

Run against a Kubernetes cluster:

```bash
make run
```

## Report a Bug

Open an [issue](https://github.com/NitriKx/provider-cloudinary/issues).
