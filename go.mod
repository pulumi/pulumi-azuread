module github.com/pulumi/pulumi-azuread

replace (
	contrib.go.opencensus.io/exporter/ocagent => contrib.go.opencensus.io/exporter/ocagent v0.4.12
	github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v31.1.0+incompatible
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2-0.20190403091019-9b3cdde74fbe
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

require (
	github.com/Azure/go-autorest/autorest/azure/auth v0.1.0 // indirect
	github.com/hashicorp/terraform v0.12.7
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.1.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20191030013051-eccdb58dc332
	github.com/terraform-providers/terraform-provider-azuread v0.6.0
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7 // indirect
)

go 1.13
