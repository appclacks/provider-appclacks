/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	command "github.com/azrod/provider-appclacks/internal/controller/healthcheck/command"
	dns "github.com/azrod/provider-appclacks/internal/controller/healthcheck/dns"
	http "github.com/azrod/provider-appclacks/internal/controller/healthcheck/http"
	tcp "github.com/azrod/provider-appclacks/internal/controller/healthcheck/tcp"
	tls "github.com/azrod/provider-appclacks/internal/controller/healthcheck/tls"
	providerconfig "github.com/azrod/provider-appclacks/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		command.Setup,
		dns.Setup,
		http.Setup,
		tcp.Setup,
		tls.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
