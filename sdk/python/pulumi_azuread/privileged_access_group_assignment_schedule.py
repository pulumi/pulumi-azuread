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

__all__ = ['PrivilegedAccessGroupAssignmentScheduleArgs', 'PrivilegedAccessGroupAssignmentSchedule']

@pulumi.input_type
class PrivilegedAccessGroupAssignmentScheduleArgs:
    def __init__(__self__, *,
                 assignment_type: pulumi.Input[str],
                 group_id: pulumi.Input[str],
                 principal_id: pulumi.Input[str],
                 duration: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 justification: Optional[pulumi.Input[str]] = None,
                 permanent_assignment: Optional[pulumi.Input[bool]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 ticket_number: Optional[pulumi.Input[str]] = None,
                 ticket_system: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrivilegedAccessGroupAssignmentSchedule resource.
        :param pulumi.Input[str] assignment_type: The type of assignment to the group. Can be either `member` or `owner`.
        :param pulumi.Input[str] group_id: The Object ID of the Azure AD group to which the principal will be assigned.
        :param pulumi.Input[str] principal_id: The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
        :param pulumi.Input[str] duration: The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
        :param pulumi.Input[str] expiration_date: The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
        :param pulumi.Input[str] justification: The justification for this assignment. May be required by the role policy.
        :param pulumi.Input[bool] permanent_assignment: Is this assigment permanently valid.
               
               At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
        :param pulumi.Input[str] start_date: The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
        :param pulumi.Input[str] ticket_number: The ticket number in the ticket system approving this assignment. May be required by the role policy.
        :param pulumi.Input[str] ticket_system: The ticket system containing the ticket number approving this assignment. May be required by the role policy.
        """
        pulumi.set(__self__, "assignment_type", assignment_type)
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "principal_id", principal_id)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if justification is not None:
            pulumi.set(__self__, "justification", justification)
        if permanent_assignment is not None:
            pulumi.set(__self__, "permanent_assignment", permanent_assignment)
        if start_date is not None:
            pulumi.set(__self__, "start_date", start_date)
        if ticket_number is not None:
            pulumi.set(__self__, "ticket_number", ticket_number)
        if ticket_system is not None:
            pulumi.set(__self__, "ticket_system", ticket_system)

    @property
    @pulumi.getter(name="assignmentType")
    def assignment_type(self) -> pulumi.Input[str]:
        """
        The type of assignment to the group. Can be either `member` or `owner`.
        """
        return pulumi.get(self, "assignment_type")

    @assignment_type.setter
    def assignment_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "assignment_type", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        The Object ID of the Azure AD group to which the principal will be assigned.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Input[str]:
        """
        The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[str]]:
        """
        The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        """
        The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter
    def justification(self) -> Optional[pulumi.Input[str]]:
        """
        The justification for this assignment. May be required by the role policy.
        """
        return pulumi.get(self, "justification")

    @justification.setter
    def justification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "justification", value)

    @property
    @pulumi.getter(name="permanentAssignment")
    def permanent_assignment(self) -> Optional[pulumi.Input[bool]]:
        """
        Is this assigment permanently valid.

        At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
        """
        return pulumi.get(self, "permanent_assignment")

    @permanent_assignment.setter
    def permanent_assignment(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "permanent_assignment", value)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[str]]:
        """
        The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter(name="ticketNumber")
    def ticket_number(self) -> Optional[pulumi.Input[str]]:
        """
        The ticket number in the ticket system approving this assignment. May be required by the role policy.
        """
        return pulumi.get(self, "ticket_number")

    @ticket_number.setter
    def ticket_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ticket_number", value)

    @property
    @pulumi.getter(name="ticketSystem")
    def ticket_system(self) -> Optional[pulumi.Input[str]]:
        """
        The ticket system containing the ticket number approving this assignment. May be required by the role policy.
        """
        return pulumi.get(self, "ticket_system")

    @ticket_system.setter
    def ticket_system(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ticket_system", value)


@pulumi.input_type
class _PrivilegedAccessGroupAssignmentScheduleState:
    def __init__(__self__, *,
                 assignment_type: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 justification: Optional[pulumi.Input[str]] = None,
                 permanent_assignment: Optional[pulumi.Input[bool]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 ticket_number: Optional[pulumi.Input[str]] = None,
                 ticket_system: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PrivilegedAccessGroupAssignmentSchedule resources.
        :param pulumi.Input[str] assignment_type: The type of assignment to the group. Can be either `member` or `owner`.
        :param pulumi.Input[str] duration: The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
        :param pulumi.Input[str] expiration_date: The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
        :param pulumi.Input[str] group_id: The Object ID of the Azure AD group to which the principal will be assigned.
        :param pulumi.Input[str] justification: The justification for this assignment. May be required by the role policy.
        :param pulumi.Input[bool] permanent_assignment: Is this assigment permanently valid.
               
               At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
        :param pulumi.Input[str] principal_id: The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
        :param pulumi.Input[str] start_date: The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
        :param pulumi.Input[str] status: (String) The provisioning status of this request.
        :param pulumi.Input[str] ticket_number: The ticket number in the ticket system approving this assignment. May be required by the role policy.
        :param pulumi.Input[str] ticket_system: The ticket system containing the ticket number approving this assignment. May be required by the role policy.
        """
        if assignment_type is not None:
            pulumi.set(__self__, "assignment_type", assignment_type)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if justification is not None:
            pulumi.set(__self__, "justification", justification)
        if permanent_assignment is not None:
            pulumi.set(__self__, "permanent_assignment", permanent_assignment)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if start_date is not None:
            pulumi.set(__self__, "start_date", start_date)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if ticket_number is not None:
            pulumi.set(__self__, "ticket_number", ticket_number)
        if ticket_system is not None:
            pulumi.set(__self__, "ticket_system", ticket_system)

    @property
    @pulumi.getter(name="assignmentType")
    def assignment_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of assignment to the group. Can be either `member` or `owner`.
        """
        return pulumi.get(self, "assignment_type")

    @assignment_type.setter
    def assignment_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "assignment_type", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[str]]:
        """
        The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[str]]:
        """
        The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Object ID of the Azure AD group to which the principal will be assigned.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def justification(self) -> Optional[pulumi.Input[str]]:
        """
        The justification for this assignment. May be required by the role policy.
        """
        return pulumi.get(self, "justification")

    @justification.setter
    def justification(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "justification", value)

    @property
    @pulumi.getter(name="permanentAssignment")
    def permanent_assignment(self) -> Optional[pulumi.Input[bool]]:
        """
        Is this assigment permanently valid.

        At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
        """
        return pulumi.get(self, "permanent_assignment")

    @permanent_assignment.setter
    def permanent_assignment(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "permanent_assignment", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[str]]:
        """
        The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        (String) The provisioning status of this request.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="ticketNumber")
    def ticket_number(self) -> Optional[pulumi.Input[str]]:
        """
        The ticket number in the ticket system approving this assignment. May be required by the role policy.
        """
        return pulumi.get(self, "ticket_number")

    @ticket_number.setter
    def ticket_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ticket_number", value)

    @property
    @pulumi.getter(name="ticketSystem")
    def ticket_system(self) -> Optional[pulumi.Input[str]]:
        """
        The ticket system containing the ticket number approving this assignment. May be required by the role policy.
        """
        return pulumi.get(self, "ticket_system")

    @ticket_system.setter
    def ticket_system(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ticket_system", value)


class PrivilegedAccessGroupAssignmentSchedule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assignment_type: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 justification: Optional[pulumi.Input[str]] = None,
                 permanent_assignment: Optional[pulumi.Input[bool]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 ticket_number: Optional[pulumi.Input[str]] = None,
                 ticket_system: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an active assignment to a privileged access group.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the `PrivilegedAssignmentSchedule.ReadWrite.AzureADGroup` Microsoft Graph API permissions.

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
        example_privileged_access_group_assignment_schedule = azuread.PrivilegedAccessGroupAssignmentSchedule("example",
            group_id=pim["id"],
            principal_id=member.id,
            assignment_type="member",
            duration="P30D",
            justification="as requested")
        ```

        ## Import

        An assignment schedule can be imported using the schedule ID, e.g.

        ```sh
        $ pulumi import azuread:index/privilegedAccessGroupAssignmentSchedule:PrivilegedAccessGroupAssignmentSchedule example 00000000-0000-0000-0000-000000000000_member_00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] assignment_type: The type of assignment to the group. Can be either `member` or `owner`.
        :param pulumi.Input[str] duration: The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
        :param pulumi.Input[str] expiration_date: The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
        :param pulumi.Input[str] group_id: The Object ID of the Azure AD group to which the principal will be assigned.
        :param pulumi.Input[str] justification: The justification for this assignment. May be required by the role policy.
        :param pulumi.Input[bool] permanent_assignment: Is this assigment permanently valid.
               
               At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
        :param pulumi.Input[str] principal_id: The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
        :param pulumi.Input[str] start_date: The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
        :param pulumi.Input[str] ticket_number: The ticket number in the ticket system approving this assignment. May be required by the role policy.
        :param pulumi.Input[str] ticket_system: The ticket system containing the ticket number approving this assignment. May be required by the role policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivilegedAccessGroupAssignmentScheduleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an active assignment to a privileged access group.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the `PrivilegedAssignmentSchedule.ReadWrite.AzureADGroup` Microsoft Graph API permissions.

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
        example_privileged_access_group_assignment_schedule = azuread.PrivilegedAccessGroupAssignmentSchedule("example",
            group_id=pim["id"],
            principal_id=member.id,
            assignment_type="member",
            duration="P30D",
            justification="as requested")
        ```

        ## Import

        An assignment schedule can be imported using the schedule ID, e.g.

        ```sh
        $ pulumi import azuread:index/privilegedAccessGroupAssignmentSchedule:PrivilegedAccessGroupAssignmentSchedule example 00000000-0000-0000-0000-000000000000_member_00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param PrivilegedAccessGroupAssignmentScheduleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivilegedAccessGroupAssignmentScheduleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assignment_type: Optional[pulumi.Input[str]] = None,
                 duration: Optional[pulumi.Input[str]] = None,
                 expiration_date: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 justification: Optional[pulumi.Input[str]] = None,
                 permanent_assignment: Optional[pulumi.Input[bool]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 ticket_number: Optional[pulumi.Input[str]] = None,
                 ticket_system: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivilegedAccessGroupAssignmentScheduleArgs.__new__(PrivilegedAccessGroupAssignmentScheduleArgs)

            if assignment_type is None and not opts.urn:
                raise TypeError("Missing required property 'assignment_type'")
            __props__.__dict__["assignment_type"] = assignment_type
            __props__.__dict__["duration"] = duration
            __props__.__dict__["expiration_date"] = expiration_date
            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["justification"] = justification
            __props__.__dict__["permanent_assignment"] = permanent_assignment
            if principal_id is None and not opts.urn:
                raise TypeError("Missing required property 'principal_id'")
            __props__.__dict__["principal_id"] = principal_id
            __props__.__dict__["start_date"] = start_date
            __props__.__dict__["ticket_number"] = ticket_number
            __props__.__dict__["ticket_system"] = ticket_system
            __props__.__dict__["status"] = None
        super(PrivilegedAccessGroupAssignmentSchedule, __self__).__init__(
            'azuread:index/privilegedAccessGroupAssignmentSchedule:PrivilegedAccessGroupAssignmentSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            assignment_type: Optional[pulumi.Input[str]] = None,
            duration: Optional[pulumi.Input[str]] = None,
            expiration_date: Optional[pulumi.Input[str]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            justification: Optional[pulumi.Input[str]] = None,
            permanent_assignment: Optional[pulumi.Input[bool]] = None,
            principal_id: Optional[pulumi.Input[str]] = None,
            start_date: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            ticket_number: Optional[pulumi.Input[str]] = None,
            ticket_system: Optional[pulumi.Input[str]] = None) -> 'PrivilegedAccessGroupAssignmentSchedule':
        """
        Get an existing PrivilegedAccessGroupAssignmentSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] assignment_type: The type of assignment to the group. Can be either `member` or `owner`.
        :param pulumi.Input[str] duration: The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
        :param pulumi.Input[str] expiration_date: The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
        :param pulumi.Input[str] group_id: The Object ID of the Azure AD group to which the principal will be assigned.
        :param pulumi.Input[str] justification: The justification for this assignment. May be required by the role policy.
        :param pulumi.Input[bool] permanent_assignment: Is this assigment permanently valid.
               
               At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
        :param pulumi.Input[str] principal_id: The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
        :param pulumi.Input[str] start_date: The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
        :param pulumi.Input[str] status: (String) The provisioning status of this request.
        :param pulumi.Input[str] ticket_number: The ticket number in the ticket system approving this assignment. May be required by the role policy.
        :param pulumi.Input[str] ticket_system: The ticket system containing the ticket number approving this assignment. May be required by the role policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivilegedAccessGroupAssignmentScheduleState.__new__(_PrivilegedAccessGroupAssignmentScheduleState)

        __props__.__dict__["assignment_type"] = assignment_type
        __props__.__dict__["duration"] = duration
        __props__.__dict__["expiration_date"] = expiration_date
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["justification"] = justification
        __props__.__dict__["permanent_assignment"] = permanent_assignment
        __props__.__dict__["principal_id"] = principal_id
        __props__.__dict__["start_date"] = start_date
        __props__.__dict__["status"] = status
        __props__.__dict__["ticket_number"] = ticket_number
        __props__.__dict__["ticket_system"] = ticket_system
        return PrivilegedAccessGroupAssignmentSchedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="assignmentType")
    def assignment_type(self) -> pulumi.Output[str]:
        """
        The type of assignment to the group. Can be either `member` or `owner`.
        """
        return pulumi.get(self, "assignment_type")

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Output[Optional[str]]:
        """
        The duration that this assignment is valid for, formatted as an ISO8601 duration (e.g. P30D for 30 days, PT3H for three hours).
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> pulumi.Output[str]:
        """
        The date that this assignment expires, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z).
        """
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        The Object ID of the Azure AD group to which the principal will be assigned.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def justification(self) -> pulumi.Output[Optional[str]]:
        """
        The justification for this assignment. May be required by the role policy.
        """
        return pulumi.get(self, "justification")

    @property
    @pulumi.getter(name="permanentAssignment")
    def permanent_assignment(self) -> pulumi.Output[bool]:
        """
        Is this assigment permanently valid.

        At least one of `expiration_date`, `duration`, or `permanent_assignment` must be supplied. The role policy may limit the maximum duration which can be supplied.
        """
        return pulumi.get(self, "permanent_assignment")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Output[str]:
        """
        The Object ID of the principal to be assigned to the above group. Can be either a user or a group.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> pulumi.Output[str]:
        """
        The date from which this assignment is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If not provided, the assignment is immediately valid.
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        (String) The provisioning status of this request.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="ticketNumber")
    def ticket_number(self) -> pulumi.Output[Optional[str]]:
        """
        The ticket number in the ticket system approving this assignment. May be required by the role policy.
        """
        return pulumi.get(self, "ticket_number")

    @property
    @pulumi.getter(name="ticketSystem")
    def ticket_system(self) -> pulumi.Output[Optional[str]]:
        """
        The ticket system containing the ticket number approving this assignment. May be required by the role policy.
        """
        return pulumi.get(self, "ticket_system")

