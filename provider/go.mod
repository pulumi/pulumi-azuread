module github.com/pulumi/pulumi-azuread/provider/v3

go 1.15

require (
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.17.1-0.20201219002337-a0c25bd3e147
	github.com/pulumi/pulumi/sdk/v2 v2.17.0
	github.com/terraform-providers/terraform-provider-azuread/shim v0.0.0
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-azuread/shim => ./shim
)
