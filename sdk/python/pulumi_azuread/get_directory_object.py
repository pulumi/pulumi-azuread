# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'GetDirectoryObjectResult',
    'AwaitableGetDirectoryObjectResult',
    'get_directory_object',
    'get_directory_object_output',
]

@pulumi.output_type
class GetDirectoryObjectResult:
    """
    A collection of values returned by getDirectoryObject.
    """
    def __init__(__self__, id=None, object_id=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if object_id and not isinstance(object_id, str):
            raise TypeError("Expected argument 'object_id' to be a str")
        pulumi.set(__self__, "object_id", object_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

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

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableGetDirectoryObjectResult(GetDirectoryObjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDirectoryObjectResult(
            id=self.id,
            object_id=self.object_id,
            type=self.type)


def get_directory_object(object_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDirectoryObjectResult:
    """
    Retrieves the OData type for a generic directory object having the provided object ID.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires either `User.Read.All`, `Group.Read.All` or `Directory.Read.All`, depending on the type of object being queried.

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    *Look up and output type of object by ID*
    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_directory_object(object_id="00000000-0000-0000-0000-000000000000")
    pulumi.export("objectType", example.type)
    ```

    ## Attributes Reference

    The following attributes are exported:

    * `object_id` - The object ID of the directory object.
    * `type` - The shortened OData type of the directory object. Possible values include: `Group`, `User` or `ServicePrincipal`.


    :param str object_id: Specifies the Object ID of the directory object to look up.
    """
    __args__ = dict()
    __args__['objectId'] = object_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuread:index/getDirectoryObject:getDirectoryObject', __args__, opts=opts, typ=GetDirectoryObjectResult).value

    return AwaitableGetDirectoryObjectResult(
        id=pulumi.get(__ret__, 'id'),
        object_id=pulumi.get(__ret__, 'object_id'),
        type=pulumi.get(__ret__, 'type'))
def get_directory_object_output(object_id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDirectoryObjectResult]:
    """
    Retrieves the OData type for a generic directory object having the provided object ID.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires either `User.Read.All`, `Group.Read.All` or `Directory.Read.All`, depending on the type of object being queried.

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    *Look up and output type of object by ID*
    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_directory_object(object_id="00000000-0000-0000-0000-000000000000")
    pulumi.export("objectType", example.type)
    ```

    ## Attributes Reference

    The following attributes are exported:

    * `object_id` - The object ID of the directory object.
    * `type` - The shortened OData type of the directory object. Possible values include: `Group`, `User` or `ServicePrincipal`.


    :param str object_id: Specifies the Object ID of the directory object to look up.
    """
    __args__ = dict()
    __args__['objectId'] = object_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuread:index/getDirectoryObject:getDirectoryObject', __args__, opts=opts, typ=GetDirectoryObjectResult)
    return __ret__.apply(lambda __response__: GetDirectoryObjectResult(
        id=pulumi.get(__response__, 'id'),
        object_id=pulumi.get(__response__, 'object_id'),
        type=pulumi.get(__response__, 'type')))
