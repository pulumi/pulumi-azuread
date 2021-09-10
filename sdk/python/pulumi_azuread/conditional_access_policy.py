# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ConditionalAccessPolicyArgs', 'ConditionalAccessPolicy']

@pulumi.input_type
class ConditionalAccessPolicyArgs:
    def __init__(__self__, *,
                 conditions: pulumi.Input['ConditionalAccessPolicyConditionsArgs'],
                 display_name: pulumi.Input[str],
                 grant_controls: pulumi.Input['ConditionalAccessPolicyGrantControlsArgs'],
                 state: pulumi.Input[str],
                 session_controls: Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']] = None):
        """
        The set of arguments for constructing a ConditionalAccessPolicy resource.
        :param pulumi.Input['ConditionalAccessPolicyConditionsArgs'] conditions: A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        :param pulumi.Input[str] display_name: The friendly name for this Conditional Access Policy.
        :param pulumi.Input['ConditionalAccessPolicyGrantControlsArgs'] grant_controls: A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        :param pulumi.Input[str] state: Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        :param pulumi.Input['ConditionalAccessPolicySessionControlsArgs'] session_controls: A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
        """
        pulumi.set(__self__, "conditions", conditions)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "grant_controls", grant_controls)
        pulumi.set(__self__, "state", state)
        if session_controls is not None:
            pulumi.set(__self__, "session_controls", session_controls)

    @property
    @pulumi.getter
    def conditions(self) -> pulumi.Input['ConditionalAccessPolicyConditionsArgs']:
        """
        A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: pulumi.Input['ConditionalAccessPolicyConditionsArgs']):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The friendly name for this Conditional Access Policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="grantControls")
    def grant_controls(self) -> pulumi.Input['ConditionalAccessPolicyGrantControlsArgs']:
        """
        A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        """
        return pulumi.get(self, "grant_controls")

    @grant_controls.setter
    def grant_controls(self, value: pulumi.Input['ConditionalAccessPolicyGrantControlsArgs']):
        pulumi.set(self, "grant_controls", value)

    @property
    @pulumi.getter
    def state(self) -> pulumi.Input[str]:
        """
        Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: pulumi.Input[str]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="sessionControls")
    def session_controls(self) -> Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']]:
        """
        A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
        """
        return pulumi.get(self, "session_controls")

    @session_controls.setter
    def session_controls(self, value: Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']]):
        pulumi.set(self, "session_controls", value)


@pulumi.input_type
class _ConditionalAccessPolicyState:
    def __init__(__self__, *,
                 conditions: Optional[pulumi.Input['ConditionalAccessPolicyConditionsArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 grant_controls: Optional[pulumi.Input['ConditionalAccessPolicyGrantControlsArgs']] = None,
                 session_controls: Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConditionalAccessPolicy resources.
        :param pulumi.Input['ConditionalAccessPolicyConditionsArgs'] conditions: A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        :param pulumi.Input[str] display_name: The friendly name for this Conditional Access Policy.
        :param pulumi.Input['ConditionalAccessPolicyGrantControlsArgs'] grant_controls: A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        :param pulumi.Input['ConditionalAccessPolicySessionControlsArgs'] session_controls: A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
        :param pulumi.Input[str] state: Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if grant_controls is not None:
            pulumi.set(__self__, "grant_controls", grant_controls)
        if session_controls is not None:
            pulumi.set(__self__, "session_controls", session_controls)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[pulumi.Input['ConditionalAccessPolicyConditionsArgs']]:
        """
        A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        """
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[pulumi.Input['ConditionalAccessPolicyConditionsArgs']]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The friendly name for this Conditional Access Policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="grantControls")
    def grant_controls(self) -> Optional[pulumi.Input['ConditionalAccessPolicyGrantControlsArgs']]:
        """
        A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        """
        return pulumi.get(self, "grant_controls")

    @grant_controls.setter
    def grant_controls(self, value: Optional[pulumi.Input['ConditionalAccessPolicyGrantControlsArgs']]):
        pulumi.set(self, "grant_controls", value)

    @property
    @pulumi.getter(name="sessionControls")
    def session_controls(self) -> Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']]:
        """
        A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
        """
        return pulumi.get(self, "session_controls")

    @session_controls.setter
    def session_controls(self, value: Optional[pulumi.Input['ConditionalAccessPolicySessionControlsArgs']]):
        pulumi.set(self, "session_controls", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


class ConditionalAccessPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 conditions: Optional[pulumi.Input[pulumi.InputType['ConditionalAccessPolicyConditionsArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 grant_controls: Optional[pulumi.Input[pulumi.InputType['ConditionalAccessPolicyGrantControlsArgs']]] = None,
                 session_controls: Optional[pulumi.Input[pulumi.InputType['ConditionalAccessPolicySessionControlsArgs']]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Conditional Access Policy within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ConditionalAccessPolicy("example",
            conditions=azuread.ConditionalAccessPolicyConditionsArgs(
                applications=azuread.ConditionalAccessPolicyConditionsApplicationsArgs(
                    excluded_applications=["00000004-0000-0ff1-ce00-000000000000"],
                    included_applications=["All"],
                ),
                client_app_types=["all"],
                locations=azuread.ConditionalAccessPolicyConditionsLocationsArgs(
                    excluded_locations=["AllTrusted"],
                    included_locations=["All"],
                ),
                platforms=azuread.ConditionalAccessPolicyConditionsPlatformsArgs(
                    excluded_platforms=["iOS"],
                    included_platforms=["android"],
                ),
                sign_in_risk_levels=["medium"],
                user_risk_levels=["medium"],
                users=azuread.ConditionalAccessPolicyConditionsUsersArgs(
                    excluded_users=["GuestsOrExternalUsers"],
                    included_users=["All"],
                ),
            ),
            display_name="example policy",
            grant_controls=azuread.ConditionalAccessPolicyGrantControlsArgs(
                built_in_controls=["mfa"],
                operator="OR",
            ),
            session_controls=azuread.ConditionalAccessPolicySessionControlsArgs(
                application_enforced_restrictions=[{
                    "enabled": True,
                }],
                cloud_app_security=[{
                    "cloudAppSecurityType": "monitorOnly",
                    "enabled": True,
                }],
                sign_in_frequency=[{
                    "enabled": True,
                    "type": "hours",
                    "value": 10,
                }],
            ),
            state="disabled")
        ```

        ## Import

        Conditional Access Policies can be imported using the `id`, e.g.

        ```sh
         $ pulumi import azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy my_location 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ConditionalAccessPolicyConditionsArgs']] conditions: A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        :param pulumi.Input[str] display_name: The friendly name for this Conditional Access Policy.
        :param pulumi.Input[pulumi.InputType['ConditionalAccessPolicyGrantControlsArgs']] grant_controls: A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        :param pulumi.Input[pulumi.InputType['ConditionalAccessPolicySessionControlsArgs']] session_controls: A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
        :param pulumi.Input[str] state: Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConditionalAccessPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Conditional Access Policy within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ConditionalAccessPolicy("example",
            conditions=azuread.ConditionalAccessPolicyConditionsArgs(
                applications=azuread.ConditionalAccessPolicyConditionsApplicationsArgs(
                    excluded_applications=["00000004-0000-0ff1-ce00-000000000000"],
                    included_applications=["All"],
                ),
                client_app_types=["all"],
                locations=azuread.ConditionalAccessPolicyConditionsLocationsArgs(
                    excluded_locations=["AllTrusted"],
                    included_locations=["All"],
                ),
                platforms=azuread.ConditionalAccessPolicyConditionsPlatformsArgs(
                    excluded_platforms=["iOS"],
                    included_platforms=["android"],
                ),
                sign_in_risk_levels=["medium"],
                user_risk_levels=["medium"],
                users=azuread.ConditionalAccessPolicyConditionsUsersArgs(
                    excluded_users=["GuestsOrExternalUsers"],
                    included_users=["All"],
                ),
            ),
            display_name="example policy",
            grant_controls=azuread.ConditionalAccessPolicyGrantControlsArgs(
                built_in_controls=["mfa"],
                operator="OR",
            ),
            session_controls=azuread.ConditionalAccessPolicySessionControlsArgs(
                application_enforced_restrictions=[{
                    "enabled": True,
                }],
                cloud_app_security=[{
                    "cloudAppSecurityType": "monitorOnly",
                    "enabled": True,
                }],
                sign_in_frequency=[{
                    "enabled": True,
                    "type": "hours",
                    "value": 10,
                }],
            ),
            state="disabled")
        ```

        ## Import

        Conditional Access Policies can be imported using the `id`, e.g.

        ```sh
         $ pulumi import azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy my_location 00000000-0000-0000-0000-000000000000
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
                 conditions: Optional[pulumi.Input[pulumi.InputType['ConditionalAccessPolicyConditionsArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 grant_controls: Optional[pulumi.Input[pulumi.InputType['ConditionalAccessPolicyGrantControlsArgs']]] = None,
                 session_controls: Optional[pulumi.Input[pulumi.InputType['ConditionalAccessPolicySessionControlsArgs']]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
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
            if grant_controls is None and not opts.urn:
                raise TypeError("Missing required property 'grant_controls'")
            __props__.__dict__["grant_controls"] = grant_controls
            __props__.__dict__["session_controls"] = session_controls
            if state is None and not opts.urn:
                raise TypeError("Missing required property 'state'")
            __props__.__dict__["state"] = state
        super(ConditionalAccessPolicy, __self__).__init__(
            'azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            conditions: Optional[pulumi.Input[pulumi.InputType['ConditionalAccessPolicyConditionsArgs']]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            grant_controls: Optional[pulumi.Input[pulumi.InputType['ConditionalAccessPolicyGrantControlsArgs']]] = None,
            session_controls: Optional[pulumi.Input[pulumi.InputType['ConditionalAccessPolicySessionControlsArgs']]] = None,
            state: Optional[pulumi.Input[str]] = None) -> 'ConditionalAccessPolicy':
        """
        Get an existing ConditionalAccessPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ConditionalAccessPolicyConditionsArgs']] conditions: A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        :param pulumi.Input[str] display_name: The friendly name for this Conditional Access Policy.
        :param pulumi.Input[pulumi.InputType['ConditionalAccessPolicyGrantControlsArgs']] grant_controls: A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        :param pulumi.Input[pulumi.InputType['ConditionalAccessPolicySessionControlsArgs']] session_controls: A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
        :param pulumi.Input[str] state: Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConditionalAccessPolicyState.__new__(_ConditionalAccessPolicyState)

        __props__.__dict__["conditions"] = conditions
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["grant_controls"] = grant_controls
        __props__.__dict__["session_controls"] = session_controls
        __props__.__dict__["state"] = state
        return ConditionalAccessPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def conditions(self) -> pulumi.Output['outputs.ConditionalAccessPolicyConditions']:
        """
        A `conditions` block as documented below, which specifies the rules that must be met for the policy to apply.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The friendly name for this Conditional Access Policy.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="grantControls")
    def grant_controls(self) -> pulumi.Output['outputs.ConditionalAccessPolicyGrantControls']:
        """
        A `grant_controls` block as documented below, which specifies the grant controls that must be fulfilled to pass the policy.
        """
        return pulumi.get(self, "grant_controls")

    @property
    @pulumi.getter(name="sessionControls")
    def session_controls(self) -> pulumi.Output[Optional['outputs.ConditionalAccessPolicySessionControls']]:
        """
        A `session_controls` block as documented below, which specifies the session controls that are enforced after sign-in.
        """
        return pulumi.get(self, "session_controls")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Specifies the state of the policy object. Possible values are: `enabled`, `disabled` and `enabledForReportingButNotEnforced`
        """
        return pulumi.get(self, "state")

