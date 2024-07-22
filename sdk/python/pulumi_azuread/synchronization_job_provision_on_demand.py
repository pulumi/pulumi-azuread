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
from . import outputs
from ._inputs import *

__all__ = ['SynchronizationJobProvisionOnDemandArgs', 'SynchronizationJobProvisionOnDemand']

@pulumi.input_type
class SynchronizationJobProvisionOnDemandArgs:
    def __init__(__self__, *,
                 parameters: pulumi.Input[Sequence[pulumi.Input['SynchronizationJobProvisionOnDemandParameterArgs']]],
                 service_principal_id: pulumi.Input[str],
                 synchronization_job_id: pulumi.Input[str],
                 triggers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SynchronizationJobProvisionOnDemand resource.
        :param pulumi.Input[Sequence[pulumi.Input['SynchronizationJobProvisionOnDemandParameterArgs']]] parameters: One or more `parameter` blocks as documented below.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for the synchronization job.
        :param pulumi.Input[str] synchronization_job_id: Identifier of the synchronization template this job is based on.
        """
        pulumi.set(__self__, "parameters", parameters)
        pulumi.set(__self__, "service_principal_id", service_principal_id)
        pulumi.set(__self__, "synchronization_job_id", synchronization_job_id)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Input[Sequence[pulumi.Input['SynchronizationJobProvisionOnDemandParameterArgs']]]:
        """
        One or more `parameter` blocks as documented below.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: pulumi.Input[Sequence[pulumi.Input['SynchronizationJobProvisionOnDemandParameterArgs']]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Input[str]:
        """
        The object ID of the service principal for the synchronization job.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_principal_id", value)

    @property
    @pulumi.getter(name="synchronizationJobId")
    def synchronization_job_id(self) -> pulumi.Input[str]:
        """
        Identifier of the synchronization template this job is based on.
        """
        return pulumi.get(self, "synchronization_job_id")

    @synchronization_job_id.setter
    def synchronization_job_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "synchronization_job_id", value)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "triggers", value)


@pulumi.input_type
class _SynchronizationJobProvisionOnDemandState:
    def __init__(__self__, *,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationJobProvisionOnDemandParameterArgs']]]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 synchronization_job_id: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SynchronizationJobProvisionOnDemand resources.
        :param pulumi.Input[Sequence[pulumi.Input['SynchronizationJobProvisionOnDemandParameterArgs']]] parameters: One or more `parameter` blocks as documented below.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for the synchronization job.
        :param pulumi.Input[str] synchronization_job_id: Identifier of the synchronization template this job is based on.
        """
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if service_principal_id is not None:
            pulumi.set(__self__, "service_principal_id", service_principal_id)
        if synchronization_job_id is not None:
            pulumi.set(__self__, "synchronization_job_id", synchronization_job_id)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationJobProvisionOnDemandParameterArgs']]]]:
        """
        One or more `parameter` blocks as documented below.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationJobProvisionOnDemandParameterArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the service principal for the synchronization job.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_principal_id", value)

    @property
    @pulumi.getter(name="synchronizationJobId")
    def synchronization_job_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of the synchronization template this job is based on.
        """
        return pulumi.get(self, "synchronization_job_id")

    @synchronization_job_id.setter
    def synchronization_job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "synchronization_job_id", value)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "triggers", value)


