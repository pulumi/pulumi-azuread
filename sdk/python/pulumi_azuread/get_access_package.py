# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAccessPackageResult',
    'AwaitableGetAccessPackageResult',
    'get_access_package',
    'get_access_package_output',
]

@pulumi.output_type
class GetAccessPackageResult:
    """
    A collection of values returned by getAccessPackage.
    """
    def __init__(__self__, catalog_id=None, description=None, display_name=None, hidden=None, id=None, object_id=None):
        if catalog_id and not isinstance(catalog_id, str):
            raise TypeError("Expected argument 'catalog_id' to be a str")
        pulumi.set(__self__, "catalog_id", catalog_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if hidden and not isinstance(hidden, bool):
            raise TypeError("Expected argument 'hidden' to be a bool")
        pulumi.set(__self__, "hidden", hidden)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if object_id and not isinstance(object_id, str):
            raise TypeError("Expected argument 'object_id' to be a str")
        pulumi.set(__self__, "object_id", object_id)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[str]:
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the access package.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def hidden(self) -> bool:
        """
        Whether the access package is hidden from the requestor.
        """
        return pulumi.get(self, "hidden")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> str:
        return pulumi.get(self, "object_id")


class AwaitableGetAccessPackageResult(GetAccessPackageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessPackageResult(
            catalog_id=self.catalog_id,
            description=self.description,
            display_name=self.display_name,
            hidden=self.hidden,
            id=self.id,
            object_id=self.object_id)


def get_access_package(catalog_id: Optional[str] = None,
                       display_name: Optional[str] = None,
                       object_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessPackageResult:
    """
    Use this data source to retrieve information for an existing access package within Identity Governance in Azure Active Directory.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All`, or `EntitlementManagement.ReadWrite.All`.

    When authenticated with a user principal, this data source requires one of the following directory roles: `Catalog owner`, `Catalog reader`, `Access package manager`, `Global Reader`, or `Global Administrator`.

    ## Example Usage

    *Look up by ID*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_access_package(object_id="00000000-0000-0000-0000-000000000000")
    ```

    *Look up by DisplayName*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_access_package(catalog_id="00000000-0000-0000-0000-000000000000",
        display_name="My access package Catalog")
    ```


    :param str catalog_id: The ID of the Catalog this access package is in.
    :param str display_name: The display name of the access package.
    :param str object_id: The ID of this access package.
    """
    __args__ = dict()
    __args__['catalogId'] = catalog_id
    __args__['displayName'] = display_name
    __args__['objectId'] = object_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuread:index/getAccessPackage:getAccessPackage', __args__, opts=opts, typ=GetAccessPackageResult).value

    return AwaitableGetAccessPackageResult(
        catalog_id=__ret__.catalog_id,
        description=__ret__.description,
        display_name=__ret__.display_name,
        hidden=__ret__.hidden,
        id=__ret__.id,
        object_id=__ret__.object_id)


@_utilities.lift_output_func(get_access_package)
def get_access_package_output(catalog_id: Optional[pulumi.Input[Optional[str]]] = None,
                              display_name: Optional[pulumi.Input[Optional[str]]] = None,
                              object_id: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccessPackageResult]:
    """
    Use this data source to retrieve information for an existing access package within Identity Governance in Azure Active Directory.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `EntitlementManagement.Read.All`, or `EntitlementManagement.ReadWrite.All`.

    When authenticated with a user principal, this data source requires one of the following directory roles: `Catalog owner`, `Catalog reader`, `Access package manager`, `Global Reader`, or `Global Administrator`.

    ## Example Usage

    *Look up by ID*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_access_package(object_id="00000000-0000-0000-0000-000000000000")
    ```

    *Look up by DisplayName*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_access_package(catalog_id="00000000-0000-0000-0000-000000000000",
        display_name="My access package Catalog")
    ```


    :param str catalog_id: The ID of the Catalog this access package is in.
    :param str display_name: The display name of the access package.
    :param str object_id: The ID of this access package.
    """
    ...
