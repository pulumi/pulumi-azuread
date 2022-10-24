module github.com/pulumi/pulumi-azuread/examples/v5

go 1.16

require github.com/pulumi/pulumi/pkg/v3 v3.43.1

// get rid of dependabot high severity security alert #14
replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
