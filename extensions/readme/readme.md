# provider-cloudinary

Crossplane provider for [Cloudinary](https://cloudinary.com/) that manages
resources via the Cloudinary **Admin API**.

## Supported Resources

| Kind | Description |
|---|---|
| `Folder` | Cloudinary folder that organises media assets |
| `Trigger` | Webhook notification trigger for asset events |

## Installation

```bash
kubectl crossplane install provider ghcr.io/nitrikx/provider-cloudinary:latest
```

## Credentials

Create a Kubernetes Secret with your Cloudinary Admin API credentials:

```bash
kubectl create secret generic cloudinary-creds -n crossplane-system \
  --from-literal=credentials='{"cloud_name":"mycloud","api_key":"KEY","api_secret":"SECRET"}'
```

Then create a ProviderConfig:

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

## Example: Create a Folder

```yaml
apiVersion: folder.cloudinary.crossplane.io/v1alpha1
kind: Folder
metadata:
  name: my-team-uploads
spec:
  forProvider:
    path: teams/my-team/uploads
  providerConfigRef:
    name: default
```

## Example: Create a Webhook Trigger

```yaml
apiVersion: trigger.cloudinary.crossplane.io/v1alpha1
kind: Trigger
metadata:
  name: on-upload
spec:
  forProvider:
    uri: https://example.com/webhooks/cloudinary
    eventType: upload
  providerConfigRef:
    name: default
```
