# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationOAuth2PermissionArgs', 'ApplicationOAuth2Permission']

@pulumi.input_type
class ApplicationOAuth2PermissionArgs:
    def __init__(__self__, *,
                 admin_consent_description: pulumi.Input[str],
                 admin_consent_display_name: pulumi.Input[str],
                 application_object_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 user_consent_description: pulumi.Input[str],
                 user_consent_display_name: pulumi.Input[str],
                 value: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 permission_id: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationOAuth2Permission resource.
        :param pulumi.Input[str] admin_consent_description: Permission help text that appears in the admin consent and app assignment experiences.
        :param pulumi.Input[str] admin_consent_display_name: Display name for the permission that appears in the admin consent and app assignment experiences.
        :param pulumi.Input[str] application_object_id: The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] type: Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
        :param pulumi.Input[str] user_consent_description: Permission help text that appears in the end user consent experience.
        :param pulumi.Input[str] user_consent_display_name: Display name for the permission that appears in the end user consent experience.
        :param pulumi.Input[str] value: The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
        :param pulumi.Input[bool] is_enabled: Determines if the Permission is enabled. Defaults to `true`.
        :param pulumi.Input[str] permission_id: Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
        """
        pulumi.set(__self__, "admin_consent_description", admin_consent_description)
        pulumi.set(__self__, "admin_consent_display_name", admin_consent_display_name)
        pulumi.set(__self__, "application_object_id", application_object_id)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "user_consent_description", user_consent_description)
        pulumi.set(__self__, "user_consent_display_name", user_consent_display_name)
        pulumi.set(__self__, "value", value)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if is_enabled is not None:
            warnings.warn("""[NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider""", DeprecationWarning)
            pulumi.log.warn("""is_enabled is deprecated: [NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider""")
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if permission_id is not None:
            warnings.warn("""[NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider""", DeprecationWarning)
            pulumi.log.warn("""permission_id is deprecated: [NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider""")
        if permission_id is not None:
            pulumi.set(__self__, "permission_id", permission_id)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)

    @property
    @pulumi.getter(name="adminConsentDescription")
    def admin_consent_description(self) -> pulumi.Input[str]:
        """
        Permission help text that appears in the admin consent and app assignment experiences.
        """
        return pulumi.get(self, "admin_consent_description")

    @admin_consent_description.setter
    def admin_consent_description(self, value: pulumi.Input[str]):
        pulumi.set(self, "admin_consent_description", value)

    @property
    @pulumi.getter(name="adminConsentDisplayName")
    def admin_consent_display_name(self) -> pulumi.Input[str]:
        """
        Display name for the permission that appears in the admin consent and app assignment experiences.
        """
        return pulumi.get(self, "admin_consent_display_name")

    @admin_consent_display_name.setter
    def admin_consent_display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "admin_consent_display_name", value)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> pulumi.Input[str]:
        """
        The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_object_id")

    @application_object_id.setter
    def application_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_object_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userConsentDescription")
    def user_consent_description(self) -> pulumi.Input[str]:
        """
        Permission help text that appears in the end user consent experience.
        """
        return pulumi.get(self, "user_consent_description")

    @user_consent_description.setter
    def user_consent_description(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_consent_description", value)

    @property
    @pulumi.getter(name="userConsentDisplayName")
    def user_consent_display_name(self) -> pulumi.Input[str]:
        """
        Display name for the permission that appears in the end user consent experience.
        """
        return pulumi.get(self, "user_consent_display_name")

    @user_consent_display_name.setter
    def user_consent_display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_consent_display_name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if the Permission is enabled. Defaults to `true`.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="permissionId")
    def permission_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "permission_id")

    @permission_id.setter
    def permission_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permission_id", value)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_id", value)


@pulumi.input_type
class _ApplicationOAuth2PermissionState:
    def __init__(__self__, *,
                 admin_consent_description: Optional[pulumi.Input[str]] = None,
                 admin_consent_display_name: Optional[pulumi.Input[str]] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 permission_id: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_consent_description: Optional[pulumi.Input[str]] = None,
                 user_consent_display_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationOAuth2Permission resources.
        :param pulumi.Input[str] admin_consent_description: Permission help text that appears in the admin consent and app assignment experiences.
        :param pulumi.Input[str] admin_consent_display_name: Display name for the permission that appears in the admin consent and app assignment experiences.
        :param pulumi.Input[str] application_object_id: The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[bool] is_enabled: Determines if the Permission is enabled. Defaults to `true`.
        :param pulumi.Input[str] permission_id: Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] type: Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
        :param pulumi.Input[str] user_consent_description: Permission help text that appears in the end user consent experience.
        :param pulumi.Input[str] user_consent_display_name: Display name for the permission that appears in the end user consent experience.
        :param pulumi.Input[str] value: The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
        """
        if admin_consent_description is not None:
            pulumi.set(__self__, "admin_consent_description", admin_consent_description)
        if admin_consent_display_name is not None:
            pulumi.set(__self__, "admin_consent_display_name", admin_consent_display_name)
        if application_object_id is not None:
            pulumi.set(__self__, "application_object_id", application_object_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if is_enabled is not None:
            warnings.warn("""[NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider""", DeprecationWarning)
            pulumi.log.warn("""is_enabled is deprecated: [NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider""")
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if permission_id is not None:
            warnings.warn("""[NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider""", DeprecationWarning)
            pulumi.log.warn("""permission_id is deprecated: [NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider""")
        if permission_id is not None:
            pulumi.set(__self__, "permission_id", permission_id)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if user_consent_description is not None:
            pulumi.set(__self__, "user_consent_description", user_consent_description)
        if user_consent_display_name is not None:
            pulumi.set(__self__, "user_consent_display_name", user_consent_display_name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="adminConsentDescription")
    def admin_consent_description(self) -> Optional[pulumi.Input[str]]:
        """
        Permission help text that appears in the admin consent and app assignment experiences.
        """
        return pulumi.get(self, "admin_consent_description")

    @admin_consent_description.setter
    def admin_consent_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_consent_description", value)

    @property
    @pulumi.getter(name="adminConsentDisplayName")
    def admin_consent_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for the permission that appears in the admin consent and app assignment experiences.
        """
        return pulumi.get(self, "admin_consent_display_name")

    @admin_consent_display_name.setter
    def admin_consent_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_consent_display_name", value)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_object_id")

    @application_object_id.setter
    def application_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_object_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if the Permission is enabled. Defaults to `true`.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter(name="permissionId")
    def permission_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "permission_id")

    @permission_id.setter
    def permission_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permission_id", value)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="userConsentDescription")
    def user_consent_description(self) -> Optional[pulumi.Input[str]]:
        """
        Permission help text that appears in the end user consent experience.
        """
        return pulumi.get(self, "user_consent_description")

    @user_consent_description.setter
    def user_consent_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_consent_description", value)

    @property
    @pulumi.getter(name="userConsentDisplayName")
    def user_consent_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for the permission that appears in the end user consent experience.
        """
        return pulumi.get(self, "user_consent_display_name")

    @user_consent_display_name.setter
    def user_consent_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_consent_display_name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class ApplicationOAuth2Permission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_consent_description: Optional[pulumi.Input[str]] = None,
                 admin_consent_display_name: Optional[pulumi.Input[str]] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 permission_id: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_consent_description: Optional[pulumi.Input[str]] = None,
                 user_consent_display_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an OAuth2 Permission (also known as a Scope) associated with an Application within Azure Active Directory.

        > This resource is deprecated in favour of ApplicationOauth2PermissionScope and will be removed in version 2.0 of the provider.

        > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication")
        example_application_o_auth2_permission = azuread.ApplicationOAuth2Permission("exampleApplicationOAuth2Permission",
            application_object_id=example_application.id,
            admin_consent_description="Administer the application",
            admin_consent_display_name="Administer",
            is_enabled=True,
            type="User",
            user_consent_description="Administer the application",
            user_consent_display_name="Administer",
            value="administer")
        ```

        ## Import

        OAuth2 Permissions can be imported using the `object id` of an Application and the `id` of the Permission, e.g.

        ```sh
         $ pulumi import azuread:index/applicationOAuth2Permission:ApplicationOAuth2Permission test 00000000-0000-0000-0000-000000000000/scope/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_consent_description: Permission help text that appears in the admin consent and app assignment experiences.
        :param pulumi.Input[str] admin_consent_display_name: Display name for the permission that appears in the admin consent and app assignment experiences.
        :param pulumi.Input[str] application_object_id: The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[bool] is_enabled: Determines if the Permission is enabled. Defaults to `true`.
        :param pulumi.Input[str] permission_id: Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] type: Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
        :param pulumi.Input[str] user_consent_description: Permission help text that appears in the end user consent experience.
        :param pulumi.Input[str] user_consent_display_name: Display name for the permission that appears in the end user consent experience.
        :param pulumi.Input[str] value: The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationOAuth2PermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an OAuth2 Permission (also known as a Scope) associated with an Application within Azure Active Directory.

        > This resource is deprecated in favour of ApplicationOauth2PermissionScope and will be removed in version 2.0 of the provider.

        > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication")
        example_application_o_auth2_permission = azuread.ApplicationOAuth2Permission("exampleApplicationOAuth2Permission",
            application_object_id=example_application.id,
            admin_consent_description="Administer the application",
            admin_consent_display_name="Administer",
            is_enabled=True,
            type="User",
            user_consent_description="Administer the application",
            user_consent_display_name="Administer",
            value="administer")
        ```

        ## Import

        OAuth2 Permissions can be imported using the `object id` of an Application and the `id` of the Permission, e.g.

        ```sh
         $ pulumi import azuread:index/applicationOAuth2Permission:ApplicationOAuth2Permission test 00000000-0000-0000-0000-000000000000/scope/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationOAuth2PermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationOAuth2PermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_consent_description: Optional[pulumi.Input[str]] = None,
                 admin_consent_display_name: Optional[pulumi.Input[str]] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 permission_id: Optional[pulumi.Input[str]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 user_consent_description: Optional[pulumi.Input[str]] = None,
                 user_consent_display_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationOAuth2PermissionArgs.__new__(ApplicationOAuth2PermissionArgs)

            if admin_consent_description is None and not opts.urn:
                raise TypeError("Missing required property 'admin_consent_description'")
            __props__.__dict__["admin_consent_description"] = admin_consent_description
            if admin_consent_display_name is None and not opts.urn:
                raise TypeError("Missing required property 'admin_consent_display_name'")
            __props__.__dict__["admin_consent_display_name"] = admin_consent_display_name
            if application_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_object_id'")
            __props__.__dict__["application_object_id"] = application_object_id
            __props__.__dict__["enabled"] = enabled
            if is_enabled is not None and not opts.urn:
                warnings.warn("""[NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider""", DeprecationWarning)
                pulumi.log.warn("""is_enabled is deprecated: [NOTE] This attribute has been renamed to `enabled` and will be removed in version 2.0 of the AzureAD provider""")
            __props__.__dict__["is_enabled"] = is_enabled
            if permission_id is not None and not opts.urn:
                warnings.warn("""[NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider""", DeprecationWarning)
                pulumi.log.warn("""permission_id is deprecated: [NOTE] This attribute has been renamed to `scope_id` and will be removed in version 2.0 of the AzureAD provider""")
            __props__.__dict__["permission_id"] = permission_id
            __props__.__dict__["scope_id"] = scope_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if user_consent_description is None and not opts.urn:
                raise TypeError("Missing required property 'user_consent_description'")
            __props__.__dict__["user_consent_description"] = user_consent_description
            if user_consent_display_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_consent_display_name'")
            __props__.__dict__["user_consent_display_name"] = user_consent_display_name
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
        super(ApplicationOAuth2Permission, __self__).__init__(
            'azuread:index/applicationOAuth2Permission:ApplicationOAuth2Permission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_consent_description: Optional[pulumi.Input[str]] = None,
            admin_consent_display_name: Optional[pulumi.Input[str]] = None,
            application_object_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            is_enabled: Optional[pulumi.Input[bool]] = None,
            permission_id: Optional[pulumi.Input[str]] = None,
            scope_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            user_consent_description: Optional[pulumi.Input[str]] = None,
            user_consent_display_name: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'ApplicationOAuth2Permission':
        """
        Get an existing ApplicationOAuth2Permission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_consent_description: Permission help text that appears in the admin consent and app assignment experiences.
        :param pulumi.Input[str] admin_consent_display_name: Display name for the permission that appears in the admin consent and app assignment experiences.
        :param pulumi.Input[str] application_object_id: The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[bool] is_enabled: Determines if the Permission is enabled. Defaults to `true`.
        :param pulumi.Input[str] permission_id: Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] type: Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
        :param pulumi.Input[str] user_consent_description: Permission help text that appears in the end user consent experience.
        :param pulumi.Input[str] user_consent_display_name: Display name for the permission that appears in the end user consent experience.
        :param pulumi.Input[str] value: The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationOAuth2PermissionState.__new__(_ApplicationOAuth2PermissionState)

        __props__.__dict__["admin_consent_description"] = admin_consent_description
        __props__.__dict__["admin_consent_display_name"] = admin_consent_display_name
        __props__.__dict__["application_object_id"] = application_object_id
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["is_enabled"] = is_enabled
        __props__.__dict__["permission_id"] = permission_id
        __props__.__dict__["scope_id"] = scope_id
        __props__.__dict__["type"] = type
        __props__.__dict__["user_consent_description"] = user_consent_description
        __props__.__dict__["user_consent_display_name"] = user_consent_display_name
        __props__.__dict__["value"] = value
        return ApplicationOAuth2Permission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminConsentDescription")
    def admin_consent_description(self) -> pulumi.Output[str]:
        """
        Permission help text that appears in the admin consent and app assignment experiences.
        """
        return pulumi.get(self, "admin_consent_description")

    @property
    @pulumi.getter(name="adminConsentDisplayName")
    def admin_consent_display_name(self) -> pulumi.Output[str]:
        """
        Display name for the permission that appears in the admin consent and app assignment experiences.
        """
        return pulumi.get(self, "admin_consent_display_name")

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> pulumi.Output[str]:
        """
        The Object ID of the Application for which this Permission should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_object_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if the Permission is enabled. Defaults to `true`.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter(name="permissionId")
    def permission_id(self) -> pulumi.Output[str]:
        """
        Specifies a custom UUID for the Permission. If omitted, a random UUID will be automatically generated. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "permission_id")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "scope_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Specifies whether this scope permission can be consented to by an end user, or whether it is a tenant-wide permission that must be consented to by an Administrator. Possible values are "User" or "Admin".
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userConsentDescription")
    def user_consent_description(self) -> pulumi.Output[str]:
        """
        Permission help text that appears in the end user consent experience.
        """
        return pulumi.get(self, "user_consent_description")

    @property
    @pulumi.getter(name="userConsentDisplayName")
    def user_consent_display_name(self) -> pulumi.Output[str]:
        """
        Display name for the permission that appears in the end user consent experience.
        """
        return pulumi.get(self, "user_consent_display_name")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The value of the scope claim that the resource application should expect in the OAuth 2.0 access token.
        """
        return pulumi.get(self, "value")

