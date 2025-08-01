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
from ._inputs import *

__all__ = ['ConditionalAccessPolicyArgs', 'ConditionalAccessPolicy']

@pulumi.input_type
class ConditionalAccessPolicyArgs:
    def __init__(__self__, *,
                 conditions: pulumi.Input['ConditionalAccessPolicyConditionsArgs'],
                 display_name: pulumi.Input[_builtins.str],
                 state: pulumi.Input[_builtins.str],
                 grant_controls: Optional[pulumi.Input['ConditionalAccessPolicyGrantControlsArgs']] = None,
                 session_controls: Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']] = None):
        """
        The set of arguments for constructing a ConditionalAccessPolicy resource.
        :param pulumi.Input['ConditionalAccessPolicyConditionsArgs'] conditions: A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        :param pulumi.Input[_builtins.str] display_name: The friendly name for this Conditional Access Policy.
        :param pulumi.Input[_builtins.str] state: Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        :param pulumi.Input['ConditionalAccessPolicyGrantControlsArgs'] grant_controls: A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        :param pulumi.Input['ConditionalAccessPolicySessionControlsArgs'] session_controls: A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
               
               > Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
        """
        pulumi.set(__self__, "conditions", conditions)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "state", state)
        if grant_controls is not None:
            pulumi.set(__self__, "grant_controls", grant_controls)
        if session_controls is not None:
            pulumi.set(__self__, "session_controls", session_controls)

    @_builtins.property
    @pulumi.getter
    def conditions(self) -> pulumi.Input['ConditionalAccessPolicyConditionsArgs']:
        """
        A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: pulumi.Input['ConditionalAccessPolicyConditionsArgs']):
        pulumi.set(self, "conditions", value)

    @_builtins.property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[_builtins.str]:
        """
        The friendly name for this Conditional Access Policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "display_name", value)

    @_builtins.property
    @pulumi.getter
    def state(self) -> pulumi.Input[_builtins.str]:
        """
        Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "state", value)

    @_builtins.property
    @pulumi.getter(name="grantControls")
    def grant_controls(self) -> Optional[pulumi.Input['ConditionalAccessPolicyGrantControlsArgs']]:
        """
        A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        """
        return pulumi.get(self, "grant_controls")

    @grant_controls.setter
    def grant_controls(self, value: Optional[pulumi.Input['ConditionalAccessPolicyGrantControlsArgs']]):
        pulumi.set(self, "grant_controls", value)

    @_builtins.property
    @pulumi.getter(name="sessionControls")
    def session_controls(self) -> Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']]:
        """
        A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.

        > Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
        """
        return pulumi.get(self, "session_controls")

    @session_controls.setter
    def session_controls(self, value: Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']]):
        pulumi.set(self, "session_controls", value)


@pulumi.input_type
class _ConditionalAccessPolicyState:
    def __init__(__self__, *,
                 conditions: Optional[pulumi.Input['ConditionalAccessPolicyConditionsArgs']] = None,
                 display_name: Optional[pulumi.Input[_builtins.str]] = None,
                 grant_controls: Optional[pulumi.Input['ConditionalAccessPolicyGrantControlsArgs']] = None,
                 object_id: Optional[pulumi.Input[_builtins.str]] = None,
                 session_controls: Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']] = None,
                 state: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering ConditionalAccessPolicy resources.
        :param pulumi.Input['ConditionalAccessPolicyConditionsArgs'] conditions: A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        :param pulumi.Input[_builtins.str] display_name: The friendly name for this Conditional Access Policy.
        :param pulumi.Input['ConditionalAccessPolicyGrantControlsArgs'] grant_controls: A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        :param pulumi.Input[_builtins.str] object_id: The object ID of the policy
        :param pulumi.Input['ConditionalAccessPolicySessionControlsArgs'] session_controls: A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
               
               > Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
        :param pulumi.Input[_builtins.str] state: Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if grant_controls is not None:
            pulumi.set(__self__, "grant_controls", grant_controls)
        if object_id is not None:
            pulumi.set(__self__, "object_id", object_id)
        if session_controls is not None:
            pulumi.set(__self__, "session_controls", session_controls)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @_builtins.property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input['ConditionalAccessPolicyConditionsArgs']]:
        """
        A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input['ConditionalAccessPolicyConditionsArgs']]):
        pulumi.set(self, "conditions", value)

    @_builtins.property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The friendly name for this Conditional Access Policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "display_name", value)

    @_builtins.property
    @pulumi.getter(name="grantControls")
    def grant_controls(self) -> Optional[pulumi.Input['ConditionalAccessPolicyGrantControlsArgs']]:
        """
        A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        """
        return pulumi.get(self, "grant_controls")

    @grant_controls.setter
    def grant_controls(self, value: Optional[pulumi.Input['ConditionalAccessPolicyGrantControlsArgs']]):
        pulumi.set(self, "grant_controls", value)

    @_builtins.property
    @pulumi.getter(name="objectId")
    def object_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The object ID of the policy
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "object_id", value)

    @_builtins.property
    @pulumi.getter(name="sessionControls")
    def session_controls(self) -> Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']]:
        """
        A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.

        > Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
        """
        return pulumi.get(self, "session_controls")

    @session_controls.setter
    def session_controls(self, value: Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']]):
        pulumi.set(self, "session_controls", value)

    @_builtins.property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "state", value)


