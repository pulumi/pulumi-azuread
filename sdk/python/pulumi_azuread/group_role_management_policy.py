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
from . import outputs
from ._inputs import *

__all__ = ['GroupRoleManagementPolicyArgs', 'GroupRoleManagementPolicy']

@pulumi.input_type
class GroupRoleManagementPolicyArgs:
    def __init__(__self__, *,
                 group_id: pulumi.Input[builtins.str],
                 role_id: pulumi.Input[builtins.str],
                 activation_rules: Optional[pulumi.Input['GroupRoleManagementPolicyActivationRulesArgs']] = None,
                 active_assignment_rules: Optional[pulumi.Input['GroupRoleManagementPolicyActiveAssignmentRulesArgs']] = None,
                 eligible_assignment_rules: Optional[pulumi.Input['GroupRoleManagementPolicyEligibleAssignmentRulesArgs']] = None,
                 notification_rules: Optional[pulumi.Input['GroupRoleManagementPolicyNotificationRulesArgs']] = None):
        """
        The set of arguments for constructing a GroupRoleManagementPolicy resource.
        :param pulumi.Input[builtins.str] group_id: The ID of the Azure AD group for which the policy applies.
        :param pulumi.Input[builtins.str] role_id: The type of assignment this policy coveres. Can be either `member` or `owner`.
        :param pulumi.Input['GroupRoleManagementPolicyActivationRulesArgs'] activation_rules: An `activation_rules` block as defined below.
        :param pulumi.Input['GroupRoleManagementPolicyActiveAssignmentRulesArgs'] active_assignment_rules: An `active_assignment_rules` block as defined below.
        :param pulumi.Input['GroupRoleManagementPolicyEligibleAssignmentRulesArgs'] eligible_assignment_rules: An `eligible_assignment_rules` block as defined below.
        :param pulumi.Input['GroupRoleManagementPolicyNotificationRulesArgs'] notification_rules: A `notification_rules` block as defined below.
        """
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "role_id", role_id)
        if activation_rules is not None:
            pulumi.set(__self__, "activation_rules", activation_rules)
        if active_assignment_rules is not None:
            pulumi.set(__self__, "active_assignment_rules", active_assignment_rules)
        if eligible_assignment_rules is not None:
            pulumi.set(__self__, "eligible_assignment_rules", eligible_assignment_rules)
        if notification_rules is not None:
            pulumi.set(__self__, "notification_rules", notification_rules)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the Azure AD group for which the policy applies.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Input[builtins.str]:
        """
        The type of assignment this policy coveres. Can be either `member` or `owner`.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter(name="activationRules")
    def activation_rules(self) -> Optional[pulumi.Input['GroupRoleManagementPolicyActivationRulesArgs']]:
        """
        An `activation_rules` block as defined below.
        """
        return pulumi.get(self, "activation_rules")

    @activation_rules.setter
    def activation_rules(self, value: Optional[pulumi.Input['GroupRoleManagementPolicyActivationRulesArgs']]):
        pulumi.set(self, "activation_rules", value)

    @property
    @pulumi.getter(name="activeAssignmentRules")
    def active_assignment_rules(self) -> Optional[pulumi.Input['GroupRoleManagementPolicyActiveAssignmentRulesArgs']]:
        """
        An `active_assignment_rules` block as defined below.
        """
        return pulumi.get(self, "active_assignment_rules")

    @active_assignment_rules.setter
    def active_assignment_rules(self, value: Optional[pulumi.Input['GroupRoleManagementPolicyActiveAssignmentRulesArgs']]):
        pulumi.set(self, "active_assignment_rules", value)

    @property
    @pulumi.getter(name="eligibleAssignmentRules")
    def eligible_assignment_rules(self) -> Optional[pulumi.Input['GroupRoleManagementPolicyEligibleAssignmentRulesArgs']]:
        """
        An `eligible_assignment_rules` block as defined below.
        """
        return pulumi.get(self, "eligible_assignment_rules")

    @eligible_assignment_rules.setter
    def eligible_assignment_rules(self, value: Optional[pulumi.Input['GroupRoleManagementPolicyEligibleAssignmentRulesArgs']]):
        pulumi.set(self, "eligible_assignment_rules", value)

    @property
    @pulumi.getter(name="notificationRules")
    def notification_rules(self) -> Optional[pulumi.Input['GroupRoleManagementPolicyNotificationRulesArgs']]:
        """
        A `notification_rules` block as defined below.
        """
        return pulumi.get(self, "notification_rules")

    @notification_rules.setter
    def notification_rules(self, value: Optional[pulumi.Input['GroupRoleManagementPolicyNotificationRulesArgs']]):
        pulumi.set(self, "notification_rules", value)


@pulumi.input_type
class _GroupRoleManagementPolicyState:
    def __init__(__self__, *,
                 activation_rules: Optional[pulumi.Input['GroupRoleManagementPolicyActivationRulesArgs']] = None,
                 active_assignment_rules: Optional[pulumi.Input['GroupRoleManagementPolicyActiveAssignmentRulesArgs']] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 eligible_assignment_rules: Optional[pulumi.Input['GroupRoleManagementPolicyEligibleAssignmentRulesArgs']] = None,
                 group_id: Optional[pulumi.Input[builtins.str]] = None,
                 notification_rules: Optional[pulumi.Input['GroupRoleManagementPolicyNotificationRulesArgs']] = None,
                 role_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering GroupRoleManagementPolicy resources.
        :param pulumi.Input['GroupRoleManagementPolicyActivationRulesArgs'] activation_rules: An `activation_rules` block as defined below.
        :param pulumi.Input['GroupRoleManagementPolicyActiveAssignmentRulesArgs'] active_assignment_rules: An `active_assignment_rules` block as defined below.
        :param pulumi.Input[builtins.str] description: (String) The description of this policy.
        :param pulumi.Input[builtins.str] display_name: (String) The display name of this policy.
        :param pulumi.Input['GroupRoleManagementPolicyEligibleAssignmentRulesArgs'] eligible_assignment_rules: An `eligible_assignment_rules` block as defined below.
        :param pulumi.Input[builtins.str] group_id: The ID of the Azure AD group for which the policy applies.
        :param pulumi.Input['GroupRoleManagementPolicyNotificationRulesArgs'] notification_rules: A `notification_rules` block as defined below.
        :param pulumi.Input[builtins.str] role_id: The type of assignment this policy coveres. Can be either `member` or `owner`.
        """
        if activation_rules is not None:
            pulumi.set(__self__, "activation_rules", activation_rules)
        if active_assignment_rules is not None:
            pulumi.set(__self__, "active_assignment_rules", active_assignment_rules)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if eligible_assignment_rules is not None:
            pulumi.set(__self__, "eligible_assignment_rules", eligible_assignment_rules)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if notification_rules is not None:
            pulumi.set(__self__, "notification_rules", notification_rules)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)

    @property
    @pulumi.getter(name="activationRules")
    def activation_rules(self) -> Optional[pulumi.Input['GroupRoleManagementPolicyActivationRulesArgs']]:
        """
        An `activation_rules` block as defined below.
        """
        return pulumi.get(self, "activation_rules")

    @activation_rules.setter
    def activation_rules(self, value: Optional[pulumi.Input['GroupRoleManagementPolicyActivationRulesArgs']]):
        pulumi.set(self, "activation_rules", value)

    @property
    @pulumi.getter(name="activeAssignmentRules")
    def active_assignment_rules(self) -> Optional[pulumi.Input['GroupRoleManagementPolicyActiveAssignmentRulesArgs']]:
        """
        An `active_assignment_rules` block as defined below.
        """
        return pulumi.get(self, "active_assignment_rules")

    @active_assignment_rules.setter
    def active_assignment_rules(self, value: Optional[pulumi.Input['GroupRoleManagementPolicyActiveAssignmentRulesArgs']]):
        pulumi.set(self, "active_assignment_rules", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        (String) The description of this policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        (String) The display name of this policy.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="eligibleAssignmentRules")
    def eligible_assignment_rules(self) -> Optional[pulumi.Input['GroupRoleManagementPolicyEligibleAssignmentRulesArgs']]:
        """
        An `eligible_assignment_rules` block as defined below.
        """
        return pulumi.get(self, "eligible_assignment_rules")

    @eligible_assignment_rules.setter
    def eligible_assignment_rules(self, value: Optional[pulumi.Input['GroupRoleManagementPolicyEligibleAssignmentRulesArgs']]):
        pulumi.set(self, "eligible_assignment_rules", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the Azure AD group for which the policy applies.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="notificationRules")
    def notification_rules(self) -> Optional[pulumi.Input['GroupRoleManagementPolicyNotificationRulesArgs']]:
        """
        A `notification_rules` block as defined below.
        """
        return pulumi.get(self, "notification_rules")

    @notification_rules.setter
    def notification_rules(self, value: Optional[pulumi.Input['GroupRoleManagementPolicyNotificationRulesArgs']]):
        pulumi.set(self, "notification_rules", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The type of assignment this policy coveres. Can be either `member` or `owner`.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "role_id", value)


class GroupRoleManagementPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activation_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyActivationRulesArgs', 'GroupRoleManagementPolicyActivationRulesArgsDict']]] = None,
                 active_assignment_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyActiveAssignmentRulesArgs', 'GroupRoleManagementPolicyActiveAssignmentRulesArgsDict']]] = None,
                 eligible_assignment_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyEligibleAssignmentRulesArgs', 'GroupRoleManagementPolicyEligibleAssignmentRulesArgsDict']]] = None,
                 group_id: Optional[pulumi.Input[builtins.str]] = None,
                 notification_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyNotificationRulesArgs', 'GroupRoleManagementPolicyNotificationRulesArgsDict']]] = None,
                 role_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manage a role policy for an Azure AD group.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the `RoleManagementPolicy.ReadWrite.AzureADGroup` Microsoft Graph API permissions.

        When authenticated with a user principal, this resource requires `Global Administrator` directory role, or the `Privileged Role Administrator` role in Identity Governance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Group("example",
            display_name="group-name",
            security_enabled=True)
        member = azuread.User("member",
            user_principal_name="jdoe@example.com",
            display_name="J. Doe",
            mail_nickname="jdoe",
            password="SecretP@sswd99!")
        example_group_role_management_policy = azuread.GroupRoleManagementPolicy("example",
            group_id=example.id,
            role_id="member",
            active_assignment_rules={
                "expire_after": "P365D",
            },
            eligible_assignment_rules={
                "expiration_required": False,
            },
            notification_rules={
                "eligible_assignments": {
                    "approver_notifications": {
                        "notification_level": "Critical",
                        "default_recipients": False,
                        "additional_recipients": [
                            "someone@example.com",
                            "someone.else@example.com",
                        ],
                    },
                },
            })
        ```

        ## Import

        Because these policies are created automatically by Entra ID, they will auto-import on first use.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['GroupRoleManagementPolicyActivationRulesArgs', 'GroupRoleManagementPolicyActivationRulesArgsDict']] activation_rules: An `activation_rules` block as defined below.
        :param pulumi.Input[Union['GroupRoleManagementPolicyActiveAssignmentRulesArgs', 'GroupRoleManagementPolicyActiveAssignmentRulesArgsDict']] active_assignment_rules: An `active_assignment_rules` block as defined below.
        :param pulumi.Input[Union['GroupRoleManagementPolicyEligibleAssignmentRulesArgs', 'GroupRoleManagementPolicyEligibleAssignmentRulesArgsDict']] eligible_assignment_rules: An `eligible_assignment_rules` block as defined below.
        :param pulumi.Input[builtins.str] group_id: The ID of the Azure AD group for which the policy applies.
        :param pulumi.Input[Union['GroupRoleManagementPolicyNotificationRulesArgs', 'GroupRoleManagementPolicyNotificationRulesArgsDict']] notification_rules: A `notification_rules` block as defined below.
        :param pulumi.Input[builtins.str] role_id: The type of assignment this policy coveres. Can be either `member` or `owner`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupRoleManagementPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage a role policy for an Azure AD group.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the `RoleManagementPolicy.ReadWrite.AzureADGroup` Microsoft Graph API permissions.

        When authenticated with a user principal, this resource requires `Global Administrator` directory role, or the `Privileged Role Administrator` role in Identity Governance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Group("example",
            display_name="group-name",
            security_enabled=True)
        member = azuread.User("member",
            user_principal_name="jdoe@example.com",
            display_name="J. Doe",
            mail_nickname="jdoe",
            password="SecretP@sswd99!")
        example_group_role_management_policy = azuread.GroupRoleManagementPolicy("example",
            group_id=example.id,
            role_id="member",
            active_assignment_rules={
                "expire_after": "P365D",
            },
            eligible_assignment_rules={
                "expiration_required": False,
            },
            notification_rules={
                "eligible_assignments": {
                    "approver_notifications": {
                        "notification_level": "Critical",
                        "default_recipients": False,
                        "additional_recipients": [
                            "someone@example.com",
                            "someone.else@example.com",
                        ],
                    },
                },
            })
        ```

        ## Import

        Because these policies are created automatically by Entra ID, they will auto-import on first use.

        :param str resource_name: The name of the resource.
        :param GroupRoleManagementPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupRoleManagementPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activation_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyActivationRulesArgs', 'GroupRoleManagementPolicyActivationRulesArgsDict']]] = None,
                 active_assignment_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyActiveAssignmentRulesArgs', 'GroupRoleManagementPolicyActiveAssignmentRulesArgsDict']]] = None,
                 eligible_assignment_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyEligibleAssignmentRulesArgs', 'GroupRoleManagementPolicyEligibleAssignmentRulesArgsDict']]] = None,
                 group_id: Optional[pulumi.Input[builtins.str]] = None,
                 notification_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyNotificationRulesArgs', 'GroupRoleManagementPolicyNotificationRulesArgsDict']]] = None,
                 role_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupRoleManagementPolicyArgs.__new__(GroupRoleManagementPolicyArgs)

            __props__.__dict__["activation_rules"] = activation_rules
            __props__.__dict__["active_assignment_rules"] = active_assignment_rules
            __props__.__dict__["eligible_assignment_rules"] = eligible_assignment_rules
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["notification_rules"] = notification_rules
            if role_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_id'")
            __props__.__dict__["role_id"] = role_id
            __props__.__dict__["description"] = None
            __props__.__dict__["display_name"] = None
        super(GroupRoleManagementPolicy, __self__).__init__(
            'azuread:index/groupRoleManagementPolicy:GroupRoleManagementPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            activation_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyActivationRulesArgs', 'GroupRoleManagementPolicyActivationRulesArgsDict']]] = None,
            active_assignment_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyActiveAssignmentRulesArgs', 'GroupRoleManagementPolicyActiveAssignmentRulesArgsDict']]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            display_name: Optional[pulumi.Input[builtins.str]] = None,
            eligible_assignment_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyEligibleAssignmentRulesArgs', 'GroupRoleManagementPolicyEligibleAssignmentRulesArgsDict']]] = None,
            group_id: Optional[pulumi.Input[builtins.str]] = None,
            notification_rules: Optional[pulumi.Input[Union['GroupRoleManagementPolicyNotificationRulesArgs', 'GroupRoleManagementPolicyNotificationRulesArgsDict']]] = None,
            role_id: Optional[pulumi.Input[builtins.str]] = None) -> 'GroupRoleManagementPolicy':
        """
        Get an existing GroupRoleManagementPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['GroupRoleManagementPolicyActivationRulesArgs', 'GroupRoleManagementPolicyActivationRulesArgsDict']] activation_rules: An `activation_rules` block as defined below.
        :param pulumi.Input[Union['GroupRoleManagementPolicyActiveAssignmentRulesArgs', 'GroupRoleManagementPolicyActiveAssignmentRulesArgsDict']] active_assignment_rules: An `active_assignment_rules` block as defined below.
        :param pulumi.Input[builtins.str] description: (String) The description of this policy.
        :param pulumi.Input[builtins.str] display_name: (String) The display name of this policy.
        :param pulumi.Input[Union['GroupRoleManagementPolicyEligibleAssignmentRulesArgs', 'GroupRoleManagementPolicyEligibleAssignmentRulesArgsDict']] eligible_assignment_rules: An `eligible_assignment_rules` block as defined below.
        :param pulumi.Input[builtins.str] group_id: The ID of the Azure AD group for which the policy applies.
        :param pulumi.Input[Union['GroupRoleManagementPolicyNotificationRulesArgs', 'GroupRoleManagementPolicyNotificationRulesArgsDict']] notification_rules: A `notification_rules` block as defined below.
        :param pulumi.Input[builtins.str] role_id: The type of assignment this policy coveres. Can be either `member` or `owner`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupRoleManagementPolicyState.__new__(_GroupRoleManagementPolicyState)

        __props__.__dict__["activation_rules"] = activation_rules
        __props__.__dict__["active_assignment_rules"] = active_assignment_rules
        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["eligible_assignment_rules"] = eligible_assignment_rules
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["notification_rules"] = notification_rules
        __props__.__dict__["role_id"] = role_id
        return GroupRoleManagementPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activationRules")
    def activation_rules(self) -> pulumi.Output['outputs.GroupRoleManagementPolicyActivationRules']:
        """
        An `activation_rules` block as defined below.
        """
        return pulumi.get(self, "activation_rules")

    @property
    @pulumi.getter(name="activeAssignmentRules")
    def active_assignment_rules(self) -> pulumi.Output['outputs.GroupRoleManagementPolicyActiveAssignmentRules']:
        """
        An `active_assignment_rules` block as defined below.
        """
        return pulumi.get(self, "active_assignment_rules")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[builtins.str]:
        """
        (String) The description of this policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[builtins.str]:
        """
        (String) The display name of this policy.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="eligibleAssignmentRules")
    def eligible_assignment_rules(self) -> pulumi.Output['outputs.GroupRoleManagementPolicyEligibleAssignmentRules']:
        """
        An `eligible_assignment_rules` block as defined below.
        """
        return pulumi.get(self, "eligible_assignment_rules")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the Azure AD group for which the policy applies.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="notificationRules")
    def notification_rules(self) -> pulumi.Output['outputs.GroupRoleManagementPolicyNotificationRules']:
        """
        A `notification_rules` block as defined below.
        """
        return pulumi.get(self, "notification_rules")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[builtins.str]:
        """
        The type of assignment this policy coveres. Can be either `member` or `owner`.
        """
        return pulumi.get(self, "role_id")

