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

__all__ = ['AccessPackageCatalogArgs', 'AccessPackageCatalog']

@pulumi.input_type
class AccessPackageCatalogArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 display_name: pulumi.Input[str],
                 externally_visible: Optional[pulumi.Input[bool]] = None,
                 published: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AccessPackageCatalog resource.
        :param pulumi.Input[str] description: The description of the access package catalog.
        :param pulumi.Input[str] display_name: The display name of the access package catalog.
        :param pulumi.Input[bool] externally_visible: Whether the access packages in this catalog can be requested by users outside the tenant.
        :param pulumi.Input[bool] published: Whether the access packages in this catalog are available for management.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "display_name", display_name)
        if externally_visible is not None:
            pulumi.set(__self__, "externally_visible", externally_visible)
        if published is not None:
            pulumi.set(__self__, "published", published)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        The description of the access package catalog.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name of the access package catalog.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="externallyVisible")
    def externally_visible(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the access packages in this catalog can be requested by users outside the tenant.
        """
        return pulumi.get(self, "externally_visible")

    @externally_visible.setter
    def externally_visible(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "externally_visible", value)

    @property
    @pulumi.getter
    def published(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the access packages in this catalog are available for management.
        """
        return pulumi.get(self, "published")

    @published.setter
    def published(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "published", value)


@pulumi.input_type
class _AccessPackageCatalogState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 externally_visible: Optional[pulumi.Input[bool]] = None,
                 published: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AccessPackageCatalog resources.
        :param pulumi.Input[str] description: The description of the access package catalog.
        :param pulumi.Input[str] display_name: The display name of the access package catalog.
        :param pulumi.Input[bool] externally_visible: Whether the access packages in this catalog can be requested by users outside the tenant.
        :param pulumi.Input[bool] published: Whether the access packages in this catalog are available for management.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if externally_visible is not None:
            pulumi.set(__self__, "externally_visible", externally_visible)
        if published is not None:
            pulumi.set(__self__, "published", published)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the access package catalog.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the access package catalog.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="externallyVisible")
    def externally_visible(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the access packages in this catalog can be requested by users outside the tenant.
        """
        return pulumi.get(self, "externally_visible")

    @externally_visible.setter
    def externally_visible(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "externally_visible", value)

    @property
    @pulumi.getter
    def published(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the access packages in this catalog are available for management.
        """
        return pulumi.get(self, "published")

    @published.setter
    def published(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "published", value)


class AccessPackageCatalog(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 externally_visible: Optional[pulumi.Input[bool]] = None,
                 published: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manages an access package catalog within Identity Governance in Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

        When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Catalog creator` or `Global Administrator`

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.AccessPackageCatalog("example",
            display_name="example-access-package-catalog",
            description="Example access package catalog")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        An Access Package Catalog can be imported using the `id`, e.g.

        ```sh
        $ pulumi import azuread:index/accessPackageCatalog:AccessPackageCatalog example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the access package catalog.
        :param pulumi.Input[str] display_name: The display name of the access package catalog.
        :param pulumi.Input[bool] externally_visible: Whether the access packages in this catalog can be requested by users outside the tenant.
        :param pulumi.Input[bool] published: Whether the access packages in this catalog are available for management.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccessPackageCatalogArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an access package catalog within Identity Governance in Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `EntitlementManagement.ReadWrite.All`.

        When authenticated with a user principal, this resource requires one of the following directory roles: `Catalog owner`, `Catalog creator` or `Global Administrator`

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.AccessPackageCatalog("example",
            display_name="example-access-package-catalog",
            description="Example access package catalog")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        An Access Package Catalog can be imported using the `id`, e.g.

        ```sh
        $ pulumi import azuread:index/accessPackageCatalog:AccessPackageCatalog example 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param AccessPackageCatalogArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccessPackageCatalogArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 externally_visible: Optional[pulumi.Input[bool]] = None,
                 published: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccessPackageCatalogArgs.__new__(AccessPackageCatalogArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["externally_visible"] = externally_visible
            __props__.__dict__["published"] = published
        super(AccessPackageCatalog, __self__).__init__(
            'azuread:index/accessPackageCatalog:AccessPackageCatalog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            externally_visible: Optional[pulumi.Input[bool]] = None,
            published: Optional[pulumi.Input[bool]] = None) -> 'AccessPackageCatalog':
        """
        Get an existing AccessPackageCatalog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the access package catalog.
        :param pulumi.Input[str] display_name: The display name of the access package catalog.
        :param pulumi.Input[bool] externally_visible: Whether the access packages in this catalog can be requested by users outside the tenant.
        :param pulumi.Input[bool] published: Whether the access packages in this catalog are available for management.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccessPackageCatalogState.__new__(_AccessPackageCatalogState)

        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["externally_visible"] = externally_visible
        __props__.__dict__["published"] = published
        return AccessPackageCatalog(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the access package catalog.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the access package catalog.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="externallyVisible")
    def externally_visible(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the access packages in this catalog can be requested by users outside the tenant.
        """
        return pulumi.get(self, "externally_visible")

    @property
    @pulumi.getter
    def published(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the access packages in this catalog are available for management.
        """
        return pulumi.get(self, "published")

