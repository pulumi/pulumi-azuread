// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azuread:index/application:Application":
		r = &Application{}
	case "azuread:index/applicationAppRole:ApplicationAppRole":
		r = &ApplicationAppRole{}
	case "azuread:index/applicationCertificate:ApplicationCertificate":
		r = &ApplicationCertificate{}
	case "azuread:index/applicationOAuth2Permission:ApplicationOAuth2Permission":
		r = &ApplicationOAuth2Permission{}
	case "azuread:index/applicationOauth2PermissionScope:ApplicationOauth2PermissionScope":
		r = &ApplicationOauth2PermissionScope{}
	case "azuread:index/applicationPassword:ApplicationPassword":
		r = &ApplicationPassword{}
	case "azuread:index/group:Group":
		r = &Group{}
	case "azuread:index/groupMember:GroupMember":
		r = &GroupMember{}
	case "azuread:index/servicePrincipal:ServicePrincipal":
		r = &ServicePrincipal{}
	case "azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate":
		r = &ServicePrincipalCertificate{}
	case "azuread:index/servicePrincipalPassword:ServicePrincipalPassword":
		r = &ServicePrincipalPassword{}
	case "azuread:index/user:User":
		r = &User{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:azuread" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"azuread",
		"index/application",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/applicationAppRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/applicationCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/applicationOAuth2Permission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/applicationOauth2PermissionScope",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/applicationPassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/groupMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/servicePrincipal",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/servicePrincipalCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/servicePrincipalPassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"azuread",
		&pkg{version},
	)
}
