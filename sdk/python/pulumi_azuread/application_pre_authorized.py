# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationPreAuthorizedArgs', 'ApplicationPreAuthorized']

@pulumi.input_type
class ApplicationPreAuthorizedArgs:
    def __init__(__self__, *,
                 application_object_id: pulumi.Input[str],
                 authorized_app_id: pulumi.Input[str],
                 permission_ids: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a ApplicationPreAuthorized resource.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] authorized_app_id: The application ID of the pre-authorized application
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permission_ids: A set of permission scope IDs required by the authorized application.
        """
        ApplicationPreAuthorizedArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            application_object_id=application_object_id,
            authorized_app_id=authorized_app_id,
            permission_ids=permission_ids,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             application_object_id: Optional[pulumi.Input[str]] = None,
             authorized_app_id: Optional[pulumi.Input[str]] = None,
             permission_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if application_object_id is None and 'applicationObjectId' in kwargs:
            application_object_id = kwargs['applicationObjectId']
        if application_object_id is None:
            raise TypeError("Missing 'application_object_id' argument")
        if authorized_app_id is None and 'authorizedAppId' in kwargs:
            authorized_app_id = kwargs['authorizedAppId']
        if authorized_app_id is None:
            raise TypeError("Missing 'authorized_app_id' argument")
        if permission_ids is None and 'permissionIds' in kwargs:
            permission_ids = kwargs['permissionIds']
        if permission_ids is None:
            raise TypeError("Missing 'permission_ids' argument")

        _setter("application_object_id", application_object_id)
        _setter("authorized_app_id", authorized_app_id)
        _setter("permission_ids", permission_ids)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_object_id")

    @application_object_id.setter
    def application_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_object_id", value)

    @property
    @pulumi.getter(name="authorizedAppId")
    def authorized_app_id(self) -> pulumi.Input[str]:
        """
        The application ID of the pre-authorized application
        """
        return pulumi.get(self, "authorized_app_id")

    @authorized_app_id.setter
    def authorized_app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorized_app_id", value)

    @property
    @pulumi.getter(name="permissionIds")
    def permission_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A set of permission scope IDs required by the authorized application.
        """
        return pulumi.get(self, "permission_ids")

    @permission_ids.setter
    def permission_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "permission_ids", value)


@pulumi.input_type
class _ApplicationPreAuthorizedState:
    def __init__(__self__, *,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 authorized_app_id: Optional[pulumi.Input[str]] = None,
                 permission_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ApplicationPreAuthorized resources.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] authorized_app_id: The application ID of the pre-authorized application
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permission_ids: A set of permission scope IDs required by the authorized application.
        """
        _ApplicationPreAuthorizedState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            application_object_id=application_object_id,
            authorized_app_id=authorized_app_id,
            permission_ids=permission_ids,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             application_object_id: Optional[pulumi.Input[str]] = None,
             authorized_app_id: Optional[pulumi.Input[str]] = None,
             permission_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if application_object_id is None and 'applicationObjectId' in kwargs:
            application_object_id = kwargs['applicationObjectId']
        if authorized_app_id is None and 'authorizedAppId' in kwargs:
            authorized_app_id = kwargs['authorizedAppId']
        if permission_ids is None and 'permissionIds' in kwargs:
            permission_ids = kwargs['permissionIds']

        if application_object_id is not None:
            _setter("application_object_id", application_object_id)
        if authorized_app_id is not None:
            _setter("authorized_app_id", authorized_app_id)
        if permission_ids is not None:
            _setter("permission_ids", permission_ids)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_object_id")

    @application_object_id.setter
    def application_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_object_id", value)

    @property
    @pulumi.getter(name="authorizedAppId")
    def authorized_app_id(self) -> Optional[pulumi.Input[str]]:
        """
        The application ID of the pre-authorized application
        """
        return pulumi.get(self, "authorized_app_id")

    @authorized_app_id.setter
    def authorized_app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorized_app_id", value)

    @property
    @pulumi.getter(name="permissionIds")
    def permission_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of permission scope IDs required by the authorized application.
        """
        return pulumi.get(self, "permission_ids")

    @permission_ids.setter
    def permission_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "permission_ids", value)


