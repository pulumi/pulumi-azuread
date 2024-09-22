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
    'GetApplicationTemplateResult',
    'AwaitableGetApplicationTemplateResult',
    'get_application_template',
    'get_application_template_output',
]

@pulumi.output_type
class GetApplicationTemplateResult:
    """
    A collection of values returned by getApplicationTemplate.
    """
    def __init__(__self__, categories=None, display_name=None, homepage_url=None, id=None, logo_url=None, publisher=None, supported_provisioning_types=None, supported_single_sign_on_modes=None, template_id=None):
        if categories and not isinstance(categories, list):
            raise TypeError("Expected argument 'categories' to be a list")
        pulumi.set(__self__, "categories", categories)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if homepage_url and not isinstance(homepage_url, str):
            raise TypeError("Expected argument 'homepage_url' to be a str")
        pulumi.set(__self__, "homepage_url", homepage_url)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if logo_url and not isinstance(logo_url, str):
            raise TypeError("Expected argument 'logo_url' to be a str")
        pulumi.set(__self__, "logo_url", logo_url)
        if publisher and not isinstance(publisher, str):
            raise TypeError("Expected argument 'publisher' to be a str")
        pulumi.set(__self__, "publisher", publisher)
        if supported_provisioning_types and not isinstance(supported_provisioning_types, list):
            raise TypeError("Expected argument 'supported_provisioning_types' to be a list")
        pulumi.set(__self__, "supported_provisioning_types", supported_provisioning_types)
        if supported_single_sign_on_modes and not isinstance(supported_single_sign_on_modes, list):
            raise TypeError("Expected argument 'supported_single_sign_on_modes' to be a list")
        pulumi.set(__self__, "supported_single_sign_on_modes", supported_single_sign_on_modes)
        if template_id and not isinstance(template_id, str):
            raise TypeError("Expected argument 'template_id' to be a str")
        pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter
    def categories(self) -> Sequence[str]:
        """
        List of categories for this templated application.
        """
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name for the templated application.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="homepageUrl")
    def homepage_url(self) -> str:
        """
        Home page URL of the templated application.
        """
        return pulumi.get(self, "homepage_url")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="logoUrl")
    def logo_url(self) -> str:
        """
        URL to retrieve the logo for this templated application.
        """
        return pulumi.get(self, "logo_url")

    @property
    @pulumi.getter
    def publisher(self) -> str:
        """
        Name of the publisher for this templated application.
        """
        return pulumi.get(self, "publisher")

    @property
    @pulumi.getter(name="supportedProvisioningTypes")
    def supported_provisioning_types(self) -> Sequence[str]:
        """
        List of provisioning modes supported by this templated application.
        """
        return pulumi.get(self, "supported_provisioning_types")

    @property
    @pulumi.getter(name="supportedSingleSignOnModes")
    def supported_single_sign_on_modes(self) -> Sequence[str]:
        """
        List of single sign on modes supported by this templated application.
        """
        return pulumi.get(self, "supported_single_sign_on_modes")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> str:
        """
        The ID of the templated application.
        """
        return pulumi.get(self, "template_id")


class AwaitableGetApplicationTemplateResult(GetApplicationTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationTemplateResult(
            categories=self.categories,
            display_name=self.display_name,
            homepage_url=self.homepage_url,
            id=self.id,
            logo_url=self.logo_url,
            publisher=self.publisher,
            supported_provisioning_types=self.supported_provisioning_types,
            supported_single_sign_on_modes=self.supported_single_sign_on_modes,
            template_id=self.template_id)


def get_application_template(display_name: Optional[str] = None,
                             template_id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationTemplateResult:
    """
    Use this data source to access information about an Application Template from the [Azure AD App Gallery](https://azuremarketplace.microsoft.com/en-US/marketplace/apps/category/azure-active-directory-apps).

    ## API Permissions

    This data source does not require any additional roles.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_application_template(display_name="Marketo")
    pulumi.export("applicationTemplateId", example.template_id)
    ```


    :param str display_name: Specifies the display name of the templated application.
    :param str template_id: Specifies the ID of the templated application.
           
           > One of `template_id` or `display_name` must be specified.
    """
    __args__ = dict()
    __args__['displayName'] = display_name
    __args__['templateId'] = template_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuread:index/getApplicationTemplate:getApplicationTemplate', __args__, opts=opts, typ=GetApplicationTemplateResult).value

    return AwaitableGetApplicationTemplateResult(
        categories=pulumi.get(__ret__, 'categories'),
        display_name=pulumi.get(__ret__, 'display_name'),
        homepage_url=pulumi.get(__ret__, 'homepage_url'),
        id=pulumi.get(__ret__, 'id'),
        logo_url=pulumi.get(__ret__, 'logo_url'),
        publisher=pulumi.get(__ret__, 'publisher'),
        supported_provisioning_types=pulumi.get(__ret__, 'supported_provisioning_types'),
        supported_single_sign_on_modes=pulumi.get(__ret__, 'supported_single_sign_on_modes'),
        template_id=pulumi.get(__ret__, 'template_id'))
def get_application_template_output(display_name: Optional[pulumi.Input[Optional[str]]] = None,
                                    template_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationTemplateResult]:
    """
    Use this data source to access information about an Application Template from the [Azure AD App Gallery](https://azuremarketplace.microsoft.com/en-US/marketplace/apps/category/azure-active-directory-apps).

    ## API Permissions

    This data source does not require any additional roles.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_application_template(display_name="Marketo")
    pulumi.export("applicationTemplateId", example.template_id)
    ```


    :param str display_name: Specifies the display name of the templated application.
    :param str template_id: Specifies the ID of the templated application.
           
           > One of `template_id` or `display_name` must be specified.
    """
    __args__ = dict()
    __args__['displayName'] = display_name
    __args__['templateId'] = template_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azuread:index/getApplicationTemplate:getApplicationTemplate', __args__, opts=opts, typ=GetApplicationTemplateResult)
    return __ret__.apply(lambda __response__: GetApplicationTemplateResult(
        categories=pulumi.get(__response__, 'categories'),
        display_name=pulumi.get(__response__, 'display_name'),
        homepage_url=pulumi.get(__response__, 'homepage_url'),
        id=pulumi.get(__response__, 'id'),
        logo_url=pulumi.get(__response__, 'logo_url'),
        publisher=pulumi.get(__response__, 'publisher'),
        supported_provisioning_types=pulumi.get(__response__, 'supported_provisioning_types'),
        supported_single_sign_on_modes=pulumi.get(__response__, 'supported_single_sign_on_modes'),
        template_id=pulumi.get(__response__, 'template_id')))
