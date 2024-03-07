# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AdministrativeUnitArgs', 'AdministrativeUnit']

@pulumi.input_type
class AdministrativeUnitArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 hidden_membership_enabled: Optional[pulumi.Input[bool]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 prevent_duplicate_names: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AdministrativeUnit resource.
        :param pulumi.Input[str] display_name: The display name of the administrative unit.
        :param pulumi.Input[str] description: The description of the administrative unit.
        :param pulumi.Input[bool] hidden_membership_enabled: Whether the administrative unit and its members are hidden or publicly viewable in the directory.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
               
               !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
        :param pulumi.Input[bool] prevent_duplicate_names: If `true`, will return an error if an existing administrative unit is found with the same name
        """
        pulumi.set(__self__, "display_name", display_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if hidden_membership_enabled is not None:
            pulumi.set(__self__, "hidden_membership_enabled", hidden_membership_enabled)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if prevent_duplicate_names is not None:
            pulumi.set(__self__, "prevent_duplicate_names", prevent_duplicate_names)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name of the administrative unit.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the administrative unit.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="hiddenMembershipEnabled")
    def hidden_membership_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the administrative unit and its members are hidden or publicly viewable in the directory.
        """
        return pulumi.get(self, "hidden_membership_enabled")

    @hidden_membership_enabled.setter
    def hidden_membership_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hidden_membership_enabled", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.

        !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="preventDuplicateNames")
    def prevent_duplicate_names(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, will return an error if an existing administrative unit is found with the same name
        """
        return pulumi.get(self, "prevent_duplicate_names")

    @prevent_duplicate_names.setter
    def prevent_duplicate_names(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "prevent_duplicate_names", value)


@pulumi.input_type
class _AdministrativeUnitState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 hidden_membership_enabled: Optional[pulumi.Input[bool]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 prevent_duplicate_names: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AdministrativeUnit resources.
        :param pulumi.Input[str] description: The description of the administrative unit.
        :param pulumi.Input[str] display_name: The display name of the administrative unit.
        :param pulumi.Input[bool] hidden_membership_enabled: Whether the administrative unit and its members are hidden or publicly viewable in the directory.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
               
               !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
        :param pulumi.Input[str] object_id: The object ID of the administrative unit.
        :param pulumi.Input[bool] prevent_duplicate_names: If `true`, will return an error if an existing administrative unit is found with the same name
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if hidden_membership_enabled is not None:
            pulumi.set(__self__, "hidden_membership_enabled", hidden_membership_enabled)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if object_id is not None:
            pulumi.set(__self__, "object_id", object_id)
        if prevent_duplicate_names is not None:
            pulumi.set(__self__, "prevent_duplicate_names", prevent_duplicate_names)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the administrative unit.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the administrative unit.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="hiddenMembershipEnabled")
    def hidden_membership_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the administrative unit and its members are hidden or publicly viewable in the directory.
        """
        return pulumi.get(self, "hidden_membership_enabled")

    @hidden_membership_enabled.setter
    def hidden_membership_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hidden_membership_enabled", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.

        !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the administrative unit.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter(name="preventDuplicateNames")
    def prevent_duplicate_names(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, will return an error if an existing administrative unit is found with the same name
        """
        return pulumi.get(self, "prevent_duplicate_names")

    @prevent_duplicate_names.setter
    def prevent_duplicate_names(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "prevent_duplicate_names", value)


class AdministrativeUnit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 hidden_membership_enabled: Optional[pulumi.Input[bool]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 prevent_duplicate_names: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manages an Administrative Unit within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.AdministrativeUnit("example",
            display_name="Example-AU",
            description="Just an example",
            hidden_membership_enabled=False)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Administrative units can be imported using their object ID, e.g.

        ```sh
        $ pulumi import azuread:index/administrativeUnit:AdministrativeUnit example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the administrative unit.
        :param pulumi.Input[str] display_name: The display name of the administrative unit.
        :param pulumi.Input[bool] hidden_membership_enabled: Whether the administrative unit and its members are hidden or publicly viewable in the directory.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
               
               !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
        :param pulumi.Input[bool] prevent_duplicate_names: If `true`, will return an error if an existing administrative unit is found with the same name
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AdministrativeUnitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Administrative Unit within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.AdministrativeUnit("example",
            display_name="Example-AU",
            description="Just an example",
            hidden_membership_enabled=False)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Administrative units can be imported using their object ID, e.g.

        ```sh
        $ pulumi import azuread:index/administrativeUnit:AdministrativeUnit example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param AdministrativeUnitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AdministrativeUnitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 hidden_membership_enabled: Optional[pulumi.Input[bool]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 prevent_duplicate_names: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AdministrativeUnitArgs.__new__(AdministrativeUnitArgs)

            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["hidden_membership_enabled"] = hidden_membership_enabled
            __props__.__dict__["members"] = members
            __props__.__dict__["prevent_duplicate_names"] = prevent_duplicate_names
            __props__.__dict__["object_id"] = None
        super(AdministrativeUnit, __self__).__init__(
            'azuread:index/administrativeUnit:AdministrativeUnit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            hidden_membership_enabled: Optional[pulumi.Input[bool]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            object_id: Optional[pulumi.Input[str]] = None,
            prevent_duplicate_names: Optional[pulumi.Input[bool]] = None) -> 'AdministrativeUnit':
        """
        Get an existing AdministrativeUnit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the administrative unit.
        :param pulumi.Input[str] display_name: The display name of the administrative unit.
        :param pulumi.Input[bool] hidden_membership_enabled: Whether the administrative unit and its members are hidden or publicly viewable in the directory.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.
               
               !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
        :param pulumi.Input[str] object_id: The object ID of the administrative unit.
        :param pulumi.Input[bool] prevent_duplicate_names: If `true`, will return an error if an existing administrative unit is found with the same name
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AdministrativeUnitState.__new__(_AdministrativeUnitState)

        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["hidden_membership_enabled"] = hidden_membership_enabled
        __props__.__dict__["members"] = members
        __props__.__dict__["object_id"] = object_id
        __props__.__dict__["prevent_duplicate_names"] = prevent_duplicate_names
        return AdministrativeUnit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the administrative unit.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the administrative unit.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="hiddenMembershipEnabled")
    def hidden_membership_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the administrative unit and its members are hidden or publicly viewable in the directory.
        """
        return pulumi.get(self, "hidden_membership_enabled")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        """
        A set of object IDs of members who should be present in this administrative unit. Supported object types are Users or Groups.

        !> **Warning** Do not use the `members` property at the same time as the AdministrativeUnitMember resource for the same administrative unit. Doing so will cause a conflict and administrative unit members will be removed.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the administrative unit.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter(name="preventDuplicateNames")
    def prevent_duplicate_names(self) -> pulumi.Output[Optional[bool]]:
        """
        If `true`, will return an error if an existing administrative unit is found with the same name
        """
        return pulumi.get(self, "prevent_duplicate_names")

