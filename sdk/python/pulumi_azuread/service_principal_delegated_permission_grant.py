# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServicePrincipalDelegatedPermissionGrantArgs', 'ServicePrincipalDelegatedPermissionGrant']

@pulumi.input_type
class ServicePrincipalDelegatedPermissionGrantArgs:
    def __init__(__self__, *,
                 claim_values: pulumi.Input[Sequence[pulumi.Input[str]]],
                 resource_service_principal_object_id: pulumi.Input[str],
                 service_principal_object_id: pulumi.Input[str],
                 user_object_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServicePrincipalDelegatedPermissionGrant resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] claim_values: A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
        :param pulumi.Input[str] resource_service_principal_object_id: The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_principal_object_id: The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_object_id: The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
               
               > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
        """
        pulumi.set(__self__, "claim_values", claim_values)
        pulumi.set(__self__, "resource_service_principal_object_id", resource_service_principal_object_id)
        pulumi.set(__self__, "service_principal_object_id", service_principal_object_id)
        if user_object_id is not None:
            pulumi.set(__self__, "user_object_id", user_object_id)

    @property
    @pulumi.getter(name="claimValues")
    def claim_values(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
        """
        return pulumi.get(self, "claim_values")

    @claim_values.setter
    def claim_values(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "claim_values", value)

    @property
    @pulumi.getter(name="resourceServicePrincipalObjectId")
    def resource_service_principal_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_service_principal_object_id")

    @resource_service_principal_object_id.setter
    def resource_service_principal_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_service_principal_object_id", value)

    @property
    @pulumi.getter(name="servicePrincipalObjectId")
    def service_principal_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_object_id")

    @service_principal_object_id.setter
    def service_principal_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_principal_object_id", value)

    @property
    @pulumi.getter(name="userObjectId")
    def user_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.

        > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
        """
        return pulumi.get(self, "user_object_id")

    @user_object_id.setter
    def user_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_object_id", value)


@pulumi.input_type
class _ServicePrincipalDelegatedPermissionGrantState:
    def __init__(__self__, *,
                 claim_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_service_principal_object_id: Optional[pulumi.Input[str]] = None,
                 service_principal_object_id: Optional[pulumi.Input[str]] = None,
                 user_object_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServicePrincipalDelegatedPermissionGrant resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] claim_values: A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
        :param pulumi.Input[str] resource_service_principal_object_id: The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_principal_object_id: The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_object_id: The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
               
               > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
        """
        if claim_values is not None:
            pulumi.set(__self__, "claim_values", claim_values)
        if resource_service_principal_object_id is not None:
            pulumi.set(__self__, "resource_service_principal_object_id", resource_service_principal_object_id)
        if service_principal_object_id is not None:
            pulumi.set(__self__, "service_principal_object_id", service_principal_object_id)
        if user_object_id is not None:
            pulumi.set(__self__, "user_object_id", user_object_id)

    @property
    @pulumi.getter(name="claimValues")
    def claim_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
        """
        return pulumi.get(self, "claim_values")

    @claim_values.setter
    def claim_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "claim_values", value)

    @property
    @pulumi.getter(name="resourceServicePrincipalObjectId")
    def resource_service_principal_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_service_principal_object_id")

    @resource_service_principal_object_id.setter
    def resource_service_principal_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_service_principal_object_id", value)

    @property
    @pulumi.getter(name="servicePrincipalObjectId")
    def service_principal_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_object_id")

    @service_principal_object_id.setter
    def service_principal_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_principal_object_id", value)

    @property
    @pulumi.getter(name="userObjectId")
    def user_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.

        > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
        """
        return pulumi.get(self, "user_object_id")

    @user_object_id.setter
    def user_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_object_id", value)


