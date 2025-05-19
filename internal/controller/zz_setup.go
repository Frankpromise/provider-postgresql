// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	database "github.com/frankpromise/provider-postgresql/internal/controller/postgresql/database"
	extension "github.com/frankpromise/provider-postgresql/internal/controller/postgresql/extension"
	grant "github.com/frankpromise/provider-postgresql/internal/controller/postgresql/grant"
	role "github.com/frankpromise/provider-postgresql/internal/controller/postgresql/role"
	schema "github.com/frankpromise/provider-postgresql/internal/controller/postgresql/schema"
	providerconfig "github.com/frankpromise/provider-postgresql/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		extension.Setup,
		grant.Setup,
		role.Setup,
		schema.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
