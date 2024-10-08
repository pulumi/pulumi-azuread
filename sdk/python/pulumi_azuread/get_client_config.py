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
    'GetClientConfigResult',
    'AwaitableGetClientConfigResult',
    'get_client_config',
    'get_client_config_output',
]

@pulumi.output_type
class GetClientConfigResult:
    """
    A collection of values returned by getClientConfig.
    """
    def __init__(__self__, client_id=None, id=None, object_id=None, tenant_id=None):
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if object_id and not isinstance(object_id, str):
            raise TypeError("Expected argument 'object_id' to be a str")
        pulumi.set(__self__, "object_id", object_id)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication.
        """
        return pulumi.get(self, "client_id")

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
        """
        The object ID of the authenticated principal.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The tenant ID of the authenticated principal.
        """
        return pulumi.get(self, "tenant_id")


class AwaitableGetClientConfigResult(GetClientConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientConfigResult(
            client_id=self.client_id,
            id=self.id,
            object_id=self.object_id,
            tenant_id=self.tenant_id)


def get_client_config(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClientConfigResult:
    """
    Use this data source to access the configuration of the AzureAD provider.

    ## API Permissions

    No additional roles are required to use this data source.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    current = azuread.get_client_config()
    pulumi.export("objectId", current.object_id)
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuread:index/getClientConfig:getClientConfig', __args__, opts=opts, typ=GetClientConfigResult).value

    return AwaitableGetClientConfigResult(
        client_id=pulumi.get(__ret__, 'client_id'),
        id=pulumi.get(__ret__, 'id'),
        object_id=pulumi.get(__ret__, 'object_id'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))
def get_client_config_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClientConfigResult]:
    """
    Use this data source to access the configuration of the AzureAD provider.

    ## API Permissions

    No additional roles are required to use this data source.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    current = azuread.get_client_config()
    pulumi.export("objectId", current.object_id)
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuread:index/getClientConfig:getClientConfig', __args__, opts=opts, typ=GetClientConfigResult)
    return __ret__.apply(lambda __response__: GetClientConfigResult(
        client_id=pulumi.get(__response__, 'client_id'),
        id=pulumi.get(__response__, 'id'),
        object_id=pulumi.get(__response__, 'object_id'),
        tenant_id=pulumi.get(__response__, 'tenant_id')))
