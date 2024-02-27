# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AuthenticationStrengthPolicyArgs', 'AuthenticationStrengthPolicy']

@pulumi.input_type
class AuthenticationStrengthPolicyArgs:
    def __init__(__self__, *,
                 allowed_combinations: pulumi.Input[Sequence[pulumi.Input[str]]],
                 display_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AuthenticationStrengthPolicy resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_combinations: List of allowed authentication methods for this authentication strength policy.
        :param pulumi.Input[str] display_name: The friendly name for this authentication strength policy.
        :param pulumi.Input[str] description: The description for this authentication strength policy.
        """
        pulumi.set(__self__, "allowed_combinations", allowed_combinations)
        pulumi.set(__self__, "display_name", display_name)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="allowedCombinations")
    def allowed_combinations(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of allowed authentication methods for this authentication strength policy.
        """
        return pulumi.get(self, "allowed_combinations")

    @allowed_combinations.setter
    def allowed_combinations(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_combinations", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The friendly name for this authentication strength policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for this authentication strength policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _AuthenticationStrengthPolicyState:
    def __init__(__self__, *,
                 allowed_combinations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AuthenticationStrengthPolicy resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_combinations: List of allowed authentication methods for this authentication strength policy.
        :param pulumi.Input[str] description: The description for this authentication strength policy.
        :param pulumi.Input[str] display_name: The friendly name for this authentication strength policy.
        """
        if allowed_combinations is not None:
            pulumi.set(__self__, "allowed_combinations", allowed_combinations)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter(name="allowedCombinations")
    def allowed_combinations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of allowed authentication methods for this authentication strength policy.
        """
        return pulumi.get(self, "allowed_combinations")

    @allowed_combinations.setter
    def allowed_combinations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_combinations", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for this authentication strength policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The friendly name for this authentication strength policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)


class AuthenticationStrengthPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_combinations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Authentication Strength Policy within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.index.authentication_strength_policy.AuthenticationStrengthPolicy("example",
            display_name=Example Authentication Strength Policy,
            description=Policy for demo purposes,
            allowed_combinations=[
                fido2,
                password,
            ])
        example2 = azuread.index.authentication_strength_policy.AuthenticationStrengthPolicy("example2",
            display_name=Example Authentication Strength Policy,
            description=Policy for demo purposes with all possible combinations,
            allowed_combinations=[
                fido2,
                password,
                deviceBasedPush,
                temporaryAccessPassOneTime,
                federatedMultiFactor,
                federatedSingleFactor,
                hardwareOath,federatedSingleFactor,
                microsoftAuthenticatorPush,federatedSingleFactor,
                password,hardwareOath,
                password,microsoftAuthenticatorPush,
                password,sms,
                password,softwareOath,
                password,voice,
                sms,
                sms,federatedSingleFactor,
                softwareOath,federatedSingleFactor,
                temporaryAccessPassMultiUse,
                voice,federatedSingleFactor,
                windowsHelloForBusiness,
                x509CertificateMultiFactor,
                x509CertificateSingleFactor,
            ])
        ```

        ## Import

        Authentication Strength Policies can be imported using the `id`, e.g.

        ```sh
        $ pulumi import azuread:index/authenticationStrengthPolicy:AuthenticationStrengthPolicy my_policy 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_combinations: List of allowed authentication methods for this authentication strength policy.
        :param pulumi.Input[str] description: The description for this authentication strength policy.
        :param pulumi.Input[str] display_name: The friendly name for this authentication strength policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthenticationStrengthPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Authentication Strength Policy within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.index.authentication_strength_policy.AuthenticationStrengthPolicy("example",
            display_name=Example Authentication Strength Policy,
            description=Policy for demo purposes,
            allowed_combinations=[
                fido2,
                password,
            ])
        example2 = azuread.index.authentication_strength_policy.AuthenticationStrengthPolicy("example2",
            display_name=Example Authentication Strength Policy,
            description=Policy for demo purposes with all possible combinations,
            allowed_combinations=[
                fido2,
                password,
                deviceBasedPush,
                temporaryAccessPassOneTime,
                federatedMultiFactor,
                federatedSingleFactor,
                hardwareOath,federatedSingleFactor,
                microsoftAuthenticatorPush,federatedSingleFactor,
                password,hardwareOath,
                password,microsoftAuthenticatorPush,
                password,sms,
                password,softwareOath,
                password,voice,
                sms,
                sms,federatedSingleFactor,
                softwareOath,federatedSingleFactor,
                temporaryAccessPassMultiUse,
                voice,federatedSingleFactor,
                windowsHelloForBusiness,
                x509CertificateMultiFactor,
                x509CertificateSingleFactor,
            ])
        ```

        ## Import

        Authentication Strength Policies can be imported using the `id`, e.g.

        ```sh
        $ pulumi import azuread:index/authenticationStrengthPolicy:AuthenticationStrengthPolicy my_policy 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param AuthenticationStrengthPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthenticationStrengthPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_combinations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthenticationStrengthPolicyArgs.__new__(AuthenticationStrengthPolicyArgs)

            if allowed_combinations is None and not opts.urn:
                raise TypeError("Missing required property 'allowed_combinations'")
            __props__.__dict__["allowed_combinations"] = allowed_combinations
            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
        super(AuthenticationStrengthPolicy, __self__).__init__(
            'azuread:index/authenticationStrengthPolicy:AuthenticationStrengthPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_combinations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None) -> 'AuthenticationStrengthPolicy':
        """
        Get an existing AuthenticationStrengthPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_combinations: List of allowed authentication methods for this authentication strength policy.
        :param pulumi.Input[str] description: The description for this authentication strength policy.
        :param pulumi.Input[str] display_name: The friendly name for this authentication strength policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthenticationStrengthPolicyState.__new__(_AuthenticationStrengthPolicyState)

        __props__.__dict__["allowed_combinations"] = allowed_combinations
        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        return AuthenticationStrengthPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedCombinations")
    def allowed_combinations(self) -> pulumi.Output[Sequence[str]]:
        """
        List of allowed authentication methods for this authentication strength policy.
        """
        return pulumi.get(self, "allowed_combinations")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for this authentication strength policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The friendly name for this authentication strength policy.
        """
        return pulumi.get(self, "display_name")

