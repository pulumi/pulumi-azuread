module github.com/pulumi/pulumi-azuread/provider/v5

go 1.16

require (
	github.com/hashicorp/terraform-provider-azuread/shim v0.0.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.21.1-0.20220429072011-9483b81086cd
	github.com/pulumi/pulumi/sdk/v3 v3.30.1-0.20220425233152-77eb530a1117
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211230170131-3a7c83bfab87
	github.com/hashicorp/terraform-provider-azuread/shim => ./shim
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
