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

__all__ = ['CustomDirectoryRoleArgs', 'CustomDirectoryRole']

@pulumi.input_type
class CustomDirectoryRoleArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 enabled: pulumi.Input[bool],
                 permissions: pulumi.Input[Sequence[pulumi.Input['CustomDirectoryRolePermissionArgs']]],
                 version: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CustomDirectoryRole resource.
        :param pulumi.Input[str] display_name: The display name of the custom directory role.
        :param pulumi.Input[bool] enabled: Indicates whether the role is enabled for assignment.
        :param pulumi.Input[Sequence[pulumi.Input['CustomDirectoryRolePermissionArgs']]] permissions: A collection of `permissions` blocks as documented below.
        :param pulumi.Input[str] version: The version of the role definition. This can be any arbitrary string between 1-128 characters.
        :param pulumi.Input[str] description: The description of the custom directory role.
        :param pulumi.Input[str] template_id: Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "version", version)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name of the custom directory role.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Indicates whether the role is enabled for assignment.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input[Sequence[pulumi.Input['CustomDirectoryRolePermissionArgs']]]:
        """
        A collection of `permissions` blocks as documented below.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input[Sequence[pulumi.Input['CustomDirectoryRolePermissionArgs']]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        The version of the role definition. This can be any arbitrary string between 1-128 characters.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the custom directory role.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)


@pulumi.input_type
class _CustomDirectoryRoleState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDirectoryRolePermissionArgs']]]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CustomDirectoryRole resources.
        :param pulumi.Input[str] description: The description of the custom directory role.
        :param pulumi.Input[str] display_name: The display name of the custom directory role.
        :param pulumi.Input[bool] enabled: Indicates whether the role is enabled for assignment.
        :param pulumi.Input[str] object_id: The object ID of the custom directory role.
        :param pulumi.Input[Sequence[pulumi.Input['CustomDirectoryRolePermissionArgs']]] permissions: A collection of `permissions` blocks as documented below.
        :param pulumi.Input[str] template_id: Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
        :param pulumi.Input[str] version: The version of the role definition. This can be any arbitrary string between 1-128 characters.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if object_id is not None:
            pulumi.set(__self__, "object_id", object_id)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the custom directory role.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the custom directory role.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the role is enabled for assignment.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the custom directory role.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CustomDirectoryRolePermissionArgs']]]]:
        """
        A collection of `permissions` blocks as documented below.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CustomDirectoryRolePermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the role definition. This can be any arbitrary string between 1-128 characters.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class CustomDirectoryRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['CustomDirectoryRolePermissionArgs', 'CustomDirectoryRolePermissionArgsDict']]]]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Custom Directory Role within Azure Active Directory.

        This resource is for managing custom directory roles. For management of built-in roles, see the DirectoryRole resource.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.CustomDirectoryRole("example",
            display_name="My Custom Role",
            description="Allows reading applications and updating groups",
            enabled=True,
            version="1.0",
            permissions=[
                {
                    "allowedResourceActions": [
                        "microsoft.directory/applications/basic/update",
                        "microsoft.directory/applications/create",
                        "microsoft.directory/applications/standard/read",
                    ],
                },
                {
                    "allowedResourceActions": [
                        "microsoft.directory/groups/allProperties/read",
                        "microsoft.directory/groups/allProperties/read",
                        "microsoft.directory/groups/basic/update",
                        "microsoft.directory/groups/create",
                        "microsoft.directory/groups/delete",
                    ],
                },
            ])
        ```

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the custom directory role.
        :param pulumi.Input[str] display_name: The display name of the custom directory role.
        :param pulumi.Input[bool] enabled: Indicates whether the role is enabled for assignment.
        :param pulumi.Input[Sequence[pulumi.Input[Union['CustomDirectoryRolePermissionArgs', 'CustomDirectoryRolePermissionArgsDict']]]] permissions: A collection of `permissions` blocks as documented below.
        :param pulumi.Input[str] template_id: Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
        :param pulumi.Input[str] version: The version of the role definition. This can be any arbitrary string between 1-128 characters.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomDirectoryRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Custom Directory Role within Azure Active Directory.

        This resource is for managing custom directory roles. For management of built-in roles, see the DirectoryRole resource.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.CustomDirectoryRole("example",
            display_name="My Custom Role",
            description="Allows reading applications and updating groups",
            enabled=True,
            version="1.0",
            permissions=[
                {
                    "allowedResourceActions": [
                        "microsoft.directory/applications/basic/update",
                        "microsoft.directory/applications/create",
                        "microsoft.directory/applications/standard/read",
                    ],
                },
                {
                    "allowedResourceActions": [
                        "microsoft.directory/groups/allProperties/read",
                        "microsoft.directory/groups/allProperties/read",
                        "microsoft.directory/groups/basic/update",
                        "microsoft.directory/groups/create",
                        "microsoft.directory/groups/delete",
                    ],
                },
            ])
        ```

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param CustomDirectoryRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomDirectoryRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['CustomDirectoryRolePermissionArgs', 'CustomDirectoryRolePermissionArgsDict']]]]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomDirectoryRoleArgs.__new__(CustomDirectoryRoleArgs)

            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["template_id"] = template_id
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
            __props__.__dict__["object_id"] = None
        super(CustomDirectoryRole, __self__).__init__(
            'azuread:index/customDirectoryRole:CustomDirectoryRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            object_id: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['CustomDirectoryRolePermissionArgs', 'CustomDirectoryRolePermissionArgsDict']]]]] = None,
            template_id: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'CustomDirectoryRole':
        """
        Get an existing CustomDirectoryRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the custom directory role.
        :param pulumi.Input[str] display_name: The display name of the custom directory role.
        :param pulumi.Input[bool] enabled: Indicates whether the role is enabled for assignment.
        :param pulumi.Input[str] object_id: The object ID of the custom directory role.
        :param pulumi.Input[Sequence[pulumi.Input[Union['CustomDirectoryRolePermissionArgs', 'CustomDirectoryRolePermissionArgsDict']]]] permissions: A collection of `permissions` blocks as documented below.
        :param pulumi.Input[str] template_id: Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
        :param pulumi.Input[str] version: The version of the role definition. This can be any arbitrary string between 1-128 characters.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomDirectoryRoleState.__new__(_CustomDirectoryRoleState)

        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["object_id"] = object_id
        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["template_id"] = template_id
        __props__.__dict__["version"] = version
        return CustomDirectoryRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the custom directory role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the custom directory role.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Indicates whether the role is enabled for assignment.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the custom directory role.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Sequence['outputs.CustomDirectoryRolePermission']]:
        """
        A collection of `permissions` blocks as documented below.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[str]:
        """
        Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The version of the role definition. This can be any arbitrary string between 1-128 characters.
        """
        return pulumi.get(self, "version")

