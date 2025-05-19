// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	database "github.com/frankpromise/provider-postgresql/internal/controller/db/database"
	extension "github.com/frankpromise/provider-postgresql/internal/controller/extension/extension"
	grant "github.com/frankpromise/provider-postgresql/internal/controller/grant/grant"
	role "github.com/frankpromise/provider-postgresql/internal/controller/grant/role"
	providerconfig "github.com/frankpromise/provider-postgresql/internal/controller/providerconfig"
	rolerole "github.com/frankpromise/provider-postgresql/internal/controller/role/role"
	schema "github.com/frankpromise/provider-postgresql/internal/controller/schema/schema"
	mapping "github.com/frankpromise/provider-postgresql/internal/controller/user/mapping"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		extension.Setup,
		grant.Setup,
		role.Setup,
		providerconfig.Setup,
		rolerole.Setup,
		schema.Setup,
		mapping.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