class ServicePrincipalDelegatedPermissionGrant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 claim_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_service_principal_object_id: Optional[pulumi.Input[str]] = None,
                 service_principal_object_id: Optional[pulumi.Input[str]] = None,
                 user_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a delegated permission grant for a service principal, on behalf of a single user, or all users.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one the following directory role: `Global Administrator`

        ## Example Usage

        *Delegated permission grant for all users*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        well_known = azuread.get_application_published_app_ids()
        msgraph = azuread.ServicePrincipal("msgraph",
            client_id=well_known.result["microsoftGraph"],
            use_existing=True)
        example = azuread.Application("example",
            display_name="example",
            required_resource_accesses=[azuread.ApplicationRequiredResourceAccessArgs(
                resource_app_id=well_known.result["microsoftGraph"],
                resource_accesses=[
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.oauth2_permission_scope_ids["openid"],
                        type="Scope",
                    ),
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.oauth2_permission_scope_ids["User.Read"],
                        type="Scope",
                    ),
                ],
            )])
        example_service_principal = azuread.ServicePrincipal("example", client_id=example.application_id)
        example_service_principal_delegated_permission_grant = azuread.ServicePrincipalDelegatedPermissionGrant("example",
            service_principal_object_id=example_service_principal.object_id,
            resource_service_principal_object_id=msgraph.object_id,
            claim_values=[
                "openid",
                "User.Read.All",
            ])
        ```

        *Delegated permission grant for a single user*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        well_known = azuread.get_application_published_app_ids()
        msgraph = azuread.ServicePrincipal("msgraph",
            client_id=well_known.result["microsoftGraph"],
            use_existing=True)
        example = azuread.Application("example",
            display_name="example",
            required_resource_accesses=[azuread.ApplicationRequiredResourceAccessArgs(
                resource_app_id=well_known.result["microsoftGraph"],
                resource_accesses=[
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.oauth2_permission_scope_ids["openid"],
                        type="Scope",
                    ),
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.oauth2_permission_scope_ids["User.Read"],
                        type="Scope",
                    ),
                ],
            )])
        example_service_principal = azuread.ServicePrincipal("example", client_id=example.application_id)
        example_user = azuread.User("example",
            display_name="J. Doe",
            user_principal_name="jdoe@example.com",
            mail_nickname="jdoe",
            password="SecretP@sswd99!")
        example_service_principal_delegated_permission_grant = azuread.ServicePrincipalDelegatedPermissionGrant("example",
            service_principal_object_id=example_service_principal.object_id,
            resource_service_principal_object_id=msgraph.object_id,
            claim_values=[
                "openid",
                "User.Read.All",
            ],
            user_object_id=example_user.object_id)
        ```

        ## Import

        Delegated permission grants can be imported using their ID, e.g.

        ```sh
        $ pulumi import azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant example aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] claim_values: A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
        :param pulumi.Input[str] resource_service_principal_object_id: The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_principal_object_id: The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_object_id: The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
               
               > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServicePrincipalDelegatedPermissionGrantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a delegated permission grant for a service principal, on behalf of a single user, or all users.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires the following application role: `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one the following directory role: `Global Administrator`

        ## Example Usage

        *Delegated permission grant for all users*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        well_known = azuread.get_application_published_app_ids()
        msgraph = azuread.ServicePrincipal("msgraph",
            client_id=well_known.result["microsoftGraph"],
            use_existing=True)
        example = azuread.Application("example",
            display_name="example",
            required_resource_accesses=[azuread.ApplicationRequiredResourceAccessArgs(
                resource_app_id=well_known.result["microsoftGraph"],
                resource_accesses=[
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.oauth2_permission_scope_ids["openid"],
                        type="Scope",
                    ),
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.oauth2_permission_scope_ids["User.Read"],
                        type="Scope",
                    ),
                ],
            )])
        example_service_principal = azuread.ServicePrincipal("example", client_id=example.application_id)
        example_service_principal_delegated_permission_grant = azuread.ServicePrincipalDelegatedPermissionGrant("example",
            service_principal_object_id=example_service_principal.object_id,
            resource_service_principal_object_id=msgraph.object_id,
            claim_values=[
                "openid",
                "User.Read.All",
            ])
        ```

        *Delegated permission grant for a single user*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        well_known = azuread.get_application_published_app_ids()
        msgraph = azuread.ServicePrincipal("msgraph",
            client_id=well_known.result["microsoftGraph"],
            use_existing=True)
        example = azuread.Application("example",
            display_name="example",
            required_resource_accesses=[azuread.ApplicationRequiredResourceAccessArgs(
                resource_app_id=well_known.result["microsoftGraph"],
                resource_accesses=[
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.oauth2_permission_scope_ids["openid"],
                        type="Scope",
                    ),
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.oauth2_permission_scope_ids["User.Read"],
                        type="Scope",
                    ),
                ],
            )])
        example_service_principal = azuread.ServicePrincipal("example", client_id=example.application_id)
        example_user = azuread.User("example",
            display_name="J. Doe",
            user_principal_name="jdoe@example.com",
            mail_nickname="jdoe",
            password="SecretP@sswd99!")
        example_service_principal_delegated_permission_grant = azuread.ServicePrincipalDelegatedPermissionGrant("example",
            service_principal_object_id=example_service_principal.object_id,
            resource_service_principal_object_id=msgraph.object_id,
            claim_values=[
                "openid",
                "User.Read.All",
            ],
            user_object_id=example_user.object_id)
        ```

        ## Import

        Delegated permission grants can be imported using their ID, e.g.

        ```sh
        $ pulumi import azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant example aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
        ```

        :param str resource_name: The name of the resource.
        :param ServicePrincipalDelegatedPermissionGrantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServicePrincipalDelegatedPermissionGrantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 claim_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_service_principal_object_id: Optional[pulumi.Input[str]] = None,
                 service_principal_object_id: Optional[pulumi.Input[str]] = None,
                 user_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServicePrincipalDelegatedPermissionGrantArgs.__new__(ServicePrincipalDelegatedPermissionGrantArgs)

            if claim_values is None and not opts.urn:
                raise TypeError("Missing required property 'claim_values'")
            __props__.__dict__["claim_values"] = claim_values
            if resource_service_principal_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_service_principal_object_id'")
            __props__.__dict__["resource_service_principal_object_id"] = resource_service_principal_object_id
            if service_principal_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_principal_object_id'")
            __props__.__dict__["service_principal_object_id"] = service_principal_object_id
            __props__.__dict__["user_object_id"] = user_object_id
        super(ServicePrincipalDelegatedPermissionGrant, __self__).__init__(
            'azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            claim_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_service_principal_object_id: Optional[pulumi.Input[str]] = None,
            service_principal_object_id: Optional[pulumi.Input[str]] = None,
            user_object_id: Optional[pulumi.Input[str]] = None) -> 'ServicePrincipalDelegatedPermissionGrant':
        """
        Get an existing ServicePrincipalDelegatedPermissionGrant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] claim_values: A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
        :param pulumi.Input[str] resource_service_principal_object_id: The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_principal_object_id: The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] user_object_id: The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
               
               > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServicePrincipalDelegatedPermissionGrantState.__new__(_ServicePrincipalDelegatedPermissionGrantState)

        __props__.__dict__["claim_values"] = claim_values
        __props__.__dict__["resource_service_principal_object_id"] = resource_service_principal_object_id
        __props__.__dict__["service_principal_object_id"] = service_principal_object_id
        __props__.__dict__["user_object_id"] = user_object_id
        return ServicePrincipalDelegatedPermissionGrant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="claimValues")
    def claim_values(self) -> pulumi.Output[Sequence[str]]:
        """
        A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
        """
        return pulumi.get(self, "claim_values")

    @property
    @pulumi.getter(name="resourceServicePrincipalObjectId")
    def resource_service_principal_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_service_principal_object_id")

    @property
    @pulumi.getter(name="servicePrincipalObjectId")
    def service_principal_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_object_id")

    @property
    @pulumi.getter(name="userObjectId")
    def user_object_id(self) -> pulumi.Output[Optional[str]]:
        """
        The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.

        > **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
        """
        return pulumi.get(self, "user_object_id")

