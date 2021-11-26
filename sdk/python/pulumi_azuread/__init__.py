# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .administrative_unit import *
from .administrative_unit_member import *
from .app_role_assignment import *
from .application import *
from .application_certificate import *
from .application_password import *
from .application_pre_authorized import *
from .conditional_access_policy import *
from .directory_role import *
from .directory_role_member import *
from .get_administrative_unit import *
from .get_application import *
from .get_application_published_app_ids import *
from .get_application_template import *
from .get_client_config import *
from .get_domains import *
from .get_group import *
from .get_groups import *
from .get_service_principal import *
from .get_service_principals import *
from .get_user import *
from .get_users import *
from .group import *
from .group_member import *
from .invitation import *
from .named_location import *
from .provider import *
from .service_principal import *
from .service_principal_certificate import *
from .service_principal_delegated_permission_grant import *
from .service_principal_password import *
from .user import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azuread.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_azuread.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "azuread",
  "mod": "index/administrativeUnit",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/administrativeUnit:AdministrativeUnit": "AdministrativeUnit"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/administrativeUnitMember",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/administrativeUnitMember:AdministrativeUnitMember": "AdministrativeUnitMember"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/appRoleAssignment",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/appRoleAssignment:AppRoleAssignment": "AppRoleAssignment"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/application",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/application:Application": "Application"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/applicationCertificate",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/applicationCertificate:ApplicationCertificate": "ApplicationCertificate"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/applicationPassword",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/applicationPassword:ApplicationPassword": "ApplicationPassword"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/applicationPreAuthorized",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/applicationPreAuthorized:ApplicationPreAuthorized": "ApplicationPreAuthorized"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/conditionalAccessPolicy",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/conditionalAccessPolicy:ConditionalAccessPolicy": "ConditionalAccessPolicy"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/directoryRole",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/directoryRole:DirectoryRole": "DirectoryRole"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/directoryRoleMember",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/directoryRoleMember:DirectoryRoleMember": "DirectoryRoleMember"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/group",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/group:Group": "Group"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/groupMember",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/groupMember:GroupMember": "GroupMember"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/invitation",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/invitation:Invitation": "Invitation"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/namedLocation",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/namedLocation:NamedLocation": "NamedLocation"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/servicePrincipal",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/servicePrincipal:ServicePrincipal": "ServicePrincipal"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/servicePrincipalCertificate",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/servicePrincipalCertificate:ServicePrincipalCertificate": "ServicePrincipalCertificate"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/servicePrincipalDelegatedPermissionGrant",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant": "ServicePrincipalDelegatedPermissionGrant"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/servicePrincipalPassword",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/servicePrincipalPassword:ServicePrincipalPassword": "ServicePrincipalPassword"
  }
 },
 {
  "pkg": "azuread",
  "mod": "index/user",
  "fqn": "pulumi_azuread",
  "classes": {
   "azuread:index/user:User": "User"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "azuread",
  "token": "pulumi:providers:azuread",
  "fqn": "pulumi_azuread",
  "class": "Provider"
 }
]
"""
)
