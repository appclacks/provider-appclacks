---
title: Configuration
weight: 2
---

# Appclacks provider documentation

## Install the provider

Install the Appclacks provider with the following configuration file

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-appclacks
spec:
  package: xpkg.upbound.io/appclacks/provider-appclacks:<version>
```

Define the provider version with `spec.package`.

Install the provider with `kubectl apply -f`.

Verify the configuration with `kubectl get providers`.

```shell
$ kubectl get providers
NAME                      INSTALLED   HEALTHY   PACKAGE                                                        AGE
provider-appclacks        True        True      xpkg.upbound.io/appclacks/provider-appclacks:v0.0.1            60s
```

View the Crossplane [Provider CRD definition](https://doc.crds.dev/github.com/Appclacks/provider-appclacks) to view all available `Provider` options.

## Configure the provider

The Appclacks provider requires credentials for authentication to Appclacks Clouds. This can be done in one of the following ways:

* Authenticating using AK/SK

### Generate a Kubernetes Secret

Create a JSON file containing the Appclacks account credentials (OrganizationID/Token).

Here is an example key file:

```json
{
    "organization_id": "your organization id",
    "token": "your token"
}
```

Use the JSON file to generate a Kubernetes secret.

```shell
kubectl -n upbound-system create secret generic appclacks-creds --from-file=credentials=./<JSON file name>
```

### Create a ProviderConfig object

Apply the secret in a `ProviderConfig` Kubernetes configuration file.

```yaml
apiVersion: appclacks.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: appclacks-creds
      namespace: crossplane-system
      key: credentials

```

**Note:** the `spec.credentials.secretRef.name` must match the `name` in the `kubectl create secret generic <name>` command.
