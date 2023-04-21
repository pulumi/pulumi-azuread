# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServicePrincipalTokenSigningCertificateArgs', 'ServicePrincipalTokenSigningCertificate']

@pulumi.input_type
class ServicePrincipalTokenSigningCertificateArgs:
    def __init__(__self__, *,
                 service_principal_id: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServicePrincipalTokenSigningCertificate resource.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] display_name: Specifies a friendly name for the certificate.
               Must start with `CN=`. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date: The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        pulumi.set(__self__, "service_principal_id", service_principal_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Input[str]:
        """
        The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_principal_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a friendly name for the certificate.
        Must start with `CN=`. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[str]]:
        """
        The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date", value)


@pulumi.input_type
class _ServicePrincipalTokenSigningCertificateState:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 thumbprint: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServicePrincipalTokenSigningCertificate resources.
        :param pulumi.Input[str] display_name: Specifies a friendly name for the certificate.
               Must start with `CN=`. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date: The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] key_id: A UUID used to uniquely identify the verify certificate.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] start_date: The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        :param pulumi.Input[str] thumbprint: A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
        :param pulumi.Input[str] value: The certificate data, which is PEM encoded but does not include the
               header `-----BEGIN CERTIFICATE-----\\n` or the footer `\\n-----END CERTIFICATE-----`.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if service_principal_id is not None:
            pulumi.set(__self__, "service_principal_id", service_principal_id)
        if start_date is not None:
            pulumi.set(__self__, "start_date", start_date)
        if thumbprint is not None:
            pulumi.set(__self__, "thumbprint", thumbprint)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a friendly name for the certificate.
        Must start with `CN=`. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[str]]:
        """
        The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        A UUID used to uniquely identify the verify certificate.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

    @service_principal_id.setter
    def service_principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_principal_id", value)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[str]]:
        """
        The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter
    def thumbprint(self) -> Optional[pulumi.Input[str]]:
        """
        A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
        """
        return pulumi.get(self, "thumbprint")

    @thumbprint.setter
    def thumbprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "thumbprint", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate data, which is PEM encoded but does not include the
        header `-----BEGIN CERTIFICATE-----\\n` or the footer `\\n-----END CERTIFICATE-----`.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class ServicePrincipalTokenSigningCertificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a token signing certificate associated with a service principal within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

        ## Example Usage

        *Using default settings*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication", display_name="example")
        example_service_principal = azuread.ServicePrincipal("exampleServicePrincipal", application_id=example_application.application_id)
        example_service_principal_token_signing_certificate = azuread.ServicePrincipalTokenSigningCertificate("exampleServicePrincipalTokenSigningCertificate", service_principal_id=example_service_principal.id)
        ```

        *Using custom settings*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication", display_name="example")
        example_service_principal = azuread.ServicePrincipal("exampleServicePrincipal", application_id=example_application.application_id)
        example_service_principal_token_signing_certificate = azuread.ServicePrincipalTokenSigningCertificate("exampleServicePrincipalTokenSigningCertificate",
            service_principal_id=example_service_principal.id,
            display_name="CN=example.com",
            end_date="2023-05-01T01:02:03Z")
        ```

        ## Import

        Token signing certificates can be imported using the object ID of the associated service principal and the key ID of the verify certificate credential, e.g.

        ```sh
         $ pulumi import azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate example 00000000-0000-0000-0000-000000000000/tokenSigningCertificate/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "tokenSigningCertificate" and the verify certificate's key ID in the format `{ServicePrincipalObjectId}/tokenSigningCertificate/{CertificateKeyId}`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Specifies a friendly name for the certificate.
               Must start with `CN=`. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date: The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServicePrincipalTokenSigningCertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a token signing certificate associated with a service principal within Azure Active Directory.

        ## API Permissions

        The following API permissions are required in order to use this resource.

        When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`

        When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`

        ## Example Usage

        *Using default settings*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication", display_name="example")
        example_service_principal = azuread.ServicePrincipal("exampleServicePrincipal", application_id=example_application.application_id)
        example_service_principal_token_signing_certificate = azuread.ServicePrincipalTokenSigningCertificate("exampleServicePrincipalTokenSigningCertificate", service_principal_id=example_service_principal.id)
        ```

        *Using custom settings*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_application = azuread.Application("exampleApplication", display_name="example")
        example_service_principal = azuread.ServicePrincipal("exampleServicePrincipal", application_id=example_application.application_id)
        example_service_principal_token_signing_certificate = azuread.ServicePrincipalTokenSigningCertificate("exampleServicePrincipalTokenSigningCertificate",
            service_principal_id=example_service_principal.id,
            display_name="CN=example.com",
            end_date="2023-05-01T01:02:03Z")
        ```

        ## Import

        Token signing certificates can be imported using the object ID of the associated service principal and the key ID of the verify certificate credential, e.g.

        ```sh
         $ pulumi import azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate example 00000000-0000-0000-0000-000000000000/tokenSigningCertificate/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "tokenSigningCertificate" and the verify certificate's key ID in the format `{ServicePrincipalObjectId}/tokenSigningCertificate/{CertificateKeyId}`.

        :param str resource_name: The name of the resource.
        :param ServicePrincipalTokenSigningCertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServicePrincipalTokenSigningCertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServicePrincipalTokenSigningCertificateArgs.__new__(ServicePrincipalTokenSigningCertificateArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["end_date"] = end_date
            if service_principal_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_principal_id'")
            __props__.__dict__["service_principal_id"] = service_principal_id
            __props__.__dict__["key_id"] = None
            __props__.__dict__["start_date"] = None
            __props__.__dict__["thumbprint"] = None
            __props__.__dict__["value"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["value"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServicePrincipalTokenSigningCertificate, __self__).__init__(
            'azuread:index/servicePrincipalTokenSigningCertificate:ServicePrincipalTokenSigningCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            end_date: Optional[pulumi.Input[str]] = None,
            key_id: Optional[pulumi.Input[str]] = None,
            service_principal_id: Optional[pulumi.Input[str]] = None,
            start_date: Optional[pulumi.Input[str]] = None,
            thumbprint: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'ServicePrincipalTokenSigningCertificate':
        """
        Get an existing ServicePrincipalTokenSigningCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Specifies a friendly name for the certificate.
               Must start with `CN=`. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date: The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] key_id: A UUID used to uniquely identify the verify certificate.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] start_date: The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        :param pulumi.Input[str] thumbprint: A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
        :param pulumi.Input[str] value: The certificate data, which is PEM encoded but does not include the
               header `-----BEGIN CERTIFICATE-----\\n` or the footer `\\n-----END CERTIFICATE-----`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServicePrincipalTokenSigningCertificateState.__new__(_ServicePrincipalTokenSigningCertificateState)

        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["end_date"] = end_date
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["service_principal_id"] = service_principal_id
        __props__.__dict__["start_date"] = start_date
        __props__.__dict__["thumbprint"] = thumbprint
        __props__.__dict__["value"] = value
        return ServicePrincipalTokenSigningCertificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Specifies a friendly name for the certificate.
        Must start with `CN=`. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> pulumi.Output[str]:
        """
        The end date until which the token signing certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        """
        A UUID used to uniquely identify the verify certificate.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="servicePrincipalId")
    def service_principal_id(self) -> pulumi.Output[str]:
        """
        The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "service_principal_id")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> pulumi.Output[str]:
        """
        The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`).
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter
    def thumbprint(self) -> pulumi.Output[str]:
        """
        A SHA-1 generated thumbprint of the token signing certificate, which can be used to set the preferred signing certificate for a service principal.
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The certificate data, which is PEM encoded but does not include the
        header `-----BEGIN CERTIFICATE-----\\n` or the footer `\\n-----END CERTIFICATE-----`.
        """
        return pulumi.get(self, "value")
