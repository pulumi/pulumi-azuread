# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationApiAccessArgs', 'ApplicationApiAccess']

@pulumi.input_type
class ApplicationApiAccessArgs:
    def __init__(__self__, *,
                 api_client_id: pulumi.Input[str],
                 application_id: pulumi.Input[str],
                 role_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scope_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ApplicationApiAccess resource.
        :param pulumi.Input[str] api_client_id: The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_ids: A set of role IDs to be granted to the application, as published by the API.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scope_ids: A set of scope IDs to be granted to the application, as published by the API.
               
               > At least one of `role_ids` or `scope_ids` must be specified.
        """
        pulumi.set(__self__, "api_client_id", api_client_id)
        pulumi.set(__self__, "application_id", application_id)
        if role_ids is not None:
            pulumi.set(__self__, "role_ids", role_ids)
        if scope_ids is not None:
            pulumi.set(__self__, "scope_ids", scope_ids)

    @property
    @pulumi.getter(name="apiClientId")
    def api_client_id(self) -> pulumi.Input[str]:
        """
        The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_client_id")

    @api_client_id.setter
    def api_client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_client_id", value)

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
    @pulumi.getter(name="roleIds")
    def role_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of role IDs to be granted to the application, as published by the API.
        """
        return pulumi.get(self, "role_ids")

    @role_ids.setter
    def role_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "role_ids", value)

    @property
    @pulumi.getter(name="scopeIds")
    def scope_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of scope IDs to be granted to the application, as published by the API.

        > At least one of `role_ids` or `scope_ids` must be specified.
        """
        return pulumi.get(self, "scope_ids")

    @scope_ids.setter
    def scope_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scope_ids", value)


@pulumi.input_type
class _ApplicationApiAccessState:
    def __init__(__self__, *,
                 api_client_id: Optional[pulumi.Input[str]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 role_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scope_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ApplicationApiAccess resources.
        :param pulumi.Input[str] api_client_id: The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_ids: A set of role IDs to be granted to the application, as published by the API.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scope_ids: A set of scope IDs to be granted to the application, as published by the API.
               
               > At least one of `role_ids` or `scope_ids` must be specified.
        """
        if api_client_id is not None:
            pulumi.set(__self__, "api_client_id", api_client_id)
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if role_ids is not None:
            pulumi.set(__self__, "role_ids", role_ids)
        if scope_ids is not None:
            pulumi.set(__self__, "scope_ids", scope_ids)

    @property
    @pulumi.getter(name="apiClientId")
    def api_client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_client_id")

    @api_client_id.setter
    def api_client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_client_id", value)

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
    @pulumi.getter(name="roleIds")
    def role_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of role IDs to be granted to the application, as published by the API.
        """
        return pulumi.get(self, "role_ids")

    @role_ids.setter
    def role_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "role_ids", value)

    @property
    @pulumi.getter(name="scopeIds")
    def scope_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of scope IDs to be granted to the application, as published by the API.

        > At least one of `role_ids` or `scope_ids` must be specified.
        """
        return pulumi.get(self, "scope_ids")

    @scope_ids.setter
    def scope_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scope_ids", value)


class ApplicationApiAccess(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_client_id: Optional[pulumi.Input[str]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 role_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scope_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        well_known = azuread.get_application_published_app_ids()
        msgraph = azuread.get_service_principal(client_id=well_known.result["MicrosoftGraph"])
        example = azuread.ApplicationRegistration("example", display_name="example")
        example_msgraph = azuread.ApplicationApiAccess("example_msgraph",
            application_id=example.id,
            api_client_id=well_known.result["MicrosoftGraph"],
            role_ids=[
                msgraph.app_role_ids["Group.Read.All"],
                msgraph.app_role_ids["User.Read.All"],
            ],
            scope_ids=[msgraph.oauth2_permission_scope_ids["User.ReadWrite"]])
        ```
        <!--End PulumiCodeChooser -->

        > **Tip** For managing permissions for an additional API, create another instance of this resource

        *Usage with Application resource*

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Application("example", display_name="example")
        example_application_api_access = azuread.ApplicationApiAccess("example", application_id=example.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Application API Access can be imported using the object ID of the application and the client ID of the API, in the following format.

        ```sh
        $ pulumi import azuread:index/applicationApiAccess:ApplicationApiAccess example /applications/00000000-0000-0000-0000-000000000000/apiAccess/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_client_id: The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_ids: A set of role IDs to be granted to the application, as published by the API.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scope_ids: A set of scope IDs to be granted to the application, as published by the API.
               
               > At least one of `role_ids` or `scope_ids` must be specified.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationApiAccessArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        well_known = azuread.get_application_published_app_ids()
        msgraph = azuread.get_service_principal(client_id=well_known.result["MicrosoftGraph"])
        example = azuread.ApplicationRegistration("example", display_name="example")
        example_msgraph = azuread.ApplicationApiAccess("example_msgraph",
            application_id=example.id,
            api_client_id=well_known.result["MicrosoftGraph"],
            role_ids=[
                msgraph.app_role_ids["Group.Read.All"],
                msgraph.app_role_ids["User.Read.All"],
            ],
            scope_ids=[msgraph.oauth2_permission_scope_ids["User.ReadWrite"]])
        ```
        <!--End PulumiCodeChooser -->

        > **Tip** For managing permissions for an additional API, create another instance of this resource

        *Usage with Application resource*

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Application("example", display_name="example")
        example_application_api_access = azuread.ApplicationApiAccess("example", application_id=example.id)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Application API Access can be imported using the object ID of the application and the client ID of the API, in the following format.

        ```sh
        $ pulumi import azuread:index/applicationApiAccess:ApplicationApiAccess example /applications/00000000-0000-0000-0000-000000000000/apiAccess/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationApiAccessArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationApiAccessArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_client_id: Optional[pulumi.Input[str]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 role_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 scope_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationApiAccessArgs.__new__(ApplicationApiAccessArgs)

            if api_client_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_client_id'")
            __props__.__dict__["api_client_id"] = api_client_id
            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            __props__.__dict__["role_ids"] = role_ids
            __props__.__dict__["scope_ids"] = scope_ids
        super(ApplicationApiAccess, __self__).__init__(
            'azuread:index/applicationApiAccess:ApplicationApiAccess',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_client_id: Optional[pulumi.Input[str]] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            role_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            scope_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ApplicationApiAccess':
        """
        Get an existing ApplicationApiAccess resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_client_id: The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] role_ids: A set of role IDs to be granted to the application, as published by the API.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scope_ids: A set of scope IDs to be granted to the application, as published by the API.
               
               > At least one of `role_ids` or `scope_ids` must be specified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationApiAccessState.__new__(_ApplicationApiAccessState)

        __props__.__dict__["api_client_id"] = api_client_id
        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["role_ids"] = role_ids
        __props__.__dict__["scope_ids"] = scope_ids
        return ApplicationApiAccess(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiClientId")
    def api_client_id(self) -> pulumi.Output[str]:
        """
        The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "api_client_id")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="roleIds")
    def role_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of role IDs to be granted to the application, as published by the API.
        """
        return pulumi.get(self, "role_ids")

    @property
    @pulumi.getter(name="scopeIds")
    def scope_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of scope IDs to be granted to the application, as published by the API.

        > At least one of `role_ids` or `scope_ids` must be specified.
        """
        return pulumi.get(self, "scope_ids")