@pulumi.type_token("azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy")
class ConditionalAccessPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 conditions: Optional[pulumi.Input[Union['ConditionalAccessPolicyConditionsArgs', 'ConditionalAccessPolicyConditionsArgsDict']]] = None,
                 display_name: Optional[pulumi.Input[_builtins.str]] = None,
                 grant_controls: Optional[pulumi.Input[Union['ConditionalAccessPolicyGrantControlsArgs', 'ConditionalAccessPolicyGrantControlsArgsDict']]] = None,
                 session_controls: Optional[pulumi.Input[Union['ConditionalAccessPolicySessionControlsArgs', 'ConditionalAccessPolicySessionControlsArgsDict']]] = None,
                 state: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### All users except guests or external users

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ConditionalAccessPolicy("example",
            display_name="example policy",
            state="disabled",
            conditions={
                "client_app_types": ["all"],
                "sign_in_risk_levels": ["medium"],
                "user_risk_levels": ["medium"],
                "applications": {
                    "included_applications": ["All"],
                    "excluded_applications": [],
                },
                "devices": {
                    "filter": {
                        "mode": "exclude",
                        "rule": "device.operatingSystem eq \\"Doors\\"",
                    },
                },
                "locations": {
                    "included_locations": ["All"],
                    "excluded_locations": ["AllTrusted"],
                },
                "platforms": {
                    "included_platforms": ["android"],
                    "excluded_platforms": ["iOS"],
                },
                "users": {
                    "included_users": ["All"],
                    "excluded_users": ["GuestsOrExternalUsers"],
                },
            },
            grant_controls={
                "operator": "OR",
                "built_in_controls": ["mfa"],
            },
            session_controls={
                "application_enforced_restrictions_enabled": True,
                "disable_resilience_defaults": False,
                "sign_in_frequency": 10,
                "sign_in_frequency_period": "hours",
                "cloud_app_security_policy": "monitorOnly",
            })
        ```

        ### Included client applications / service principals

        ```python
        import pulumi
        import pulumi_azuread as azuread

        current = azuread.get_client_config()
        example = azuread.ConditionalAccessPolicy("example",
            display_name="example policy",
            state="disabled",
            conditions={
                "client_app_types": ["all"],
                "applications": {
                    "included_applications": ["All"],
                },
                "client_applications": {
                    "included_service_principals": [current.object_id],
                    "excluded_service_principals": [],
                },
                "users": {
                    "included_users": ["None"],
                },
            },
            grant_controls={
                "operator": "OR",
                "built_in_controls": ["block"],
            })
        ```

        ### Excluded client applications / service principals

        ```python
        import pulumi
        import pulumi_azuread as azuread

        current = azuread.get_client_config()
        example = azuread.ConditionalAccessPolicy("example",
            display_name="example policy",
            state="disabled",
            conditions={
                "client_app_types": ["all"],
                "applications": {
                    "included_applications": ["All"],
                },
                "client_applications": {
                    "included_service_principals": ["ServicePrincipalsInMyTenant"],
                    "excluded_service_principals": [current.object_id],
                },
                "users": {
                    "included_users": ["None"],
                },
            },
            grant_controls={
                "operator": "OR",
                "built_in_controls": ["block"],
            })
        ```

        ## Import

        Conditional Access Policies can be imported using the `id`, e.g.

        ```sh
        $ pulumi import azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy my_location /identity/conditionalAccess/policies/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ConditionalAccessPolicyConditionsArgs', 'ConditionalAccessPolicyConditionsArgsDict']] conditions: A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        :param pulumi.Input[_builtins.str] display_name: The friendly name for this Conditional Access Policy.
        :param pulumi.Input[Union['ConditionalAccessPolicyGrantControlsArgs', 'ConditionalAccessPolicyGrantControlsArgsDict']] grant_controls: A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        :param pulumi.Input[Union['ConditionalAccessPolicySessionControlsArgs', 'ConditionalAccessPolicySessionControlsArgsDict']] session_controls: A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
               
               > Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
        :param pulumi.Input[_builtins.str] state: Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConditionalAccessPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### All users except guests or external users

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ConditionalAccessPolicy("example",
            display_name="example policy",
            state="disabled",
            conditions={
                "client_app_types": ["all"],
                "sign_in_risk_levels": ["medium"],
                "user_risk_levels": ["medium"],
                "applications": {
                    "included_applications": ["All"],
                    "excluded_applications": [],
                },
                "devices": {
                    "filter": {
                        "mode": "exclude",
                        "rule": "device.operatingSystem eq \\"Doors\\"",
                    },
                },
                "locations": {
                    "included_locations": ["All"],
                    "excluded_locations": ["AllTrusted"],
                },
                "platforms": {
                    "included_platforms": ["android"],
                    "excluded_platforms": ["iOS"],
                },
                "users": {
                    "included_users": ["All"],
                    "excluded_users": ["GuestsOrExternalUsers"],
                },
            },
            grant_controls={
                "operator": "OR",
                "built_in_controls": ["mfa"],
            },
            session_controls={
                "application_enforced_restrictions_enabled": True,
                "disable_resilience_defaults": False,
                "sign_in_frequency": 10,
                "sign_in_frequency_period": "hours",
                "cloud_app_security_policy": "monitorOnly",
            })
        ```

        ### Included client applications / service principals

        ```python
        import pulumi
        import pulumi_azuread as azuread

        current = azuread.get_client_config()
        example = azuread.ConditionalAccessPolicy("example",
            display_name="example policy",
            state="disabled",
            conditions={
                "client_app_types": ["all"],
                "applications": {
                    "included_applications": ["All"],
                },
                "client_applications": {
                    "included_service_principals": [current.object_id],
                    "excluded_service_principals": [],
                },
                "users": {
                    "included_users": ["None"],
                },
            },
            grant_controls={
                "operator": "OR",
                "built_in_controls": ["block"],
            })
        ```

        ### Excluded client applications / service principals

        ```python
        import pulumi
        import pulumi_azuread as azuread

        current = azuread.get_client_config()
        example = azuread.ConditionalAccessPolicy("example",
            display_name="example policy",
            state="disabled",
            conditions={
                "client_app_types": ["all"],
                "applications": {
                    "included_applications": ["All"],
                },
                "client_applications": {
                    "included_service_principals": ["ServicePrincipalsInMyTenant"],
                    "excluded_service_principals": [current.object_id],
                },
                "users": {
                    "included_users": ["None"],
                },
            },
            grant_controls={
                "operator": "OR",
                "built_in_controls": ["block"],
            })
        ```

        ## Import

        Conditional Access Policies can be imported using the `id`, e.g.

        ```sh
        $ pulumi import azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy my_location /identity/conditionalAccess/policies/00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param ConditionalAccessPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConditionalAccessPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 conditions: Optional[pulumi.Input[Union['ConditionalAccessPolicyConditionsArgs', 'ConditionalAccessPolicyConditionsArgsDict']]] = None,
                 display_name: Optional[pulumi.Input[_builtins.str]] = None,
                 grant_controls: Optional[pulumi.Input[Union['ConditionalAccessPolicyGrantControlsArgs', 'ConditionalAccessPolicyGrantControlsArgsDict']]] = None,
                 session_controls: Optional[pulumi.Input[Union['ConditionalAccessPolicySessionControlsArgs', 'ConditionalAccessPolicySessionControlsArgsDict']]] = None,
                 state: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConditionalAccessPolicyArgs.__new__(ConditionalAccessPolicyArgs)

            if conditions is None and not opts.urn:
                raise TypeError("Missing required property 'conditions'")
            __props__.__dict__["conditions"] = conditions
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["grant_controls"] = grant_controls
            __props__.__dict__["session_controls"] = session_controls
            if state is None and not opts.urn:
                raise TypeError("Missing required property 'state'")
            __props__.__dict__["state"] = state
            __props__.__dict__["object_id"] = None
        super(ConditionalAccessPolicy, __self__).__init__(
            'azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            conditions: Optional[pulumi.Input[Union['ConditionalAccessPolicyConditionsArgs', 'ConditionalAccessPolicyConditionsArgsDict']]] = None,
            display_name: Optional[pulumi.Input[_builtins.str]] = None,
            grant_controls: Optional[pulumi.Input[Union['ConditionalAccessPolicyGrantControlsArgs', 'ConditionalAccessPolicyGrantControlsArgsDict']]] = None,
            object_id: Optional[pulumi.Input[_builtins.str]] = None,
            session_controls: Optional[pulumi.Input[Union['ConditionalAccessPolicySessionControlsArgs', 'ConditionalAccessPolicySessionControlsArgsDict']]] = None,
            state: Optional[pulumi.Input[_builtins.str]] = None) -> 'ConditionalAccessPolicy':
        """
        Get an existing ConditionalAccessPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ConditionalAccessPolicyConditionsArgs', 'ConditionalAccessPolicyConditionsArgsDict']] conditions: A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        :param pulumi.Input[_builtins.str] display_name: The friendly name for this Conditional Access Policy.
        :param pulumi.Input[Union['ConditionalAccessPolicyGrantControlsArgs', 'ConditionalAccessPolicyGrantControlsArgsDict']] grant_controls: A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        :param pulumi.Input[_builtins.str] object_id: The object ID of the policy
        :param pulumi.Input[Union['ConditionalAccessPolicySessionControlsArgs', 'ConditionalAccessPolicySessionControlsArgsDict']] session_controls: A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
               
               > Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
        :param pulumi.Input[_builtins.str] state: Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConditionalAccessPolicyState.__new__(_ConditionalAccessPolicyState)

        __props__.__dict__["conditions"] = conditions
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["grant_controls"] = grant_controls
        __props__.__dict__["object_id"] = object_id
        __props__.__dict__["session_controls"] = session_controls
        __props__.__dict__["state"] = state
        return ConditionalAccessPolicy(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def conditions(self) -> pulumi.Output['outputs.ConditionalAccessPolicyConditions']:
        """
        A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        """
        return pulumi.get(self, "conditions")

    @_builtins.property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[_builtins.str]:
        """
        The friendly name for this Conditional Access Policy.
        """
        return pulumi.get(self, "display_name")

    @_builtins.property
    @pulumi.getter(name="grantControls")
    def grant_controls(self) -> pulumi.Output[Optional['outputs.ConditionalAccessPolicyGrantControls']]:
        """
        A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        """
        return pulumi.get(self, "grant_controls")

    @_builtins.property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Output[_builtins.str]:
        """
        The object ID of the policy
        """
        return pulumi.get(self, "object_id")

    @_builtins.property
    @pulumi.getter(name="sessionControls")
    def session_controls(self) -> pulumi.Output[Optional['outputs.ConditionalAccessPolicySessionControls']]:
        """
        A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.

        > Note: At least one of `grant_controls` and/or `session_controls` blocks must be specified.
        """
        return pulumi.get(self, "session_controls")

    @_builtins.property
    @pulumi.getter
    def state(self) -> pulumi.Output[_builtins.str]:
        """
        Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        return pulumi.get(self, "state")

