# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SynchronizationSecretArgs', 'SynchronizationSecret']

@pulumi.input_type
class SynchronizationSecretArgs:
    def __init__(__self__, *,
                 service_principal_id: pulumi.Input[str],
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationSecretCredentialArgs']]]] = None):
        """
        The set of arguments for constructing a SynchronizationSecret resource.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input['SynchronizationSecretCredentialArgs']]] credentials: One or more `credential` blocks as documented below.
        """
        pulumi.set(__self__, "service_principal_id", service_principal_id)
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Input[str]:
        """
        The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_principal_id", value)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationSecretCredentialArgs']]]]:
        """
        One or more `credential` blocks as documented below.
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationSecretCredentialArgs']]]]):
        pulumi.set(self, "credentials", value)


@pulumi.input_type
class _SynchronizationSecretState:
    def __init__(__self__, *,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationSecretCredentialArgs']]]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SynchronizationSecret resources.
        :param pulumi.Input[Sequence[pulumi.Input['SynchronizationSecretCredentialArgs']]] credentials: One or more `credential` blocks as documented below.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
        """
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if service_principal_id is not None:
            pulumi.set(__self__, "service_principal_id", service_principal_id)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationSecretCredentialArgs']]]]:
        """
        One or more `credential` blocks as documented below.
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SynchronizationSecretCredentialArgs']]]]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_principal_id", value)


class SynchronizationSecret(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SynchronizationSecretCredentialArgs']]]]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages synchronization secrets associated with a service principal (enterprise application) within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SynchronizationSecretCredentialArgs']]]] credentials: One or more `credential` blocks as documented below.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SynchronizationSecretArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages synchronization secrets associated with a service principal (enterprise application) within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`

        ## Import

        This resource does not support importing.

        :param str resource_name: The name of the resource.
        :param SynchronizationSecretArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SynchronizationSecretArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SynchronizationSecretCredentialArgs']]]]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SynchronizationSecretArgs.__new__(SynchronizationSecretArgs)

            __props__.__dict__["credentials"] = credentials
            if service_principal_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_principal_id'")
            __props__.__dict__["service_principal_id"] = service_principal_id
        super(SynchronizationSecret, __self__).__init__(
            'azuread:index/synchronizationSecret:SynchronizationSecret',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SynchronizationSecretCredentialArgs']]]]] = None,
            service_principal_id: Optional[pulumi.Input[str]] = None) -> 'SynchronizationSecret':
        """
        Get an existing SynchronizationSecret resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SynchronizationSecretCredentialArgs']]]] credentials: One or more `credential` blocks as documented below.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SynchronizationSecretState.__new__(_SynchronizationSecretState)

        __props__.__dict__["credentials"] = credentials
        __props__.__dict__["service_principal_id"] = service_principal_id
        return SynchronizationSecret(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Optional[Sequence['outputs.SynchronizationSecretCredential']]]:
        """
        One or more `credential` blocks as documented below.
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Output[str]:
        """
        The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

