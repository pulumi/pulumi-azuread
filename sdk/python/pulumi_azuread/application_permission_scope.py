# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationPermissionScopeArgs', 'ApplicationPermissionScope']

@pulumi.input_type
class ApplicationPermissionScopeArgs:
    def __init__(__self__, *,
                 admin_consent_description: pulumi.Input[str],
                 admin_consent_display_name: pulumi.Input[str],
                 application_id: pulumi.Input[str],
                 scope_id: pulumi.Input[str],
                 value: pulumi.Input[str],
                 type: Optional[pulumi.Input[str]] = None,
                 user_consent_description: Optional[pulumi.Input[str]] = None,
                 user_consent_display_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationPermissionScope resource.
        :param pulumi.Input[str] admin_consent_description: Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
        :param pulumi.Input[str] admin_consent_display_name: Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scope_id: The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
        :param pulumi.Input[str] value: The value that is used for the `scp` claim in OAuth access tokens.
               
               > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        :param pulumi.Input[str] type: Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
        :param pulumi.Input[str] user_consent_description: Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
        :param pulumi.Input[str] user_consent_display_name: Display name for the delegated permission that appears in the end user consent experience
        """
        ApplicationPermissionScopeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            admin_consent_description=admin_consent_description,
            admin_consent_display_name=admin_consent_display_name,
            application_id=application_id,
            scope_id=scope_id,
            value=value,
            type=type,
            user_consent_description=user_consent_description,
            user_consent_display_name=user_consent_display_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             admin_consent_description: Optional[pulumi.Input[str]] = None,
             admin_consent_display_name: Optional[pulumi.Input[str]] = None,
             application_id: Optional[pulumi.Input[str]] = None,
             scope_id: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[str]] = None,
             type: Optional[pulumi.Input[str]] = None,
             user_consent_description: Optional[pulumi.Input[str]] = None,
             user_consent_display_name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if admin_consent_description is None and 'adminConsentDescription' in kwargs:
            admin_consent_description = kwargs['adminConsentDescription']
        if admin_consent_description is None:
            raise TypeError("Missing 'admin_consent_description' argument")
        if admin_consent_display_name is None and 'adminConsentDisplayName' in kwargs:
            admin_consent_display_name = kwargs['adminConsentDisplayName']
        if admin_consent_display_name is None:
            raise TypeError("Missing 'admin_consent_display_name' argument")
        if application_id is None and 'applicationId' in kwargs:
            application_id = kwargs['applicationId']
        if application_id is None:
            raise TypeError("Missing 'application_id' argument")
        if scope_id is None and 'scopeId' in kwargs:
            scope_id = kwargs['scopeId']
        if scope_id is None:
            raise TypeError("Missing 'scope_id' argument")
        if value is None:
            raise TypeError("Missing 'value' argument")
        if user_consent_description is None and 'userConsentDescription' in kwargs:
            user_consent_description = kwargs['userConsentDescription']
        if user_consent_display_name is None and 'userConsentDisplayName' in kwargs:
            user_consent_display_name = kwargs['userConsentDisplayName']

        _setter("admin_consent_description", admin_consent_description)
        _setter("admin_consent_display_name", admin_consent_display_name)
        _setter("application_id", application_id)
        _setter("scope_id", scope_id)
        _setter("value", value)
        if type is not None:
            _setter("type", type)
        if user_consent_description is not None:
            _setter("user_consent_description", user_consent_description)
        if user_consent_display_name is not None:
            _setter("user_consent_display_name", user_consent_display_name)

    @property
    @pulumi.getter(name="adminConsentDescription")
    def admin_consent_description(self) -> pulumi.Input[str]:
        """
        Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
        """
        return pulumi.get(self, "admin_consent_description")

    @admin_consent_description.setter
    def admin_consent_description(self, value: pulumi.Input[str]):
        pulumi.set(self, "admin_consent_description", value)

    @property
    @pulumi.getter(name="adminConsentDisplayName")
    def admin_consent_display_name(self) -> pulumi.Input[str]:
        """
        Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
        """
        return pulumi.get(self, "admin_consent_display_name")

    @admin_consent_display_name.setter
    def admin_consent_display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "admin_consent_display_name", value)

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
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Input[str]:
        """
        The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value that is used for the `scp` claim in OAuth access tokens.

        > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userConsentDescription")
    def user_consent_description(self) -> Optional[pulumi.Input[str]]:
        """
        Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
        """
        return pulumi.get(self, "user_consent_description")

    @user_consent_description.setter
    def user_consent_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_consent_description", value)

    @property
    @pulumi.getter(name="userConsentDisplayName")
    def user_consent_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for the delegated permission that appears in the end user consent experience
        """
        return pulumi.get(self, "user_consent_display_name")

    @user_consent_display_name.setter
    def user_consent_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_consent_display_name", value)


@pulumi.input_type
class _ApplicationPermissionScopeState:
    def __init__(__self__, *,
                 admin_consent_description: Optional[pulumi.Input[str]] = None,
                 admin_consent_display_name: Optional[pulumi.Input[str]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_consent_description: Optional[pulumi.Input[str]] = None,
                 user_consent_display_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationPermissionScope resources.
        :param pulumi.Input[str] admin_consent_description: Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
        :param pulumi.Input[str] admin_consent_display_name: Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scope_id: The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type: Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
        :param pulumi.Input[str] user_consent_description: Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
        :param pulumi.Input[str] user_consent_display_name: Display name for the delegated permission that appears in the end user consent experience
        :param pulumi.Input[str] value: The value that is used for the `scp` claim in OAuth access tokens.
               
               > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        _ApplicationPermissionScopeState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            admin_consent_description=admin_consent_description,
            admin_consent_display_name=admin_consent_display_name,
            application_id=application_id,
            scope_id=scope_id,
            type=type,
            user_consent_description=user_consent_description,
            user_consent_display_name=user_consent_display_name,
            value=value,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             admin_consent_description: Optional[pulumi.Input[str]] = None,
             admin_consent_display_name: Optional[pulumi.Input[str]] = None,
             application_id: Optional[pulumi.Input[str]] = None,
             scope_id: Optional[pulumi.Input[str]] = None,
             type: Optional[pulumi.Input[str]] = None,
             user_consent_description: Optional[pulumi.Input[str]] = None,
             user_consent_display_name: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if admin_consent_description is None and 'adminConsentDescription' in kwargs:
            admin_consent_description = kwargs['adminConsentDescription']
        if admin_consent_display_name is None and 'adminConsentDisplayName' in kwargs:
            admin_consent_display_name = kwargs['adminConsentDisplayName']
        if application_id is None and 'applicationId' in kwargs:
            application_id = kwargs['applicationId']
        if scope_id is None and 'scopeId' in kwargs:
            scope_id = kwargs['scopeId']
        if user_consent_description is None and 'userConsentDescription' in kwargs:
            user_consent_description = kwargs['userConsentDescription']
        if user_consent_display_name is None and 'userConsentDisplayName' in kwargs:
            user_consent_display_name = kwargs['userConsentDisplayName']

        if admin_consent_description is not None:
            _setter("admin_consent_description", admin_consent_description)
        if admin_consent_display_name is not None:
            _setter("admin_consent_display_name", admin_consent_display_name)
        if application_id is not None:
            _setter("application_id", application_id)
        if scope_id is not None:
            _setter("scope_id", scope_id)
        if type is not None:
            _setter("type", type)
        if user_consent_description is not None:
            _setter("user_consent_description", user_consent_description)
        if user_consent_display_name is not None:
            _setter("user_consent_display_name", user_consent_display_name)
        if value is not None:
            _setter("value", value)

    @property
    @pulumi.getter(name="adminConsentDescription")
    def admin_consent_description(self) -> Optional[pulumi.Input[str]]:
        """
        Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
        """
        return pulumi.get(self, "admin_consent_description")

    @admin_consent_description.setter
    def admin_consent_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_consent_description", value)

    @property
    @pulumi.getter(name="adminConsentDisplayName")
    def admin_consent_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
        """
        return pulumi.get(self, "admin_consent_display_name")

    @admin_consent_display_name.setter
    def admin_consent_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_consent_display_name", value)

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
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userConsentDescription")
    def user_consent_description(self) -> Optional[pulumi.Input[str]]:
        """
        Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
        """
        return pulumi.get(self, "user_consent_description")

    @user_consent_description.setter
    def user_consent_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_consent_description", value)

    @property
    @pulumi.getter(name="userConsentDisplayName")
    def user_consent_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for the delegated permission that appears in the end user consent experience
        """
        return pulumi.get(self, "user_consent_display_name")

    @user_consent_display_name.setter
    def user_consent_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_consent_display_name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value that is used for the `scp` claim in OAuth access tokens.

        > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class ApplicationPermissionScope(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_consent_description: Optional[pulumi.Input[str]] = None,
                 admin_consent_display_name: Optional[pulumi.Input[str]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_consent_description: Optional[pulumi.Input[str]] = None,
                 user_consent_display_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread
        import pulumi_random as random

        example_application_registration = azuread.ApplicationRegistration("exampleApplicationRegistration", display_name="example")
        example_administer = random.RandomUuid("exampleAdminister")
        example_application_permission_scope = azuread.ApplicationPermissionScope("exampleApplicationPermissionScope",
            application_id=azuread_application_registration["test"]["id"],
            scope_id=example_administer.id,
            value="administer",
            admin_consent_description="Administer the application",
            admin_consent_display_name="Administer")
        ```

        > **Tip** For managing more permissions scopes, create additional instances of this resource

        *Usage with Application resource*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication", display_name="example")
        example_application_permission_scope = azuread.ApplicationPermissionScope("exampleApplicationPermissionScope", application_id=example_application.id)
        # ...
        ```

        ## Import

        Application App Roles can be imported using the object ID of the application and the ID of the permission scope, in the following format.

        ```sh
         $ pulumi import azuread:index/applicationPermissionScope:ApplicationPermissionScope example /applications/00000000-0000-0000-0000-000000000000/permissionScopes/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_consent_description: Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
        :param pulumi.Input[str] admin_consent_display_name: Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scope_id: The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type: Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
        :param pulumi.Input[str] user_consent_description: Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
        :param pulumi.Input[str] user_consent_display_name: Display name for the delegated permission that appears in the end user consent experience
        :param pulumi.Input[str] value: The value that is used for the `scp` claim in OAuth access tokens.
               
               > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationPermissionScopeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread
        import pulumi_random as random

        example_application_registration = azuread.ApplicationRegistration("exampleApplicationRegistration", display_name="example")
        example_administer = random.RandomUuid("exampleAdminister")
        example_application_permission_scope = azuread.ApplicationPermissionScope("exampleApplicationPermissionScope",
            application_id=azuread_application_registration["test"]["id"],
            scope_id=example_administer.id,
            value="administer",
            admin_consent_description="Administer the application",
            admin_consent_display_name="Administer")
        ```

        > **Tip** For managing more permissions scopes, create additional instances of this resource

        *Usage with Application resource*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication", display_name="example")
        example_application_permission_scope = azuread.ApplicationPermissionScope("exampleApplicationPermissionScope", application_id=example_application.id)
        # ...
        ```

        ## Import

        Application App Roles can be imported using the object ID of the application and the ID of the permission scope, in the following format.

        ```sh
         $ pulumi import azuread:index/applicationPermissionScope:ApplicationPermissionScope example /applications/00000000-0000-0000-0000-000000000000/permissionScopes/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationPermissionScopeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationPermissionScopeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ApplicationPermissionScopeArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_consent_description: Optional[pulumi.Input[str]] = None,
                 admin_consent_display_name: Optional[pulumi.Input[str]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_consent_description: Optional[pulumi.Input[str]] = None,
                 user_consent_display_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationPermissionScopeArgs.__new__(ApplicationPermissionScopeArgs)

            if admin_consent_description is None and not opts.urn:
                raise TypeError("Missing required property 'admin_consent_description'")
            __props__.__dict__["admin_consent_description"] = admin_consent_description
            if admin_consent_display_name is None and not opts.urn:
                raise TypeError("Missing required property 'admin_consent_display_name'")
            __props__.__dict__["admin_consent_display_name"] = admin_consent_display_name
            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            if scope_id is None and not opts.urn:
                raise TypeError("Missing required property 'scope_id'")
            __props__.__dict__["scope_id"] = scope_id
            __props__.__dict__["type"] = type
            __props__.__dict__["user_consent_description"] = user_consent_description
            __props__.__dict__["user_consent_display_name"] = user_consent_display_name
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
        super(ApplicationPermissionScope, __self__).__init__(
            'azuread:index/applicationPermissionScope:ApplicationPermissionScope',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_consent_description: Optional[pulumi.Input[str]] = None,
            admin_consent_display_name: Optional[pulumi.Input[str]] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            scope_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            user_consent_description: Optional[pulumi.Input[str]] = None,
            user_consent_display_name: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'ApplicationPermissionScope':
        """
        Get an existing ApplicationPermissionScope resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_consent_description: Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
        :param pulumi.Input[str] admin_consent_display_name: Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] scope_id: The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type: Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
        :param pulumi.Input[str] user_consent_description: Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
        :param pulumi.Input[str] user_consent_display_name: Display name for the delegated permission that appears in the end user consent experience
        :param pulumi.Input[str] value: The value that is used for the `scp` claim in OAuth access tokens.
               
               > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationPermissionScopeState.__new__(_ApplicationPermissionScopeState)

        __props__.__dict__["admin_consent_description"] = admin_consent_description
        __props__.__dict__["admin_consent_display_name"] = admin_consent_display_name
        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["scope_id"] = scope_id
        __props__.__dict__["type"] = type
        __props__.__dict__["user_consent_description"] = user_consent_description
        __props__.__dict__["user_consent_display_name"] = user_consent_display_name
        __props__.__dict__["value"] = value
        return ApplicationPermissionScope(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminConsentDescription")
    def admin_consent_description(self) -> pulumi.Output[str]:
        """
        Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
        """
        return pulumi.get(self, "admin_consent_description")

    @property
    @pulumi.getter(name="adminConsentDisplayName")
    def admin_consent_display_name(self) -> pulumi.Output[str]:
        """
        Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
        """
        return pulumi.get(self, "admin_consent_display_name")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the permission scope. Must be a valid UUID. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "scope_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userConsentDescription")
    def user_consent_description(self) -> pulumi.Output[Optional[str]]:
        """
        Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
        """
        return pulumi.get(self, "user_consent_description")

    @property
    @pulumi.getter(name="userConsentDisplayName")
    def user_consent_display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Display name for the delegated permission that appears in the end user consent experience
        """
        return pulumi.get(self, "user_consent_display_name")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The value that is used for the `scp` claim in OAuth access tokens.

        > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        return pulumi.get(self, "value")

