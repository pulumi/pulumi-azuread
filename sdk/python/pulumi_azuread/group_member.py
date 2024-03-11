# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GroupMemberArgs', 'GroupMember']

@pulumi.input_type
class GroupMemberArgs:
    def __init__(__self__, *,
                 group_object_id: pulumi.Input[str],
                 member_object_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a GroupMember resource.
        :param pulumi.Input[str] group_object_id: The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] member_object_id: The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "group_object_id", group_object_id)
        pulumi.set(__self__, "member_object_id", member_object_id)

    @property
    @pulumi.getter(name="groupObjectId")
    def group_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "group_object_id")

    @group_object_id.setter
    def group_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_object_id", value)

    @property
    @pulumi.getter(name="memberObjectId")
    def member_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "member_object_id")

    @member_object_id.setter
    def member_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "member_object_id", value)


@pulumi.input_type
class _GroupMemberState:
    def __init__(__self__, *,
                 group_object_id: Optional[pulumi.Input[str]] = None,
                 member_object_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupMember resources.
        :param pulumi.Input[str] group_object_id: The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] member_object_id: The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        if group_object_id is not None:
            pulumi.set(__self__, "group_object_id", group_object_id)
        if member_object_id is not None:
            pulumi.set(__self__, "member_object_id", member_object_id)

    @property
    @pulumi.getter(name="groupObjectId")
    def group_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "group_object_id")

    @group_object_id.setter
    def group_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_object_id", value)

    @property
    @pulumi.getter(name="memberObjectId")
    def member_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "member_object_id")

    @member_object_id.setter
    def member_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member_object_id", value)


class GroupMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_object_id: Optional[pulumi.Input[str]] = None,
                 member_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a single group membership within Azure Active Directory.

        > **Warning** Do not use this resource at the same time as the `members` property of the `Group` resource for the same group. Doing so will cause a conflict and group members will be removed.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Group.ReadWrite.All` or `Directory.ReadWrite.All`.

        However, if the authenticated service principal is an owner of the group being managed, an application role is not required.

        When authenticated with a user principal, this resource requires one of the following directory roles: `Groups Administrator`, `User Administrator` or `Global Administrator`

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.get_user(user_principal_name="jdoe@example.com")
        example_group = azuread.Group("example",
            display_name="my_group",
            security_enabled=True)
        example_group_member = azuread.GroupMember("example",
            group_object_id=example_group.id,
            member_object_id=example.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Group members can be imported using the object ID of the group and the object ID of the member, e.g.

        ```sh
        $ pulumi import azuread:index/groupMember:GroupMember example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
        ```

        -> This ID format is unique to Terraform and is composed of the Azure AD Group Object ID and the target Member Object ID in the format `{GroupObjectID}/member/{MemberObjectID}`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_object_id: The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] member_object_id: The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a single group membership within Azure Active Directory.

        > **Warning** Do not use this resource at the same time as the `members` property of the `Group` resource for the same group. Doing so will cause a conflict and group members will be removed.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Group.ReadWrite.All` or `Directory.ReadWrite.All`.

        However, if the authenticated service principal is an owner of the group being managed, an application role is not required.

        When authenticated with a user principal, this resource requires one of the following directory roles: `Groups Administrator`, `User Administrator` or `Global Administrator`

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.get_user(user_principal_name="jdoe@example.com")
        example_group = azuread.Group("example",
            display_name="my_group",
            security_enabled=True)
        example_group_member = azuread.GroupMember("example",
            group_object_id=example_group.id,
            member_object_id=example.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Group members can be imported using the object ID of the group and the object ID of the member, e.g.

        ```sh
        $ pulumi import azuread:index/groupMember:GroupMember example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
        ```

        -> This ID format is unique to Terraform and is composed of the Azure AD Group Object ID and the target Member Object ID in the format `{GroupObjectID}/member/{MemberObjectID}`.

        :param str resource_name: The name of the resource.
        :param GroupMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_object_id: Optional[pulumi.Input[str]] = None,
                 member_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupMemberArgs.__new__(GroupMemberArgs)

            if group_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_object_id'")
            __props__.__dict__["group_object_id"] = group_object_id
            if member_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'member_object_id'")
            __props__.__dict__["member_object_id"] = member_object_id
        super(GroupMember, __self__).__init__(
            'azuread:index/groupMember:GroupMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_object_id: Optional[pulumi.Input[str]] = None,
            member_object_id: Optional[pulumi.Input[str]] = None) -> 'GroupMember':
        """
        Get an existing GroupMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_object_id: The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] member_object_id: The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupMemberState.__new__(_GroupMemberState)

        __props__.__dict__["group_object_id"] = group_object_id
        __props__.__dict__["member_object_id"] = member_object_id
        return GroupMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupObjectId")
    def group_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "group_object_id")

    @property
    @pulumi.getter(name="memberObjectId")
    def member_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "member_object_id")

