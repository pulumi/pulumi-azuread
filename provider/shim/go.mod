module github.com/hashicorp/terraform-provider-azuread/shim

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.8.0
	github.com/hashicorp/terraform-provider-azuread v1.6.1-0.20220107004207-eeb7181dd0a9
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
