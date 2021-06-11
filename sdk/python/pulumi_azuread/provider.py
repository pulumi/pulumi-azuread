# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 metadata_host: pulumi.Input[str],
                 client_certificate_password: Optional[pulumi.Input[str]] = None,
                 client_certificate_path: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 disable_terraform_partner_id: Optional[pulumi.Input[bool]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 msi_endpoint: Optional[pulumi.Input[str]] = None,
                 partner_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 use_cli: Optional[pulumi.Input[bool]] = None,
                 use_microsoft_graph: Optional[pulumi.Input[bool]] = None,
                 use_msi: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] metadata_host: [DEPRECATED] The Hostname which should be used for the Azure Metadata Service.
        :param pulumi.Input[str] client_certificate_path: The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
               Principal using a Client Certificate.
        :param pulumi.Input[str] client_id: The Client ID which should be used for service principal authentication.
        :param pulumi.Input[str] client_secret: The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
               Certificate
        :param pulumi.Input[bool] disable_terraform_partner_id: Disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
        :param pulumi.Input[str] environment: The cloud environment which should be used. Possible values are `global` (formerly `public`), `usgovernment`, `dod`,
               `germany`, and `china`. Defaults to `global`.
        :param pulumi.Input[str] msi_endpoint: The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically.
        :param pulumi.Input[str] partner_id: A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        :param pulumi.Input[str] tenant_id: The Tenant ID which should be used. Works with all authentication methods except Managed Identity.
        :param pulumi.Input[bool] use_cli: Allow Azure CLI to be used for Authentication.
        :param pulumi.Input[bool] use_microsoft_graph: Beta: Use the Microsoft Graph API, instead of the legacy Azure Active Directory Graph API, where supported.
        :param pulumi.Input[bool] use_msi: Allow Managed Identity to be used for Authentication.
        """
        pulumi.set(__self__, "metadata_host", metadata_host)
        if client_certificate_password is not None:
            pulumi.set(__self__, "client_certificate_password", client_certificate_password)
        if client_certificate_path is not None:
            pulumi.set(__self__, "client_certificate_path", client_certificate_path)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if disable_terraform_partner_id is not None:
            pulumi.set(__self__, "disable_terraform_partner_id", disable_terraform_partner_id)
        if environment is None:
            environment = (_utilities.get_env('ARM_ENVIRONMENT') or 'public')
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if msi_endpoint is None:
            msi_endpoint = _utilities.get_env('ARM_MSI_ENDPOINT')
        if msi_endpoint is not None:
            pulumi.set(__self__, "msi_endpoint", msi_endpoint)
        if partner_id is not None:
            pulumi.set(__self__, "partner_id", partner_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if use_cli is not None:
            pulumi.set(__self__, "use_cli", use_cli)
        if use_microsoft_graph is not None:
            pulumi.set(__self__, "use_microsoft_graph", use_microsoft_graph)
        if use_msi is None:
            use_msi = (_utilities.get_env_bool('ARM_USE_MSI') or False)
        if use_msi is not None:
            pulumi.set(__self__, "use_msi", use_msi)

    @property
    @pulumi.getter(name="metadataHost")
    def metadata_host(self) -> pulumi.Input[str]:
        """
        [DEPRECATED] The Hostname which should be used for the Azure Metadata Service.
        """
        return pulumi.get(self, "metadata_host")

    @metadata_host.setter
    def metadata_host(self, value: pulumi.Input[str]):
        pulumi.set(self, "metadata_host", value)

    @property
    @pulumi.getter(name="clientCertificatePassword")
    def client_certificate_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_certificate_password")

    @client_certificate_password.setter
    def client_certificate_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_certificate_password", value)

    @property
    @pulumi.getter(name="clientCertificatePath")
    def client_certificate_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
        Principal using a Client Certificate.
        """
        return pulumi.get(self, "client_certificate_path")

    @client_certificate_path.setter
    def client_certificate_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_certificate_path", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Client ID which should be used for service principal authentication.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
        Certificate
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="disableTerraformPartnerId")
    def disable_terraform_partner_id(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
        """
        return pulumi.get(self, "disable_terraform_partner_id")

    @disable_terraform_partner_id.setter
    def disable_terraform_partner_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_terraform_partner_id", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        """
        The cloud environment which should be used. Possible values are `global` (formerly `public`), `usgovernment`, `dod`,
        `germany`, and `china`. Defaults to `global`.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="msiEndpoint")
    def msi_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically.
        """
        return pulumi.get(self, "msi_endpoint")

    @msi_endpoint.setter
    def msi_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "msi_endpoint", value)

    @property
    @pulumi.getter(name="partnerId")
    def partner_id(self) -> Optional[pulumi.Input[str]]:
        """
        A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        """
        return pulumi.get(self, "partner_id")

    @partner_id.setter
    def partner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partner_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Tenant ID which should be used. Works with all authentication methods except Managed Identity.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="useCli")
    def use_cli(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow Azure CLI to be used for Authentication.
        """
        return pulumi.get(self, "use_cli")

    @use_cli.setter
    def use_cli(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_cli", value)

    @property
    @pulumi.getter(name="useMicrosoftGraph")
    def use_microsoft_graph(self) -> Optional[pulumi.Input[bool]]:
        """
        Beta: Use the Microsoft Graph API, instead of the legacy Azure Active Directory Graph API, where supported.
        """
        return pulumi.get(self, "use_microsoft_graph")

    @use_microsoft_graph.setter
    def use_microsoft_graph(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_microsoft_graph", value)

    @property
    @pulumi.getter(name="useMsi")
    def use_msi(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow Managed Identity to be used for Authentication.
        """
        return pulumi.get(self, "use_msi")

    @use_msi.setter
    def use_msi(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_msi", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_certificate_password: Optional[pulumi.Input[str]] = None,
                 client_certificate_path: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 disable_terraform_partner_id: Optional[pulumi.Input[bool]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 metadata_host: Optional[pulumi.Input[str]] = None,
                 msi_endpoint: Optional[pulumi.Input[str]] = None,
                 partner_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 use_cli: Optional[pulumi.Input[bool]] = None,
                 use_microsoft_graph: Optional[pulumi.Input[bool]] = None,
                 use_msi: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        The provider type for the azuread package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_certificate_path: The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
               Principal using a Client Certificate.
        :param pulumi.Input[str] client_id: The Client ID which should be used for service principal authentication.
        :param pulumi.Input[str] client_secret: The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
               Certificate
        :param pulumi.Input[bool] disable_terraform_partner_id: Disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
        :param pulumi.Input[str] environment: The cloud environment which should be used. Possible values are `global` (formerly `public`), `usgovernment`, `dod`,
               `germany`, and `china`. Defaults to `global`.
        :param pulumi.Input[str] metadata_host: [DEPRECATED] The Hostname which should be used for the Azure Metadata Service.
        :param pulumi.Input[str] msi_endpoint: The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically.
        :param pulumi.Input[str] partner_id: A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        :param pulumi.Input[str] tenant_id: The Tenant ID which should be used. Works with all authentication methods except Managed Identity.
        :param pulumi.Input[bool] use_cli: Allow Azure CLI to be used for Authentication.
        :param pulumi.Input[bool] use_microsoft_graph: Beta: Use the Microsoft Graph API, instead of the legacy Azure Active Directory Graph API, where supported.
        :param pulumi.Input[bool] use_msi: Allow Managed Identity to be used for Authentication.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the azuread package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_certificate_password: Optional[pulumi.Input[str]] = None,
                 client_certificate_path: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 disable_terraform_partner_id: Optional[pulumi.Input[bool]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 metadata_host: Optional[pulumi.Input[str]] = None,
                 msi_endpoint: Optional[pulumi.Input[str]] = None,
                 partner_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 use_cli: Optional[pulumi.Input[bool]] = None,
                 use_microsoft_graph: Optional[pulumi.Input[bool]] = None,
                 use_msi: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["client_certificate_password"] = client_certificate_password
            __props__.__dict__["client_certificate_path"] = client_certificate_path
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["client_secret"] = client_secret
            __props__.__dict__["disable_terraform_partner_id"] = pulumi.Output.from_input(disable_terraform_partner_id).apply(pulumi.runtime.to_json) if disable_terraform_partner_id is not None else None
            if environment is None:
                environment = (_utilities.get_env('ARM_ENVIRONMENT') or 'public')
            __props__.__dict__["environment"] = environment
            if metadata_host is None and not opts.urn:
                raise TypeError("Missing required property 'metadata_host'")
            __props__.__dict__["metadata_host"] = metadata_host
            if msi_endpoint is None:
                msi_endpoint = _utilities.get_env('ARM_MSI_ENDPOINT')
            __props__.__dict__["msi_endpoint"] = msi_endpoint
            __props__.__dict__["partner_id"] = partner_id
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["use_cli"] = pulumi.Output.from_input(use_cli).apply(pulumi.runtime.to_json) if use_cli is not None else None
            __props__.__dict__["use_microsoft_graph"] = pulumi.Output.from_input(use_microsoft_graph).apply(pulumi.runtime.to_json) if use_microsoft_graph is not None else None
            if use_msi is None:
                use_msi = (_utilities.get_env_bool('ARM_USE_MSI') or False)
            __props__.__dict__["use_msi"] = pulumi.Output.from_input(use_msi).apply(pulumi.runtime.to_json) if use_msi is not None else None
        super(Provider, __self__).__init__(
            'azuread',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="clientCertificatePassword")
    def client_certificate_password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "client_certificate_password")

    @property
    @pulumi.getter(name="clientCertificatePath")
    def client_certificate_path(self) -> pulumi.Output[Optional[str]]:
        """
        The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
        Principal using a Client Certificate.
        """
        return pulumi.get(self, "client_certificate_path")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Client ID which should be used for service principal authentication.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[Optional[str]]:
        """
        The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
        Certificate
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Optional[str]]:
        """
        The cloud environment which should be used. Possible values are `global` (formerly `public`), `usgovernment`, `dod`,
        `germany`, and `china`. Defaults to `global`.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="metadataHost")
    def metadata_host(self) -> pulumi.Output[str]:
        """
        [DEPRECATED] The Hostname which should be used for the Azure Metadata Service.
        """
        return pulumi.get(self, "metadata_host")

    @property
    @pulumi.getter(name="msiEndpoint")
    def msi_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically.
        """
        return pulumi.get(self, "msi_endpoint")

    @property
    @pulumi.getter(name="partnerId")
    def partner_id(self) -> pulumi.Output[Optional[str]]:
        """
        A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        """
        return pulumi.get(self, "partner_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Tenant ID which should be used. Works with all authentication methods except Managed Identity.
        """
        return pulumi.get(self, "tenant_id")

