# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ApplicationPasswordArgs', 'ApplicationPassword']

@pulumi.input_type
class ApplicationPasswordArgs:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 rotate_when_changed: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_date: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApplicationPassword resource.
        :param pulumi.Input[str] application_id: The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which this password should be created
        :param pulumi.Input[str] display_name: A display name for the password. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date: The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date_relative: A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] rotate_when_changed: A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        :param pulumi.Input[str] start_date: The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if application_object_id is not None:
            warnings.warn("""The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider""", DeprecationWarning)
            pulumi.log.warn("""application_object_id is deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider""")
        if application_object_id is not None:
            pulumi.set(__self__, "application_object_id", application_object_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)
        if end_date_relative is not None:
            pulumi.set(__self__, "end_date_relative", end_date_relative)
        if rotate_when_changed is not None:
            pulumi.set(__self__, "rotate_when_changed", rotate_when_changed)
        if start_date is not None:
            pulumi.set(__self__, "start_date", start_date)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the application for which this password should be created
        """
        warnings.warn("""The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider""", DeprecationWarning)
        pulumi.log.warn("""application_object_id is deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider""")

        return pulumi.get(self, "application_object_id")

    @application_object_id.setter
    def application_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_object_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A display name for the password. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[str]]:
        """
        The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date", value)

    @property
    @pulumi.getter(name="endDateRelative")
    def end_date_relative(self) -> Optional[pulumi.Input[str]]:
        """
        A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date_relative")

    @end_date_relative.setter
    def end_date_relative(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date_relative", value)

    @property
    @pulumi.getter(name="rotateWhenChanged")
    def rotate_when_changed(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "rotate_when_changed")

    @rotate_when_changed.setter
    def rotate_when_changed(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "rotate_when_changed", value)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[str]]:
        """
        The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_date", value)


