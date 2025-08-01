# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['ServicePrincipalPasswordArgs', 'ServicePrincipalPassword']

@pulumi.input_type
class ServicePrincipalPasswordArgs:
    def __init__(__self__, *,
                 service_principal_id: pulumi.Input[_builtins.str],
                 display_name: Optional[pulumi.Input[_builtins.str]] = None,
                 end_date: Optional[pulumi.Input[_builtins.str]] = None,
                 end_date_relative: Optional[pulumi.Input[_builtins.str]] = None,
                 rotate_when_changed: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 start_date: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a ServicePrincipalPassword resource.
        :param pulumi.Input[_builtins.str] service_principal_id: The ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] display_name: A display name for the password.
        :param pulumi.Input[_builtins.str] end_date: The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] end_date_relative: A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] rotate_when_changed: A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        :param pulumi.Input[_builtins.str] start_date: The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        """
        pulumi.set(__self__, "service_principal_id", service_principal_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)
        if end_date_relative is not None:
            warnings.warn("""The `end_date_relative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `end_date` property.""", DeprecationWarning)
            pulumi.log.warn("""end_date_relative is deprecated: The `end_date_relative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `end_date` property.""")
        if end_date_relative is not None:
            pulumi.set(__self__, "end_date_relative", end_date_relative)
        if rotate_when_changed is not None:
            pulumi.set(__self__, "rotate_when_changed", rotate_when_changed)
        if start_date is not None:
            pulumi.set(__self__, "start_date", start_date)

    @_builtins.property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Input[_builtins.str]:
        """
        The ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "service_principal_id", value)

    @_builtins.property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A display name for the password.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "display_name", value)

    @_builtins.property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "end_date", value)

    @_builtins.property
    @pulumi.getter(name="endDateRelative")
    @_utilities.deprecated("""The `end_date_relative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `end_date` property.""")
    def end_date_relative(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date_relative")

    @end_date_relative.setter
    def end_date_relative(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "end_date_relative", value)

    @_builtins.property
    @pulumi.getter(name="rotateWhenChanged")
    def rotate_when_changed(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "rotate_when_changed")

    @rotate_when_changed.setter
    def rotate_when_changed(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "rotate_when_changed", value)

    @_builtins.property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "start_date", value)


@pulumi.input_type
class _ServicePrincipalPasswordState:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[_builtins.str]] = None,
                 end_date: Optional[pulumi.Input[_builtins.str]] = None,
                 end_date_relative: Optional[pulumi.Input[_builtins.str]] = None,
                 key_id: Optional[pulumi.Input[_builtins.str]] = None,
                 rotate_when_changed: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 service_principal_id: Optional[pulumi.Input[_builtins.str]] = None,
                 start_date: Optional[pulumi.Input[_builtins.str]] = None,
                 value: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering ServicePrincipalPassword resources.
        :param pulumi.Input[_builtins.str] display_name: A display name for the password.
        :param pulumi.Input[_builtins.str] end_date: The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] end_date_relative: A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] key_id: A UUID used to uniquely identify this password credential.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] rotate_when_changed: A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        :param pulumi.Input[_builtins.str] service_principal_id: The ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] start_date: The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] value: The password for this service principal, which is generated by Azure Active Directory.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)
        if end_date_relative is not None:
            warnings.warn("""The `end_date_relative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `end_date` property.""", DeprecationWarning)
            pulumi.log.warn("""end_date_relative is deprecated: The `end_date_relative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `end_date` property.""")
        if end_date_relative is not None:
            pulumi.set(__self__, "end_date_relative", end_date_relative)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if rotate_when_changed is not None:
            pulumi.set(__self__, "rotate_when_changed", rotate_when_changed)
        if service_principal_id is not None:
            pulumi.set(__self__, "service_principal_id", service_principal_id)
        if start_date is not None:
            pulumi.set(__self__, "start_date", start_date)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @_builtins.property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A display name for the password.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "display_name", value)

    @_builtins.property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "end_date", value)

    @_builtins.property
    @pulumi.getter(name="endDateRelative")
    @_utilities.deprecated("""The `end_date_relative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `end_date` property.""")
    def end_date_relative(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date_relative")

    @end_date_relative.setter
    def end_date_relative(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "end_date_relative", value)

    @_builtins.property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A UUID used to uniquely identify this password credential.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "key_id", value)

    @_builtins.property
    @pulumi.getter(name="rotateWhenChanged")
    def rotate_when_changed(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "rotate_when_changed")

    @rotate_when_changed.setter
    def rotate_when_changed(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "rotate_when_changed", value)

    @_builtins.property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "service_principal_id", value)

    @_builtins.property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "start_date", value)

    @_builtins.property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The password for this service principal, which is generated by Azure Active Directory.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "value", value)


