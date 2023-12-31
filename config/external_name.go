/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/config"
)

func qemuVMConf() config.ExternalName {
	e := config.NameAsIdentifier
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		if name, ok := tfstate["name"].(string); ok && name != "" {
			return name, nil
		}
		return "", errors.New("cannot find name in tfstate")
	}
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		targetNode, ok := parameters["target_node"]
		if !ok {
			return "", errors.New("cannot get target_node")
		}
		vmId, ok := parameters["vmid"]
		if !ok {
			return "", errors.New("cannot get vmid")
		}
		return fmt.Sprintf("%s/qemu/%.0f", targetNode, vmId), nil
	}
	return e
}

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"proxmox_vm_qemu": qemuVMConf(),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
