# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppRoleAssignmentArgs', 'AppRoleAssignment']

@pulumi.input_type
class AppRoleAssignmentArgs:
    def __init__(__self__, *,
                 app_role_id: pulumi.Input[str],
                 principal_object_id: pulumi.Input[str],
                 resource_object_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a AppRoleAssignment resource.
        :param pulumi.Input[str] app_role_id: The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_object_id: The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_object_id: The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "app_role_id", app_role_id)
        pulumi.set(__self__, "principal_object_id", principal_object_id)
        pulumi.set(__self__, "resource_object_id", resource_object_id)

    @property
    @pulumi.getter(name="appRoleId")
    def app_role_id(self) -> pulumi.Input[str]:
        """
        The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "app_role_id")

    @app_role_id.setter
    def app_role_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_role_id", value)

    @property
    @pulumi.getter(name="principalObjectId")
    def principal_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_object_id")

    @principal_object_id.setter
    def principal_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal_object_id", value)

    @property
    @pulumi.getter(name="resourceObjectId")
    def resource_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_object_id")

    @resource_object_id.setter
    def resource_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_object_id", value)


@pulumi.input_type
class _AppRoleAssignmentState:
    def __init__(__self__, *,
                 app_role_id: Optional[pulumi.Input[str]] = None,
                 principal_display_name: Optional[pulumi.Input[str]] = None,
                 principal_object_id: Optional[pulumi.Input[str]] = None,
                 principal_type: Optional[pulumi.Input[str]] = None,
                 resource_display_name: Optional[pulumi.Input[str]] = None,
                 resource_object_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppRoleAssignment resources.
        :param pulumi.Input[str] app_role_id: The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_display_name: The display name of the principal to which the app role is assigned.
        :param pulumi.Input[str] principal_object_id: The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_type: The object type of the principal to which the app role is assigned.
        :param pulumi.Input[str] resource_display_name: The display name of the application representing the resource.
        :param pulumi.Input[str] resource_object_id: The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
        """
        if app_role_id is not None:
            pulumi.set(__self__, "app_role_id", app_role_id)
        if principal_display_name is not None:
            pulumi.set(__self__, "principal_display_name", principal_display_name)
        if principal_object_id is not None:
            pulumi.set(__self__, "principal_object_id", principal_object_id)
        if principal_type is not None:
            pulumi.set(__self__, "principal_type", principal_type)
        if resource_display_name is not None:
            pulumi.set(__self__, "resource_display_name", resource_display_name)
        if resource_object_id is not None:
            pulumi.set(__self__, "resource_object_id", resource_object_id)

    @property
    @pulumi.getter(name="appRoleId")
    def app_role_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "app_role_id")

    @app_role_id.setter
    def app_role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_role_id", value)

    @property
    @pulumi.getter(name="principalDisplayName")
    def principal_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the principal to which the app role is assigned.
        """
        return pulumi.get(self, "principal_display_name")

    @principal_display_name.setter
    def principal_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_display_name", value)

    @property
    @pulumi.getter(name="principalObjectId")
    def principal_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_object_id")

    @principal_object_id.setter
    def principal_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_object_id", value)

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> Optional[pulumi.Input[str]]:
        """
        The object type of the principal to which the app role is assigned.
        """
        return pulumi.get(self, "principal_type")

    @principal_type.setter
    def principal_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_type", value)

    @property
    @pulumi.getter(name="resourceDisplayName")
    def resource_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the application representing the resource.
        """
        return pulumi.get(self, "resource_display_name")

    @resource_display_name.setter
    def resource_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_display_name", value)

    @property
    @pulumi.getter(name="resourceObjectId")
    def resource_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_object_id")

    @resource_object_id.setter
    def resource_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_object_id", value)


class AppRoleAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_role_id: Optional[pulumi.Input[str]] = None,
                 principal_object_id: Optional[pulumi.Input[str]] = None,
                 resource_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an app role assignment for a group, user or service principal. Can be used to grant admin consent for application permissions.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `AppRoleAssignment.ReadWrite.All` and `Application.Read.All`, or `AppRoleAssignment.ReadWrite.All` and `Directory.Read.All`, or `Application.ReadWrite.All`, or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

        ## Example Usage

        *App role assignment for accessing Microsoft Graph*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        well_known = azuread.get_application_published_app_ids()
        msgraph = azuread.ServicePrincipal("msgraph",
            application_id=well_known.result["MicrosoftGraph"],
            use_existing=True)
        example_application = azuread.Application("exampleApplication",
            display_name="example",
            required_resource_accesses=[azuread.ApplicationRequiredResourceAccessArgs(
                resource_app_id=well_known.result["MicrosoftGraph"],
                resource_accesses=[
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.app_role_ids["User.Read.All"],
                        type="Role",
                    ),
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.oauth2_permission_scope_ids["User.ReadWrite"],
                        type="Scope",
                    ),
                ],
            )])
        example_service_principal = azuread.ServicePrincipal("exampleServicePrincipal", application_id=example_application.application_id)
        example_app_role_assignment = azuread.AppRoleAssignment("exampleAppRoleAssignment",
            app_role_id=msgraph.app_role_ids["User.Read.All"],
            principal_object_id=example_service_principal.object_id,
            resource_object_id=msgraph.object_id)
        ```

        *App role assignment for internal application*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        internal_application = azuread.Application("internalApplication",
            display_name="internal",
            app_roles=[azuread.ApplicationAppRoleArgs(
                allowed_member_types=["Application"],
                description="Apps can query the database",
                display_name="Query",
                enabled=True,
                id="00000000-0000-0000-0000-111111111111",
                value="Query.All",
            )])
        internal_service_principal = azuread.ServicePrincipal("internalServicePrincipal", application_id=internal_application.application_id)
        example_application = azuread.Application("exampleApplication",
            display_name="example",
            required_resource_accesses=[azuread.ApplicationRequiredResourceAccessArgs(
                resource_app_id=internal_application.application_id,
                resource_accesses=[azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                    id=internal_service_principal.app_role_ids["Query.All"],
                    type="Role",
                )],
            )])
        example_service_principal = azuread.ServicePrincipal("exampleServicePrincipal", application_id=example_application.application_id)
        example_app_role_assignment = azuread.AppRoleAssignment("exampleAppRoleAssignment",
            app_role_id=internal_service_principal.app_role_ids["Query.All"],
            principal_object_id=example_service_principal.object_id,
            resource_object_id=internal_service_principal.object_id)
        ```

        *Assign a user and group to an internal application*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_domains = azuread.get_domains(only_initial=True)
        internal_application = azuread.Application("internalApplication",
            display_name="internal",
            app_roles=[azuread.ApplicationAppRoleArgs(
                allowed_member_types=[
                    "Application",
                    "User",
                ],
                description="Admins can perform all task actions",
                display_name="Admin",
                enabled=True,
                id="00000000-0000-0000-0000-222222222222",
                value="Admin.All",
            )])
        internal_service_principal = azuread.ServicePrincipal("internalServicePrincipal", application_id=internal_application.application_id)
        example_group = azuread.Group("exampleGroup",
            display_name="example",
            security_enabled=True)
        example_app_role_assignment = azuread.AppRoleAssignment("exampleAppRoleAssignment",
            app_role_id=internal_service_principal.app_role_ids["Admin.All"],
            principal_object_id=example_group.object_id,
            resource_object_id=internal_service_principal.object_id)
        example_user = azuread.User("exampleUser",
            display_name="D. Duck",
            password="SecretP@sswd99!",
            user_principal_name=f"d.duck@{example_domains.domains[0].domain_name}")
        example_index_app_role_assignment_app_role_assignment = azuread.AppRoleAssignment("exampleIndex/appRoleAssignmentAppRoleAssignment",
            app_role_id=internal_service_principal.app_role_ids["Admin.All"],
            principal_object_id=example_user.object_id,
            resource_object_id=internal_service_principal.object_id)
        ```

        *Assign a group to the default app role for an internal application*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        internal_application = azuread.Application("internalApplication", display_name="internal")
        internal_service_principal = azuread.ServicePrincipal("internalServicePrincipal", application_id=internal_application.application_id)
        example_group = azuread.Group("exampleGroup",
            display_name="example",
            security_enabled=True)
        example_app_role_assignment = azuread.AppRoleAssignment("exampleAppRoleAssignment",
            app_role_id="00000000-0000-0000-0000-000000000000",
            principal_object_id=example_group.object_id,
            resource_object_id=internal_service_principal.object_id)
        ```

        ## Import

        App role assignments can be imported using the object ID of the service principal representing the resource and the ID of the app role assignment (note_not_ the ID of the app role), e.g.

        ```sh
         $ pulumi import azuread:index/appRoleAssignment:AppRoleAssignment example 00000000-0000-0000-0000-000000000000/appRoleAssignment/aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
        ```

         -> This ID format is unique to Terraform and is composed of the Resource Service Principal Object ID and the ID of the App Role Assignment in the format `{ResourcePrincipalID}/appRoleAssignment/{AppRoleAssignmentID}`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_role_id: The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_object_id: The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_object_id: The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppRoleAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an app role assignment for a group, user or service principal. Can be used to grant admin consent for application permissions.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `AppRoleAssignment.ReadWrite.All` and `Application.Read.All`, or `AppRoleAssignment.ReadWrite.All` and `Directory.Read.All`, or `Application.ReadWrite.All`, or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

        ## Example Usage

        *App role assignment for accessing Microsoft Graph*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        well_known = azuread.get_application_published_app_ids()
        msgraph = azuread.ServicePrincipal("msgraph",
            application_id=well_known.result["MicrosoftGraph"],
            use_existing=True)
        example_application = azuread.Application("exampleApplication",
            display_name="example",
            required_resource_accesses=[azuread.ApplicationRequiredResourceAccessArgs(
                resource_app_id=well_known.result["MicrosoftGraph"],
                resource_accesses=[
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.app_role_ids["User.Read.All"],
                        type="Role",
                    ),
                    azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                        id=msgraph.oauth2_permission_scope_ids["User.ReadWrite"],
                        type="Scope",
                    ),
                ],
            )])
        example_service_principal = azuread.ServicePrincipal("exampleServicePrincipal", application_id=example_application.application_id)
        example_app_role_assignment = azuread.AppRoleAssignment("exampleAppRoleAssignment",
            app_role_id=msgraph.app_role_ids["User.Read.All"],
            principal_object_id=example_service_principal.object_id,
            resource_object_id=msgraph.object_id)
        ```

        *App role assignment for internal application*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        internal_application = azuread.Application("internalApplication",
            display_name="internal",
            app_roles=[azuread.ApplicationAppRoleArgs(
                allowed_member_types=["Application"],
                description="Apps can query the database",
                display_name="Query",
                enabled=True,
                id="00000000-0000-0000-0000-111111111111",
                value="Query.All",
            )])
        internal_service_principal = azuread.ServicePrincipal("internalServicePrincipal", application_id=internal_application.application_id)
        example_application = azuread.Application("exampleApplication",
            display_name="example",
            required_resource_accesses=[azuread.ApplicationRequiredResourceAccessArgs(
                resource_app_id=internal_application.application_id,
                resource_accesses=[azuread.ApplicationRequiredResourceAccessResourceAccessArgs(
                    id=internal_service_principal.app_role_ids["Query.All"],
                    type="Role",
                )],
            )])
        example_service_principal = azuread.ServicePrincipal("exampleServicePrincipal", application_id=example_application.application_id)
        example_app_role_assignment = azuread.AppRoleAssignment("exampleAppRoleAssignment",
            app_role_id=internal_service_principal.app_role_ids["Query.All"],
            principal_object_id=example_service_principal.object_id,
            resource_object_id=internal_service_principal.object_id)
        ```

        *Assign a user and group to an internal application*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_domains = azuread.get_domains(only_initial=True)
        internal_application = azuread.Application("internalApplication",
            display_name="internal",
            app_roles=[azuread.ApplicationAppRoleArgs(
                allowed_member_types=[
                    "Application",
                    "User",
                ],
                description="Admins can perform all task actions",
                display_name="Admin",
                enabled=True,
                id="00000000-0000-0000-0000-222222222222",
                value="Admin.All",
            )])
        internal_service_principal = azuread.ServicePrincipal("internalServicePrincipal", application_id=internal_application.application_id)
        example_group = azuread.Group("exampleGroup",
            display_name="example",
            security_enabled=True)
        example_app_role_assignment = azuread.AppRoleAssignment("exampleAppRoleAssignment",
            app_role_id=internal_service_principal.app_role_ids["Admin.All"],
            principal_object_id=example_group.object_id,
            resource_object_id=internal_service_principal.object_id)
        example_user = azuread.User("exampleUser",
            display_name="D. Duck",
            password="SecretP@sswd99!",
            user_principal_name=f"d.duck@{example_domains.domains[0].domain_name}")
        example_index_app_role_assignment_app_role_assignment = azuread.AppRoleAssignment("exampleIndex/appRoleAssignmentAppRoleAssignment",
            app_role_id=internal_service_principal.app_role_ids["Admin.All"],
            principal_object_id=example_user.object_id,
            resource_object_id=internal_service_principal.object_id)
        ```

        *Assign a group to the default app role for an internal application*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        internal_application = azuread.Application("internalApplication", display_name="internal")
        internal_service_principal = azuread.ServicePrincipal("internalServicePrincipal", application_id=internal_application.application_id)
        example_group = azuread.Group("exampleGroup",
            display_name="example",
            security_enabled=True)
        example_app_role_assignment = azuread.AppRoleAssignment("exampleAppRoleAssignment",
            app_role_id="00000000-0000-0000-0000-000000000000",
            principal_object_id=example_group.object_id,
            resource_object_id=internal_service_principal.object_id)
        ```

        ## Import

        App role assignments can be imported using the object ID of the service principal representing the resource and the ID of the app role assignment (note_not_ the ID of the app role), e.g.

        ```sh
         $ pulumi import azuread:index/appRoleAssignment:AppRoleAssignment example 00000000-0000-0000-0000-000000000000/appRoleAssignment/aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
        ```

         -> This ID format is unique to Terraform and is composed of the Resource Service Principal Object ID and the ID of the App Role Assignment in the format `{ResourcePrincipalID}/appRoleAssignment/{AppRoleAssignmentID}`.

        :param str resource_name: The name of the resource.
        :param AppRoleAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppRoleAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_role_id: Optional[pulumi.Input[str]] = None,
                 principal_object_id: Optional[pulumi.Input[str]] = None,
                 resource_object_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppRoleAssignmentArgs.__new__(AppRoleAssignmentArgs)

            if app_role_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_role_id'")
            __props__.__dict__["app_role_id"] = app_role_id
            if principal_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'principal_object_id'")
            __props__.__dict__["principal_object_id"] = principal_object_id
            if resource_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_object_id'")
            __props__.__dict__["resource_object_id"] = resource_object_id
            __props__.__dict__["principal_display_name"] = None
            __props__.__dict__["principal_type"] = None
            __props__.__dict__["resource_display_name"] = None
        super(AppRoleAssignment, __self__).__init__(
            'azuread:index/appRoleAssignment:AppRoleAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_role_id: Optional[pulumi.Input[str]] = None,
            principal_display_name: Optional[pulumi.Input[str]] = None,
            principal_object_id: Optional[pulumi.Input[str]] = None,
            principal_type: Optional[pulumi.Input[str]] = None,
            resource_display_name: Optional[pulumi.Input[str]] = None,
            resource_object_id: Optional[pulumi.Input[str]] = None) -> 'AppRoleAssignment':
        """
        Get an existing AppRoleAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_role_id: The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_display_name: The display name of the principal to which the app role is assigned.
        :param pulumi.Input[str] principal_object_id: The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] principal_type: The object type of the principal to which the app role is assigned.
        :param pulumi.Input[str] resource_display_name: The display name of the application representing the resource.
        :param pulumi.Input[str] resource_object_id: The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppRoleAssignmentState.__new__(_AppRoleAssignmentState)

        __props__.__dict__["app_role_id"] = app_role_id
        __props__.__dict__["principal_display_name"] = principal_display_name
        __props__.__dict__["principal_object_id"] = principal_object_id
        __props__.__dict__["principal_type"] = principal_type
        __props__.__dict__["resource_display_name"] = resource_display_name
        __props__.__dict__["resource_object_id"] = resource_object_id
        return AppRoleAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appRoleId")
    def app_role_id(self) -> pulumi.Output[str]:
        """
        The ID of the app role to be assigned, or the default role ID `00000000-0000-0000-0000-000000000000`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "app_role_id")

    @property
    @pulumi.getter(name="principalDisplayName")
    def principal_display_name(self) -> pulumi.Output[str]:
        """
        The display name of the principal to which the app role is assigned.
        """
        return pulumi.get(self, "principal_display_name")

    @property
    @pulumi.getter(name="principalObjectId")
    def principal_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the user, group or service principal to be assigned this app role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_object_id")

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> pulumi.Output[str]:
        """
        The object type of the principal to which the app role is assigned.
        """
        return pulumi.get(self, "principal_type")

    @property
    @pulumi.getter(name="resourceDisplayName")
    def resource_display_name(self) -> pulumi.Output[str]:
        """
        The display name of the application representing the resource.
        """
        return pulumi.get(self, "resource_display_name")

    @property
    @pulumi.getter(name="resourceObjectId")
    def resource_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the service principal representing the resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_object_id")

