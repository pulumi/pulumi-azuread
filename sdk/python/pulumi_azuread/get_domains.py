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
    'GetDomainsResult',
    'AwaitableGetDomainsResult',
    'get_domains',
    'get_domains_output',
]

@pulumi.output_type
class GetDomainsResult:
    """
    A collection of values returned by getDomains.
    """
    def __init__(__self__, admin_managed=None, domains=None, id=None, include_unverified=None, only_default=None, only_initial=None, only_root=None, supports_services=None):
        if admin_managed and not isinstance(admin_managed, bool):
            raise TypeError("Expected argument 'admin_managed' to be a bool")
        pulumi.set(__self__, "admin_managed", admin_managed)
        if domains and not isinstance(domains, list):
            raise TypeError("Expected argument 'domains' to be a list")
        pulumi.set(__self__, "domains", domains)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_unverified and not isinstance(include_unverified, bool):
            raise TypeError("Expected argument 'include_unverified' to be a bool")
        pulumi.set(__self__, "include_unverified", include_unverified)
        if only_default and not isinstance(only_default, bool):
            raise TypeError("Expected argument 'only_default' to be a bool")
        pulumi.set(__self__, "only_default", only_default)
        if only_initial and not isinstance(only_initial, bool):
            raise TypeError("Expected argument 'only_initial' to be a bool")
        pulumi.set(__self__, "only_initial", only_initial)
        if only_root and not isinstance(only_root, bool):
            raise TypeError("Expected argument 'only_root' to be a bool")
        pulumi.set(__self__, "only_root", only_root)
        if supports_services and not isinstance(supports_services, list):
            raise TypeError("Expected argument 'supports_services' to be a list")
        pulumi.set(__self__, "supports_services", supports_services)

    @_builtins.property
    @pulumi.getter(name="adminManaged")
    def admin_managed(self) -> Optional[_builtins.bool]:
        """
        Whether the DNS for the domain is managed by Microsoft 365.
        """
        return pulumi.get(self, "admin_managed")

    @_builtins.property
    @pulumi.getter
    def domains(self) -> Sequence['outputs.GetDomainsDomainResult']:
        """
        A list of tenant domains. Each `domain` object provides the attributes documented below.
        """
        return pulumi.get(self, "domains")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter(name="includeUnverified")
    def include_unverified(self) -> Optional[_builtins.bool]:
        return pulumi.get(self, "include_unverified")

    @_builtins.property
    @pulumi.getter(name="onlyDefault")
    def only_default(self) -> Optional[_builtins.bool]:
        return pulumi.get(self, "only_default")

    @_builtins.property
    @pulumi.getter(name="onlyInitial")
    def only_initial(self) -> Optional[_builtins.bool]:
        return pulumi.get(self, "only_initial")

    @_builtins.property
    @pulumi.getter(name="onlyRoot")
    def only_root(self) -> Optional[_builtins.bool]:
        return pulumi.get(self, "only_root")

    @_builtins.property
    @pulumi.getter(name="supportsServices")
    def supports_services(self) -> Optional[Sequence[_builtins.str]]:
        return pulumi.get(self, "supports_services")