class SynchronizationJobProvisionOnDemand(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SynchronizationJobProvisionOnDemandParameterArgs', 'SynchronizationJobProvisionOnDemandParameterArgsDict']]]]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 synchronization_job_id: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manages synchronization job on demand provisioning associated with a service principal (enterprise application) within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Synchronization.ReadWrite.All`

        ## Example Usage

        *Basic example*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        current = azuread.get_client_config()
        example_group = azuread.Group("example",
            display_name="example",
            owners=[current.object_id],
            security_enabled=True)
        example = azuread.get_application_template(display_name="Azure Databricks SCIM Provisioning Connector")
        example_application = azuread.Application("example",
            display_name="example",
            template_id=example.template_id,
            feature_tags=[{
                "enterprise": True,
                "gallery": True,
            }])
        example_service_principal = azuread.ServicePrincipal("example",
            client_id=example_application.client_id,
            use_existing=True)
        example_synchronization_secret = azuread.SynchronizationSecret("example",
            service_principal_id=example_service_principal.id,
            credentials=[
                {
                    "key": "BaseAddress",
                    "value": "https://adb-example.azuredatabricks.net/api/2.0/preview/scim",
                },
                {
                    "key": "SecretToken",
                    "value": "some-token",
                },
            ])
        example_synchronization_job = azuread.SynchronizationJob("example",
            service_principal_id=example_service_principal.id,
            template_id="dataBricks",
            enabled=True)
        example_synchronization_job_provision_on_demand = azuread.SynchronizationJobProvisionOnDemand("example",
            service_principal_id=example_service_principal.id,
            synchronization_job_id=example_synchronization_job.id,
            parameters=[{
                "rule_id": "",
                "subjects": [{
                    "object_id": example_group.object_id,
                    "object_type_name": "Group",
                }],
            }])
        ```

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['SynchronizationJobProvisionOnDemandParameterArgs', 'SynchronizationJobProvisionOnDemandParameterArgsDict']]]] parameters: One or more `parameter` blocks as documented below.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for the synchronization job.
        :param pulumi.Input[str] synchronization_job_id: Identifier of the synchronization template this job is based on.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SynchronizationJobProvisionOnDemandArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages synchronization job on demand provisioning associated with a service principal (enterprise application) within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Synchronization.ReadWrite.All`

        ## Example Usage

        *Basic example*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        current = azuread.get_client_config()
        example_group = azuread.Group("example",
            display_name="example",
            owners=[current.object_id],
            security_enabled=True)
        example = azuread.get_application_template(display_name="Azure Databricks SCIM Provisioning Connector")
        example_application = azuread.Application("example",
            display_name="example",
            template_id=example.template_id,
            feature_tags=[{
                "enterprise": True,
                "gallery": True,
            }])
        example_service_principal = azuread.ServicePrincipal("example",
            client_id=example_application.client_id,
            use_existing=True)
        example_synchronization_secret = azuread.SynchronizationSecret("example",
            service_principal_id=example_service_principal.id,
            credentials=[
                {
                    "key": "BaseAddress",
                    "value": "https://adb-example.azuredatabricks.net/api/2.0/preview/scim",
                },
                {
                    "key": "SecretToken",
                    "value": "some-token",
                },
            ])
        example_synchronization_job = azuread.SynchronizationJob("example",
            service_principal_id=example_service_principal.id,
            template_id="dataBricks",
            enabled=True)
        example_synchronization_job_provision_on_demand = azuread.SynchronizationJobProvisionOnDemand("example",
            service_principal_id=example_service_principal.id,
            synchronization_job_id=example_synchronization_job.id,
            parameters=[{
                "rule_id": "",
                "subjects": [{
                    "object_id": example_group.object_id,
                    "object_type_name": "Group",
                }],
            }])
        ```

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param SynchronizationJobProvisionOnDemandArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SynchronizationJobProvisionOnDemandArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SynchronizationJobProvisionOnDemandParameterArgs', 'SynchronizationJobProvisionOnDemandParameterArgsDict']]]]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 synchronization_job_id: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SynchronizationJobProvisionOnDemandArgs.__new__(SynchronizationJobProvisionOnDemandArgs)

            if parameters is None and not opts.urn:
                raise TypeError("Missing required property 'parameters'")
            __props__.__dict__["parameters"] = parameters
            if service_principal_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_principal_id'")
            __props__.__dict__["service_principal_id"] = service_principal_id
            if synchronization_job_id is None and not opts.urn:
                raise TypeError("Missing required property 'synchronization_job_id'")
            __props__.__dict__["synchronization_job_id"] = synchronization_job_id
            __props__.__dict__["triggers"] = triggers
        super(SynchronizationJobProvisionOnDemand, __self__).__init__(
            'azuread:index/synchronizationJobProvisionOnDemand:SynchronizationJobProvisionOnDemand',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            parameters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SynchronizationJobProvisionOnDemandParameterArgs', 'SynchronizationJobProvisionOnDemandParameterArgsDict']]]]] = None,
            service_principal_id: Optional[pulumi.Input[str]] = None,
            synchronization_job_id: Optional[pulumi.Input[str]] = None,
            triggers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'SynchronizationJobProvisionOnDemand':
        """
        Get an existing SynchronizationJobProvisionOnDemand resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['SynchronizationJobProvisionOnDemandParameterArgs', 'SynchronizationJobProvisionOnDemandParameterArgsDict']]]] parameters: One or more `parameter` blocks as documented below.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for the synchronization job.
        :param pulumi.Input[str] synchronization_job_id: Identifier of the synchronization template this job is based on.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SynchronizationJobProvisionOnDemandState.__new__(_SynchronizationJobProvisionOnDemandState)

        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["service_principal_id"] = service_principal_id
        __props__.__dict__["synchronization_job_id"] = synchronization_job_id
        __props__.__dict__["triggers"] = triggers
        return SynchronizationJobProvisionOnDemand(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Sequence['outputs.SynchronizationJobProvisionOnDemandParameter']]:
        """
        One or more `parameter` blocks as documented below.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Output[str]:
        """
        The object ID of the service principal for the synchronization job.
        """
        return pulumi.get(self, "service_principal_id")

    @property
    @pulumi.getter(name="synchronizationJobId")
    def synchronization_job_id(self) -> pulumi.Output[str]:
        """
        Identifier of the synchronization template this job is based on.
        """
        return pulumi.get(self, "synchronization_job_id")

    @property
    @pulumi.getter
    def triggers(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "triggers")

