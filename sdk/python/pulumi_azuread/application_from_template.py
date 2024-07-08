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

__all__ = ['ApplicationFromTemplateArgs', 'ApplicationFromTemplate']

@pulumi.input_type
class ApplicationFromTemplateArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 template_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ApplicationFromTemplate resource.
        :param pulumi.Input[str] display_name: The display name for the application.
        :param pulumi.Input[str] template_id: Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name for the application.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Input[str]:
        """
        Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_id", value)


@pulumi.input_type
class _ApplicationFromTemplateState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 service_principal_object_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationFromTemplate resources.
        :param pulumi.Input[str] application_id: The resource ID for the application.
        :param pulumi.Input[str] application_object_id: The object ID for the application.
        :param pulumi.Input[str] display_name: The display name for the application.
        :param pulumi.Input[str] service_principal_id: The resource ID for the service principal.
        :param pulumi.Input[str] service_principal_object_id: The object ID for the service principal.
        :param pulumi.Input[str] template_id: Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if application_object_id is not None:
            pulumi.set(__self__, "application_object_id", application_object_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if service_principal_id is not None:
            pulumi.set(__self__, "service_principal_id", service_principal_id)
        if service_principal_object_id is not None:
            pulumi.set(__self__, "service_principal_object_id", service_principal_object_id)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID for the application.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID for the application.
        """
        return pulumi.get(self, "application_object_id")

    @application_object_id.setter
    def application_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_object_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name for the application.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID for the service principal.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_principal_id", value)

    @property
    @pulumi.getter(name="servicePrincipalObjectId")
    def service_principal_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID for the service principal.
        """
        return pulumi.get(self, "service_principal_object_id")

    @service_principal_object_id.setter
    def service_principal_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_principal_object_id", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)


class ApplicationFromTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an application registration and associated service principal from a gallery template.

        > The Application resource can also be used to instantiate a gallery application, however unlike the `Application` resource, this resource does not attempt to manage any properties of the resulting application.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

        When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.get_application_template(display_name="Marketo")
        example_application_from_template = azuread.ApplicationFromTemplate("example",
            display_name="Example Application",
            template_id=example.template_id)
        example_get_application = azuread.get_application_output(object_id=example_application_from_template.application_object_id)
        example_get_service_principal = azuread.get_service_principal_output(object_id=example_application_from_template.service_principal_object_id)
        ```

        ## Import

        Templated Applications can be imported using the template ID, the object ID of the application, and the object ID of the service principal, in the following format.

        ```sh
        $ pulumi import azuread:index/applicationFromTemplate:ApplicationFromTemplate example /applicationTemplates/00000000-0000-0000-0000-000000000000/instantiate/11111111-1111-1111-1111-111111111111/22222222-2222-2222-2222-222222222222
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The display name for the application.
        :param pulumi.Input[str] template_id: Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationFromTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an application registration and associated service principal from a gallery template.

        > The Application resource can also be used to instantiate a gallery application, however unlike the `Application` resource, this resource does not attempt to manage any properties of the resulting application.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`

        When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.get_application_template(display_name="Marketo")
        example_application_from_template = azuread.ApplicationFromTemplate("example",
            display_name="Example Application",
            template_id=example.template_id)
        example_get_application = azuread.get_application_output(object_id=example_application_from_template.application_object_id)
        example_get_service_principal = azuread.get_service_principal_output(object_id=example_application_from_template.service_principal_object_id)
        ```

        ## Import

        Templated Applications can be imported using the template ID, the object ID of the application, and the object ID of the service principal, in the following format.

        ```sh
        $ pulumi import azuread:index/applicationFromTemplate:ApplicationFromTemplate example /applicationTemplates/00000000-0000-0000-0000-000000000000/instantiate/11111111-1111-1111-1111-111111111111/22222222-2222-2222-2222-222222222222
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationFromTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationFromTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationFromTemplateArgs.__new__(ApplicationFromTemplateArgs)

            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            if template_id is None and not opts.urn:
                raise TypeError("Missing required property 'template_id'")
            __props__.__dict__["template_id"] = template_id
            __props__.__dict__["application_id"] = None
            __props__.__dict__["application_object_id"] = None
            __props__.__dict__["service_principal_id"] = None
            __props__.__dict__["service_principal_object_id"] = None
        super(ApplicationFromTemplate, __self__).__init__(
            'azuread:index/applicationFromTemplate:ApplicationFromTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            application_object_id: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            service_principal_id: Optional[pulumi.Input[str]] = None,
            service_principal_object_id: Optional[pulumi.Input[str]] = None,
            template_id: Optional[pulumi.Input[str]] = None) -> 'ApplicationFromTemplate':
        """
        Get an existing ApplicationFromTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The resource ID for the application.
        :param pulumi.Input[str] application_object_id: The object ID for the application.
        :param pulumi.Input[str] display_name: The display name for the application.
        :param pulumi.Input[str] service_principal_id: The resource ID for the service principal.
        :param pulumi.Input[str] service_principal_object_id: The object ID for the service principal.
        :param pulumi.Input[str] template_id: Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationFromTemplateState.__new__(_ApplicationFromTemplateState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["application_object_id"] = application_object_id
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["service_principal_id"] = service_principal_id
        __props__.__dict__["service_principal_object_id"] = service_principal_object_id
        __props__.__dict__["template_id"] = template_id
        return ApplicationFromTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The resource ID for the application.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> pulumi.Output[str]:
        """
        The object ID for the application.
        """
        return pulumi.get(self, "application_object_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name for the application.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Output[str]:
        """
        The resource ID for the service principal.
        """
        return pulumi.get(self, "service_principal_id")

    @property
    @pulumi.getter(name="servicePrincipalObjectId")
    def service_principal_object_id(self) -> pulumi.Output[str]:
        """
        The object ID for the service principal.
        """
        return pulumi.get(self, "service_principal_object_id")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[str]:
        """
        Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "template_id")

