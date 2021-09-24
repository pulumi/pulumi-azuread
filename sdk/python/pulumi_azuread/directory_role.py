# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DirectoryRoleArgs', 'DirectoryRole']

@pulumi.input_type
class DirectoryRoleArgs:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DirectoryRole resource.
        :param pulumi.Input[str] display_name: The display name of the directory role to activate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] template_id: The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the directory role to activate. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)


@pulumi.input_type
class _DirectoryRoleState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DirectoryRole resources.
        :param pulumi.Input[str] description: The description of the directory role.
        :param pulumi.Input[str] display_name: The display name of the directory role to activate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] object_id: The object ID of the directory role.
        :param pulumi.Input[str] template_id: The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if object_id is not None:
            pulumi.set(__self__, "object_id", object_id)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the directory role.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the directory role to activate. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the directory role.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)


class DirectoryRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a Directory Role within Azure Active Directory. Directory Roles are also known as Administrator Roles.

        Directory Roles are built-in to Azure Active Directory and are immutable. However, by default they are not activated in a tenant (except for the Global Administrator role). This resource ensures a directory role is activated from its associated role template, and exports the object ID of the role, so that role assignments can be made for it.

        Once activated, directory roles cannot be deactivated and so this resource does not perform any actions on destroy.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Example Usage

        *Activate a directory role by its template ID*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.DirectoryRole("example", template_id="00000000-0000-0000-0000-000000000000")
        ```

        *Activate a directory role by display name*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.DirectoryRole("example", display_name="Printer administrator")
        ```

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The display name of the directory role to activate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] template_id: The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DirectoryRoleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Directory Role within Azure Active Directory. Directory Roles are also known as Administrator Roles.

        Directory Roles are built-in to Azure Active Directory and are immutable. However, by default they are not activated in a tenant (except for the Global Administrator role). This resource ensures a directory role is activated from its associated role template, and exports the object ID of the role, so that role assignments can be made for it.

        Once activated, directory roles cannot be deactivated and so this resource does not perform any actions on destroy.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Example Usage

        *Activate a directory role by its template ID*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.DirectoryRole("example", template_id="00000000-0000-0000-0000-000000000000")
        ```

        *Activate a directory role by display name*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.DirectoryRole("example", display_name="Printer administrator")
        ```

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param DirectoryRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DirectoryRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = DirectoryRoleArgs.__new__(DirectoryRoleArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["template_id"] = template_id
            __props__.__dict__["description"] = None
            __props__.__dict__["object_id"] = None
        super(DirectoryRole, __self__).__init__(
            'azuread:index/directoryRole:DirectoryRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            object_id: Optional[pulumi.Input[str]] = None,
            template_id: Optional[pulumi.Input[str]] = None) -> 'DirectoryRole':
        """
        Get an existing DirectoryRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the directory role.
        :param pulumi.Input[str] display_name: The display name of the directory role to activate. Changing this forces a new resource to be created.
        :param pulumi.Input[str] object_id: The object ID of the directory role.
        :param pulumi.Input[str] template_id: The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DirectoryRoleState.__new__(_DirectoryRoleState)

        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["object_id"] = object_id
        __props__.__dict__["template_id"] = template_id
        return DirectoryRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the directory role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the directory role to activate. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the directory role.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[str]:
        """
        The object ID of the role template from which to activate the directory role. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "template_id")

