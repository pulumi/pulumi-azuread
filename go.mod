module github.com/pulumi/pulumi-azuread

go 1.12

require (
	github.com/hashicorp/terraform-plugin-sdk v1.1.0
	github.com/pkg/errors v0.8.1
<<<<<<< HEAD
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform-bridge v1.5.3-0.20200116214742-68c8f56da38a
=======
	github.com/pulumi/pulumi v1.9.1
	github.com/pulumi/pulumi-terraform-bridge v1.6.5
>>>>>>> master
	github.com/terraform-providers/terraform-provider-azuread v0.7.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
