module github.com/hashicorp/terraform-provider-azuread/shim

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.8.0
	github.com/hashicorp/terraform-provider-azuread v1.6.1-0.20220311140254-fef2a3399fc1
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211230170131-3a7c83bfab87
