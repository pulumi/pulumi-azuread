module github.com/pulumi/pulumi-azuread/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.15.1
	github.com/pulumi/pulumi/sdk/v2 v2.15.1-0.20201202214525-260620430c4c
	github.com/terraform-providers/terraform-provider-azuread v0.11.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
