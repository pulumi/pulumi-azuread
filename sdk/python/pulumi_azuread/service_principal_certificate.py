# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServicePrincipalCertificateArgs', 'ServicePrincipalCertificate']

@pulumi.input_type
class ServicePrincipalCertificateArgs:
    def __init__(__self__, *,
                 service_principal_id: pulumi.Input[str],
                 value: pulumi.Input[str],
                 encoding: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServicePrincipalCertificate resource.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] value: The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
        :param pulumi.Input[str] encoding: Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
               
               > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
        :param pulumi.Input[str] end_date: The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date_relative: A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
               
               > One of `end_date` or `end_date_relative` must be set. The maximum duration is determined by Azure AD.
        :param pulumi.Input[str] key_id: A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] start_date: The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] type: The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        """
        ServicePrincipalCertificateArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            service_principal_id=service_principal_id,
            value=value,
            encoding=encoding,
            end_date=end_date,
            end_date_relative=end_date_relative,
            key_id=key_id,
            start_date=start_date,
            type=type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             service_principal_id: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[str]] = None,
             encoding: Optional[pulumi.Input[str]] = None,
             end_date: Optional[pulumi.Input[str]] = None,
             end_date_relative: Optional[pulumi.Input[str]] = None,
             key_id: Optional[pulumi.Input[str]] = None,
             start_date: Optional[pulumi.Input[str]] = None,
             type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if service_principal_id is None and 'servicePrincipalId' in kwargs:
            service_principal_id = kwargs['servicePrincipalId']
        if service_principal_id is None:
            raise TypeError("Missing 'service_principal_id' argument")
        if value is None:
            raise TypeError("Missing 'value' argument")
        if end_date is None and 'endDate' in kwargs:
            end_date = kwargs['endDate']
        if end_date_relative is None and 'endDateRelative' in kwargs:
            end_date_relative = kwargs['endDateRelative']
        if key_id is None and 'keyId' in kwargs:
            key_id = kwargs['keyId']
        if start_date is None and 'startDate' in kwargs:
            start_date = kwargs['startDate']

        _setter("service_principal_id", service_principal_id)
        _setter("value", value)
        if encoding is not None:
            _setter("encoding", encoding)
        if end_date is not None:
            _setter("end_date", end_date)
        if end_date_relative is not None:
            _setter("end_date_relative", end_date_relative)
        if key_id is not None:
            _setter("key_id", key_id)
        if start_date is not None:
            _setter("start_date", start_date)
        if type is not None:
            _setter("type", type)

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
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.

        > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[str]]:
        """
        The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date", value)

    @property
    @pulumi.getter(name="endDateRelative")
    def end_date_relative(self) -> Optional[pulumi.Input[str]]:
        """
        A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.

        > One of `end_date` or `end_date_relative` must be set. The maximum duration is determined by Azure AD.
        """
        return pulumi.get(self, "end_date_relative")

    @end_date_relative.setter
    def end_date_relative(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date_relative", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[str]]:
        """
        The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _ServicePrincipalCertificateState:
    def __init__(__self__, *,
                 encoding: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServicePrincipalCertificate resources.
        :param pulumi.Input[str] encoding: Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
               
               > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
        :param pulumi.Input[str] end_date: The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date_relative: A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
               
               > One of `end_date` or `end_date_relative` must be set. The maximum duration is determined by Azure AD.
        :param pulumi.Input[str] key_id: A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] start_date: The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] type: The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        :param pulumi.Input[str] value: The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
        """
        _ServicePrincipalCertificateState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            encoding=encoding,
            end_date=end_date,
            end_date_relative=end_date_relative,
            key_id=key_id,
            service_principal_id=service_principal_id,
            start_date=start_date,
            type=type,
            value=value,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             encoding: Optional[pulumi.Input[str]] = None,
             end_date: Optional[pulumi.Input[str]] = None,
             end_date_relative: Optional[pulumi.Input[str]] = None,
             key_id: Optional[pulumi.Input[str]] = None,
             service_principal_id: Optional[pulumi.Input[str]] = None,
             start_date: Optional[pulumi.Input[str]] = None,
             type: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if end_date is None and 'endDate' in kwargs:
            end_date = kwargs['endDate']
        if end_date_relative is None and 'endDateRelative' in kwargs:
            end_date_relative = kwargs['endDateRelative']
        if key_id is None and 'keyId' in kwargs:
            key_id = kwargs['keyId']
        if service_principal_id is None and 'servicePrincipalId' in kwargs:
            service_principal_id = kwargs['servicePrincipalId']
        if start_date is None and 'startDate' in kwargs:
            start_date = kwargs['startDate']

        if encoding is not None:
            _setter("encoding", encoding)
        if end_date is not None:
            _setter("end_date", end_date)
        if end_date_relative is not None:
            _setter("end_date_relative", end_date_relative)
        if key_id is not None:
            _setter("key_id", key_id)
        if service_principal_id is not None:
            _setter("service_principal_id", service_principal_id)
        if start_date is not None:
            _setter("start_date", start_date)
        if type is not None:
            _setter("type", type)
        if value is not None:
            _setter("value", value)

    @property
    @pulumi.getter
    def encoding(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.

        > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
        """
        return pulumi.get(self, "encoding")

    @encoding.setter
    def encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encoding", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[str]]:
        """
        The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date", value)

    @property
    @pulumi.getter(name="endDateRelative")
    def end_date_relative(self) -> Optional[pulumi.Input[str]]:
        """
        A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.

        > One of `end_date` or `end_date_relative` must be set. The maximum duration is determined by Azure AD.
        """
        return pulumi.get(self, "end_date_relative")

    @end_date_relative.setter
    def end_date_relative(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date_relative", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
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
        The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class ServicePrincipalCertificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encoding: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Certificates can be imported using the object ID of the associated service principal and the key ID of the certificate credential, e.g.

        ```sh
         $ pulumi import azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate example 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "certificate" and the certificate's key ID in the format `{ServicePrincipalObjectId}/certificate/{CertificateKeyId}`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] encoding: Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
               
               > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
        :param pulumi.Input[str] end_date: The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date_relative: A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
               
               > One of `end_date` or `end_date_relative` must be set. The maximum duration is determined by Azure AD.
        :param pulumi.Input[str] key_id: A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] start_date: The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] type: The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        :param pulumi.Input[str] value: The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServicePrincipalCertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Certificates can be imported using the object ID of the associated service principal and the key ID of the certificate credential, e.g.

        ```sh
         $ pulumi import azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate example 00000000-0000-0000-0000-000000000000/certificate/11111111-1111-1111-1111-111111111111
        ```

         -> This ID format is unique to Terraform and is composed of the service principal's object ID, the string "certificate" and the certificate's key ID in the format `{ServicePrincipalObjectId}/certificate/{CertificateKeyId}`.

        :param str resource_name: The name of the resource.
        :param ServicePrincipalCertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServicePrincipalCertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ServicePrincipalCertificateArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 encoding: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 end_date_relative: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 service_principal_id: Optional[pulumi.Input[str]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServicePrincipalCertificateArgs.__new__(ServicePrincipalCertificateArgs)

            __props__.__dict__["encoding"] = encoding
            __props__.__dict__["end_date"] = end_date
            __props__.__dict__["end_date_relative"] = end_date_relative
            __props__.__dict__["key_id"] = key_id
            if service_principal_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_principal_id'")
            __props__.__dict__["service_principal_id"] = service_principal_id
            __props__.__dict__["start_date"] = start_date
            __props__.__dict__["type"] = type
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = None if value is None else pulumi.Output.secret(value)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["value"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ServicePrincipalCertificate, __self__).__init__(
            'azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            encoding: Optional[pulumi.Input[str]] = None,
            end_date: Optional[pulumi.Input[str]] = None,
            end_date_relative: Optional[pulumi.Input[str]] = None,
            key_id: Optional[pulumi.Input[str]] = None,
            service_principal_id: Optional[pulumi.Input[str]] = None,
            start_date: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'ServicePrincipalCertificate':
        """
        Get an existing ServicePrincipalCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] encoding: Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.
               
               > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
        :param pulumi.Input[str] end_date: The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        :param pulumi.Input[str] end_date_relative: A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
               
               > One of `end_date` or `end_date_relative` must be set. The maximum duration is determined by Azure AD.
        :param pulumi.Input[str] key_id: A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] service_principal_id: The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] start_date: The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
        :param pulumi.Input[str] type: The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        :param pulumi.Input[str] value: The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServicePrincipalCertificateState.__new__(_ServicePrincipalCertificateState)

        __props__.__dict__["encoding"] = encoding
        __props__.__dict__["end_date"] = end_date
        __props__.__dict__["end_date_relative"] = end_date_relative
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["service_principal_id"] = service_principal_id
        __props__.__dict__["start_date"] = start_date
        __props__.__dict__["type"] = type
        __props__.__dict__["value"] = value
        return ServicePrincipalCertificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def encoding(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the encoding used for the supplied certificate data. Must be one of `pem`, `base64` or `hex`. Defaults to `pem`.

        > **Tip for Azure Key Vault** The `hex` encoding option is useful for consuming certificate data from the azurerm_key_vault_certificate resource.
        """
        return pulumi.get(self, "encoding")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> pulumi.Output[str]:
        """
        The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter(name="endDateRelative")
    def end_date_relative(self) -> pulumi.Output[Optional[str]]:
        """
        A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.

        > One of `end_date` or `end_date_relative` must be set. The maximum duration is determined by Azure AD.
        """
        return pulumi.get(self, "end_date_relative")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        """
        A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
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
        The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of key/certificate. Must be one of `AsymmetricX509Cert` or `Symmetric`. Changing this fields forces a new resource to be created.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the `encoding` argument.
        """
        return pulumi.get(self, "value")

