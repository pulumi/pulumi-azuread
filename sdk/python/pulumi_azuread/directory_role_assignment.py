# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DirectoryRoleAssignmentArgs', 'DirectoryRoleAssignment']

@pulumi.input_type
class DirectoryRoleAssignmentArgs:
    def __init__(__self__, *,
                 principal_object_id: pulumi.Input[str],
                 role_id: pulumi.Input[str],
                 app_scope_id: Optional[pulumi.Input[str]] = None,
                 app_scope_object_id: Optional[pulumi.Input[str]] = None,
                 directory_scope_id: Optional[pulumi.Input[str]] = None,
                 directory_scope_object_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DirectoryRoleAssignment resource.
        :param pulumi.Input[str] principal_object_id: The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_id: The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        :param pulumi.Input[str] app_scope_id: Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        :param pulumi.Input[str] app_scope_object_id: Identifier of the app-specific scope when the assignment scope is app-specific
        :param pulumi.Input[str] directory_scope_id: Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        :param pulumi.Input[str] directory_scope_object_id: Identifier of the directory object representing the scope of the assignment
        """
        pulumi.set(__self__, "principal_object_id", principal_object_id)
        pulumi.set(__self__, "role_id", role_id)
        if app_scope_id is not None:
            pulumi.set(__self__, "app_scope_id", app_scope_id)
        if app_scope_object_id is not None:
            warnings.warn("""`app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider""", DeprecationWarning)
            pulumi.log.warn("""app_scope_object_id is deprecated: `app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider""")
        if app_scope_object_id is not None:
            pulumi.set(__self__, "app_scope_object_id", app_scope_object_id)
        if directory_scope_id is not None:
            pulumi.set(__self__, "directory_scope_id", directory_scope_id)
        if directory_scope_object_id is not None:
            pulumi.set(__self__, "directory_scope_object_id", directory_scope_object_id)

    @property
    @pulumi.getter(name="principalObjectId")
    def principal_object_id(self) -> pulumi.Input[str]:
        """
        The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_object_id")

    @principal_object_id.setter
    def principal_object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal_object_id", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Input[str]:
        """
        The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter(name="appScopeId")
    def app_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "app_scope_id")

    @app_scope_id.setter
    def app_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_scope_id", value)

    @property
    @pulumi.getter(name="appScopeObjectId")
    def app_scope_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the app-specific scope when the assignment scope is app-specific
        """
        return pulumi.get(self, "app_scope_object_id")

    @app_scope_object_id.setter
    def app_scope_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_scope_object_id", value)

    @property
    @pulumi.getter(name="directoryScopeId")
    def directory_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "directory_scope_id")

    @directory_scope_id.setter
    def directory_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_scope_id", value)

    @property
    @pulumi.getter(name="directoryScopeObjectId")
    def directory_scope_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the directory object representing the scope of the assignment
        """
        return pulumi.get(self, "directory_scope_object_id")

    @directory_scope_object_id.setter
    def directory_scope_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_scope_object_id", value)


@pulumi.input_type
class _DirectoryRoleAssignmentState:
    def __init__(__self__, *,
                 app_scope_id: Optional[pulumi.Input[str]] = None,
                 app_scope_object_id: Optional[pulumi.Input[str]] = None,
                 directory_scope_id: Optional[pulumi.Input[str]] = None,
                 directory_scope_object_id: Optional[pulumi.Input[str]] = None,
                 principal_object_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DirectoryRoleAssignment resources.
        :param pulumi.Input[str] app_scope_id: Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        :param pulumi.Input[str] app_scope_object_id: Identifier of the app-specific scope when the assignment scope is app-specific
        :param pulumi.Input[str] directory_scope_id: Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        :param pulumi.Input[str] directory_scope_object_id: Identifier of the directory object representing the scope of the assignment
        :param pulumi.Input[str] principal_object_id: The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_id: The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        if app_scope_id is not None:
            pulumi.set(__self__, "app_scope_id", app_scope_id)
        if app_scope_object_id is not None:
            warnings.warn("""`app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider""", DeprecationWarning)
            pulumi.log.warn("""app_scope_object_id is deprecated: `app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider""")
        if app_scope_object_id is not None:
            pulumi.set(__self__, "app_scope_object_id", app_scope_object_id)
        if directory_scope_id is not None:
            pulumi.set(__self__, "directory_scope_id", directory_scope_id)
        if directory_scope_object_id is not None:
            pulumi.set(__self__, "directory_scope_object_id", directory_scope_object_id)
        if principal_object_id is not None:
            pulumi.set(__self__, "principal_object_id", principal_object_id)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)

    @property
    @pulumi.getter(name="appScopeId")
    def app_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "app_scope_id")

    @app_scope_id.setter
    def app_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_scope_id", value)

    @property
    @pulumi.getter(name="appScopeObjectId")
    def app_scope_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the app-specific scope when the assignment scope is app-specific
        """
        return pulumi.get(self, "app_scope_object_id")

    @app_scope_object_id.setter
    def app_scope_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_scope_object_id", value)

    @property
    @pulumi.getter(name="directoryScopeId")
    def directory_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "directory_scope_id")

    @directory_scope_id.setter
    def directory_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_scope_id", value)

    @property
    @pulumi.getter(name="directoryScopeObjectId")
    def directory_scope_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the directory object representing the scope of the assignment
        """
        return pulumi.get(self, "directory_scope_object_id")

    @directory_scope_object_id.setter
    def directory_scope_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_scope_object_id", value)

    @property
    @pulumi.getter(name="principalObjectId")
    def principal_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_object_id")

    @principal_object_id.setter
    def principal_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_object_id", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[str]]:
        """
        The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_id", value)


class DirectoryRoleAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_scope_id: Optional[pulumi.Input[str]] = None,
                 app_scope_object_id: Optional[pulumi.Input[str]] = None,
                 directory_scope_id: Optional[pulumi.Input[str]] = None,
                 directory_scope_object_id: Optional[pulumi.Input[str]] = None,
                 principal_object_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a single directory role assignment within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Import

        Directory role assignments can be imported using the ID of the assignment, e.g.

        ```sh
         $ pulumi import azuread:index/directoryRoleAssignment:DirectoryRoleAssignment test ePROZI_iKE653D_d6aoLHyr-lKgHI8ZGiIdz8CLVcng-1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_scope_id: Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        :param pulumi.Input[str] app_scope_object_id: Identifier of the app-specific scope when the assignment scope is app-specific
        :param pulumi.Input[str] directory_scope_id: Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        :param pulumi.Input[str] directory_scope_object_id: Identifier of the directory object representing the scope of the assignment
        :param pulumi.Input[str] principal_object_id: The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_id: The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DirectoryRoleAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a single directory role assignment within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`

        ## Import

        Directory role assignments can be imported using the ID of the assignment, e.g.

        ```sh
         $ pulumi import azuread:index/directoryRoleAssignment:DirectoryRoleAssignment test ePROZI_iKE653D_d6aoLHyr-lKgHI8ZGiIdz8CLVcng-1
        ```

        :param str resource_name: The name of the resource.
        :param DirectoryRoleAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DirectoryRoleAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_scope_id: Optional[pulumi.Input[str]] = None,
                 app_scope_object_id: Optional[pulumi.Input[str]] = None,
                 directory_scope_id: Optional[pulumi.Input[str]] = None,
                 directory_scope_object_id: Optional[pulumi.Input[str]] = None,
                 principal_object_id: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DirectoryRoleAssignmentArgs.__new__(DirectoryRoleAssignmentArgs)

            __props__.__dict__["app_scope_id"] = app_scope_id
            if app_scope_object_id is not None and not opts.urn:
                warnings.warn("""`app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider""", DeprecationWarning)
                pulumi.log.warn("""app_scope_object_id is deprecated: `app_scope_object_id` has been renamed to `app_scope_id` and will be removed in version 3.0 or the AzureAD Provider""")
            __props__.__dict__["app_scope_object_id"] = app_scope_object_id
            __props__.__dict__["directory_scope_id"] = directory_scope_id
            __props__.__dict__["directory_scope_object_id"] = directory_scope_object_id
            if principal_object_id is None and not opts.urn:
                raise TypeError("Missing required property 'principal_object_id'")
            __props__.__dict__["principal_object_id"] = principal_object_id
            if role_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_id'")
            __props__.__dict__["role_id"] = role_id
        super(DirectoryRoleAssignment, __self__).__init__(
            'azuread:index/directoryRoleAssignment:DirectoryRoleAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_scope_id: Optional[pulumi.Input[str]] = None,
            app_scope_object_id: Optional[pulumi.Input[str]] = None,
            directory_scope_id: Optional[pulumi.Input[str]] = None,
            directory_scope_object_id: Optional[pulumi.Input[str]] = None,
            principal_object_id: Optional[pulumi.Input[str]] = None,
            role_id: Optional[pulumi.Input[str]] = None) -> 'DirectoryRoleAssignment':
        """
        Get an existing DirectoryRoleAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_scope_id: Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        :param pulumi.Input[str] app_scope_object_id: Identifier of the app-specific scope when the assignment scope is app-specific
        :param pulumi.Input[str] directory_scope_id: Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        :param pulumi.Input[str] directory_scope_object_id: Identifier of the directory object representing the scope of the assignment
        :param pulumi.Input[str] principal_object_id: The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        :param pulumi.Input[str] role_id: The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DirectoryRoleAssignmentState.__new__(_DirectoryRoleAssignmentState)

        __props__.__dict__["app_scope_id"] = app_scope_id
        __props__.__dict__["app_scope_object_id"] = app_scope_object_id
        __props__.__dict__["directory_scope_id"] = directory_scope_id
        __props__.__dict__["directory_scope_object_id"] = directory_scope_object_id
        __props__.__dict__["principal_object_id"] = principal_object_id
        __props__.__dict__["role_id"] = role_id
        return DirectoryRoleAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appScopeId")
    def app_scope_id(self) -> pulumi.Output[str]:
        """
        Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "app_scope_id")

    @property
    @pulumi.getter(name="appScopeObjectId")
    def app_scope_object_id(self) -> pulumi.Output[str]:
        """
        Identifier of the app-specific scope when the assignment scope is app-specific
        """
        return pulumi.get(self, "app_scope_object_id")

    @property
    @pulumi.getter(name="directoryScopeId")
    def directory_scope_id(self) -> pulumi.Output[str]:
        """
        Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&tabs=http) for example usage. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "directory_scope_id")

    @property
    @pulumi.getter(name="directoryScopeObjectId")
    def directory_scope_object_id(self) -> pulumi.Output[str]:
        """
        Identifier of the directory object representing the scope of the assignment
        """
        return pulumi.get(self, "directory_scope_object_id")

    @property
    @pulumi.getter(name="principalObjectId")
    def principal_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "principal_object_id")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[str]:
        """
        The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_id")

