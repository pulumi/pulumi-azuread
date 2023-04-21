# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AccessPackageResourceCatalogAssociationArgs', 'AccessPackageResourceCatalogAssociation']

@pulumi.input_type
class AccessPackageResourceCatalogAssociationArgs:
    def __init__(__self__, *,
                 catalog_id: pulumi.Input[str],
                 resource_origin_id: pulumi.Input[str],
                 resource_origin_system: pulumi.Input[str]):
        """
        The set of arguments for constructing a AccessPackageResourceCatalogAssociation resource.
        :param pulumi.Input[str] catalog_id: The unique ID of the access package catalog. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_origin_id: The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_origin_system: The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "catalog_id", catalog_id)
        pulumi.set(__self__, "resource_origin_id", resource_origin_id)
        pulumi.set(__self__, "resource_origin_system", resource_origin_system)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Input[str]:
        """
        The unique ID of the access package catalog. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="resourceOriginId")
    def resource_origin_id(self) -> pulumi.Input[str]:
        """
        The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_origin_id")

    @resource_origin_id.setter
    def resource_origin_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_origin_id", value)

    @property
    @pulumi.getter(name="resourceOriginSystem")
    def resource_origin_system(self) -> pulumi.Input[str]:
        """
        The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_origin_system")

    @resource_origin_system.setter
    def resource_origin_system(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_origin_system", value)


@pulumi.input_type
class _AccessPackageResourceCatalogAssociationState:
    def __init__(__self__, *,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 resource_origin_id: Optional[pulumi.Input[str]] = None,
                 resource_origin_system: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccessPackageResourceCatalogAssociation resources.
        :param pulumi.Input[str] catalog_id: The unique ID of the access package catalog. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_origin_id: The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_origin_system: The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
        """
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if resource_origin_id is not None:
            pulumi.set(__self__, "resource_origin_id", resource_origin_id)
        if resource_origin_system is not None:
            pulumi.set(__self__, "resource_origin_system", resource_origin_system)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID of the access package catalog. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="resourceOriginId")
    def resource_origin_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_origin_id")

    @resource_origin_id.setter
    def resource_origin_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_origin_id", value)

    @property
    @pulumi.getter(name="resourceOriginSystem")
    def resource_origin_system(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_origin_system")

    @resource_origin_system.setter
    def resource_origin_system(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_origin_system", value)


class AccessPackageResourceCatalogAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 resource_origin_id: Optional[pulumi.Input[str]] = None,
                 resource_origin_system: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages the resources added to access package catalogs within Identity Governance in Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

        When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_group = azuread.Group("exampleGroup",
            display_name="example-group",
            security_enabled=True)
        example_access_package_catalog = azuread.AccessPackageCatalog("exampleAccessPackageCatalog",
            display_name="example-catalog",
            description="Example catalog")
        example_access_package_resource_catalog_association = azuread.AccessPackageResourceCatalogAssociation("exampleAccessPackageResourceCatalogAssociation",
            catalog_id=azuread_access_package_catalog["example_catalog"]["id"],
            resource_origin_id=azuread_group["example_group"]["object_id"],
            resource_origin_system="AadGroup")
        ```

        ## Import

        The resource and catalog association can be imported using the catalog ID and the resource origin ID, e.g.

        ```sh
         $ pulumi import azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the Catalog ID and the Resource Origin ID in the format `{CatalogID}/{ResourceOriginID}`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The unique ID of the access package catalog. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_origin_id: The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_origin_system: The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessPackageResourceCatalogAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages the resources added to access package catalogs within Identity Governance in Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

        When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_group = azuread.Group("exampleGroup",
            display_name="example-group",
            security_enabled=True)
        example_access_package_catalog = azuread.AccessPackageCatalog("exampleAccessPackageCatalog",
            display_name="example-catalog",
            description="Example catalog")
        example_access_package_resource_catalog_association = azuread.AccessPackageResourceCatalogAssociation("exampleAccessPackageResourceCatalogAssociation",
            catalog_id=azuread_access_package_catalog["example_catalog"]["id"],
            resource_origin_id=azuread_group["example_group"]["object_id"],
            resource_origin_system="AadGroup")
        ```

        ## Import

        The resource and catalog association can be imported using the catalog ID and the resource origin ID, e.g.

        ```sh
         $ pulumi import azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation example 00000000-0000-0000-0000-000000000000/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the Catalog ID and the Resource Origin ID in the format `{CatalogID}/{ResourceOriginID}`.

        :param str resource_name: The name of the resource.
        :param AccessPackageResourceCatalogAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessPackageResourceCatalogAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 resource_origin_id: Optional[pulumi.Input[str]] = None,
                 resource_origin_system: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessPackageResourceCatalogAssociationArgs.__new__(AccessPackageResourceCatalogAssociationArgs)

            if catalog_id is None and not opts.urn:
                raise TypeError("Missing required property 'catalog_id'")
            __props__.__dict__["catalog_id"] = catalog_id
            if resource_origin_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_origin_id'")
            __props__.__dict__["resource_origin_id"] = resource_origin_id
            if resource_origin_system is None and not opts.urn:
                raise TypeError("Missing required property 'resource_origin_system'")
            __props__.__dict__["resource_origin_system"] = resource_origin_system
        super(AccessPackageResourceCatalogAssociation, __self__).__init__(
            'azuread:index/accessPackageResourceCatalogAssociation:AccessPackageResourceCatalogAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            catalog_id: Optional[pulumi.Input[str]] = None,
            resource_origin_id: Optional[pulumi.Input[str]] = None,
            resource_origin_system: Optional[pulumi.Input[str]] = None) -> 'AccessPackageResourceCatalogAssociation':
        """
        Get an existing AccessPackageResourceCatalogAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The unique ID of the access package catalog. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_origin_id: The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_origin_system: The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccessPackageResourceCatalogAssociationState.__new__(_AccessPackageResourceCatalogAssociationState)

        __props__.__dict__["catalog_id"] = catalog_id
        __props__.__dict__["resource_origin_id"] = resource_origin_id
        __props__.__dict__["resource_origin_system"] = resource_origin_system
        return AccessPackageResourceCatalogAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[str]:
        """
        The unique ID of the access package catalog. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="resourceOriginId")
    def resource_origin_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the resource in the origin system. In the case of an Azure AD group, this is the identifier of the group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_origin_id")

    @property
    @pulumi.getter(name="resourceOriginSystem")
    def resource_origin_system(self) -> pulumi.Output[str]:
        """
        The type of the resource in the origin system, such as `SharePointOnline`, `AadApplication` or `AadGroup`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_origin_system")
