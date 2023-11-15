# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetDirectoryRolesResult',
    'AwaitableGetDirectoryRolesResult',
    'get_directory_roles',
    'get_directory_roles_output',
]

@pulumi.output_type
calass GetDirectoryRolesResult:
    """
    A collection of values returned by getDirectoryRoles.
    """
    def __init__(__self__, id=None, object_ids=None, roles=None, template_ids=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if object_ids and not isinstance(object_ids, list):
            raise TypeError("Expected argument 'object_ids' to be a list")
        pulumi.set(__self__, "object_ids", object_ids)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if template_ids and not isinstance(template_ids, list):
            raise TypeError("Expected argument 'template_ids' to be a list")
        pulumi.set(__self__, "template_ids", template_ids)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="objectIds")
    def object_ids(self) -> Sequence[str]:
        """
        The object IDs of the roles.
        """
        return pulumi.get(self, "object_ids")

    @property
    @pulumi.getter
    def roles(self) -> Sequence['outputs.GetDirectoryRolesRoleResult']:
        """
        A list of users. Each `role` object provides the attributes documented below.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="templateIds")
    def template_ids(self) -> Sequence[str]:
        """
        The template IDs of the roles.
        """
        return pulumi.get(self, "template_ids")


calass AwaitableGetDirectoryRolesResult(GetDirectoryRolesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDirectoryRolesResult(
            id=self.id,
            object_ids=self.object_ids,
            roles=self.roles,
            template_ids=self.template_ids)


def get_directory_roles(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDirectoryRolesResult:
    """
    Use this data source to access information about activated directory roles within Azure Active Directory.

    ## API Permissions

    The following API permissions are required in order to use this resource.

    When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.Read.Directory` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    current = azuread.get_directory_roles()
    pulumi.export("roles", current.object_ids)
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuread:index/getDirectoryRoles:getDirectoryRoles', __args__, opts=opts, typ=GetDirectoryRolesResult).value

    return AwaitableGetDirectoryRolesResult(
        id=pulumi.get(__ret__, 'id'),
        object_ids=pulumi.get(__ret__, 'object_ids'),
        roles=pulumi.get(__ret__, 'roles'),
        template_ids=pulumi.get(__ret__, 'template_ids'))


@_utilities.lift_output_func(get_directory_roles)
def get_directory_roles_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDirectoryRolesResult]:
    """
    Use this data source to access information about activated directory roles within Azure Active Directory.

    ## API Permissions

    The following API permissions are required in order to use this resource.

    When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.Read.Directory` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    current = azuread.get_directory_roles()
    pulumi.export("roles", current.object_ids)
    ```
    """
    ...