@pulumi.input_type
class _ApplicationPasswordState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 rotate_when_changed: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApplicationPassword resources.
        :param pulumi.Input[str] application_id: The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which this password should be created
        :param pulumi.Input[str] display_name: A display name for the password. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date: The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date_relative: A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] key_id: A UUID used to uniquely identify this password credential.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] rotate_when_changed: A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        :param pulumi.Input[str] start_date: The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        :param pulumi.Input[str] value: The password for this application, which is generated by Azure Active Directory.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if application_object_id is not None:
            warnings.warn("""The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider""", DeprecationWarning)
            pulumi.log.warn("""application_object_id is deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider""")
        if application_object_id is not None:
            pulumi.set(__self__, "application_object_id", application_object_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)
        if end_date_relative is not None:
            pulumi.set(__self__, "end_date_relative", end_date_relative)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if rotate_when_changed is not None:
            pulumi.set(__self__, "rotate_when_changed", rotate_when_changed)
        if start_date is not None:
            pulumi.set(__self__, "start_date", start_date)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the application for which this password should be created
        """
        warnings.warn("""The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider""", DeprecationWarning)
        pulumi.log.warn("""application_object_id is deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider""")

        return pulumi.get(self, "application_object_id")

    @application_object_id.setter
    def application_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_object_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A display name for the password. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[str]]:
        """
        The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date", value)

    @property
    @pulumi.getter(name="endDateRelative")
    def end_date_relative(self) -> Optional[pulumi.Input[str]]:
        """
        A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date_relative")

    @end_date_relative.setter
    def end_date_relative(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date_relative", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        A UUID used to uniquely identify this password credential.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="rotateWhenChanged")
    def rotate_when_changed(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "rotate_when_changed")

    @rotate_when_changed.setter
    def rotate_when_changed(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "rotate_when_changed", value)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[str]]:
        """
        The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The password for this application, which is generated by Azure Active Directory.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class ApplicationPassword(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 rotate_when_changed: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        *Basic example*

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ApplicationRegistration("example", display_name="example")
        example_application_password = azuread.ApplicationPassword("example", application_id=example.id)
        ```
        <!--End PulumiCodeChooser -->

        *Time-based rotation*

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread
        import pulumiverse_time as time

        example = azuread.ApplicationRegistration("example", display_name="example")
        example_rotating = time.Rotating("example", rotation_days=7)
        example_application_password = azuread.ApplicationPassword("example",
            application_id=example.id,
            rotate_when_changed={
                "rotation": example_rotating.id,
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which this password should be created
        :param pulumi.Input[str] display_name: A display name for the password. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date: The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date_relative: A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] rotate_when_changed: A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        :param pulumi.Input[str] start_date: The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ApplicationPasswordArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        *Basic example*

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.ApplicationRegistration("example", display_name="example")
        example_application_password = azuread.ApplicationPassword("example", application_id=example.id)
        ```
        <!--End PulumiCodeChooser -->

        *Time-based rotation*

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_azuread as azuread
        import pulumiverse_time as time

        example = azuread.ApplicationRegistration("example", display_name="example")
        example_rotating = time.Rotating("example", rotation_days=7)
        example_application_password = azuread.ApplicationPassword("example",
            application_id=example.id,
            rotate_when_changed={
                "rotation": example_rotating.id,
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param ApplicationPasswordArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationPasswordArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 rotate_when_changed: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationPasswordArgs.__new__(ApplicationPasswordArgs)

            __props__.__dict__["application_id"] = application_id
            __props__.__dict__["application_object_id"] = application_object_id
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["end_date"] = end_date
            __props__.__dict__["end_date_relative"] = end_date_relative
            __props__.__dict__["rotate_when_changed"] = rotate_when_changed
            __props__.__dict__["start_date"] = start_date
            __props__.__dict__["key_id"] = None
            __props__.__dict__["value"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["value"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ApplicationPassword, __self__).__init__(
            'azuread:index/applicationPassword:ApplicationPassword',
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
            end_date: Optional[pulumi.Input[str]] = None,
            end_date_relative: Optional[pulumi.Input[str]] = None,
            key_id: Optional[pulumi.Input[str]] = None,
            rotate_when_changed: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            start_date: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'ApplicationPassword':
        """
        Get an existing ApplicationPassword resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] application_object_id: The object ID of the application for which this password should be created
        :param pulumi.Input[str] display_name: A display name for the password. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date: The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date_relative: A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] key_id: A UUID used to uniquely identify this password credential.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] rotate_when_changed: A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        :param pulumi.Input[str] start_date: The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        :param pulumi.Input[str] value: The password for this application, which is generated by Azure Active Directory.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationPasswordState.__new__(_ApplicationPasswordState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["application_object_id"] = application_object_id
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["end_date"] = end_date
        __props__.__dict__["end_date_relative"] = end_date_relative
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["rotate_when_changed"] = rotate_when_changed
        __props__.__dict__["start_date"] = start_date
        __props__.__dict__["value"] = value
        return ApplicationPassword(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The resource ID of the application for which this password should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> pulumi.Output[str]:
        """
        The object ID of the application for which this password should be created
        """
        warnings.warn("""The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider""", DeprecationWarning)
        pulumi.log.warn("""application_object_id is deprecated: The `application_object_id` property has been replaced with the `application_id` property and will be removed in version 3.0 of the AzureAD provider""")

        return pulumi.get(self, "application_object_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        A display name for the password. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> pulumi.Output[str]:
        """
        The end date until which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter(name="endDateRelative")
    def end_date_relative(self) -> pulumi.Output[Optional[str]]:
        """
        A relative duration for which the password is valid until, for example `240h` (10 days) or `2400h30m`. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date_relative")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        """
        A UUID used to uniquely identify this password credential.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="rotateWhenChanged")
    def rotate_when_changed(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of arbitrary key/value pairs that will force recreation of the password when they change, enabling password rotation based on external conditions such as a rotating timestamp. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "rotate_when_changed")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> pulumi.Output[str]:
        """
        The start date from which the password is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used.  Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The password for this application, which is generated by Azure Active Directory.
        """
        return pulumi.get(self, "value")

