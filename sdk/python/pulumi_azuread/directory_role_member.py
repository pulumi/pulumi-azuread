# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DirectoryRoleMemberArgs', 'DirectoryRoleMember']

@pulumi.input_type
class DirectoryRoleMemberArgs:
    def __init__(__self__, *,
                 member_object_id: Optional[pulumi.Input[str]] = None,
                 role_object_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DirectoryRoleMember resource.
        :param pulumi.Input[str] member_object_id: The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_object_id: The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
        """
        DirectoryRoleMemberArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            member_object_id=member_object_id,
            role_object_id=role_object_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             member_object_id: Optional[pulumi.Input[str]] = None,
             role_object_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if member_object_id is None and 'memberObjectId' in kwargs:
            member_object_id = kwargs['memberObjectId']
        if role_object_id is None and 'roleObjectId' in kwargs:
            role_object_id = kwargs['roleObjectId']

        if member_object_id is not None:
            _setter("member_object_id", member_object_id)
        if role_object_id is not None:
            _setter("role_object_id", role_object_id)

    @property
    @pulumi.getter(name="memberObjectId")
    def member_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "member_object_id")

    @member_object_id.setter
    def member_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member_object_id", value)

    @property
    @pulumi.getter(name="roleObjectId")
    def role_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_object_id")

    @role_object_id.setter
    def role_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_object_id", value)


@pulumi.input_type
class _DirectoryRoleMemberState:
    def __init__(__self__, *,
                 member_object_id: Optional[pulumi.Input[str]] = None,
                 role_object_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DirectoryRoleMember resources.
        :param pulumi.Input[str] member_object_id: The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_object_id: The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
        """
        _DirectoryRoleMemberState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            member_object_id=member_object_id,
            role_object_id=role_object_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             member_object_id: Optional[pulumi.Input[str]] = None,
             role_object_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if member_object_id is None and 'memberObjectId' in kwargs:
            member_object_id = kwargs['memberObjectId']
        if role_object_id is None and 'roleObjectId' in kwargs:
            role_object_id = kwargs['roleObjectId']

        if member_object_id is not None:
            _setter("member_object_id", member_object_id)
        if role_object_id is not None:
            _setter("role_object_id", role_object_id)

    @property
    @pulumi.getter(name="memberObjectId")
    def member_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "member_object_id")

    @member_object_id.setter
    def member_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member_object_id", value)

    @property
    @pulumi.getter(name="roleObjectId")
    def role_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_object_id")

    @role_object_id.setter
    def role_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_object_id", value)


class DirectoryRoleMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 member_object_id: Optional[pulumi.Input[str]] = None,
                 role_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a single directory role membership (assignment) within Azure Active Directory.

        > **Deprecation Warning:** This resource has been superseded by the DirectoryRoleAssignment resource and will be removed in version 3.0 of the AzureAD provider

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_user = azuread.get_user(user_principal_name="jdoe@hashicorp.com")
        example_directory_role = azuread.DirectoryRole("exampleDirectoryRole", display_name="Security administrator")
        example_directory_role_member = azuread.DirectoryRoleMember("exampleDirectoryRoleMember",
            role_object_id=example_directory_role.object_id,
            member_object_id=example_user.object_id)
        ```

        ## Import

        Directory role members can be imported using the object ID of the role and the object ID of the member, e.g.

        ```sh
         $ pulumi import azuread:index/directoryRoleMember:DirectoryRoleMember test 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the Directory Role Object ID and the target Member Object ID in the format `{RoleObjectID}/member/{MemberObjectID}`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] member_object_id: The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_object_id: The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DirectoryRoleMemberArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a single directory role membership (assignment) within Azure Active Directory.

        > **Deprecation Warning:** This resource has been superseded by the DirectoryRoleAssignment resource and will be removed in version 3.0 of the AzureAD provider

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_user = azuread.get_user(user_principal_name="jdoe@hashicorp.com")
        example_directory_role = azuread.DirectoryRole("exampleDirectoryRole", display_name="Security administrator")
        example_directory_role_member = azuread.DirectoryRoleMember("exampleDirectoryRoleMember",
            role_object_id=example_directory_role.object_id,
            member_object_id=example_user.object_id)
        ```

        ## Import

        Directory role members can be imported using the object ID of the role and the object ID of the member, e.g.

        ```sh
         $ pulumi import azuread:index/directoryRoleMember:DirectoryRoleMember test 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the Directory Role Object ID and the target Member Object ID in the format `{RoleObjectID}/member/{MemberObjectID}`.

        :param str resource_name: The name of the resource.
        :param DirectoryRoleMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DirectoryRoleMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            DirectoryRoleMemberArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 member_object_id: Optional[pulumi.Input[str]] = None,
                 role_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DirectoryRoleMemberArgs.__new__(DirectoryRoleMemberArgs)

            __props__.__dict__["member_object_id"] = member_object_id
            __props__.__dict__["role_object_id"] = role_object_id
        super(DirectoryRoleMember, __self__).__init__(
            'azuread:index/directoryRoleMember:DirectoryRoleMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            member_object_id: Optional[pulumi.Input[str]] = None,
            role_object_id: Optional[pulumi.Input[str]] = None) -> 'DirectoryRoleMember':
        """
        Get an existing DirectoryRoleMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] member_object_id: The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_object_id: The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DirectoryRoleMemberState.__new__(_DirectoryRoleMemberState)

        __props__.__dict__["member_object_id"] = member_object_id
        __props__.__dict__["role_object_id"] = role_object_id
        return DirectoryRoleMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="memberObjectId")
    def member_object_id(self) -> pulumi.Output[Optional[str]]:
        """
        The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "member_object_id")

    @property
    @pulumi.getter(name="roleObjectId")
    def role_object_id(self) -> pulumi.Output[Optional[str]]:
        """
        The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_object_id")