class AwaitableGetDomainsResult(GetDomainsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainsResult(
            admin_managed=self.admin_managed,
            domains=self.domains,
            id=self.id,
            include_unverified=self.include_unverified,
            only_default=self.only_default,
            only_initial=self.only_initial,
            only_root=self.only_root,
            supports_services=self.supports_services)


def get_domains(admin_managed: Optional[_builtins.bool] = None,
                include_unverified: Optional[_builtins.bool] = None,
                only_default: Optional[_builtins.bool] = None,
                only_initial: Optional[_builtins.bool] = None,
                only_root: Optional[_builtins.bool] = None,
                supports_services: Optional[Sequence[_builtins.str]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainsResult:
    """
    Use this data source to access information about existing Domains within Azure Active Directory.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `Domain.Read.All` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    aad_domains = azuread.get_domains()
    pulumi.export("domainNames", [__item.domain_name for __item in aad_domains.domains])
    ```


    :param _builtins.bool admin_managed: Set to `true` to only return domains whose DNS is managed by Microsoft 365. Defaults to `false`.
    :param _builtins.bool include_unverified: Set to `true` if unverified Azure AD domains should be included. Defaults to `false`.
    :param _builtins.bool only_default: Set to `true` to only return the default domain.
    :param _builtins.bool only_initial: Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Defaults to `false`.
    :param _builtins.bool only_root: Set to `true` to only return verified root domains. Excludes subdomains and unverified domains.
    :param Sequence[_builtins.str] supports_services: A list of supported services that must be supported by a domain. Possible values include `Email`, `Sharepoint`, `EmailInternalRelayOnly`, `OfficeCommunicationsOnline`, `SharePointDefaultDomain`, `FullRedelegation`, `SharePointPublic`, `OrgIdAuthentication`, `Yammer` and `Intune`.
           
           > **Note on filters** If `include_unverified` is set to `true`, you cannot specify `only_default` or `only_initial`. Additionally, you cannot combine `only_default` with `only_initial`.
    """
    __args__ = dict()
    __args__['adminManaged'] = admin_managed
    __args__['includeUnverified'] = include_unverified
    __args__['onlyDefault'] = only_default
    __args__['onlyInitial'] = only_initial
    __args__['onlyRoot'] = only_root
    __args__['supportsServices'] = supports_services
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuread:index/getDomains:getDomains', __args__, opts=opts, typ=GetDomainsResult).value

    return AwaitableGetDomainsResult(
        admin_managed=pulumi.get(__ret__, 'admin_managed'),
        domains=pulumi.get(__ret__, 'domains'),
        id=pulumi.get(__ret__, 'id'),
        include_unverified=pulumi.get(__ret__, 'include_unverified'),
        only_default=pulumi.get(__ret__, 'only_default'),
        only_initial=pulumi.get(__ret__, 'only_initial'),
        only_root=pulumi.get(__ret__, 'only_root'),
        supports_services=pulumi.get(__ret__, 'supports_services'))
def get_domains_output(admin_managed: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                       include_unverified: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                       only_default: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                       only_initial: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                       only_root: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                       supports_services: Optional[pulumi.Input[Optional[Sequence[_builtins.str]]]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDomainsResult]:
    """
    Use this data source to access information about existing Domains within Azure Active Directory.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `Domain.Read.All` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    aad_domains = azuread.get_domains()
    pulumi.export("domainNames", [__item.domain_name for __item in aad_domains.domains])
    ```


    :param _builtins.bool admin_managed: Set to `true` to only return domains whose DNS is managed by Microsoft 365. Defaults to `false`.
    :param _builtins.bool include_unverified: Set to `true` if unverified Azure AD domains should be included. Defaults to `false`.
    :param _builtins.bool only_default: Set to `true` to only return the default domain.
    :param _builtins.bool only_initial: Set to `true` to only return the initial domain, which is your primary Azure Active Directory tenant domain. Defaults to `false`.
    :param _builtins.bool only_root: Set to `true` to only return verified root domains. Excludes subdomains and unverified domains.
    :param Sequence[_builtins.str] supports_services: A list of supported services that must be supported by a domain. Possible values include `Email`, `Sharepoint`, `EmailInternalRelayOnly`, `OfficeCommunicationsOnline`, `SharePointDefaultDomain`, `FullRedelegation`, `SharePointPublic`, `OrgIdAuthentication`, `Yammer` and `Intune`.
           
           > **Note on filters** If `include_unverified` is set to `true`, you cannot specify `only_default` or `only_initial`. Additionally, you cannot combine `only_default` with `only_initial`.
    """
    __args__ = dict()
    __args__['adminManaged'] = admin_managed
    __args__['includeUnverified'] = include_unverified
    __args__['onlyDefault'] = only_default
    __args__['onlyInitial'] = only_initial
    __args__['onlyRoot'] = only_root
    __args__['supportsServices'] = supports_services
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuread:index/getDomains:getDomains', __args__, opts=opts, typ=GetDomainsResult)
    return __ret__.apply(lambda __response__: GetDomainsResult(
        admin_managed=pulumi.get(__response__, 'admin_managed'),
        domains=pulumi.get(__response__, 'domains'),
        id=pulumi.get(__response__, 'id'),
        include_unverified=pulumi.get(__response__, 'include_unverified'),
        only_default=pulumi.get(__response__, 'only_default'),
        only_initial=pulumi.get(__response__, 'only_initial'),
        only_root=pulumi.get(__response__, 'only_root'),
        supports_services=pulumi.get(__response__, 'supports_services')))
