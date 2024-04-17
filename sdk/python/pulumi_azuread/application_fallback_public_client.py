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
    from typing import NotRequired, TypedDict
else:
    from typing_extensions import NotRequired, TypedDict
from . import _utilities

__all__ = ['ApplicationFallbackPublicClientArgs', 'ApplicationFallbackPublicClient']

@pulumi.input_type
class ApplicationFallbackPublicClientArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ApplicationFallbackPublicClient resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enabled: Whether to enable the application as a fallback public client.
               
               > Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
        """
        pulumi.set(__self__, "application_id", application_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the application as a fallback public client.

        > Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class _ApplicationFallbackPublicClientState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering ApplicationFallbackPublicClient resources.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enabled: Whether to enable the application as a fallback public client.
               
               > Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the application as a fallback public client.

        > Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


class ApplicationFallbackPublicClient(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ApplicationRegistration("example", display_name="example")
        example_application_fallback_public_client = azuread.ApplicationFallbackPublicClient("example",
            application_id=example.id,
            enabled=True)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        The Application Fallback Public Client setting can be imported using the object ID of the application, in the following format.

        ```sh
        $ pulumi import azuread:index/applicationFallbackPublicClient:ApplicationFallbackPublicClient example /applications/00000000-0000-0000-0000-000000000000/fallbackPublicClient
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enabled: Whether to enable the application as a fallback public client.
               
               > Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationFallbackPublicClientArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ApplicationRegistration("example", display_name="example")
        example_application_fallback_public_client = azuread.ApplicationFallbackPublicClient("example",
            application_id=example.id,
            enabled=True)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        The Application Fallback Public Client setting can be imported using the object ID of the application, in the following format.

        ```sh
        $ pulumi import azuread:index/applicationFallbackPublicClient:ApplicationFallbackPublicClient example /applications/00000000-0000-0000-0000-000000000000/fallbackPublicClient
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationFallbackPublicClientArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationFallbackPublicClientArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationFallbackPublicClientArgs.__new__(ApplicationFallbackPublicClientArgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            __props__.__dict__["enabled"] = enabled
        super(ApplicationFallbackPublicClient, __self__).__init__(
            'azuread:index/applicationFallbackPublicClient:ApplicationFallbackPublicClient',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None) -> 'ApplicationFallbackPublicClient':
        """
        Get an existing ApplicationFallbackPublicClient resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enabled: Whether to enable the application as a fallback public client.
               
               > Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationFallbackPublicClientState.__new__(_ApplicationFallbackPublicClientState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["enabled"] = enabled
        return ApplicationFallbackPublicClient(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable the application as a fallback public client.

        > Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
        """
        return pulumi.get(self, "enabled")

