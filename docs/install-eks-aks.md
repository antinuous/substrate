# Installing on EKS and AKS

Use `ate-setup deploy ate-system` with `--platform eks` or `--platform aks`.
GKE remains the default. `--kind` is supported only with GKE.
EKS and AKS use the Envoy dataplane. `agentgateway` is supported only on GKE.

EKS requires an S3 region and the exact service-account JWT issuer published by
the cluster. AKS also requires the AWS role ARN used for S3 workload identity.
For example:

```sh
ate-setup deploy ate-system \
  --platform eks \
  --s3-region us-east-1 \
  --expected-jwt-issuer "$EXPECTED_JWT_ISSUER" \
  --external-store-secret \
  --credential-provider='{"name":"k8s.io"}' \
  --image-repo ghcr.io/antinuous/substrate \
  --image-tag vX.Y.Z

ate-setup deploy ate-system \
  --platform aks \
  --s3-region us-east-1 \
  --s3-role-arn "$AWS_ROLE_ARN" \
  --expected-jwt-issuer "$EXPECTED_JWT_ISSUER" \
  --external-store-secret \
  --credential-provider='{"name":"k8s.io"}' \
  --image-repo ghcr.io/antinuous/substrate \
  --image-tag vX.Y.Z
```

The AKS workload token uses audience `sts.amazonaws.com`. EKS credentials come
from an EKS Pod Identity association configured outside Substrate.

## External MySQL Secret

Create `ate-api-server-secret-envvars` in the install namespace before running
ate-setup with `--external-store-secret`. It must contain non-empty keys:

```text
ATE_API_STORE_BACKEND=mysql
ATE_API_MYSQL_READ_WRITE_CONNECTION_STRING=<read-write DSN>
ATE_API_MYSQL_OWNER_CONNECTION_STRING=<owner DSN>
```

The installer validates this Secret and leaves it unchanged. The DSNs must use
an unsharded keyspace. For PlanetScale, set `tls=true` and leave the server CA
file unset. See [MySQL storage](mysql.md) for connection and TLS requirements.

## Autoscaled sandbox nodes

`ate-setup` labels nodes that exist during installation with
`ate.dev/substrate-version`. It does not label nodes created later, and
atecontroller does not add the version label to autoscaled nodes. Configure
the node-pool template to create sandbox nodes with
`ate.dev/substrate-version` set to the installed build's
`versionlabel.Value(version)` and `antinuous.io/pool=sandbox`.

Sandbox nodes should carry the taint
`ate.dev/sandboxClass=gvisor:NoSchedule`. Worker pods and atelet tolerate this
taint.
