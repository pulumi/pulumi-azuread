module github.com/pulumi/pulumi-azuread

go 1.12

replace (
	contrib.go.opencensus.io/exporter/ocagent => contrib.go.opencensus.io/exporter/ocagent v0.4.12
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2-0.20190403091019-9b3cdde74fbe
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
	github.com/terraform-providers/terraform-provider-azurerm => github.com/pulumi/terraform-provider-azurerm v0.0.0-20190417123607-dd01e8265e07
)

require (
	github.com/hashicorp/terraform v0.12.0-alpha4.0.20190417210818-177a7afb781f
	github.com/pulumi/pulumi v0.17.6-0.20190410045519-ef5e148a73c0
	github.com/pulumi/pulumi-terraform v0.14.1-dev.0.20190423021607-1090d14a55ea
	github.com/stretchr/testify v1.3.0
	github.com/terraform-providers/terraform-provider-azuread v0.3.1
)
