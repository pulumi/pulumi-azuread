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
    from typing import NotRequired, TypedDict
else:
    from typing_extensions import NotRequired, TypedDict
from . import _utilities

__all__ = ['AccessPackageResourcePackageAssociationArgs', 'AccessPackageResourcePackageAssociation']

@pulumi.input_type
class AccessPackageResourcePackageAssociationArgs:
    def __init__(__self__, *,
                 access_package_id: pulumi.Input[str],
                 catalog_resource_association_id: pulumi.Input[str],
                 access_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AccessPackageResourcePackageAssociation resource.
        :param pulumi.Input[str] access_package_id: The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] catalog_resource_association_id: The ID of the catalog association from the `AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] access_type: The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "access_package_id", access_package_id)
        pulumi.set(__self__, "catalog_resource_association_id", catalog_resource_association_id)
        if access_type is not None:
            pulumi.set(__self__, "access_type", access_type)

    @property
    @pulumi.getter(name="accessPackageId")
    def access_package_id(self) -> pulumi.Input[str]:
        """
        The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "access_package_id")

    @access_package_id.setter
    def access_package_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_package_id", value)

    @property
    @pulumi.getter(name="catalogResourceAssociationId")
    def catalog_resource_association_id(self) -> pulumi.Input[str]:
        """
        The ID of the catalog association from the `AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "catalog_resource_association_id")

    @catalog_resource_association_id.setter
    def catalog_resource_association_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "catalog_resource_association_id", value)

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> Optional[pulumi.Input[str]]:
        """
        The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "access_type")

    @access_type.setter
    def access_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_type", value)


@pulumi.input_type
class _AccessPackageResourcePackageAssociationState:
    def __init__(__self__, *,
                 access_package_id: Optional[pulumi.Input[str]] = None,
                 access_type: Optional[pulumi.Input[str]] = None,
                 catalog_resource_association_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccessPackageResourcePackageAssociation resources.
        :param pulumi.Input[str] access_package_id: The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] access_type: The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] catalog_resource_association_id: The ID of the catalog association from the `AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
        """
        if access_package_id is not None:
            pulumi.set(__self__, "access_package_id", access_package_id)
        if access_type is not None:
            pulumi.set(__self__, "access_type", access_type)
        if catalog_resource_association_id is not None:
            pulumi.set(__self__, "catalog_resource_association_id", catalog_resource_association_id)

    @property
    @pulumi.getter(name="accessPackageId")
    def access_package_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "access_package_id")

    @access_package_id.setter
    def access_package_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_package_id", value)

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> Optional[pulumi.Input[str]]:
        """
        The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "access_type")

    @access_type.setter
    def access_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_type", value)

    @property
    @pulumi.getter(name="catalogResourceAssociationId")
    def catalog_resource_association_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the catalog association from the `AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "catalog_resource_association_id")

    @catalog_resource_association_id.setter
    def catalog_resource_association_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_resource_association_id", value)


class AccessPackageResourcePackageAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_package_id: Optional[pulumi.Input[str]] = None,
                 access_type: Optional[pulumi.Input[str]] = None,
                 catalog_resource_association_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages the resources added to access packages within Identity Governance in Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

        When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Group("example",
            display_name="example-group",
            security_enabled=True)
        example_access_package_catalog = azuread.AccessPackageCatalog("example",
            display_name="example-catalog",
            description="Example catalog")
        example_access_package_resource_catalog_association = azuread.AccessPackageResourceCatalogAssociation("example",
            catalog_id=example_catalog["id"],
            resource_origin_id=example_group["objectId"],
            resource_origin_system="AadGroup")
        example_access_package = azuread.AccessPackage("example",
            display_name="example-package",
            description="Example Package",
            catalog_id=example_catalog["id"])
        example_access_package_resource_package_association = azuread.AccessPackageResourcePackageAssociation("example",
            access_package_id=example_access_package.id,
            catalog_resource_association_id=example_access_package_resource_catalog_association.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        The resource and catalog association can be imported using the access package ID, the access package ResourceRoleScope, the resource origin ID, and the access type, e.g.

        ```sh
        $ pulumi import azuread:index/accessPackageResourcePackageAssociation:AccessPackageResourcePackageAssociation example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-22222222/33333333-3333-3333-3333-33333333/Member
        ```

        -> This ID format is unique to Terraform and is composed of the Access Package ID, the access package ResourceRoleScope (in the format Role_Scope), the Resource Origin ID, and the Access Type, in the format `{AccessPackageID}/{ResourceRoleScope}/{ResourceOriginID}/{AccessType}`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_package_id: The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] access_type: The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] catalog_resource_association_id: The ID of the catalog association from the `AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessPackageResourcePackageAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages the resources added to access packages within Identity Governance in Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

        When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Group("example",
            display_name="example-group",
            security_enabled=True)
        example_access_package_catalog = azuread.AccessPackageCatalog("example",
            display_name="example-catalog",
            description="Example catalog")
        example_access_package_resource_catalog_association = azuread.AccessPackageResourceCatalogAssociation("example",
            catalog_id=example_catalog["id"],
            resource_origin_id=example_group["objectId"],
            resource_origin_system="AadGroup")
        example_access_package = azuread.AccessPackage("example",
            display_name="example-package",
            description="Example Package",
            catalog_id=example_catalog["id"])
        example_access_package_resource_package_association = azuread.AccessPackageResourcePackageAssociation("example",
            access_package_id=example_access_package.id,
            catalog_resource_association_id=example_access_package_resource_catalog_association.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        The resource and catalog association can be imported using the access package ID, the access package ResourceRoleScope, the resource origin ID, and the access type, e.g.

        ```sh
        $ pulumi import azuread:index/accessPackageResourcePackageAssociation:AccessPackageResourcePackageAssociation example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111_22222222-2222-2222-2222-22222222/33333333-3333-3333-3333-33333333/Member
        ```

        -> This ID format is unique to Terraform and is composed of the Access Package ID, the access package ResourceRoleScope (in the format Role_Scope), the Resource Origin ID, and the Access Type, in the format `{AccessPackageID}/{ResourceRoleScope}/{ResourceOriginID}/{AccessType}`.

        :param str resource_name: The name of the resource.
        :param AccessPackageResourcePackageAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessPackageResourcePackageAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_package_id: Optional[pulumi.Input[str]] = None,
                 access_type: Optional[pulumi.Input[str]] = None,
                 catalog_resource_association_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessPackageResourcePackageAssociationArgs.__new__(AccessPackageResourcePackageAssociationArgs)

            if access_package_id is None and not opts.urn:
                raise TypeError("Missing required property 'access_package_id'")
            __props__.__dict__["access_package_id"] = access_package_id
            __props__.__dict__["access_type"] = access_type
            if catalog_resource_association_id is None and not opts.urn:
                raise TypeError("Missing required property 'catalog_resource_association_id'")
            __props__.__dict__["catalog_resource_association_id"] = catalog_resource_association_id
        super(AccessPackageResourcePackageAssociation, __self__).__init__(
            'azuread:index/accessPackageResourcePackageAssociation:AccessPackageResourcePackageAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_package_id: Optional[pulumi.Input[str]] = None,
            access_type: Optional[pulumi.Input[str]] = None,
            catalog_resource_association_id: Optional[pulumi.Input[str]] = None) -> 'AccessPackageResourcePackageAssociation':
        """
        Get an existing AccessPackageResourcePackageAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_package_id: The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] access_type: The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] catalog_resource_association_id: The ID of the catalog association from the `AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccessPackageResourcePackageAssociationState.__new__(_AccessPackageResourcePackageAssociationState)

        __props__.__dict__["access_package_id"] = access_package_id
        __props__.__dict__["access_type"] = access_type
        __props__.__dict__["catalog_resource_association_id"] = catalog_resource_association_id
        return AccessPackageResourcePackageAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessPackageId")
    def access_package_id(self) -> pulumi.Output[str]:
        """
        The ID of access package this resource association is configured to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "access_package_id")

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> pulumi.Output[Optional[str]]:
        """
        The role of access type to the specified resource. Valid values are `Member`, or `Owner` The default is `Member`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "access_type")

    @property
    @pulumi.getter(name="catalogResourceAssociationId")
    def catalog_resource_association_id(self) -> pulumi.Output[str]:
        """
        The ID of the catalog association from the `AccessPackageResourceCatalogAssociation` resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "catalog_resource_association_id")

