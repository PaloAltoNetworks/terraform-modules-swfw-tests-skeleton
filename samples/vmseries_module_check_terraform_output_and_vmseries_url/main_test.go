package vmseries_module_check_terraform_output_and_vmseries_url

import (
	"testing"

	"github.com/PaloAltoNetworks/terraform-modules-swfw-tests-skeleton/pkg/helpers"
	"github.com/PaloAltoNetworks/terraform-modules-swfw-tests-skeleton/pkg/testskeleton"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestVmseriesModuleCheckTerraformOutputAndVmseriesUrl(t *testing.T) {
	// define options for Terraform
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: ".",
		VarFiles:     []string{"test.tfvars"},
		Vars: map[string]interface{}{
			"name_prefix": "test_mod_vm_out_url_",
		},
		Logger:               logger.Default,
		Lock:                 true,
		Upgrade:              true,
		SetVarsAfterVarFiles: true,
	})

	// prepare list of items to check
	assertList := []testskeleton.AssertExpression{
		// check VM-Series URL
		{
			OutputName: "vmseries_url",
			Operation:  testskeleton.NotEmpty,
		},
		// check VM-Series SSH
		{
			OutputName: "vmseries_ssh",
			Operation:  testskeleton.NotEmpty,
		},
		// check access to login page in web UI for VM-Series
		{
			Operation:  testskeleton.CheckFunctionWithOutput,
			Check:      helpers.CheckHttpGetWebApp,
			OutputName: "vmseries_url",
			Message:    "After bootstrapping, which takes few minutes, web UI for VM-Series should be accessible",
		},
		// check access via SSH to VM-Series
		{
			Operation:  testskeleton.CheckFunctionWithOutput,
			Check:      helpers.CheckTcpPortOpened,
			OutputName: "vmseries_ssh",
			Message:    "After bootstrapping, which takes few minutes, SSH for VM-Series should be accessible",
		},
	}

	// deploy test infrastructure and verify outputs and check if there are no planned changes after deployment
	testskeleton.DeployInfraCheckOutputsVerifyChanges(t, terraformOptions, assertList)
}
