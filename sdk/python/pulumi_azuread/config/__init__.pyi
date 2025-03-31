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
from .. import _utilities

adoPipelineServiceConnectionId: Optional[str]
"""
The Azure DevOps Pipeline Service Connection ID.
"""

clientCertificate: Optional[str]
"""
Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
"""

clientCertificatePassword: Optional[str]
"""
The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
Certificate
"""

clientCertificatePath: Optional[str]
"""
The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
Principal using a Client Certificate
"""

clientId: Optional[str]
"""
The Client ID which should be used for service principal authentication
"""

clientIdFilePath: Optional[str]
"""
The path to a file containing the Client ID which should be used for service principal authentication
"""

clientSecret: Optional[str]
"""
The application password to use when authenticating as a Service Principal using a Client Secret
"""

clientSecretFilePath: Optional[str]
"""
The path to a file containing the application password to use when authenticating as a Service Principal using a Client
Secret
"""

disableTerraformPartnerId: Optional[bool]

environment: str
"""
The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
`usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`. Not used and should not be specified
when `metadata_host` is specified.
"""

metadataHost: Optional[str]
"""
The Hostname which should be used for the Azure Metadata Service.
"""

msiEndpoint: Optional[str]
"""
The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
"""

oidcRequestToken: Optional[str]
"""
The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
Connect.
"""

oidcRequestUrl: Optional[str]
"""
The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
using OpenID Connect.
"""

oidcToken: Optional[str]
"""
The ID token for use when authenticating as a Service Principal using OpenID Connect.
"""

oidcTokenFilePath: Optional[str]
"""
The path to a file containing an ID token for use when authenticating as a Service Principal using OpenID Connect.
"""

partnerId: Optional[str]
"""
A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
"""

tenantId: Optional[str]
"""
The Tenant ID which should be used. Works with all authentication methods except Managed Identity
"""

useAksWorkloadIdentity: Optional[bool]
"""
Allow Azure AKS Workload Identity to be used for Authentication.
"""

useCli: Optional[bool]
"""
Allow Azure CLI to be used for Authentication
"""

useMsi: bool
"""
Allow Managed Identity to be used for Authentication
"""

useOidc: Optional[bool]
"""
Allow OpenID Connect to be used for authentication
"""