class ApplicationPreAuthorized(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 authorized_app_id: Optional[pulumi.Input[str]] = None,
                 permission_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        authorized = azuread.Application("authorized", display_name="example-authorized-app")
        authorizer = azuread.Application("authorizer",
            display_name="example-authorizing-app",
            api=azuread.ApplicationApiArgs(
                oauth2_permission_scopes=[
                    azuread.ApplicationApiOauth2PermissionScopeArgs(
                        admin_consent_description="Administer the application",
                        admin_consent_display_name="Administer",
                        enabled=True,
                        id="ced9c4c3-c273-4f0f-ac71-a20377b90f9c",
                        type="Admin",
                        value="administer",
                    ),
                    azuread.ApplicationApiOauth2PermissionScopeArgs(
                        admin_consent_description="Access the application",
                        admin_consent_display_name="Access",
                        enabled=True,
                        id="2d5e07ca-664d-4d9b-ad61-ec07fd215213",
                        type="User",
                        user_consent_description="Access the application",
                        user_consent_display_name="Access",
                        value="user_impersonation",
                    ),
                ],
            ))
        example = azuread.ApplicationPreAuthorized("example",
            application_object_id=authorizer.object_id,
            authorized_app_id=authorized.application_id,
            permission_ids=[
                "ced9c4c3-c273-4f0f-ac71-a20377b90f9c",
                "2d5e07ca-664d-4d9b-ad61-ec07fd215213",
            ])
        ```

        ## Import

        Pre-authorized applications can be imported using the object ID of the authorizing application and the application ID of the application being authorized, e.g.

        ```sh
         $ pulumi import azuread:index/applicationPreAuthorized:ApplicationPreAuthorized example 00000000-0000-0000-0000-000000000000/preAuthorizedApplication/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the authorizing application's object ID, the string "preAuthorizedApplication" and the authorized application's application ID (client ID) in the format `{ObjectId}/preAuthorizedApplication/{ApplicationId}`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] authorized_app_id: The application ID of the pre-authorized application
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permission_ids: A set of permission scope IDs required by the authorized application.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationPreAuthorizedArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        authorized = azuread.Application("authorized", display_name="example-authorized-app")
        authorizer = azuread.Application("authorizer",
            display_name="example-authorizing-app",
            api=azuread.ApplicationApiArgs(
                oauth2_permission_scopes=[
                    azuread.ApplicationApiOauth2PermissionScopeArgs(
                        admin_consent_description="Administer the application",
                        admin_consent_display_name="Administer",
                        enabled=True,
                        id="ced9c4c3-c273-4f0f-ac71-a20377b90f9c",
                        type="Admin",
                        value="administer",
                    ),
                    azuread.ApplicationApiOauth2PermissionScopeArgs(
                        admin_consent_description="Access the application",
                        admin_consent_display_name="Access",
                        enabled=True,
                        id="2d5e07ca-664d-4d9b-ad61-ec07fd215213",
                        type="User",
                        user_consent_description="Access the application",
                        user_consent_display_name="Access",
                        value="user_impersonation",
                    ),
                ],
            ))
        example = azuread.ApplicationPreAuthorized("example",
            application_object_id=authorizer.object_id,
            authorized_app_id=authorized.application_id,
            permission_ids=[
                "ced9c4c3-c273-4f0f-ac71-a20377b90f9c",
                "2d5e07ca-664d-4d9b-ad61-ec07fd215213",
            ])
        ```

        ## Import

        Pre-authorized applications can be imported using the object ID of the authorizing application and the application ID of the application being authorized, e.g.

        ```sh
         $ pulumi import azuread:index/applicationPreAuthorized:ApplicationPreAuthorized example 00000000-0000-0000-0000-000000000000/preAuthorizedApplication/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the authorizing application's object ID, the string "preAuthorizedApplication" and the authorized application's application ID (client ID) in the format `{ObjectId}/preAuthorizedApplication/{ApplicationId}`.

        :param str resource_name: The name of the resource.
        :param ApplicationPreAuthorizedArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationPreAuthorizedArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ApplicationPreAuthorizedArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 authorized_app_id: Optional[pulumi.Input[str]] = None,
                 permission_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationPreAuthorizedArgs.__new__(ApplicationPreAuthorizedArgs)

            if application_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_object_id'")
            __props__.__dict__["application_object_id"] = application_object_id
            if authorized_app_id is None and not opts.urn:
                raise TypeError("Missing required property 'authorized_app_id'")
            __props__.__dict__["authorized_app_id"] = authorized_app_id
            if permission_ids is None and not opts.urn:
                raise TypeError("Missing required property 'permission_ids'")
            __props__.__dict__["permission_ids"] = permission_ids
        super(ApplicationPreAuthorized, __self__).__init__(
            'azuread:index/applicationPreAuthorized:ApplicationPreAuthorized',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_object_id: Optional[pulumi.Input[str]] = None,
            authorized_app_id: Optional[pulumi.Input[str]] = None,
            permission_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ApplicationPreAuthorized':
        """
        Get an existing ApplicationPreAuthorized resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] authorized_app_id: The application ID of the pre-authorized application
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permission_ids: A set of permission scope IDs required by the authorized application.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationPreAuthorizedState.__new__(_ApplicationPreAuthorizedState)

        __props__.__dict__["application_object_id"] = application_object_id
        __props__.__dict__["authorized_app_id"] = authorized_app_id
        __props__.__dict__["permission_ids"] = permission_ids
        return ApplicationPreAuthorized(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_object_id")

    @property
    @pulumi.getter(name="authorizedAppId")
    def authorized_app_id(self) -> pulumi.Output[str]:
        """
        The application ID of the pre-authorized application
        """
        return pulumi.get(self, "authorized_app_id")

    @property
    @pulumi.getter(name="permissionIds")
    def permission_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        A set of permission scope IDs required by the authorized application.
        """
        return pulumi.get(self, "permission_ids")

