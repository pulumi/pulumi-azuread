module github.com/pulumi/pulumi-azuread/provider/v5

go 1.16

require (
	github.com/hashicorp/terraform-provider-azuread/shim v0.0.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.15.0
	github.com/pulumi/pulumi/sdk/v3 v3.21.0
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/hashicorp/terraform-provider-azuread/shim => ./shim
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
