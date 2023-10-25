# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SynchronizationJobArgs', 'SynchronizationJob']

@pulumi.input_type
class SynchronizationJobArgs:
    def __init__(__self__, *,
                 service_principal_id: pulumi.Input[str],
                 template_id: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a SynchronizationJob resource.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] template_id: Identifier of the synchronization template this job is based on.
        :param pulumi.Input[bool] enabled: Whether or not the provisioning job is enabled. Default state is `true`.
        """
        SynchronizationJobArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            service_principal_id=service_principal_id,
            template_id=template_id,
            enabled=enabled,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             service_principal_id: Optional[pulumi.Input[str]] = None,
             template_id: Optional[pulumi.Input[str]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if service_principal_id is None and 'servicePrincipalId' in kwargs:
            service_principal_id = kwargs['servicePrincipalId']
        if service_principal_id is None:
            raise TypeError("Missing 'service_principal_id' argument")
        if template_id is None and 'templateId' in kwargs:
            template_id = kwargs['templateId']
        if template_id is None:
            raise TypeError("Missing 'template_id' argument")

        _setter("service_principal_id", service_principal_id)
        _setter("template_id", template_id)
        if enabled is not None:
            _setter("enabled", enabled)

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Input[str]:
        """
        The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_principal_id", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Input[str]:
        """
        Identifier of the synchronization template this job is based on.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the provisioning job is enabled. Default state is `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class _SynchronizationJobState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 schedules: Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationJobScheduleArgs']]]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SynchronizationJob resources.
        :param pulumi.Input[bool] enabled: Whether or not the provisioning job is enabled. Default state is `true`.
        :param pulumi.Input[Sequence[pulumi.Input['SynchronizationJobScheduleArgs']]] schedules: A `schedule` list as documented below.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] template_id: Identifier of the synchronization template this job is based on.
        """
        _SynchronizationJobState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            enabled=enabled,
            schedules=schedules,
            service_principal_id=service_principal_id,
            template_id=template_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             enabled: Optional[pulumi.Input[bool]] = None,
             schedules: Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationJobScheduleArgs']]]] = None,
             service_principal_id: Optional[pulumi.Input[str]] = None,
             template_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if service_principal_id is None and 'servicePrincipalId' in kwargs:
            service_principal_id = kwargs['servicePrincipalId']
        if template_id is None and 'templateId' in kwargs:
            template_id = kwargs['templateId']

        if enabled is not None:
            _setter("enabled", enabled)
        if schedules is not None:
            _setter("schedules", schedules)
        if service_principal_id is not None:
            _setter("service_principal_id", service_principal_id)
        if template_id is not None:
            _setter("template_id", template_id)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not the provisioning job is enabled. Default state is `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def schedules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationJobScheduleArgs']]]]:
        """
        A `schedule` list as documented below.
        """
        return pulumi.get(self, "schedules")

    @schedules.setter
    def schedules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationJobScheduleArgs']]]]):
        pulumi.set(self, "schedules", value)

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_principal_id", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the synchronization template this job is based on.
        """
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_id", value)


class SynchronizationJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a synchronization job associated with a service principal (enterprise application) within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`

        ## Import

        Synchronization jobs can be imported using the `id`, e.g.

        ```sh
         $ pulumi import azuread:index/synchronizationJob:SynchronizationJob example 00000000-0000-0000-0000-000000000000/job/dataBricks.f5532fc709734b1a90e8a1fa9fd03a82.8442fd39-2183-419c-8732-74b6ce866bd5
        ```

         -> This ID format is unique to Terraform and is composed of the Service Principal Object ID and the ID of the Synchronization Job Id in the format `{servicePrincipalId}/job/{jobId}`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Whether or not the provisioning job is enabled. Default state is `true`.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] template_id: Identifier of the synchronization template this job is based on.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SynchronizationJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a synchronization job associated with a service principal (enterprise application) within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`

        ## Import

        Synchronization jobs can be imported using the `id`, e.g.

        ```sh
         $ pulumi import azuread:index/synchronizationJob:SynchronizationJob example 00000000-0000-0000-0000-000000000000/job/dataBricks.f5532fc709734b1a90e8a1fa9fd03a82.8442fd39-2183-419c-8732-74b6ce866bd5
        ```

         -> This ID format is unique to Terraform and is composed of the Service Principal Object ID and the ID of the Synchronization Job Id in the format `{servicePrincipalId}/job/{jobId}`.

        :param str resource_name: The name of the resource.
        :param SynchronizationJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SynchronizationJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            SynchronizationJobArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SynchronizationJobArgs.__new__(SynchronizationJobArgs)

            __props__.__dict__["enabled"] = enabled
            if service_principal_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_principal_id'")
            __props__.__dict__["service_principal_id"] = service_principal_id
            if template_id is None and not opts.urn:
                raise TypeError("Missing required property 'template_id'")
            __props__.__dict__["template_id"] = template_id
            __props__.__dict__["schedules"] = None
        super(SynchronizationJob, __self__).__init__(
            'azuread:index/synchronizationJob:SynchronizationJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            schedules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SynchronizationJobScheduleArgs']]]]] = None,
            service_principal_id: Optional[pulumi.Input[str]] = None,
            template_id: Optional[pulumi.Input[str]] = None) -> 'SynchronizationJob':
        """
        Get an existing SynchronizationJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Whether or not the provisioning job is enabled. Default state is `true`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SynchronizationJobScheduleArgs']]]] schedules: A `schedule` list as documented below.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] template_id: Identifier of the synchronization template this job is based on.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SynchronizationJobState.__new__(_SynchronizationJobState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["schedules"] = schedules
        __props__.__dict__["service_principal_id"] = service_principal_id
        __props__.__dict__["template_id"] = template_id
        return SynchronizationJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not the provisioning job is enabled. Default state is `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def schedules(self) -> pulumi.Output[Sequence['outputs.SynchronizationJobSchedule']]:
        """
        A `schedule` list as documented below.
        """
        return pulumi.get(self, "schedules")

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Output[str]:
        """
        The object ID of the service principal for which this synchronization job should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[str]:
        """
        Identifier of the synchronization template this job is based on.
        """
        return pulumi.get(self, "template_id")

