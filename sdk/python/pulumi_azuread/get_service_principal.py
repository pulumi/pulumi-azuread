# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetServicePrincipalResult',
    'AwaitableGetServicePrincipalResult',
    'get_service_principal',
]

@pulumi.output_type
class GetServicePrincipalResult:
    """
    A collection of values returned by getServicePrincipal.
    """
    def __init__(__self__, app_roles=None, application_id=None, display_name=None, id=None, oauth2_permission_scopes=None, oauth2_permissions=None, object_id=None):
        if app_roles and not isinstance(app_roles, list):
            raise TypeError("Expected argument 'app_roles' to be a list")
        pulumi.set(__self__, "app_roles", app_roles)
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if oauth2_permission_scopes and not isinstance(oauth2_permission_scopes, list):
            raise TypeError("Expected argument 'oauth2_permission_scopes' to be a list")
        pulumi.set(__self__, "oauth2_permission_scopes", oauth2_permission_scopes)
        if oauth2_permissions and not isinstance(oauth2_permissions, list):
            raise TypeError("Expected argument 'oauth2_permissions' to be a list")
        if oauth2_permissions is not None:
            warnings.warn("""[NOTE] The `oauth2_permissions` block has been renamed to `oauth2_permission_scopes` and moved to the `api` block. `oauth2_permissions` will be removed in version 2.0 of the AzureAD provider.""", DeprecationWarning)
            pulumi.log.warn("""oauth2_permissions is deprecated: [NOTE] The `oauth2_permissions` block has been renamed to `oauth2_permission_scopes` and moved to the `api` block. `oauth2_permissions` will be removed in version 2.0 of the AzureAD provider.""")

        pulumi.set(__self__, "oauth2_permissions", oauth2_permissions)
        if object_id and not isinstance(object_id, str):
            raise TypeError("Expected argument 'object_id' to be a str")
        pulumi.set(__self__, "object_id", object_id)

    @property
    @pulumi.getter(name="appRoles")
    def app_roles(self) -> Sequence['outputs.GetServicePrincipalAppRoleResult']:
        """
        A collection of `app_roles` blocks as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
        """
        return pulumi.get(self, "app_roles")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> str:
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name for the permission that appears in the admin consent and app assignment experiences.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="oauth2PermissionScopes")
    def oauth2_permission_scopes(self) -> Sequence['outputs.GetServicePrincipalOauth2PermissionScopeResult']:
        """
        A collection of OAuth 2.0 delegated permissions exposed by the associated Application. Each permission is covered by an `oauth2_permission_scopes` block as documented below.
        """
        return pulumi.get(self, "oauth2_permission_scopes")

    @property
    @pulumi.getter(name="oauth2Permissions")
    def oauth2_permissions(self) -> Sequence['outputs.GetServicePrincipalOauth2PermissionResult']:
        """
        (**Deprecated**) A collection of OAuth 2.0 permissions exposed by the associated Application. Each permission is covered by an `oauth2_permissions` block as documented below. Deprecated in favour of `oauth2_permission_scopes`.
        """
        return pulumi.get(self, "oauth2_permissions")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> str:
        """
        The Object ID for the Service Principal.
        """
        return pulumi.get(self, "object_id")


class AwaitableGetServicePrincipalResult(GetServicePrincipalResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServicePrincipalResult(
            app_roles=self.app_roles,
            application_id=self.application_id,
            display_name=self.display_name,
            id=self.id,
            oauth2_permission_scopes=self.oauth2_permission_scopes,
            oauth2_permissions=self.oauth2_permissions,
            object_id=self.object_id)


def get_service_principal(application_id: Optional[str] = None,
                          display_name: Optional[str] = None,
                          oauth2_permission_scopes: Optional[Sequence[pulumi.InputType['GetServicePrincipalOauth2PermissionScopeArgs']]] = None,
                          oauth2_permissions: Optional[Sequence[pulumi.InputType['GetServicePrincipalOauth2PermissionArgs']]] = None,
                          object_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServicePrincipalResult:
    """
    Gets information about an existing Service Principal associated with an Application within Azure Active Directory.

    > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

    ## Example Usage
    ### By Application Display Name)

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principal(display_name="my-awesome-application")
    ```
    ### By Application ID)

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principal(application_id="00000000-0000-0000-0000-000000000000")
    ```
    ### By Object ID)

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principal(object_id="00000000-0000-0000-0000-000000000000")
    ```


    :param str application_id: The ID of the Azure AD Application.
    :param str display_name: The Display Name of the Azure AD Application associated with this Service Principal.
    :param Sequence[pulumi.InputType['GetServicePrincipalOauth2PermissionScopeArgs']] oauth2_permission_scopes: A collection of OAuth 2.0 delegated permissions exposed by the associated Application. Each permission is covered by an `oauth2_permission_scopes` block as documented below.
    :param Sequence[pulumi.InputType['GetServicePrincipalOauth2PermissionArgs']] oauth2_permissions: (**Deprecated**) A collection of OAuth 2.0 permissions exposed by the associated Application. Each permission is covered by an `oauth2_permissions` block as documented below. Deprecated in favour of `oauth2_permission_scopes`.
    :param str object_id: The ID of the Azure AD Service Principal.
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    __args__['displayName'] = display_name
    __args__['oauth2PermissionScopes'] = oauth2_permission_scopes
    __args__['oauth2Permissions'] = oauth2_permissions
    __args__['objectId'] = object_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuread:index/getServicePrincipal:getServicePrincipal', __args__, opts=opts, typ=GetServicePrincipalResult).value

    return AwaitableGetServicePrincipalResult(
        app_roles=__ret__.app_roles,
        application_id=__ret__.application_id,
        display_name=__ret__.display_name,
        id=__ret__.id,
        oauth2_permission_scopes=__ret__.oauth2_permission_scopes,
        oauth2_permissions=__ret__.oauth2_permissions,
        object_id=__ret__.object_id)
