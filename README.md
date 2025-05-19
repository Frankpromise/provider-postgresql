# Generating a Crossplane Provider for PostgreSQL

This beginner-friendly guide walks you through generating a Crossplane provider using [Upjet](https://github.com/crossplane/upjet), based on the [Terraform PostgreSQL provider](https://registry.terraform.io/providers/cyrilgdn/postgresql/latest/docs).

---

## ✨ Overview

**Crossplane** is a Kubernetes-native infrastructure management tool. Instead of managing your infrastructure with Terraform directly, you define everything as Kubernetes manifests. **Upjet** helps you generate Crossplane providers from Terraform providers, reducing the effort to integrate with existing Terraform ecosystems.

This guide adapts the official Upjet postgresql to build a PostgreSQL Crossplane provider.

---

## ✏️ 1. Prepare Your Provider Repository

### What we’re doing:

You’re forking a boilerplate repo that scaffolds everything needed to create a new Crossplane provider based on a Terraform provider.

1. Click "**Use this postgresql**" on [provider-postgresql](https://github.com/Frankpromise/provider-postgresql) to create a new repo named `provider-postgresql`.

2. Clone your repo locally and navigate into it:

   ```bash
   git clone https://github.com/<your-org>/provider-postgresql.git
   cd provider-postgresql
   ```

3. Fetch the required build submodules (used for tooling and code generation):

   ```bash
   make submodules
   ```

4. Run the prepare script to initialize naming:

   ```bash
   ./hack/prepare.sh
   ```

   This will ask for the Terraform source, provider name, and API group. It renames placeholders in the codebase accordingly.

5. Edit your `Makefile` and update the following variables:

| Variable                           | Explanation                                               |
| ---------------------------------- | --------------------------------------------------------- |
| `TERRAFORM_PROVIDER_SOURCE`        | Terraform registry source (e.g. `cyrilgdn/postgresql`)    |
| `TERRAFORM_PROVIDER_REPO`          | GitHub repo of the Terraform provider                     |
| `TERRAFORM_PROVIDER_VERSION`       | Specific version to generate against                      |
| `TERRAFORM_PROVIDER_DOWNLOAD_NAME` | Name of the Terraform binary                              |
| `TERRAFORM_NATIVE_PROVIDER_BINARY` | Full name of binary including version                     |
| `TERRAFORM_DOCS_PATH`              | Optional: relative path to Terraform docs (not mandatory) |

---

## 📂 2. Configure Provider Resources

### 2.1 Add ProviderConfig logic

**What is ProviderConfig?**
This is the Kubernetes resource that provides credentials and connection details to the Terraform provider (e.g. your PostgreSQL DB info).

In `internal/clients/postgresql.go`:

```go
const (
  keyHost     = "host"
  keyPort     = "port"
  keyUsername = "username"
  keyPassword = "password"
  keySSLMode  = "sslmode"
)

func TerraformSetupBuilder(...) terraform.SetupFn {
  return func(...) (terraform.Setup, error) {
    creds, err := terraform.CredentialsFromProviderConfig(...)
    s.Configuration = map[string]any{
      "host": creds[keyHost],
      ...
    }
    return *s, nil
  }
}
```

### 2.2 Configure external names

**What is an external name?**
The "external name" is how Crossplane knows what name to use in the external system (e.g., the name of the DB in PostgreSQL).

In `config/external_name.go`, add:

```go
var ExternalNameConfigs = map[string]config.ExternalName{
  "postgresql_database": config.NameAsIdentifier,
  "postgresql_role": config.NameAsIdentifier,
}
```

### 2.3 Configure resource groups

**Why?** Upjet by default groups all resources under the main provider group. For clarity and modularity, we move resources like `Database` to their own group (e.g., `db.postgresql.crossplane.io`).

```bash
mkdir -p config/database config/role
```

Then create `config/role/config.go`:

```go
func Configure(p *config.Provider) {
  p.AddResourceConfigurator("postgresql_role", func(r *config.Resource) {
    r.ShortGroup = "role"
  })
}
```

In `config/provider.go`, register these configurations.

---

## ⚡ 3. Generate the Provider

**Why?** Upjet will generate:

* Kubernetes CRDs
* Go types for CRDs
* Reconciler controllers

Ensure `goimports` is installed:

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

Then run:

```bash
make generate
```

---

## ✅ 4. Test Your Provider

Testing can be done using a **local KinD (Kubernetes in Docker) cluster**, leveraging the automatically generated example manifests located in the `example-generated` folder. These files demonstrate how to provision PostgreSQL resources (like databases and roles) using your new provider. how to provision PostgreSQL resources (like databases and roles) using your new provider.

### 4.1 Start KinD

Make sure you have [KinD](https://kind.sigs.k8s.io/) installed and create a local cluster:

```bash
kind create cluster --name crossplane-provider-test
```

### 4.2 Apply the CRDs

These are the custom resource definitions your provider generated:

```bash
kubectl apply -f package/crds
```

### 4.4 Install the provider

> Note: If you're using pre-generated resources from `example-generated`, you don't need to deploy a controller manually.
> You can skip this step if no `install.yaml` exists or you're not testing a controller build.

### 4.5 Apply sample resources

Apply sample resources like a test database or role:

```bash
kubectl apply -f example-generated/database/database.yaml
kubectl apply -f example-generated/role/role.yaml
```

### 4.6 Observe resources

Check the state of managed resources:

```bash
kubectl get databases.db.postgresql.crossplane.io
kubectl describe database testdb
```

Once resources show `READY=True`, your provider is working!

---

## 💡 Using This Provider with Azure Cosmos DB for PostgreSQL

> ⚠️ Note: The `provider-postgresql` only manages **resources inside an existing PostgreSQL instance**. It does **not** create PostgreSQL servers.

### ✅ Creating Users and Passwords

To create PostgreSQL users and assign them to roles, you will typically need two separate resources:

1. A `Role` resource to create the PostgreSQL user (with login and password)
2. A `Mapping` resource to associate the user with a role (if role-based access is used)

Here’s an example for creating a user:

```yaml
apiVersion: role.postgresql.crossplane.io/v1alpha1
kind: Role
metadata:
  name: user123
  annotations:
    crossplane.io/external-name: user123
spec:
  forProvider:
    login: true
    password: StrongP@ssw0rd123
  providerConfigRef:
    name: azure-postgres
```

You can also create the role `readonly` before mapping:

```yaml
apiVersion: role.postgresql.crossplane.io/v1alpha1
kind: Role
metadata:
  name: readonly
  annotations:
    crossplane.io/external-name: readonly
spec:
  forProvider:
    login: false
  providerConfigRef:
    name: azure-postgres
```

Then you can map that user to the role:

```yamlyaml
apiVersion: user.postgresql.crossplane.io/v1alpha1
kind: Mapping
metadata:
  name: testmapping
  annotations:
    crossplane.io/external-name: user123:readonly
spec:
  forProvider:
    user: user123
    role: readonly
  providerConfigRef:
    name: azure-postgres
```

> ⚠️ Ensure the role `readonly` already exists or is also managed by a `Role` resource. You can create it using another manifest similar to the user one.

This setup will create a PostgreSQL user and assign it to the `readonly` role inside your Azure Cosmos DB for PostgreSQL instance.> ⚠️ Note: The `provider-postgresql` only manages **resources inside an existing PostgreSQL instance**. It does **not** create PostgreSQL servers.

To use this provider with **Azure Cosmos DB for PostgreSQL**, you must point it to your existing provisioned CosmosDB PostgreSQL instance:

### 1. Ensure Azure Cosmos DB for PostgreSQL is ready

* Your instance must be provisioned (this guide assumes you've already done this).
* Enable public access.
* Allow incoming traffic from your Crossplane cluster (e.g., KinD or remote IP range).

### 2. Create a `Secret` with credentials

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: azure-postgres-creds
  namespace: crossplane-system
stringData:
  creds: |
    {
      "host": "<your-postgres-host>.postgres.database.azure.com",
      "port": "5432",
      "username": "your-admin-user@<your-postgres-host>",
      "password": "your-password",
      "sslmode": "require"
    }
```

### 3. Create a `ProviderConfig`

```yaml
apiVersion: postgresql.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: azure-postgres
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: azure-postgres-creds
      key: creds
```

### 4. Create a PostgreSQL `Database` resource

```yaml
apiVersion: db.postgresql.crossplane.io/v1alpha1
kind: Database
metadata:
  name: exampledb
  annotations:
    crossplane.io/external-name: exampledb
spec:
  forProvider:
    owner: your-admin-user
  providerConfigRef:
    name: azure-postgres
```

Once applied, this will create a new database named `exampledb` inside your Azure Cosmos DB for PostgreSQL instance.

---

## 💪 Next Steps

* Add more resources (grants, schemas, extensions)
* Push image to registry and update `install.yaml`
