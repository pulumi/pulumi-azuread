module github.com/pulumi/pulumi-azuread

go 1.12

require (
	github.com/hashicorp/terraform-plugin-sdk v1.1.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform-bridge v1.5.0
	github.com/terraform-providers/terraform-provider-azuread v0.7.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
