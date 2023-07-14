# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ClaimsMappingPolicyArgs', 'ClaimsMappingPolicy']

@pulumi.input_type
class ClaimsMappingPolicyArgs:
    def __init__(__self__, *,
                 definitions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 display_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a ClaimsMappingPolicy resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] definitions: A string collection containing a JSON string that defines the rules and settings for this policy
        :param pulumi.Input[str] display_name: Display name for this policy
        """
        pulumi.set(__self__, "definitions", definitions)
        pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter
    def definitions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A string collection containing a JSON string that defines the rules and settings for this policy
        """
        return pulumi.get(self, "definitions")

    @definitions.setter
    def definitions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "definitions", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        Display name for this policy
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)


@pulumi.input_type
class _ClaimsMappingPolicyState:
    def __init__(__self__, *,
                 definitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ClaimsMappingPolicy resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] definitions: A string collection containing a JSON string that defines the rules and settings for this policy
        :param pulumi.Input[str] display_name: Display name for this policy
        """
        if definitions is not None:
            pulumi.set(__self__, "definitions", definitions)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter
    def definitions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A string collection containing a JSON string that defines the rules and settings for this policy
        """
        return pulumi.get(self, "definitions")

    @definitions.setter
    def definitions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "definitions", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for this policy
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)


class ClaimsMappingPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ClaimsMappingPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] definitions: A string collection containing a JSON string that defines the rules and settings for this policy
        :param pulumi.Input[str] display_name: Display name for this policy
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClaimsMappingPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ClaimsMappingPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ClaimsMappingPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClaimsMappingPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClaimsMappingPolicyArgs.__new__(ClaimsMappingPolicyArgs)

            if definitions is None and not opts.urn:
                raise TypeError("Missing required property 'definitions'")
            __props__.__dict__["definitions"] = definitions
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
        super(ClaimsMappingPolicy, __self__).__init__(
            'azuread:index/claimsMappingPolicy:ClaimsMappingPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            definitions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            display_name: Optional[pulumi.Input[str]] = None) -> 'ClaimsMappingPolicy':
        """
        Get an existing ClaimsMappingPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] definitions: A string collection containing a JSON string that defines the rules and settings for this policy
        :param pulumi.Input[str] display_name: Display name for this policy
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClaimsMappingPolicyState.__new__(_ClaimsMappingPolicyState)

        __props__.__dict__["definitions"] = definitions
        __props__.__dict__["display_name"] = display_name
        return ClaimsMappingPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def definitions(self) -> pulumi.Output[Sequence[str]]:
        """
        A string collection containing a JSON string that defines the rules and settings for this policy
        """
        return pulumi.get(self, "definitions")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Display name for this policy
        """
        return pulumi.get(self, "display_name")

