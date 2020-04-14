module github.com/pulumi/pulumi-azuread/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.6.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.0.0-20200414133247-94746eebdf3b
	github.com/pulumi/pulumi/sdk/v2 v2.0.0-beta.3
	github.com/terraform-providers/terraform-provider-azuread v0.8.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
