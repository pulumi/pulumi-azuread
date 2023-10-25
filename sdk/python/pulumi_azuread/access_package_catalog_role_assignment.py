# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AccessPackageCatalogRoleAssignmentArgs', 'AccessPackageCatalogRoleAssignment']

@pulumi.input_type
class AccessPackageCatalogRoleAssignmentArgs:
    def __init__(__self__, *,
                 catalog_id: pulumi.Input[str],
                 principal_object_id: pulumi.Input[str],
                 role_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a AccessPackageCatalogRoleAssignment resource.
        :param pulumi.Input[str] catalog_id: The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_object_id: The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_id: The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
        """
        AccessPackageCatalogRoleAssignmentArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            catalog_id=catalog_id,
            principal_object_id=principal_object_id,
            role_id=role_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             catalog_id: Optional[pulumi.Input[str]] = None,
             principal_object_id: Optional[pulumi.Input[str]] = None,
             role_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if catalog_id is None and 'catalogId' in kwargs:
            catalog_id = kwargs['catalogId']
        if catalog_id is None:
            raise TypeError("Missing 'catalog_id' argument")
        if principal_object_id is None and 'principalObjectId' in kwargs:
            principal_object_id = kwargs['principalObjectId']
        if principal_object_id is None:
            raise TypeError("Missing 'principal_object_id' argument")
        if role_id is None and 'roleId' in kwargs:
            role_id = kwargs['roleId']
        if role_id is None:
            raise TypeError("Missing 'role_id' argument")

        _setter("catalog_id", catalog_id)
        _setter("principal_object_id", principal_object_id)
        _setter("role_id", role_id)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Input[str]:
        """
        The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="principalObjectId")
    def principal_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_object_id")

    @principal_object_id.setter
    def principal_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal_object_id", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Input[str]:
        """
        The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_id", value)


@pulumi.input_type
class _AccessPackageCatalogRoleAssignmentState:
    def __init__(__self__, *,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 principal_object_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccessPackageCatalogRoleAssignment resources.
        :param pulumi.Input[str] catalog_id: The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_object_id: The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_id: The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
        """
        _AccessPackageCatalogRoleAssignmentState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            catalog_id=catalog_id,
            principal_object_id=principal_object_id,
            role_id=role_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             catalog_id: Optional[pulumi.Input[str]] = None,
             principal_object_id: Optional[pulumi.Input[str]] = None,
             role_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if catalog_id is None and 'catalogId' in kwargs:
            catalog_id = kwargs['catalogId']
        if principal_object_id is None and 'principalObjectId' in kwargs:
            principal_object_id = kwargs['principalObjectId']
        if role_id is None and 'roleId' in kwargs:
            role_id = kwargs['roleId']

        if catalog_id is not None:
            _setter("catalog_id", catalog_id)
        if principal_object_id is not None:
            _setter("principal_object_id", principal_object_id)
        if role_id is not None:
            _setter("role_id", role_id)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="principalObjectId")
    def principal_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_object_id")

    @principal_object_id.setter
    def principal_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_object_id", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_id", value)


class AccessPackageCatalogRoleAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 principal_object_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a single catalog role assignment within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `EntitlementManagement.ReadWrite.All` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Identity Governance administrator` or `Global Administrator`

        ## Import

        Catalog role assignments can be imported using the ID of the assignment, e.g.

        ```sh
         $ pulumi import azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_object_id: The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_id: The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessPackageCatalogRoleAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a single catalog role assignment within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `EntitlementManagement.ReadWrite.All` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Identity Governance administrator` or `Global Administrator`

        ## Import

        Catalog role assignments can be imported using the ID of the assignment, e.g.

        ```sh
         $ pulumi import azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param AccessPackageCatalogRoleAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessPackageCatalogRoleAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            AccessPackageCatalogRoleAssignmentArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 principal_object_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessPackageCatalogRoleAssignmentArgs.__new__(AccessPackageCatalogRoleAssignmentArgs)

            if catalog_id is None and not opts.urn:
                raise TypeError("Missing required property 'catalog_id'")
            __props__.__dict__["catalog_id"] = catalog_id
            if principal_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'principal_object_id'")
            __props__.__dict__["principal_object_id"] = principal_object_id
            if role_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_id'")
            __props__.__dict__["role_id"] = role_id
        super(AccessPackageCatalogRoleAssignment, __self__).__init__(
            'azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            catalog_id: Optional[pulumi.Input[str]] = None,
            principal_object_id: Optional[pulumi.Input[str]] = None,
            role_id: Optional[pulumi.Input[str]] = None) -> 'AccessPackageCatalogRoleAssignment':
        """
        Get an existing AccessPackageCatalogRoleAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_object_id: The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_id: The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccessPackageCatalogRoleAssignmentState.__new__(_AccessPackageCatalogRoleAssignmentState)

        __props__.__dict__["catalog_id"] = catalog_id
        __props__.__dict__["principal_object_id"] = principal_object_id
        __props__.__dict__["role_id"] = role_id
        return AccessPackageCatalogRoleAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[str]:
        """
        The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="principalObjectId")
    def principal_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_object_id")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[str]:
        """
        The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_id")

