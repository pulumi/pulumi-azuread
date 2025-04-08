# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['ApplicationRedirectUrisArgs', 'ApplicationRedirectUris']

@pulumi.input_type
class ApplicationRedirectUrisArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[builtins.str],
                 redirect_uris: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 type: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a ApplicationRedirectUris resource.
        :param pulumi.Input[builtins.str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] redirect_uris: A set of redirect URIs to assign to the application.
        :param pulumi.Input[builtins.str] type: The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "application_id", application_id)
        pulumi.set(__self__, "redirect_uris", redirect_uris)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[builtins.str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="redirectUris")
    def redirect_uris(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        A set of redirect URIs to assign to the application.
        """
        return pulumi.get(self, "redirect_uris")

    @redirect_uris.setter
    def redirect_uris(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "redirect_uris", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[builtins.str]:
        """
        The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _ApplicationRedirectUrisState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[builtins.str]] = None,
                 redirect_uris: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ApplicationRedirectUris resources.
        :param pulumi.Input[builtins.str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] redirect_uris: A set of redirect URIs to assign to the application.
        :param pulumi.Input[builtins.str] type: The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if redirect_uris is not None:
            pulumi.set(__self__, "redirect_uris", redirect_uris)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="redirectUris")
    def redirect_uris(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A set of redirect URIs to assign to the application.
        """
        return pulumi.get(self, "redirect_uris")

    @redirect_uris.setter
    def redirect_uris(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "redirect_uris", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


class ApplicationRedirectUris(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[builtins.str]] = None,
                 redirect_uris: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ApplicationRegistration("example", display_name="example")
        example_public = azuread.ApplicationRedirectUris("example_public",
            application_id=example.id,
            type="PublicClient",
            redirect_uris=[
                "myapp://auth",
                "sample.mobile.app.bundie.id://auth",
                "https://login.microsoftonline.com/common/oauth2/nativeclient",
                "https://login.live.com/oauth20_desktop.srf",
                "ms-appx-web://Microsoft.AAD.BrokerPlugin/00000000-1111-1111-1111-222222222222",
                "urn:ietf:wg:oauth:2.0:foo",
            ])
        example_spa = azuread.ApplicationRedirectUris("example_spa",
            application_id=example.id,
            type="SPA",
            redirect_uris=[
                "https://mobile.example.com/",
                "https://beta.example.com/",
            ])
        example_web = azuread.ApplicationRedirectUris("example_web",
            application_id=example.id,
            type="Web",
            redirect_uris=[
                "https://app.example.com/",
                "https://classic.example.com/",
                "urn:ietf:wg:oauth:2.0:oob",
            ])
        ```

        ## Import

        Application API Access can be imported using the object ID of the application and the URI type, in the following format.

        ```sh
        $ pulumi import azuread:index/applicationRedirectUris:ApplicationRedirectUris example /applications/00000000-0000-0000-0000-000000000000/redirectUris/Web
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] redirect_uris: A set of redirect URIs to assign to the application.
        :param pulumi.Input[builtins.str] type: The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationRedirectUrisArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ApplicationRegistration("example", display_name="example")
        example_public = azuread.ApplicationRedirectUris("example_public",
            application_id=example.id,
            type="PublicClient",
            redirect_uris=[
                "myapp://auth",
                "sample.mobile.app.bundie.id://auth",
                "https://login.microsoftonline.com/common/oauth2/nativeclient",
                "https://login.live.com/oauth20_desktop.srf",
                "ms-appx-web://Microsoft.AAD.BrokerPlugin/00000000-1111-1111-1111-222222222222",
                "urn:ietf:wg:oauth:2.0:foo",
            ])
        example_spa = azuread.ApplicationRedirectUris("example_spa",
            application_id=example.id,
            type="SPA",
            redirect_uris=[
                "https://mobile.example.com/",
                "https://beta.example.com/",
            ])
        example_web = azuread.ApplicationRedirectUris("example_web",
            application_id=example.id,
            type="Web",
            redirect_uris=[
                "https://app.example.com/",
                "https://classic.example.com/",
                "urn:ietf:wg:oauth:2.0:oob",
            ])
        ```

        ## Import

        Application API Access can be imported using the object ID of the application and the URI type, in the following format.

        ```sh
        $ pulumi import azuread:index/applicationRedirectUris:ApplicationRedirectUris example /applications/00000000-0000-0000-0000-000000000000/redirectUris/Web
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationRedirectUrisArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationRedirectUrisArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[builtins.str]] = None,
                 redirect_uris: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationRedirectUrisArgs.__new__(ApplicationRedirectUrisArgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            if redirect_uris is None and not opts.urn:
                raise TypeError("Missing required property 'redirect_uris'")
            __props__.__dict__["redirect_uris"] = redirect_uris
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(ApplicationRedirectUris, __self__).__init__(
            'azuread:index/applicationRedirectUris:ApplicationRedirectUris',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[builtins.str]] = None,
            redirect_uris: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None) -> 'ApplicationRedirectUris':
        """
        Get an existing ApplicationRedirectUris resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] redirect_uris: A set of redirect URIs to assign to the application.
        :param pulumi.Input[builtins.str] type: The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationRedirectUrisState.__new__(_ApplicationRedirectUrisState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["redirect_uris"] = redirect_uris
        __props__.__dict__["type"] = type
        return ApplicationRedirectUris(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[builtins.str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="redirectUris")
    def redirect_uris(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        A set of redirect URIs to assign to the application.
        """
        return pulumi.get(self, "redirect_uris")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "type")

