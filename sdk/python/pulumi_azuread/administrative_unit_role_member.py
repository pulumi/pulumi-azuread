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

__all__ = ['AdministrativeUnitRoleMemberArgs', 'AdministrativeUnitRoleMember']

@pulumi.input_type
class AdministrativeUnitRoleMemberArgs:
    def __init__(__self__, *,
                 administrative_unit_object_id: pulumi.Input[str],
                 member_object_id: pulumi.Input[str],
                 role_object_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a AdministrativeUnitRoleMember resource.
        :param pulumi.Input[str] administrative_unit_object_id: The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] member_object_id: The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_object_id: The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "administrative_unit_object_id", administrative_unit_object_id)
        pulumi.set(__self__, "member_object_id", member_object_id)
        pulumi.set(__self__, "role_object_id", role_object_id)

    @property
    @pulumi.getter(name="administrativeUnitObjectId")
    def administrative_unit_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "administrative_unit_object_id")

    @administrative_unit_object_id.setter
    def administrative_unit_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "administrative_unit_object_id", value)

    @property
    @pulumi.getter(name="memberObjectId")
    def member_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "member_object_id")

    @member_object_id.setter
    def member_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "member_object_id", value)

    @property
    @pulumi.getter(name="roleObjectId")
    def role_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_object_id")

    @role_object_id.setter
    def role_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_object_id", value)


@pulumi.input_type
class _AdministrativeUnitRoleMemberState:
    def __init__(__self__, *,
                 administrative_unit_object_id: Optional[pulumi.Input[str]] = None,
                 member_object_id: Optional[pulumi.Input[str]] = None,
                 role_object_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AdministrativeUnitRoleMember resources.
        :param pulumi.Input[str] administrative_unit_object_id: The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] member_object_id: The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_object_id: The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        if administrative_unit_object_id is not None:
            pulumi.set(__self__, "administrative_unit_object_id", administrative_unit_object_id)
        if member_object_id is not None:
            pulumi.set(__self__, "member_object_id", member_object_id)
        if role_object_id is not None:
            pulumi.set(__self__, "role_object_id", role_object_id)

    @property
    @pulumi.getter(name="administrativeUnitObjectId")
    def administrative_unit_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "administrative_unit_object_id")

    @administrative_unit_object_id.setter
    def administrative_unit_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "administrative_unit_object_id", value)

    @property
    @pulumi.getter(name="memberObjectId")
    def member_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "member_object_id")

    @member_object_id.setter
    def member_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member_object_id", value)

    @property
    @pulumi.getter(name="roleObjectId")
    def role_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_object_id")

    @role_object_id.setter
    def role_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_object_id", value)


class AdministrativeUnitRoleMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 administrative_unit_object_id: Optional[pulumi.Input[str]] = None,
                 member_object_id: Optional[pulumi.Input[str]] = None,
                 role_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a single directory role assignment scoped to an administrative unit within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` and `RoleManagement.ReadWrite.Directory`, or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.get_user(user_principal_name="jdoe@example.com")
        example_administrative_unit = azuread.AdministrativeUnit("example", display_name="Example-AU")
        example_directory_role = azuread.DirectoryRole("example", display_name="Security administrator")
        example_administrative_unit_role_member = azuread.AdministrativeUnitRoleMember("example",
            role_object_id=example_directory_role.object_id,
            administrative_unit_object_id=example_administrative_unit.id,
            member_object_id=example.id)
        ```

        ## Import

        Administrative unit role members can be imported using the object ID of the administrative unit and the unique ID of the role assignment, e.g.

        ```sh
        $ pulumi import azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember example 
        ```

        /directory/administrativeUnits/00000000-0000-0000-0000-000000000000/scopedRoleMembers/zX37MRLyF0uvE-xf2WH4B7x-6CPLfudNnxFGj800htpBXqkxW7bITqGb6Rj4kuTuS

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] administrative_unit_object_id: The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] member_object_id: The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_object_id: The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AdministrativeUnitRoleMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a single directory role assignment scoped to an administrative unit within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` and `RoleManagement.ReadWrite.Directory`, or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.get_user(user_principal_name="jdoe@example.com")
        example_administrative_unit = azuread.AdministrativeUnit("example", display_name="Example-AU")
        example_directory_role = azuread.DirectoryRole("example", display_name="Security administrator")
        example_administrative_unit_role_member = azuread.AdministrativeUnitRoleMember("example",
            role_object_id=example_directory_role.object_id,
            administrative_unit_object_id=example_administrative_unit.id,
            member_object_id=example.id)
        ```

        ## Import

        Administrative unit role members can be imported using the object ID of the administrative unit and the unique ID of the role assignment, e.g.

        ```sh
        $ pulumi import azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember example 
        ```

        /directory/administrativeUnits/00000000-0000-0000-0000-000000000000/scopedRoleMembers/zX37MRLyF0uvE-xf2WH4B7x-6CPLfudNnxFGj800htpBXqkxW7bITqGb6Rj4kuTuS

        :param str resource_name: The name of the resource.
        :param AdministrativeUnitRoleMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AdministrativeUnitRoleMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 administrative_unit_object_id: Optional[pulumi.Input[str]] = None,
                 member_object_id: Optional[pulumi.Input[str]] = None,
                 role_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AdministrativeUnitRoleMemberArgs.__new__(AdministrativeUnitRoleMemberArgs)

            if administrative_unit_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'administrative_unit_object_id'")
            __props__.__dict__["administrative_unit_object_id"] = administrative_unit_object_id
            if member_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'member_object_id'")
            __props__.__dict__["member_object_id"] = member_object_id
            if role_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_object_id'")
            __props__.__dict__["role_object_id"] = role_object_id
        super(AdministrativeUnitRoleMember, __self__).__init__(
            'azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            administrative_unit_object_id: Optional[pulumi.Input[str]] = None,
            member_object_id: Optional[pulumi.Input[str]] = None,
            role_object_id: Optional[pulumi.Input[str]] = None) -> 'AdministrativeUnitRoleMember':
        """
        Get an existing AdministrativeUnitRoleMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] administrative_unit_object_id: The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] member_object_id: The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_object_id: The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AdministrativeUnitRoleMemberState.__new__(_AdministrativeUnitRoleMemberState)

        __props__.__dict__["administrative_unit_object_id"] = administrative_unit_object_id
        __props__.__dict__["member_object_id"] = member_object_id
        __props__.__dict__["role_object_id"] = role_object_id
        return AdministrativeUnitRoleMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="administrativeUnitObjectId")
    def administrative_unit_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "administrative_unit_object_id")

    @property
    @pulumi.getter(name="memberObjectId")
    def member_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "member_object_id")

    @property
    @pulumi.getter(name="roleObjectId")
    def role_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_object_id")

