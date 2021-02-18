# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetGroupResult',
    'AwaitableGetGroupResult',
    'get_group',
]

@pulumi.output_type
class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, description=None, display_name=None, id=None, mail_enabled=None, members=None, name=None, object_id=None, owners=None, security_enabled=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mail_enabled and not isinstance(mail_enabled, bool):
            raise TypeError("Expected argument 'mail_enabled' to be a bool")
        pulumi.set(__self__, "mail_enabled", mail_enabled)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        if name is not None:
            warnings.warn("""This property has been renamed to `display_name` and will be removed in v2.0 of this provider.""", DeprecationWarning)
            pulumi.log.warn("name is deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.")

        pulumi.set(__self__, "name", name)
        if object_id and not isinstance(object_id, str):
            raise TypeError("Expected argument 'object_id' to be a str")
        pulumi.set(__self__, "object_id", object_id)
        if owners and not isinstance(owners, list):
            raise TypeError("Expected argument 'owners' to be a list")
        pulumi.set(__self__, "owners", owners)
        if security_enabled and not isinstance(security_enabled, bool):
            raise TypeError("Expected argument 'security_enabled' to be a bool")
        pulumi.set(__self__, "security_enabled", security_enabled)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The optional description of the Group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name for the Group.
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
    @pulumi.getter(name="mailEnabled")
    def mail_enabled(self) -> bool:
        """
        Whether the group is mail-enabled.
        """
        return pulumi.get(self, "mail_enabled")

    @property
    @pulumi.getter
    def members(self) -> Sequence[str]:
        """
        The Object IDs of the Group members.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> str:
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter
    def owners(self) -> Sequence[str]:
        """
        The Object IDs of the Group owners.
        """
        return pulumi.get(self, "owners")

    @property
    @pulumi.getter(name="securityEnabled")
    def security_enabled(self) -> bool:
        """
        Whether the group is a security group.
        """
        return pulumi.get(self, "security_enabled")


class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            description=self.description,
            display_name=self.display_name,
            id=self.id,
            mail_enabled=self.mail_enabled,
            members=self.members,
            name=self.name,
            object_id=self.object_id,
            owners=self.owners,
            security_enabled=self.security_enabled)


def get_group(display_name: Optional[str] = None,
              mail_enabled: Optional[bool] = None,
              name: Optional[str] = None,
              object_id: Optional[str] = None,
              security_enabled: Optional[bool] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupResult:
    """
    Gets information about an Azure Active Directory group.

    > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

    ## Example Usage
    ### By Group Display Name)

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_group(display_name="MyGroupName",
        security_enabled=True)
    ```


    :param str display_name: The display name for the Group.
    :param bool mail_enabled: Whether the group is mail-enabled.
    :param str object_id: Specifies the Object ID of the Group.
    :param bool security_enabled: Whether the group is a security group.
    """
    __args__ = dict()
    __args__['displayName'] = display_name
    __args__['mailEnabled'] = mail_enabled
    __args__['name'] = name
    __args__['objectId'] = object_id
    __args__['securityEnabled'] = security_enabled
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuread:index/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult).value

    return AwaitableGetGroupResult(
        description=__ret__.description,
        display_name=__ret__.display_name,
        id=__ret__.id,
        mail_enabled=__ret__.mail_enabled,
        members=__ret__.members,
        name=__ret__.name,
        object_id=__ret__.object_id,
        owners=__ret__.owners,
        security_enabled=__ret__.security_enabled)
