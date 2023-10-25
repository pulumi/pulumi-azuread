# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationAppRoleInitArgs', 'ApplicationAppRole']

@pulumi.input_type
class ApplicationAppRoleInitArgs:
    def __init__(__self__, *,
                 allowed_member_types: pulumi.Input[Sequence[pulumi.Input[str]]],
                 application_id: pulumi.Input[str],
                 description: pulumi.Input[str],
                 display_name: pulumi.Input[str],
                 role_id: pulumi.Input[str],
                 value: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationAppRole resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_member_types: A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
        :param pulumi.Input[str] display_name: Display name for the app role that appears during app role assignment and in consent experiences.
        :param pulumi.Input[str] role_id: The unique identifier of the app role
        :param pulumi.Input[str] value: The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
               
               > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        ApplicationAppRoleInitArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            allowed_member_types=allowed_member_types,
            application_id=application_id,
            description=description,
            display_name=display_name,
            role_id=role_id,
            value=value,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             allowed_member_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             application_id: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             display_name: Optional[pulumi.Input[str]] = None,
             role_id: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if allowed_member_types is None and 'allowedMemberTypes' in kwargs:
            allowed_member_types = kwargs['allowedMemberTypes']
        if allowed_member_types is None:
            raise TypeError("Missing 'allowed_member_types' argument")
        if application_id is None and 'applicationId' in kwargs:
            application_id = kwargs['applicationId']
        if application_id is None:
            raise TypeError("Missing 'application_id' argument")
        if description is None:
            raise TypeError("Missing 'description' argument")
        if display_name is None and 'displayName' in kwargs:
            display_name = kwargs['displayName']
        if display_name is None:
            raise TypeError("Missing 'display_name' argument")
        if role_id is None and 'roleId' in kwargs:
            role_id = kwargs['roleId']
        if role_id is None:
            raise TypeError("Missing 'role_id' argument")

        _setter("allowed_member_types", allowed_member_types)
        _setter("application_id", application_id)
        _setter("description", description)
        _setter("display_name", display_name)
        _setter("role_id", role_id)
        if value is not None:
            _setter("value", value)

    @property
    @pulumi.getter(name="allowedMemberTypes")
    def allowed_member_types(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
        """
        return pulumi.get(self, "allowed_member_types")

    @allowed_member_types.setter
    def allowed_member_types(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_member_types", value)

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
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        Display name for the app role that appears during app role assignment and in consent experiences.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Input[str]:
        """
        The unique identifier of the app role
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.

        > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class _ApplicationAppRoleState:
    def __init__(__self__, *,
                 allowed_member_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationAppRole resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_member_types: A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
        :param pulumi.Input[str] display_name: Display name for the app role that appears during app role assignment and in consent experiences.
        :param pulumi.Input[str] role_id: The unique identifier of the app role
        :param pulumi.Input[str] value: The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
               
               > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        _ApplicationAppRoleState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            allowed_member_types=allowed_member_types,
            application_id=application_id,
            description=description,
            display_name=display_name,
            role_id=role_id,
            value=value,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             allowed_member_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             application_id: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             display_name: Optional[pulumi.Input[str]] = None,
             role_id: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if allowed_member_types is None and 'allowedMemberTypes' in kwargs:
            allowed_member_types = kwargs['allowedMemberTypes']
        if application_id is None and 'applicationId' in kwargs:
            application_id = kwargs['applicationId']
        if display_name is None and 'displayName' in kwargs:
            display_name = kwargs['displayName']
        if role_id is None and 'roleId' in kwargs:
            role_id = kwargs['roleId']

        if allowed_member_types is not None:
            _setter("allowed_member_types", allowed_member_types)
        if application_id is not None:
            _setter("application_id", application_id)
        if description is not None:
            _setter("description", description)
        if display_name is not None:
            _setter("display_name", display_name)
        if role_id is not None:
            _setter("role_id", role_id)
        if value is not None:
            _setter("value", value)

    @property
    @pulumi.getter(name="allowedMemberTypes")
    def allowed_member_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
        """
        return pulumi.get(self, "allowed_member_types")

    @allowed_member_types.setter
    def allowed_member_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_member_types", value)

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
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name for the app role that appears during app role assignment and in consent experiences.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the app role
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.

        > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class ApplicationAppRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_member_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread
        import pulumi_random as random

        example = azuread.ApplicationRegistration("example", display_name="example")
        example_administrator = random.RandomUuid("exampleAdministrator")
        example_administer = azuread.ApplicationAppRole("exampleAdminister",
            application_id=example.id,
            role_id=example_administrator.id,
            allowed_member_types=["User"],
            description="My role description",
            display_name="Administer",
            value="admin")
        ```

        > **Tip** For managing more app roles, create additional instances of this resource

        *Usage with Application resource*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Application("example", display_name="example")
        example_administer = azuread.ApplicationAppRole("exampleAdminister", application_id=example.id)
        # ...
        ```

        ## Import

        Application App Roles can be imported using the object ID of the application and the ID of the app role, in the following format.

        ```sh
         $ pulumi import azuread:index/applicationAppRole:ApplicationAppRole example /applications/00000000-0000-0000-0000-000000000000/appRoles/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_member_types: A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
        :param pulumi.Input[str] display_name: Display name for the app role that appears during app role assignment and in consent experiences.
        :param pulumi.Input[str] role_id: The unique identifier of the app role
        :param pulumi.Input[str] value: The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
               
               > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationAppRoleInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread
        import pulumi_random as random

        example = azuread.ApplicationRegistration("example", display_name="example")
        example_administrator = random.RandomUuid("exampleAdministrator")
        example_administer = azuread.ApplicationAppRole("exampleAdminister",
            application_id=example.id,
            role_id=example_administrator.id,
            allowed_member_types=["User"],
            description="My role description",
            display_name="Administer",
            value="admin")
        ```

        > **Tip** For managing more app roles, create additional instances of this resource

        *Usage with Application resource*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Application("example", display_name="example")
        example_administer = azuread.ApplicationAppRole("exampleAdminister", application_id=example.id)
        # ...
        ```

        ## Import

        Application App Roles can be imported using the object ID of the application and the ID of the app role, in the following format.

        ```sh
         $ pulumi import azuread:index/applicationAppRole:ApplicationAppRole example /applications/00000000-0000-0000-0000-000000000000/appRoles/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationAppRoleInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationAppRoleInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ApplicationAppRoleInitArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_member_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationAppRoleInitArgs.__new__(ApplicationAppRoleInitArgs)

            if allowed_member_types is None and not opts.urn:
                raise TypeError("Missing required property 'allowed_member_types'")
            __props__.__dict__["allowed_member_types"] = allowed_member_types
            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            if role_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_id'")
            __props__.__dict__["role_id"] = role_id
            __props__.__dict__["value"] = value
        super(ApplicationAppRole, __self__).__init__(
            'azuread:index/applicationAppRole:ApplicationAppRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_member_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            role_id: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'ApplicationAppRole':
        """
        Get an existing ApplicationAppRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_member_types: A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] description: Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
        :param pulumi.Input[str] display_name: Display name for the app role that appears during app role assignment and in consent experiences.
        :param pulumi.Input[str] role_id: The unique identifier of the app role
        :param pulumi.Input[str] value: The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
               
               > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationAppRoleState.__new__(_ApplicationAppRoleState)

        __props__.__dict__["allowed_member_types"] = allowed_member_types
        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["role_id"] = role_id
        __props__.__dict__["value"] = value
        return ApplicationAppRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedMemberTypes")
    def allowed_member_types(self) -> pulumi.Output[Sequence[str]]:
        """
        A set of values to specify whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications by setting to `Application`, or to both.
        """
        return pulumi.get(self, "allowed_member_types")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the app role that appears when the role is being assigned, and if the role functions as an application permissions, during the consent experiences.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Display name for the app role that appears during app role assignment and in consent experiences.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the app role
        """
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[Optional[str]]:
        """
        The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.

        > **Roles and Permission Scopes** In Azure Active Directory, application roles and permission scopes exported by an application share the same namespace and cannot contain duplicate values.
        """
        return pulumi.get(self, "value")
