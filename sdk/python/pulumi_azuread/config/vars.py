# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
<<<<<<< HEAD
from .. import _utilities, _tables
=======
from .. import _utilities
>>>>>>> 1e7e750 (Upgrade to Pulumi v3.0.0-beta.2)

__all__ = [
    'client_certificate_password',
    'client_certificate_path',
    'client_id',
    'client_secret',
    'disable_terraform_partner_id',
    'environment',
    'metadata_host',
    'msi_endpoint',
    'partner_id',
    'tenant_id',
    'use_msi',
]

__config__ = pulumi.Config('azuread')

client_certificate_password = __config__.get('clientCertificatePassword')

client_certificate_path = __config__.get('clientCertificatePath')
"""
The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
Principal using a Client Certificate.
"""

client_id = __config__.get('clientId')
"""
The Client ID which should be used for service principal authentication.
"""

client_secret = __config__.get('clientSecret')
"""
The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
Certificate
"""

disable_terraform_partner_id = __config__.get('disableTerraformPartnerId')
"""
Disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
"""

environment = __config__.get('environment') or (_utilities.get_env('ARM_ENVIRONMENT') or 'public')
"""
The Cloud Environment which should be used. Possible values are `public`, `usgovernment`, `german`, and `china`.
Defaults to `public`.
"""

metadata_host = __config__.get('metadataHost')
"""
The Hostname which should be used for the Azure Metadata Service.
"""

msi_endpoint = __config__.get('msiEndpoint') or (_utilities.get_env('ARM_MSI_ENDPOINT') or '')
"""
The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected
automatically.
"""

partner_id = __config__.get('partnerId')
"""
A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
"""

tenant_id = __config__.get('tenantId')
"""
The Tenant ID which should be used. Works with all authentication methods except MSI.
"""

use_msi = __config__.get('useMsi') or (_utilities.get_env_bool('ARM_USE_MSI') or False)
"""
Allow Managed Service Identity to be used for Authentication.
"""

