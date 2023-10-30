# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ApplicationOptionalClaimsInitArgs', 'ApplicationOptionalClaims']

@pulumi.input_type
class ApplicationOptionalClaimsInitArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 access_tokens: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsAccessTokenArgs']]]] = None,
                 id_tokens: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsIdTokenArgs']]]] = None,
                 saml2_tokens: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsSaml2TokenArgs']]]] = None):
        """
        The set of arguments for constructing a ApplicationOptionalClaims resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsAccessTokenArgs']]] access_tokens: One or more `access_token` blocks as documented below.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsIdTokenArgs']]] id_tokens: One or more `id_token` blocks as documented below.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsSaml2TokenArgs']]] saml2_tokens: One or more `saml2_token` blocks as documented below.
               
               > At least one of `access_token`, `id_token` or `saml2_token` must be specified
        """
        pulumi.set(__self__, "application_id", application_id)
        if access_tokens is not None:
            pulumi.set(__self__, "access_tokens", access_tokens)
        if id_tokens is not None:
            pulumi.set(__self__, "id_tokens", id_tokens)
        if saml2_tokens is not None:
            pulumi.set(__self__, "saml2_tokens", saml2_tokens)

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
    @pulumi.getter(name="accessTokens")
    def access_tokens(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsAccessTokenArgs']]]]:
        """
        One or more `access_token` blocks as documented below.
        """
        return pulumi.get(self, "access_tokens")

    @access_tokens.setter
    def access_tokens(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsAccessTokenArgs']]]]):
        pulumi.set(self, "access_tokens", value)

    @property
    @pulumi.getter(name="idTokens")
    def id_tokens(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsIdTokenArgs']]]]:
        """
        One or more `id_token` blocks as documented below.
        """
        return pulumi.get(self, "id_tokens")

    @id_tokens.setter
    def id_tokens(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsIdTokenArgs']]]]):
        pulumi.set(self, "id_tokens", value)

    @property
    @pulumi.getter(name="saml2Tokens")
    def saml2_tokens(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsSaml2TokenArgs']]]]:
        """
        One or more `saml2_token` blocks as documented below.

        > At least one of `access_token`, `id_token` or `saml2_token` must be specified
        """
        return pulumi.get(self, "saml2_tokens")

    @saml2_tokens.setter
    def saml2_tokens(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsSaml2TokenArgs']]]]):
        pulumi.set(self, "saml2_tokens", value)


@pulumi.input_type
class _ApplicationOptionalClaimsState:
    def __init__(__self__, *,
                 access_tokens: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsAccessTokenArgs']]]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 id_tokens: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsIdTokenArgs']]]] = None,
                 saml2_tokens: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsSaml2TokenArgs']]]] = None):
        """
        Input properties used for looking up and filtering ApplicationOptionalClaims resources.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsAccessTokenArgs']]] access_tokens: One or more `access_token` blocks as documented below.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsIdTokenArgs']]] id_tokens: One or more `id_token` blocks as documented below.
        :param pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsSaml2TokenArgs']]] saml2_tokens: One or more `saml2_token` blocks as documented below.
               
               > At least one of `access_token`, `id_token` or `saml2_token` must be specified
        """
        if access_tokens is not None:
            pulumi.set(__self__, "access_tokens", access_tokens)
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if id_tokens is not None:
            pulumi.set(__self__, "id_tokens", id_tokens)
        if saml2_tokens is not None:
            pulumi.set(__self__, "saml2_tokens", saml2_tokens)

    @property
    @pulumi.getter(name="accessTokens")
    def access_tokens(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsAccessTokenArgs']]]]:
        """
        One or more `access_token` blocks as documented below.
        """
        return pulumi.get(self, "access_tokens")

    @access_tokens.setter
    def access_tokens(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsAccessTokenArgs']]]]):
        pulumi.set(self, "access_tokens", value)

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
    @pulumi.getter(name="idTokens")
    def id_tokens(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsIdTokenArgs']]]]:
        """
        One or more `id_token` blocks as documented below.
        """
        return pulumi.get(self, "id_tokens")

    @id_tokens.setter
    def id_tokens(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsIdTokenArgs']]]]):
        pulumi.set(self, "id_tokens", value)

    @property
    @pulumi.getter(name="saml2Tokens")
    def saml2_tokens(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsSaml2TokenArgs']]]]:
        """
        One or more `saml2_token` blocks as documented below.

        > At least one of `access_token`, `id_token` or `saml2_token` must be specified
        """
        return pulumi.get(self, "saml2_tokens")

    @saml2_tokens.setter
    def saml2_tokens(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationOptionalClaimsSaml2TokenArgs']]]]):
        pulumi.set(self, "saml2_tokens", value)


class ApplicationOptionalClaims(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsAccessTokenArgs']]]]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 id_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsIdTokenArgs']]]]] = None,
                 saml2_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsSaml2TokenArgs']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application_registration = azuread.ApplicationRegistration("exampleApplicationRegistration", display_name="example")
        example_application_optional_claims = azuread.ApplicationOptionalClaims("exampleApplicationOptionalClaims",
            application_id=example_application_registration.id,
            access_tokens=[
                azuread.ApplicationOptionalClaimsAccessTokenArgs(
                    name="myclaim",
                ),
                azuread.ApplicationOptionalClaimsAccessTokenArgs(
                    name="otherclaim",
                ),
            ],
            id_tokens=[azuread.ApplicationOptionalClaimsIdTokenArgs(
                name="userclaim",
                source="user",
                essential=True,
                additional_properties=["emit_as_roles"],
            )],
            saml2_tokens=[azuread.ApplicationOptionalClaimsSaml2TokenArgs(
                name="samlexample",
            )])
        ```

        ## Import

        Application Optional Claims can be imported using the object ID of the application, in the following format.

        ```sh
         $ pulumi import azuread:index/applicationOptionalClaims:ApplicationOptionalClaims example /applications/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsAccessTokenArgs']]]] access_tokens: One or more `access_token` blocks as documented below.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsIdTokenArgs']]]] id_tokens: One or more `id_token` blocks as documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsSaml2TokenArgs']]]] saml2_tokens: One or more `saml2_token` blocks as documented below.
               
               > At least one of `access_token`, `id_token` or `saml2_token` must be specified
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationOptionalClaimsInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application_registration = azuread.ApplicationRegistration("exampleApplicationRegistration", display_name="example")
        example_application_optional_claims = azuread.ApplicationOptionalClaims("exampleApplicationOptionalClaims",
            application_id=example_application_registration.id,
            access_tokens=[
                azuread.ApplicationOptionalClaimsAccessTokenArgs(
                    name="myclaim",
                ),
                azuread.ApplicationOptionalClaimsAccessTokenArgs(
                    name="otherclaim",
                ),
            ],
            id_tokens=[azuread.ApplicationOptionalClaimsIdTokenArgs(
                name="userclaim",
                source="user",
                essential=True,
                additional_properties=["emit_as_roles"],
            )],
            saml2_tokens=[azuread.ApplicationOptionalClaimsSaml2TokenArgs(
                name="samlexample",
            )])
        ```

        ## Import

        Application Optional Claims can be imported using the object ID of the application, in the following format.

        ```sh
         $ pulumi import azuread:index/applicationOptionalClaims:ApplicationOptionalClaims example /applications/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationOptionalClaimsInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationOptionalClaimsInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsAccessTokenArgs']]]]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 id_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsIdTokenArgs']]]]] = None,
                 saml2_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsSaml2TokenArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationOptionalClaimsInitArgs.__new__(ApplicationOptionalClaimsInitArgs)

            __props__.__dict__["access_tokens"] = access_tokens
            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            __props__.__dict__["id_tokens"] = id_tokens
            __props__.__dict__["saml2_tokens"] = saml2_tokens
        super(ApplicationOptionalClaims, __self__).__init__(
            'azuread:index/applicationOptionalClaims:ApplicationOptionalClaims',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsAccessTokenArgs']]]]] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            id_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsIdTokenArgs']]]]] = None,
            saml2_tokens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsSaml2TokenArgs']]]]] = None) -> 'ApplicationOptionalClaims':
        """
        Get an existing ApplicationOptionalClaims resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsAccessTokenArgs']]]] access_tokens: One or more `access_token` blocks as documented below.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsIdTokenArgs']]]] id_tokens: One or more `id_token` blocks as documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationOptionalClaimsSaml2TokenArgs']]]] saml2_tokens: One or more `saml2_token` blocks as documented below.
               
               > At least one of `access_token`, `id_token` or `saml2_token` must be specified
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationOptionalClaimsState.__new__(_ApplicationOptionalClaimsState)

        __props__.__dict__["access_tokens"] = access_tokens
        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["id_tokens"] = id_tokens
        __props__.__dict__["saml2_tokens"] = saml2_tokens
        return ApplicationOptionalClaims(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessTokens")
    def access_tokens(self) -> pulumi.Output[Optional[Sequence['outputs.ApplicationOptionalClaimsAccessToken']]]:
        """
        One or more `access_token` blocks as documented below.
        """
        return pulumi.get(self, "access_tokens")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="idTokens")
    def id_tokens(self) -> pulumi.Output[Optional[Sequence['outputs.ApplicationOptionalClaimsIdToken']]]:
        """
        One or more `id_token` blocks as documented below.
        """
        return pulumi.get(self, "id_tokens")

    @property
    @pulumi.getter(name="saml2Tokens")
    def saml2_tokens(self) -> pulumi.Output[Optional[Sequence['outputs.ApplicationOptionalClaimsSaml2Token']]]:
        """
        One or more `saml2_token` blocks as documented below.

        > At least one of `access_token`, `id_token` or `saml2_token` must be specified
        """
        return pulumi.get(self, "saml2_tokens")

