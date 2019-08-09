module github.com/pulumi/pulumi-azuread

go 1.12

replace (
	contrib.go.opencensus.io/exporter/ocagent => contrib.go.opencensus.io/exporter/ocagent v0.4.12
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v11.1.2+incompatible
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2-0.20190403091019-9b3cdde74fbe
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)

require (
	github.com/Nvveen/Gotty v0.0.0-20170406111628-a8b993ba6abd // indirect
	github.com/apache/thrift v0.12.0 // indirect
	github.com/go-ini/ini v1.31.0 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/hashicorp/terraform v0.12.6
	github.com/openzipkin/zipkin-go v0.1.6 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v0.17.28-0.20190731182900-6804d640fc7c
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190809001837-9db2fc93cdf5
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-azuread v0.5.1
	golang.org/x/build v0.0.0-20190314133821-5284462c4bec // indirect
	gopkg.in/vmihailenco/msgpack.v2 v2.9.1 // indirect
	labix.org/v2/mgo v0.0.0-20140701140051-000000000287 // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)
