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
]

@pulumi.output_type
class GetDirectoryRolesResult:
    """
    A collection of values returned by getDirectoryRoles.
    """
    def __init__(__self__, id=None, object_ids=None, roles=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if object_ids and not isinstance(object_ids, list):
            raise TypeError("Expected argument 'object_ids' to be a list")
        pulumi.set(__self__, "object_ids", object_ids)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)

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


class AwaitableGetDirectoryRolesResult(GetDirectoryRolesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDirectoryRolesResult(
            id=self.id,
            object_ids=self.object_ids,
            roles=self.roles)


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
        id=__ret__.id,
        object_ids=__ret__.object_ids,
        roles=__ret__.roles)