@pulumi.type_token("azuread:index/servicePrincipalPassword:ServicePrincipalPassword")
class ServicePrincipalPassword(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[_builtins.str]] = None,
                 end_date: Optional[pulumi.Input[_builtins.str]] = None,
                 end_date_relative: Optional[pulumi.Input[_builtins.str]] = None,
                 rotate_when_changed: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 service_principal_id: Optional[pulumi.Input[_builtins.str]] = None,
                 start_date: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        ## Example Usage

        *Basic example*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Application("example", display_name="example")
        example_service_principal = azuread.ServicePrincipal("example", client_id=example.client_id)
        example_service_principal_password = azuread.ServicePrincipalPassword("example", service_principal_id=example_service_principal.id)
        ```

        *Time-based rotation*

        ```python
        import pulumi
        import pulumi_azuread as azuread
        import pulumiverse_time as time

        example = azuread.Application("example", display_name="example")
        example_service_principal = azuread.ServicePrincipal("example", client_id=example.client_id)
        example_rotating = time.Rotating("example", rotation_days=7)
        example_service_principal_password = azuread.ServicePrincipalPassword("example",
            service_principal_id=example_service_principal.id,
            rotate_when_changed={
                "rotation": example_rotating.id,
            })
        ```

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] display_name: A display name for the password.
        :param pulumi.Input[_builtins.str] end_date: The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] end_date_relative: A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] rotate_when_changed: A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        :param pulumi.Input[_builtins.str] service_principal_id: The ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] start_date: The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServicePrincipalPasswordArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        *Basic example*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Application("example", display_name="example")
        example_service_principal = azuread.ServicePrincipal("example", client_id=example.client_id)
        example_service_principal_password = azuread.ServicePrincipalPassword("example", service_principal_id=example_service_principal.id)
        ```

        *Time-based rotation*

        ```python
        import pulumi
        import pulumi_azuread as azuread
        import pulumiverse_time as time

        example = azuread.Application("example", display_name="example")
        example_service_principal = azuread.ServicePrincipal("example", client_id=example.client_id)
        example_rotating = time.Rotating("example", rotation_days=7)
        example_service_principal_password = azuread.ServicePrincipalPassword("example",
            service_principal_id=example_service_principal.id,
            rotate_when_changed={
                "rotation": example_rotating.id,
            })
        ```

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param ServicePrincipalPasswordArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServicePrincipalPasswordArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[_builtins.str]] = None,
                 end_date: Optional[pulumi.Input[_builtins.str]] = None,
                 end_date_relative: Optional[pulumi.Input[_builtins.str]] = None,
                 rotate_when_changed: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 service_principal_id: Optional[pulumi.Input[_builtins.str]] = None,
                 start_date: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServicePrincipalPasswordArgs.__new__(ServicePrincipalPasswordArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["end_date"] = end_date
            __props__.__dict__["end_date_relative"] = end_date_relative
            __props__.__dict__["rotate_when_changed"] = rotate_when_changed
            if service_principal_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_principal_id'")
            __props__.__dict__["service_principal_id"] = service_principal_id
            __props__.__dict__["start_date"] = start_date
            __props__.__dict__["key_id"] = None
            __props__.__dict__["value"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["value"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServicePrincipalPassword, __self__).__init__(
            'azuread:index/servicePrincipalPassword:ServicePrincipalPassword',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[_builtins.str]] = None,
            end_date: Optional[pulumi.Input[_builtins.str]] = None,
            end_date_relative: Optional[pulumi.Input[_builtins.str]] = None,
            key_id: Optional[pulumi.Input[_builtins.str]] = None,
            rotate_when_changed: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
            service_principal_id: Optional[pulumi.Input[_builtins.str]] = None,
            start_date: Optional[pulumi.Input[_builtins.str]] = None,
            value: Optional[pulumi.Input[_builtins.str]] = None) -> 'ServicePrincipalPassword':
        """
        Get an existing ServicePrincipalPassword resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] display_name: A display name for the password.
        :param pulumi.Input[_builtins.str] end_date: The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] end_date_relative: A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] key_id: A UUID used to uniquely identify this password credential.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] rotate_when_changed: A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        :param pulumi.Input[_builtins.str] service_principal_id: The ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] start_date: The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        :param pulumi.Input[_builtins.str] value: The password for this service principal, which is generated by Azure Active Directory.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServicePrincipalPasswordState.__new__(_ServicePrincipalPasswordState)

        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["end_date"] = end_date
        __props__.__dict__["end_date_relative"] = end_date_relative
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["rotate_when_changed"] = rotate_when_changed
        __props__.__dict__["service_principal_id"] = service_principal_id
        __props__.__dict__["start_date"] = start_date
        __props__.__dict__["value"] = value
        return ServicePrincipalPassword(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[_builtins.str]:
        """
        A display name for the password.
        """
        return pulumi.get(self, "display_name")

    @_builtins.property
    @pulumi.getter(name="endDate")
    def end_date(self) -> pulumi.Output[_builtins.str]:
        """
        The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @_builtins.property
    @pulumi.getter(name="endDateRelative")
    @_utilities.deprecated("""The `end_date_relative` property is deprecated and will be removed in a future version of the AzureAD provider. Please instead use the Terraform `timeadd()` function to calculate a value for the `end_date` property.""")
    def end_date_relative(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date_relative")

    @_builtins.property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[_builtins.str]:
        """
        A UUID used to uniquely identify this password credential.
        """
        return pulumi.get(self, "key_id")

    @_builtins.property
    @pulumi.getter(name="rotateWhenChanged")
    def rotate_when_changed(self) -> pulumi.Output[Optional[Mapping[str, _builtins.str]]]:
        """
        A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "rotate_when_changed")

    @_builtins.property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Output[_builtins.str]:
        """
        The ID of the service principal for which this password should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

    @_builtins.property
    @pulumi.getter(name="startDate")
    def start_date(self) -> pulumi.Output[_builtins.str]:
        """
        The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "start_date")

    @_builtins.property
    @pulumi.getter
    def value(self) -> pulumi.Output[_builtins.str]:
        """
        The password for this service principal, which is generated by Azure Active Directory.
        """
        return pulumi.get(self, "value")

