module github.com/terraform-providers/terraform-provider-azuread/shim

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.3.0
	github.com/terraform-providers/terraform-provider-azuread v1.2.1
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
