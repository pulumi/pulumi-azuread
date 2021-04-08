# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
<<<<<<< HEAD
from . import _utilities, _tables
=======
from . import _utilities
>>>>>>> 1e7e750 (Upgrade to Pulumi v3.0.0-beta.2)

__all__ = [
    'GetGroupsResult',
    'AwaitableGetGroupsResult',
    'get_groups',
]

@pulumi.output_type
class GetGroupsResult:
    """
    A collection of values returned by getGroups.
    """
    def __init__(__self__, display_names=None, id=None, names=None, object_ids=None):
        if display_names and not isinstance(display_names, list):
            raise TypeError("Expected argument 'display_names' to be a list")
        pulumi.set(__self__, "display_names", display_names)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        if names is not None:
            warnings.warn("""This property has been renamed to `display_names` and will be removed in v2.0 of this provider.""", DeprecationWarning)
            pulumi.log.warn("""names is deprecated: This property has been renamed to `display_names` and will be removed in v2.0 of this provider.""")

        pulumi.set(__self__, "names", names)
        if object_ids and not isinstance(object_ids, list):
            raise TypeError("Expected argument 'object_ids' to be a list")
        pulumi.set(__self__, "object_ids", object_ids)

    @property
    @pulumi.getter(name="displayNames")
    def display_names(self) -> Sequence[str]:
        return pulumi.get(self, "display_names")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        The Display Names of the Azure AD Groups.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="objectIds")
    def object_ids(self) -> Sequence[str]:
        """
        The Object IDs of the Azure AD Groups.
        """
        return pulumi.get(self, "object_ids")


class AwaitableGetGroupsResult(GetGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupsResult(
            display_names=self.display_names,
            id=self.id,
            names=self.names,
            object_ids=self.object_ids)


def get_groups(display_names: Optional[Sequence[str]] = None,
               names: Optional[Sequence[str]] = None,
               object_ids: Optional[Sequence[str]] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupsResult:
    """
    Gets Object IDs or Display Names for multiple Azure Active Directory groups.

    > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    groups = azuread.get_groups(names=[
        "group-a",
        "group-b",
    ])
    ```


    :param Sequence[str] names: The Display Names of the Azure AD Groups.
    :param Sequence[str] object_ids: The Object IDs of the Azure AD Groups.
    """
    __args__ = dict()
    __args__['displayNames'] = display_names
    __args__['names'] = names
    __args__['objectIds'] = object_ids
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azuread:index/getGroups:getGroups', __args__, opts=opts, typ=GetGroupsResult).value

    return AwaitableGetGroupsResult(
        display_names=__ret__.display_names,
        id=__ret__.id,
        names=__ret__.names,
        object_ids=__ret__.object_ids)
