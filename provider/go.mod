module github.com/pulumi/pulumi-azuread/provider/v3

go 1.15

require (
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.17.0
	github.com/pulumi/pulumi/sdk/v2 v2.17.0
	github.com/terraform-providers/terraform-provider-azuread/shim v0.0.0
)

replace (
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-azuread/shim => ./shim
)
