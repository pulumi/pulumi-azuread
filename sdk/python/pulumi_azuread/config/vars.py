# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('azuread')


class _ExportableConfig(types.ModuleType):
    @property
    def client_certificate_password(self) -> Optional[str]:
        return __config__.get('clientCertificatePassword')

    @property
    def client_certificate_path(self) -> Optional[str]:
        """
        The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
        Principal using a Client Certificate.
        """
        return __config__.get('clientCertificatePath')

    @property
    def client_id(self) -> Optional[str]:
        """
        The Client ID which should be used for service principal authentication.
        """
        return __config__.get('clientId')

    @property
    def client_secret(self) -> Optional[str]:
        """
        The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
        Certificate
        """
        return __config__.get('clientSecret')

    @property
    def disable_terraform_partner_id(self) -> Optional[str]:
        """
        Disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
        """
        return __config__.get('disableTerraformPartnerId')

    @property
    def environment(self) -> str:
        """
        The cloud environment which should be used. Possible values are `global` (formerly `public`), `usgovernment`, `dod`,
        `germany`, and `china`. Defaults to `global`.
        """
        return __config__.get('environment') or (_utilities.get_env('ARM_ENVIRONMENT') or 'public')

    @property
    def metadata_host(self) -> str:
        """
        [DEPRECATED] The Hostname which should be used for the Azure Metadata Service.
        """
        return __config__.get('metadataHost')

    @property
    def msi_endpoint(self) -> Optional[str]:
        """
        The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically.
        """
        return __config__.get('msiEndpoint') or _utilities.get_env('ARM_MSI_ENDPOINT')

    @property
    def partner_id(self) -> Optional[str]:
        """
        A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
        """
        return __config__.get('partnerId')

    @property
    def tenant_id(self) -> Optional[str]:
        """
        The Tenant ID which should be used. Works with all authentication methods except Managed Identity.
        """
        return __config__.get('tenantId')

    @property
    def use_cli(self) -> Optional[str]:
        """
        Allow Azure CLI to be used for Authentication.
        """
        return __config__.get('useCli')

    @property
    def use_microsoft_graph(self) -> Optional[str]:
        """
        Beta: Use the Microsoft Graph API, instead of the legacy Azure Active Directory Graph API, where supported.
        """
        return __config__.get('useMicrosoftGraph')

    @property
    def use_msi(self) -> Optional[str]:
        """
        Allow Managed Identity to be used for Authentication.
        """
        return __config__.get('useMsi') or (_utilities.get_env_bool('ARM_USE_MSI') or False)

