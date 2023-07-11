# Test skeleton with Terratest for Terraform modules for VM-Series on AWS, Azure, GCP

## Description

Test skeleton with Terratest in Go used to execute integration and e2e tests for VM-Series modules:
- [terraform-aws-vmseries-modules](https://github.com/PaloAltoNetworks/terraform-aws-vmseries-modules)
- [terraform-azurerm-vmseries-modules](https://github.com/PaloAltoNetworks/terraform-azurerm-vmseries-modules)
- [terraform-google-vmseries-modules](https://github.com/PaloAltoNetworks/terraform-google-vmseries-modules)

## Usage

Terratest can be used to check whole examples and single modules. Below there are few samples of usage:
- [execute only with plan for example](samples/vmseries_example_plan_and_deploy/)

```
cd samples/vmseries_example_plan_and_deploy
go test -v -timeout 30m -count=1 ./...
```

- [deploy whole example](samples/vmseries_example_plan_and_deploy/)

```
cd samples/vmseries_example_plan_and_deploy
DO_APPLY=true go test -v -timeout 30m -count=1 ./...
```

- [verify module by checking errors from Terraform plan](samples/vmseries_module_check_terraform_plan_errors/)

```
cd samples/vmseries_module_check_terraform_plan_errors
go test -v -count=1 ./...
```

- [verify module by checking output and accessing URL after deploying VM-Series](samples/vmseries_module_check_terraform_output_and_vmseries_url/)

```
cd samples/vmseries_module_check_terraform_output_and_vmseries_url
go test -v -timeout 30m -count=1 ./...
```


- [verify module by deplyoing additional changes after initial deployment](samples/vmseries_module_check_additional_changes_after_deployment/)

```
cd samples/vmseries_module_check_additional_changes_after_deployment
go test -v -timeout 30m -count=1 ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details