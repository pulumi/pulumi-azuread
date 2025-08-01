# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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
from . import outputs

__all__ = [
    'GetServicePrincipalsResult',
    'AwaitableGetServicePrincipalsResult',
    'get_service_principals',
    'get_service_principals_output',
]

@pulumi.output_type
class GetServicePrincipalsResult:
    """
    A collection of values returned by getServicePrincipals.
    """
    def __init__(__self__, client_ids=None, display_names=None, id=None, ignore_missing=None, object_ids=None, return_all=None, service_principals=None):
        if client_ids and not isinstance(client_ids, list):
            raise TypeError("Expected argument 'client_ids' to be a list")
        pulumi.set(__self__, "client_ids", client_ids)
        if display_names and not isinstance(display_names, list):
            raise TypeError("Expected argument 'display_names' to be a list")
        pulumi.set(__self__, "display_names", display_names)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ignore_missing and not isinstance(ignore_missing, bool):
            raise TypeError("Expected argument 'ignore_missing' to be a bool")
        pulumi.set(__self__, "ignore_missing", ignore_missing)
        if object_ids and not isinstance(object_ids, list):
            raise TypeError("Expected argument 'object_ids' to be a list")
        pulumi.set(__self__, "object_ids", object_ids)
        if return_all and not isinstance(return_all, bool):
            raise TypeError("Expected argument 'return_all' to be a bool")
        pulumi.set(__self__, "return_all", return_all)
        if service_principals and not isinstance(service_principals, list):
            raise TypeError("Expected argument 'service_principals' to be a list")
        pulumi.set(__self__, "service_principals", service_principals)

    @_builtins.property
    @pulumi.getter(name="clientIds")
    def client_ids(self) -> Sequence[_builtins.str]:
        """
        The client ID of the application associated with this service principal.
        """
        return pulumi.get(self, "client_ids")

    @_builtins.property
    @pulumi.getter(name="displayNames")
    def display_names(self) -> Sequence[_builtins.str]:
        """
        A list of display names of the applications associated with the service principals.
        """
        return pulumi.get(self, "display_names")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter(name="ignoreMissing")
    def ignore_missing(self) -> Optional[_builtins.bool]:
        return pulumi.get(self, "ignore_missing")

    @_builtins.property
    @pulumi.getter(name="objectIds")
    def object_ids(self) -> Sequence[_builtins.str]:
        """
        The object IDs of the service principals.
        """
        return pulumi.get(self, "object_ids")

    @_builtins.property
    @pulumi.getter(name="returnAll")
    def return_all(self) -> Optional[_builtins.bool]:
        return pulumi.get(self, "return_all")

    @_builtins.property
    @pulumi.getter(name="servicePrincipals")
    def service_principals(self) -> Sequence['outputs.GetServicePrincipalsServicePrincipalResult']:
        """
        A list of service principals. Each `service_principal` object provides the attributes documented below.
        """
        return pulumi.get(self, "service_principals")


