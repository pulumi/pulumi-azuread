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

__all__ = ['ApplicationOwnerArgs', 'ApplicationOwner']

@pulumi.input_type
class ApplicationOwnerArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 owner_object_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ApplicationOwner resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] owner_object_id: The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "application_id", application_id)
        pulumi.set(__self__, "owner_object_id", owner_object_id)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="ownerObjectId")
    def owner_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "owner_object_id")

    @owner_object_id.setter
    def owner_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "owner_object_id", value)


@pulumi.input_type
class _ApplicationOwnerState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 owner_object_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationOwner resources.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] owner_object_id: The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if owner_object_id is not None:
            pulumi.set(__self__, "owner_object_id", owner_object_id)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="ownerObjectId")
    def owner_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "owner_object_id")

    @owner_object_id.setter
    def owner_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_object_id", value)


class ApplicationOwner(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 owner_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ApplicationRegistration("example", display_name="example")
        jane = azuread.User("jane",
            user_principal_name="jane.fischer@example.com",
            display_name="Jane Fischer",
            password="Ch@ngeMe")
        example_jane = azuread.ApplicationOwner("example_jane",
            application_id=example.id,
            owner_object_id=jane.object_id)
        ```

        > **Tip** For managing more application owners, create additional instances of this resource

        ## Import

        Application Owners can be imported using the object ID of the application and the object ID of the owner, in the following format.

        ```sh
        $ pulumi import azuread:index/applicationOwner:ApplicationOwner example /applications/00000000-0000-0000-0000-000000000000/owners/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] owner_object_id: The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationOwnerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ApplicationRegistration("example", display_name="example")
        jane = azuread.User("jane",
            user_principal_name="jane.fischer@example.com",
            display_name="Jane Fischer",
            password="Ch@ngeMe")
        example_jane = azuread.ApplicationOwner("example_jane",
            application_id=example.id,
            owner_object_id=jane.object_id)
        ```

        > **Tip** For managing more application owners, create additional instances of this resource

        ## Import

        Application Owners can be imported using the object ID of the application and the object ID of the owner, in the following format.

        ```sh
        $ pulumi import azuread:index/applicationOwner:ApplicationOwner example /applications/00000000-0000-0000-0000-000000000000/owners/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationOwnerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationOwnerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 owner_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationOwnerArgs.__new__(ApplicationOwnerArgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            if owner_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'owner_object_id'")
            __props__.__dict__["owner_object_id"] = owner_object_id
        super(ApplicationOwner, __self__).__init__(
            'azuread:index/applicationOwner:ApplicationOwner',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            owner_object_id: Optional[pulumi.Input[str]] = None) -> 'ApplicationOwner':
        """
        Get an existing ApplicationOwner resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] owner_object_id: The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationOwnerState.__new__(_ApplicationOwnerState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["owner_object_id"] = owner_object_id
        return ApplicationOwner(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="ownerObjectId")
    def owner_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "owner_object_id")

