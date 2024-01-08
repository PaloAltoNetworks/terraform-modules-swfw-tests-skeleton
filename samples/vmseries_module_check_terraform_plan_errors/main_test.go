package vmseries_module_check_terraform_plan_errors

import (
	"testing"

	"github.com/PaloAltoNetworks/terraform-modules-vmseries-tests-skeleton/pkg/testskeleton"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestVmseriesModuleCheckTerraformPlanError(t *testing.T) {
	// define options for Terraform
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: ".",
		VarFiles:     []string{"test.tfvars"},
		Vars: map[string]interface{}{
			"name_prefix": "test_mod_vm_plan_error_",
		},
		Logger:               logger.Default,
		Lock:                 true,
		Upgrade:              true,
		SetVarsAfterVarFiles: true,
	})

	// prepare list of items to check
	assertList := []testskeleton.AssertExpression{
		{
			Operation:     testskeleton.ErrorContains,
			ExpectedValue: "No value for required variable",
			Message:       "3 variables are required: vmseries, vmseries_version, bootstrap_options",
		},
	}

	// deploy test infrastructure and verify outputs and check if there are no planned changes after deployment
	testskeleton.PlanInfraCheckErrors(t, terraformOptions, assertList, "VM-Series plan deployment should fail without VM-Series configuration")
}
