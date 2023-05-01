// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

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
	case "azuread:index/accessPackage:AccessPackage":
		r = &AccessPackage{}
	case "azuread:index/accessPackageAssignmentPolicy:AccessPackageAssignmentPolicy":
		r = &AccessPackageAssignmentPolicy{}
	case "azuread:index/accessPackageCatalog:AccessPackageCatalog":
		r = &AccessPackageCatalog{}
	case "azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment":
		r = &AccessPackageCatalogRoleAssignment{}
	case "azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation":
		r = &AccessPackageResourceCatalogAssociation{}
	case "azuread:index/accessPackageResourcePackageAssociation:AccessPackageResourcePackageAssociation":
		r = &AccessPackageResourcePackageAssociation{}
	case "azuread:index/administrativeUnit:AdministrativeUnit":
		r = &AdministrativeUnit{}
	case "azuread:index/administrativeUnitMember:AdministrativeUnitMember":
		r = &AdministrativeUnitMember{}
	case "azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember":
		r = &AdministrativeUnitRoleMember{}
	case "azuread:index/appRoleAssignment:AppRoleAssignment":
		r = &AppRoleAssignment{}
	case "azuread:index/application:Application":
		r = &Application{}
	case "azuread:index/applicationCertificate:ApplicationCertificate":
		r = &ApplicationCertificate{}
	case "azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential":
		r = &ApplicationFederatedIdentityCredential{}
	case "azuread:index/applicationPassword:ApplicationPassword":
		r = &ApplicationPassword{}
	case "azuread:index/applicationPreAuthorized:ApplicationPreAuthorized":
		r = &ApplicationPreAuthorized{}
	case "azuread:index/claimsMappingPolicy:ClaimsMappingPolicy":
		r = &ClaimsMappingPolicy{}
	case "azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy":
		r = &ConditionalAccessPolicy{}
	case "azuread:index/customDirectoryRole:CustomDirectoryRole":
		r = &CustomDirectoryRole{}
	case "azuread:index/directoryRole:DirectoryRole":
		r = &DirectoryRole{}
	case "azuread:index/directoryRoleAssignment:DirectoryRoleAssignment":
		r = &DirectoryRoleAssignment{}
	case "azuread:index/directoryRoleMember:DirectoryRoleMember":
		r = &DirectoryRoleMember{}
	case "azuread:index/group:Group":
		r = &Group{}
	case "azuread:index/groupMember:GroupMember":
		r = &GroupMember{}
	case "azuread:index/invitation:Invitation":
		r = &Invitation{}
	case "azuread:index/namedLocation:NamedLocation":
		r = &NamedLocation{}
	case "azuread:index/servicePrincipal:ServicePrincipal":
		r = &ServicePrincipal{}
	case "azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate":
		r = &ServicePrincipalCertificate{}
	case "azuread:index/servicePrincipalClaimsMappingPolicyAssignment:ServicePrincipalClaimsMappingPolicyAssignment":
		r = &ServicePrincipalClaimsMappingPolicyAssignment{}
	case "azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant":
		r = &ServicePrincipalDelegatedPermissionGrant{}
	case "azuread:index/servicePrincipalPassword:ServicePrincipalPassword":
		r = &ServicePrincipalPassword{}
	case "azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate":
		r = &ServicePrincipalTokenSigningCertificate{}
	case "azuread:index/synchronizationJob:SynchronizationJob":
		r = &SynchronizationJob{}
	case "azuread:index/synchronizationSecret:SynchronizationSecret":
		r = &SynchronizationSecret{}
	case "azuread:index/user:User":
		r = &User{}
	case "azuread:index/userFlowAttribute:UserFlowAttribute":
		r = &UserFlowAttribute{}
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
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"azuread",
		"index/accessPackage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/accessPackageAssignmentPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/accessPackageCatalog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/accessPackageCatalogRoleAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/accessPackageResourceCatalogAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/accessPackageResourcePackageAssociation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/administrativeUnit",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/administrativeUnitMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/administrativeUnitRoleMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/appRoleAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/application",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/applicationCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/applicationFederatedIdentityCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/applicationPassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/applicationPreAuthorized",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/claimsMappingPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/conditionalAccessPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/customDirectoryRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/directoryRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/directoryRoleAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/directoryRoleMember",
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
		"index/invitation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/namedLocation",
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
		"index/servicePrincipalClaimsMappingPolicyAssignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/servicePrincipalDelegatedPermissionGrant",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/servicePrincipalPassword",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/servicePrincipalTokenSigningCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/synchronizationJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/synchronizationSecret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azuread",
		"index/userFlowAttribute",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"azuread",
		&pkg{version},
	)
}
