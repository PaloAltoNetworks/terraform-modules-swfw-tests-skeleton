package vmseries_module_check_additional_changes_after_deployment

import (
	"testing"

	"github.com/PaloAltoNetworks/terraform-modules-vmseries-tests-skeleton/pkg/testskeleton"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	tfjson "github.com/hashicorp/terraform-json"
)

func TestVmseriesModuleCheckTerraformAdditionalChangesAfterDeployment(t *testing.T) {
	// define options for Terraform
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: ".",
		VarFiles:     []string{"test.tfvars"},
		Vars: map[string]interface{}{
			"name_prefix": "test_mod_vm_add_changes_",
		},
		Logger:               logger.Default,
		Lock:                 true,
		Upgrade:              true,
		SetVarsAfterVarFiles: true,
	})

	// prepare list of items to check
	assertList := []testskeleton.AssertExpression{}

	// prepare additional changes deployed after
	additionalChangesAfterDeployment := []testskeleton.AdditionalChangesAfterDeployment{
		// check adding new route
		{
			AdditionalVarsValues: map[string]interface{}{
				"name_sufix": "_terratest",
			},
			FileNameWithTfCode: "panorama_routes.tf.temp",
			ChangedResources: []testskeleton.ChangedResource{
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_network_interface.this[\"mgmt\"]",
					Action: tfjson.ActionUpdate,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_eip.this[\"mgmt\"]",
					Action: tfjson.ActionUpdate,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_instance.this",
					Action: tfjson.ActionUpdate,
				},
				{
					Name:   "module.panorama_vpc_routes[\"mgmt_10.80.10.0/24\"].aws_route.this[\"us-east-1a\"]",
					Action: tfjson.ActionCreate,
				},
			},
		},
		// check removing route
		{
			AdditionalVarsValues: map[string]interface{}{
				"name_sufix": "",
			},
			ChangedResources: []testskeleton.ChangedResource{
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_network_interface.this[\"mgmt\"]",
					Action: tfjson.ActionUpdate,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_eip.this[\"mgmt\"]",
					Action: tfjson.ActionUpdate,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_instance.this",
					Action: tfjson.ActionUpdate,
				},
				{
					Name:   "module.panorama_vpc_routes[\"mgmt_10.80.10.0/24\"].aws_route.this[\"us-east-1a\"]",
					Action: tfjson.ActionDelete,
				},
			},
		},
		// check removing public IP from mgmt interface
		{
			AdditionalVarsValues: map[string]interface{}{
				"override_and_disable_mgmt_create_public_ip": true,
			},
			ChangedResources: []testskeleton.ChangedResource{
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_eip.this[\"mgmt\"]",
					Action: tfjson.ActionDelete,
				},
			},
		},
		// check adding public IP from mgmt interface
		{
			AdditionalVarsValues: map[string]interface{}{
				"override_and_disable_mgmt_create_public_ip": false,
			},
			ChangedResources: []testskeleton.ChangedResource{
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_eip.this[\"mgmt\"]",
					Action: tfjson.ActionCreate,
				},
			},
		},
		// remove security group rules
		{
			UseVarFiles: []string{"test.tfvars", "security_groups.tfvars"},
			ChangedResources: []testskeleton.ChangedResource{
				{
					Name:   "module.security_vpc.aws_security_group.this[\"vmseries_mgmt\"]",
					Action: tfjson.ActionUpdate,
				},
			},
		},
		// add security group rules
		{
			UseVarFiles: []string{"test.tfvars"},
			ChangedResources: []testskeleton.ChangedResource{
				{
					Name:   "module.security_vpc.aws_security_group.this[\"vmseries_mgmt\"]",
					Action: tfjson.ActionUpdate,
				},
			},
		},
		// add interfaces to the firewall
		{
			UseVarFiles: []string{"test.tfvars", "network_interfaces.tfvars"},
			ChangedResources: []testskeleton.ChangedResource{
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_network_interface.this[\"data1\"]",
					Action: tfjson.ActionCreate,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_network_interface_attachment.this[\"mgmt\"]",
					Action: tfjson.ActionCreate,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_instance.this",
					Action: tfjson.ActionCreate,
				},
			},
		},
		// remove interfaces to the firewall
		{
			UseVarFiles: []string{"test.tfvars"},
			ChangedResources: []testskeleton.ChangedResource{
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_network_interface.this[\"data1\"]",
					Action: tfjson.ActionDelete,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_network_interface_attachment.this[\"mgmt\"]",
					Action: tfjson.ActionDelete,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_instance.this",
					Action: tfjson.ActionCreate,
				},
			},
		},
		// change userdata parameters - currently by default user_data_replace_on_change is set to false,
		// so changing user data do not trigger replacing EC2 instancee
		{
			AdditionalVarsValues: map[string]interface{}{
				"bootstrap_options": "plugin-op-commands=aws-gwlb-inspect:enable,aws-gwlb-overlay-routing:enable;type=dhcp-client",
			},
			ChangedResources: []testskeleton.ChangedResource{},
		},
		// add tags
		{
			UseVarFiles: []string{"test.tfvars", "global_tags.tfvars"},
			ChangedResources: []testskeleton.ChangedResource{
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_network_interface.this[\"mgmt\"]",
					Action: tfjson.ActionUpdate,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_eip.this[\"mgmt\"]",
					Action: tfjson.ActionUpdate,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_instance.this",
					Action: tfjson.ActionUpdate,
				},
			},
		},
		// remove tags
		{
			UseVarFiles: []string{"test.tfvars"},
			ChangedResources: []testskeleton.ChangedResource{
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_network_interface.this[\"mgmt\"]",
					Action: tfjson.ActionUpdate,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_eip.this[\"mgmt\"]",
					Action: tfjson.ActionUpdate,
				},
				{
					Name:   "module.vmseries[\"vmseries01\"].aws_instance.this",
					Action: tfjson.ActionUpdate,
				},
			},
		},
	}

	// deploy test infrastructure and verify outputs and check if there are no planned changes after deployment
	testskeleton.DeployInfraCheckOutputsVerifyChangesDeployChanges(t, terraformOptions, assertList, additionalChangesAfterDeployment)
}
