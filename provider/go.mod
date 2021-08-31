module github.com/pulumi/pulumi-azuread/provider/v5

go 1.16

require (
	github.com/hashicorp/terraform-provider-azuread/shim v0.0.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.4.1-0.20210714215802-5020116ac4e6
	github.com/pulumi/pulumi/pkg/v3 v3.7.1-0.20210714212650-083fc64ff547 // indirect
	github.com/pulumi/pulumi/sdk/v3 v3.6.2-0.20210712183851-30278f8e4c08
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/hashicorp/terraform-provider-azuread/shim => ./shim
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
