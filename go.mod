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
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/hashicorp/terraform v0.12.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v0.17.22-0.20190702185104-ebceea93a5da
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190703150544-a9a9ca8157ca
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-azuread v0.4.0
	labix.org/v2/mgo v0.0.0-20140701140051-000000000287 // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)
