package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/frankpromise/provider-postgresql/apis/v1beta1"
)

const (
	host              = "host"
	port              = "port"
	username          = "username"
	password          = "password"
	database          = "database"
	sslmode           = "sslmode"
	azureIdentityAuth = "azure_identity_auth"
	azureTenantId     = "azure_tenant_id"

	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal PostgreSQL credentials as JSON"
	errMissingTenantID      = "azure_identity_auth is true but azure_tenant_id is missing"
)

// TerraformSetupBuilder returns a terraform.SetupFn for Cosmos DB for PostgreSQL
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, k8s client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		ref := mg.GetProviderConfigReference()
		if ref == nil {
			return ps, errors.New(errNoProviderConfig)
		}

		pc := &v1beta1.ProviderConfig{}
		if err := k8s.Get(ctx, types.NamespacedName{Name: ref.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		if err := resource.NewProviderConfigUsageTracker(k8s, &v1beta1.ProviderConfigUsage{}).Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, k8s, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Optional Azure identity validation
		if creds[azureIdentityAuth] == "true" && creds[azureTenantId] == "" {
			return ps, fmt.Errorf(errMissingTenantID)
		}

		// Only relevant Cosmos DB PostgreSQL config
		ps.Configuration = map[string]any{}
		for _, key := range []string{
			host, port, database, username, password, sslmode,
			azureIdentityAuth, azureTenantId,
		} {
			if val, ok := creds[key]; ok {
				ps.Configuration[key] = val
			}
		}

		return ps, nil
	}
}
