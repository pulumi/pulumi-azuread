# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AccessPackageArgs', 'AccessPackage']

@pulumi.input_type
class AccessPackageArgs:
    def __init__(__self__, *,
                 catalog_id: pulumi.Input[str],
                 description: pulumi.Input[str],
                 display_name: pulumi.Input[str],
                 hidden: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AccessPackage resource.
        :param pulumi.Input[str] catalog_id: The ID of the Catalog this access package will be created in.
        :param pulumi.Input[str] description: The description of the access package.
        :param pulumi.Input[str] display_name: The display name of the access package.
        :param pulumi.Input[bool] hidden: Whether the access package is hidden from the requestor.
        """
        pulumi.set(__self__, "catalog_id", catalog_id)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "display_name", display_name)
        if hidden is not None:
            pulumi.set(__self__, "hidden", hidden)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Input[str]:
        """
        The ID of the Catalog this access package will be created in.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        The description of the access package.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name of the access package.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def hidden(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the access package is hidden from the requestor.
        """
        return pulumi.get(self, "hidden")

    @hidden.setter
    def hidden(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hidden", value)


@pulumi.input_type
class _AccessPackageState:
    def __init__(__self__, *,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 hidden: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AccessPackage resources.
        :param pulumi.Input[str] catalog_id: The ID of the Catalog this access package will be created in.
        :param pulumi.Input[str] description: The description of the access package.
        :param pulumi.Input[str] display_name: The display name of the access package.
        :param pulumi.Input[bool] hidden: Whether the access package is hidden from the requestor.
        """
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if hidden is not None:
            pulumi.set(__self__, "hidden", hidden)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Catalog this access package will be created in.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the access package.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the access package.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def hidden(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the access package is hidden from the requestor.
        """
        return pulumi.get(self, "hidden")

    @hidden.setter
    def hidden(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hidden", value)


class AccessPackage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 hidden: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manages an Access Package within Identity Governance in Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

        When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.index.access_package_catalog.AccessPackageCatalog("example",
            display_name=example-catalog,
            description=Example catalog)
        example_access_package = azuread.index.access_package.AccessPackage("example",
            catalog_id=example.id,
            display_name=access-package,
            description=Access Package)
        ```

        ## Import

        Access Packages can be imported using the `id`, e.g.

        ```sh
        $ pulumi import azuread:index/accessPackage:AccessPackage example_package 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The ID of the Catalog this access package will be created in.
        :param pulumi.Input[str] description: The description of the access package.
        :param pulumi.Input[str] display_name: The display name of the access package.
        :param pulumi.Input[bool] hidden: Whether the access package is hidden from the requestor.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessPackageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Access Package within Identity Governance in Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

        When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Access package manager` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.index.access_package_catalog.AccessPackageCatalog("example",
            display_name=example-catalog,
            description=Example catalog)
        example_access_package = azuread.index.access_package.AccessPackage("example",
            catalog_id=example.id,
            display_name=access-package,
            description=Access Package)
        ```

        ## Import

        Access Packages can be imported using the `id`, e.g.

        ```sh
        $ pulumi import azuread:index/accessPackage:AccessPackage example_package 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param AccessPackageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessPackageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 hidden: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessPackageArgs.__new__(AccessPackageArgs)

            if catalog_id is None and not opts.urn:
                raise TypeError("Missing required property 'catalog_id'")
            __props__.__dict__["catalog_id"] = catalog_id
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["hidden"] = hidden
        super(AccessPackage, __self__).__init__(
            'azuread:index/accessPackage:AccessPackage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            catalog_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            hidden: Optional[pulumi.Input[bool]] = None) -> 'AccessPackage':
        """
        Get an existing AccessPackage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The ID of the Catalog this access package will be created in.
        :param pulumi.Input[str] description: The description of the access package.
        :param pulumi.Input[str] display_name: The display name of the access package.
        :param pulumi.Input[bool] hidden: Whether the access package is hidden from the requestor.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccessPackageState.__new__(_AccessPackageState)

        __props__.__dict__["catalog_id"] = catalog_id
        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["hidden"] = hidden
        return AccessPackage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[str]:
        """
        The ID of the Catalog this access package will be created in.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the access package.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the access package.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def hidden(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the access package is hidden from the requestor.
        """
        return pulumi.get(self, "hidden")

