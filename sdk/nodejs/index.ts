// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./administrativeUnit";
export * from "./administrativeUnitMember";
export * from "./appRoleAssignment";
export * from "./application";
export * from "./applicationCertificate";
export * from "./applicationFederatedIdentityCredential";
export * from "./applicationPassword";
export * from "./applicationPreAuthorized";
export * from "./conditionalAccessPolicy";
export * from "./customDirectoryRole";
export * from "./directoryRole";
export * from "./directoryRoleMember";
export * from "./getAdministrativeUnit";
export * from "./getApplication";
export * from "./getApplicationPublishedAppIds";
export * from "./getApplicationTemplate";
export * from "./getClientConfig";
export * from "./getDomains";
export * from "./getGroup";
export * from "./getGroups";
export * from "./getServicePrincipal";
export * from "./getServicePrincipals";
export * from "./getUser";
export * from "./getUsers";
export * from "./group";
export * from "./groupMember";
export * from "./invitation";
export * from "./namedLocation";
export * from "./provider";
export * from "./servicePrincipal";
export * from "./servicePrincipalCertificate";
export * from "./servicePrincipalDelegatedPermissionGrant";
export * from "./servicePrincipalPassword";
export * from "./user";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { AdministrativeUnit } from "./administrativeUnit";
import { AdministrativeUnitMember } from "./administrativeUnitMember";
import { AppRoleAssignment } from "./appRoleAssignment";
import { Application } from "./application";
import { ApplicationCertificate } from "./applicationCertificate";
import { ApplicationFederatedIdentityCredential } from "./applicationFederatedIdentityCredential";
import { ApplicationPassword } from "./applicationPassword";
import { ApplicationPreAuthorized } from "./applicationPreAuthorized";
import { ConditionalAccessPolicy } from "./conditionalAccessPolicy";
import { CustomDirectoryRole } from "./customDirectoryRole";
import { DirectoryRole } from "./directoryRole";
import { DirectoryRoleMember } from "./directoryRoleMember";
import { Group } from "./group";
import { GroupMember } from "./groupMember";
import { Invitation } from "./invitation";
import { NamedLocation } from "./namedLocation";
import { ServicePrincipal } from "./servicePrincipal";
import { ServicePrincipalCertificate } from "./servicePrincipalCertificate";
import { ServicePrincipalDelegatedPermissionGrant } from "./servicePrincipalDelegatedPermissionGrant";
import { ServicePrincipalPassword } from "./servicePrincipalPassword";
import { User } from "./user";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azuread:index/administrativeUnit:AdministrativeUnit":
                return new AdministrativeUnit(name, <any>undefined, { urn })
            case "azuread:index/administrativeUnitMember:AdministrativeUnitMember":
                return new AdministrativeUnitMember(name, <any>undefined, { urn })
            case "azuread:index/appRoleAssignment:AppRoleAssignment":
                return new AppRoleAssignment(name, <any>undefined, { urn })
            case "azuread:index/application:Application":
                return new Application(name, <any>undefined, { urn })
            case "azuread:index/applicationCertificate:ApplicationCertificate":
                return new ApplicationCertificate(name, <any>undefined, { urn })
            case "azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential":
                return new ApplicationFederatedIdentityCredential(name, <any>undefined, { urn })
            case "azuread:index/applicationPassword:ApplicationPassword":
                return new ApplicationPassword(name, <any>undefined, { urn })
            case "azuread:index/applicationPreAuthorized:ApplicationPreAuthorized":
                return new ApplicationPreAuthorized(name, <any>undefined, { urn })
            case "azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy":
                return new ConditionalAccessPolicy(name, <any>undefined, { urn })
            case "azuread:index/customDirectoryRole:CustomDirectoryRole":
                return new CustomDirectoryRole(name, <any>undefined, { urn })
            case "azuread:index/directoryRole:DirectoryRole":
                return new DirectoryRole(name, <any>undefined, { urn })
            case "azuread:index/directoryRoleMember:DirectoryRoleMember":
                return new DirectoryRoleMember(name, <any>undefined, { urn })
            case "azuread:index/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "azuread:index/groupMember:GroupMember":
                return new GroupMember(name, <any>undefined, { urn })
            case "azuread:index/invitation:Invitation":
                return new Invitation(name, <any>undefined, { urn })
            case "azuread:index/namedLocation:NamedLocation":
                return new NamedLocation(name, <any>undefined, { urn })
            case "azuread:index/servicePrincipal:ServicePrincipal":
                return new ServicePrincipal(name, <any>undefined, { urn })
            case "azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate":
                return new ServicePrincipalCertificate(name, <any>undefined, { urn })
            case "azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant":
                return new ServicePrincipalDelegatedPermissionGrant(name, <any>undefined, { urn })
            case "azuread:index/servicePrincipalPassword:ServicePrincipalPassword":
                return new ServicePrincipalPassword(name, <any>undefined, { urn })
            case "azuread:index/user:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azuread", "index/administrativeUnit", _module)
pulumi.runtime.registerResourceModule("azuread", "index/administrativeUnitMember", _module)
pulumi.runtime.registerResourceModule("azuread", "index/appRoleAssignment", _module)
pulumi.runtime.registerResourceModule("azuread", "index/application", _module)
pulumi.runtime.registerResourceModule("azuread", "index/applicationCertificate", _module)
pulumi.runtime.registerResourceModule("azuread", "index/applicationFederatedIdentityCredential", _module)
pulumi.runtime.registerResourceModule("azuread", "index/applicationPassword", _module)
pulumi.runtime.registerResourceModule("azuread", "index/applicationPreAuthorized", _module)
pulumi.runtime.registerResourceModule("azuread", "index/conditionalAccessPolicy", _module)
pulumi.runtime.registerResourceModule("azuread", "index/customDirectoryRole", _module)
pulumi.runtime.registerResourceModule("azuread", "index/directoryRole", _module)
pulumi.runtime.registerResourceModule("azuread", "index/directoryRoleMember", _module)
pulumi.runtime.registerResourceModule("azuread", "index/group", _module)
pulumi.runtime.registerResourceModule("azuread", "index/groupMember", _module)
pulumi.runtime.registerResourceModule("azuread", "index/invitation", _module)
pulumi.runtime.registerResourceModule("azuread", "index/namedLocation", _module)
pulumi.runtime.registerResourceModule("azuread", "index/servicePrincipal", _module)
pulumi.runtime.registerResourceModule("azuread", "index/servicePrincipalCertificate", _module)
pulumi.runtime.registerResourceModule("azuread", "index/servicePrincipalDelegatedPermissionGrant", _module)
pulumi.runtime.registerResourceModule("azuread", "index/servicePrincipalPassword", _module)
pulumi.runtime.registerResourceModule("azuread", "index/user", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("azuread", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:azuread") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