class AwaitableGetServicePrincipalsResult(GetServicePrincipalsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServicePrincipalsResult(
            client_ids=self.client_ids,
            display_names=self.display_names,
            id=self.id,
            ignore_missing=self.ignore_missing,
            object_ids=self.object_ids,
            return_all=self.return_all,
            service_principals=self.service_principals)


def get_service_principals(client_ids: Optional[Sequence[_builtins.str]] = None,
                           display_names: Optional[Sequence[_builtins.str]] = None,
                           ignore_missing: Optional[_builtins.bool] = None,
                           object_ids: Optional[Sequence[_builtins.str]] = None,
                           return_all: Optional[_builtins.bool] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServicePrincipalsResult:
    """
    Gets basic information for multiple Azure Active Directory service principals.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    *Look up by application display names*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principals(display_names=[
        "example-app",
        "another-app",
    ])
    ```

    *Look up by application IDs (client IDs)*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principals(client_ids=[
        "11111111-0000-0000-0000-000000000000",
        "22222222-0000-0000-0000-000000000000",
        "33333333-0000-0000-0000-000000000000",
    ])
    ```

    *Look up by service principal object IDs*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principals(object_ids=[
        "00000000-0000-0000-0000-000000000000",
        "00000000-0000-0000-0000-111111111111",
        "00000000-0000-0000-0000-222222222222",
    ])
    ```


    :param Sequence[_builtins.str] client_ids: A list of client IDs of the applications associated with the service principals.
    :param Sequence[_builtins.str] display_names: A list of display names of the applications associated with the service principals.
    :param _builtins.bool ignore_missing: Ignore missing service principals and return all service principals that are found. The data source will still fail if no service principals are found. Defaults to false.
    :param Sequence[_builtins.str] object_ids: The object IDs of the service principals.
    :param _builtins.bool return_all: When `true`, the data source will return all service principals. Cannot be used with `ignore_missing`. Defaults to false.
           
           > Either `return_all`, or one of `client_ids`, `display_names` or `object_ids` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.
    """
    __args__ = dict()
    __args__['clientIds'] = client_ids
    __args__['displayNames'] = display_names
    __args__['ignoreMissing'] = ignore_missing
    __args__['objectIds'] = object_ids
    __args__['returnAll'] = return_all
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuread:index/getServicePrincipals:getServicePrincipals', __args__, opts=opts, typ=GetServicePrincipalsResult).value

    return AwaitableGetServicePrincipalsResult(
        client_ids=pulumi.get(__ret__, 'client_ids'),
        display_names=pulumi.get(__ret__, 'display_names'),
        id=pulumi.get(__ret__, 'id'),
        ignore_missing=pulumi.get(__ret__, 'ignore_missing'),
        object_ids=pulumi.get(__ret__, 'object_ids'),
        return_all=pulumi.get(__ret__, 'return_all'),
        service_principals=pulumi.get(__ret__, 'service_principals'))
def get_service_principals_output(client_ids: Optional[pulumi.Input[Optional[Sequence[_builtins.str]]]] = None,
                                  display_names: Optional[pulumi.Input[Optional[Sequence[_builtins.str]]]] = None,
                                  ignore_missing: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                                  object_ids: Optional[pulumi.Input[Optional[Sequence[_builtins.str]]]] = None,
                                  return_all: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetServicePrincipalsResult]:
    """
    Gets basic information for multiple Azure Active Directory service principals.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    *Look up by application display names*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principals(display_names=[
        "example-app",
        "another-app",
    ])
    ```

    *Look up by application IDs (client IDs)*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principals(client_ids=[
        "11111111-0000-0000-0000-000000000000",
        "22222222-0000-0000-0000-000000000000",
        "33333333-0000-0000-0000-000000000000",
    ])
    ```

    *Look up by service principal object IDs*

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_service_principals(object_ids=[
        "00000000-0000-0000-0000-000000000000",
        "00000000-0000-0000-0000-111111111111",
        "00000000-0000-0000-0000-222222222222",
    ])
    ```


    :param Sequence[_builtins.str] client_ids: A list of client IDs of the applications associated with the service principals.
    :param Sequence[_builtins.str] display_names: A list of display names of the applications associated with the service principals.
    :param _builtins.bool ignore_missing: Ignore missing service principals and return all service principals that are found. The data source will still fail if no service principals are found. Defaults to false.
    :param Sequence[_builtins.str] object_ids: The object IDs of the service principals.
    :param _builtins.bool return_all: When `true`, the data source will return all service principals. Cannot be used with `ignore_missing`. Defaults to false.
           
           > Either `return_all`, or one of `client_ids`, `display_names` or `object_ids` must be specified. These _may_ be specified as an empty list, in which case no results will be returned.
    """
    __args__ = dict()
    __args__['clientIds'] = client_ids
    __args__['displayNames'] = display_names
    __args__['ignoreMissing'] = ignore_missing
    __args__['objectIds'] = object_ids
    __args__['returnAll'] = return_all
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuread:index/getServicePrincipals:getServicePrincipals', __args__, opts=opts, typ=GetServicePrincipalsResult)
    return __ret__.apply(lambda __response__: GetServicePrincipalsResult(
        client_ids=pulumi.get(__response__, 'client_ids'),
        display_names=pulumi.get(__response__, 'display_names'),
        id=pulumi.get(__response__, 'id'),
        ignore_missing=pulumi.get(__response__, 'ignore_missing'),
        object_ids=pulumi.get(__response__, 'object_ids'),
        return_all=pulumi.get(__response__, 'return_all'),
        service_principals=pulumi.get(__response__, 'service_principals')))
