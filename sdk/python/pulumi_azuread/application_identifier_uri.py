# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationIdentifierUriArrgs', 'ApplicationIdentifierUri']

@pulumi.input_type
calass ApplicationIdentifierUriArrgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 identifier_uri: pulumi.Input[str]):
        """
        The set of arguments for constructing a ApplicationIdentifierUri resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] identifier_uri: The user-defined URI that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "application_id", application_id)
        pulumi.set(__self__, "identifier_uri", identifier_uri)

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
    @pulumi.getter(name="identifierUri")
    def identifier_uri(self) -> pulumi.Input[str]:
        """
        The user-defined URI that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "identifier_uri")

    @identifier_uri.setter
    def identifier_uri(self, value: pulumi.Input[str]):
        pulumi.set(self, "identifier_uri", value)


@pulumi.input_type
calass _ApplicationIdentifierUriState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 identifier_uri: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationIdentifierUri resources.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] identifier_uri: The user-defined URI that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant. Changing this forces a new resource to be created.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if identifier_uri is not None:
            pulumi.set(__self__, "identifier_uri", identifier_uri)

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
    @pulumi.getter(name="identifierUri")
    def identifier_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The user-defined URI that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "identifier_uri")

    @identifier_uri.setter
    def identifier_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identifier_uri", value)


calass ApplicationIdentifierUri(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 identifier_uri: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application_registration = azuread.ApplicationRegistration("exampleApplicationRegistration", display_name="example")
        example_application_identifier_uri = azuread.ApplicationIdentifierUri("exampleApplicationIdentifierUri",
            application_id=example_application_registration.id,
            identifier_uri="https://app.hashitown.com")
        ```

        > **Tip** For managing multiple identifier URIs for the same application, create another instance of this resource

        *Usage with Application resource*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication", display_name="example")
        example_application_identifier_uri = azuread.ApplicationIdentifierUri("exampleApplicationIdentifierUri", application_id=example_application.id)
        # ...
        ```

        ## Import

        Application Identifier URIs can be imported using the object ID of the application and the base64-encoded identifier URI, in the following format.

        ```sh
         $ pulumi import azuread:index/applicationIdentifierUri:ApplicationIdentifierUri example /applications/00000000-0000-0000-0000-000000000000/identifierUris/aHR0cHM6Ly9leGFtcGxlLm5ldC8=
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] identifier_uri: The user-defined URI that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationIdentifierUriArrgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application_registration = azuread.ApplicationRegistration("exampleApplicationRegistration", display_name="example")
        example_application_identifier_uri = azuread.ApplicationIdentifierUri("exampleApplicationIdentifierUri",
            application_id=example_application_registration.id,
            identifier_uri="https://app.hashitown.com")
        ```

        > **Tip** For managing multiple identifier URIs for the same application, create another instance of this resource

        *Usage with Application resource*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication", display_name="example")
        example_application_identifier_uri = azuread.ApplicationIdentifierUri("exampleApplicationIdentifierUri", application_id=example_application.id)
        # ...
        ```

        ## Import

        Application Identifier URIs can be imported using the object ID of the application and the base64-encoded identifier URI, in the following format.

        ```sh
         $ pulumi import azuread:index/applicationIdentifierUri:ApplicationIdentifierUri example /applications/00000000-0000-0000-0000-000000000000/identifierUris/aHR0cHM6Ly9leGFtcGxlLm5ldC8=
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationIdentifierUriArrgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationIdentifierUriArrgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 identifier_uri: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationIdentifierUriArrgs.__new__(ApplicationIdentifierUriArrgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            if identifier_uri is None and not opts.urn:
                raise TypeError("Missing required property 'identifier_uri'")
            __props__.__dict__["identifier_uri"] = identifier_uri
        super(ApplicationIdentifierUri, __self__).__init__(
            'azuread:index/applicationIdentifierUri:ApplicationIdentifierUri',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            identifier_uri: Optional[pulumi.Input[str]] = None) -> 'ApplicationIdentifierUri':
        """
        Get an existing ApplicationIdentifierUri resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The resource ID of the application registration. Changing this forces a new resource to be created.
        :param pulumi.Input[str] identifier_uri: The user-defined URI that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationIdentifierUriState.__new__(_ApplicationIdentifierUriState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["identifier_uri"] = identifier_uri
        return ApplicationIdentifierUri(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the application registration. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="identifierUri")
    def identifier_uri(self) -> pulumi.Output[str]:
        """
        The user-defined URI that uniquely identifies an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "identifier_uri")

