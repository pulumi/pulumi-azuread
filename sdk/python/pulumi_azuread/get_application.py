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

__all__ = [
    'GetApplicationResult',
    'AwaitableGetApplicationResult',
    'get_application',
    'get_application_output',
]

@pulumi.output_type
class GetApplicationResult:
    """
    A collection of values returned by getApplication.
    """
    def __init__(__self__, apis=None, app_role_ids=None, app_roles=None, application_id=None, client_id=None, description=None, device_only_auth_enabled=None, disabled_by_microsoft=None, display_name=None, fallback_public_client_enabled=None, feature_tags=None, group_membership_claims=None, id=None, identifier_uri=None, identifier_uris=None, logo_url=None, marketing_url=None, notes=None, oauth2_permission_scope_ids=None, oauth2_post_response_required=None, object_id=None, optional_claims=None, owners=None, privacy_statement_url=None, public_clients=None, publisher_domain=None, required_resource_accesses=None, service_management_reference=None, sign_in_audience=None, single_page_applications=None, support_url=None, tags=None, terms_of_service_url=None, webs=None):
        if apis and not isinstance(apis, list):
            raise TypeError("Expected argument 'apis' to be a list")
        pulumi.set(__self__, "apis", apis)
        if app_role_ids and not isinstance(app_role_ids, dict):
            raise TypeError("Expected argument 'app_role_ids' to be a dict")
        pulumi.set(__self__, "app_role_ids", app_role_ids)
        if app_roles and not isinstance(app_roles, list):
            raise TypeError("Expected argument 'app_roles' to be a list")
        pulumi.set(__self__, "app_roles", app_roles)
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if device_only_auth_enabled and not isinstance(device_only_auth_enabled, bool):
            raise TypeError("Expected argument 'device_only_auth_enabled' to be a bool")
        pulumi.set(__self__, "device_only_auth_enabled", device_only_auth_enabled)
        if disabled_by_microsoft and not isinstance(disabled_by_microsoft, str):
            raise TypeError("Expected argument 'disabled_by_microsoft' to be a str")
        pulumi.set(__self__, "disabled_by_microsoft", disabled_by_microsoft)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if fallback_public_client_enabled and not isinstance(fallback_public_client_enabled, bool):
            raise TypeError("Expected argument 'fallback_public_client_enabled' to be a bool")
        pulumi.set(__self__, "fallback_public_client_enabled", fallback_public_client_enabled)
        if feature_tags and not isinstance(feature_tags, list):
            raise TypeError("Expected argument 'feature_tags' to be a list")
        pulumi.set(__self__, "feature_tags", feature_tags)
        if group_membership_claims and not isinstance(group_membership_claims, list):
            raise TypeError("Expected argument 'group_membership_claims' to be a list")
        pulumi.set(__self__, "group_membership_claims", group_membership_claims)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identifier_uri and not isinstance(identifier_uri, str):
            raise TypeError("Expected argument 'identifier_uri' to be a str")
        pulumi.set(__self__, "identifier_uri", identifier_uri)
        if identifier_uris and not isinstance(identifier_uris, list):
            raise TypeError("Expected argument 'identifier_uris' to be a list")
        pulumi.set(__self__, "identifier_uris", identifier_uris)
        if logo_url and not isinstance(logo_url, str):
            raise TypeError("Expected argument 'logo_url' to be a str")
        pulumi.set(__self__, "logo_url", logo_url)
        if marketing_url and not isinstance(marketing_url, str):
            raise TypeError("Expected argument 'marketing_url' to be a str")
        pulumi.set(__self__, "marketing_url", marketing_url)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if oauth2_permission_scope_ids and not isinstance(oauth2_permission_scope_ids, dict):
            raise TypeError("Expected argument 'oauth2_permission_scope_ids' to be a dict")
        pulumi.set(__self__, "oauth2_permission_scope_ids", oauth2_permission_scope_ids)
        if oauth2_post_response_required and not isinstance(oauth2_post_response_required, bool):
            raise TypeError("Expected argument 'oauth2_post_response_required' to be a bool")
        pulumi.set(__self__, "oauth2_post_response_required", oauth2_post_response_required)
        if object_id and not isinstance(object_id, str):
            raise TypeError("Expected argument 'object_id' to be a str")
        pulumi.set(__self__, "object_id", object_id)
        if optional_claims and not isinstance(optional_claims, list):
            raise TypeError("Expected argument 'optional_claims' to be a list")
        pulumi.set(__self__, "optional_claims", optional_claims)
        if owners and not isinstance(owners, list):
            raise TypeError("Expected argument 'owners' to be a list")
        pulumi.set(__self__, "owners", owners)
        if privacy_statement_url and not isinstance(privacy_statement_url, str):
            raise TypeError("Expected argument 'privacy_statement_url' to be a str")
        pulumi.set(__self__, "privacy_statement_url", privacy_statement_url)
        if public_clients and not isinstance(public_clients, list):
            raise TypeError("Expected argument 'public_clients' to be a list")
        pulumi.set(__self__, "public_clients", public_clients)
        if publisher_domain and not isinstance(publisher_domain, str):
            raise TypeError("Expected argument 'publisher_domain' to be a str")
        pulumi.set(__self__, "publisher_domain", publisher_domain)
        if required_resource_accesses and not isinstance(required_resource_accesses, list):
            raise TypeError("Expected argument 'required_resource_accesses' to be a list")
        pulumi.set(__self__, "required_resource_accesses", required_resource_accesses)
        if service_management_reference and not isinstance(service_management_reference, str):
            raise TypeError("Expected argument 'service_management_reference' to be a str")
        pulumi.set(__self__, "service_management_reference", service_management_reference)
        if sign_in_audience and not isinstance(sign_in_audience, str):
            raise TypeError("Expected argument 'sign_in_audience' to be a str")
        pulumi.set(__self__, "sign_in_audience", sign_in_audience)
        if single_page_applications and not isinstance(single_page_applications, list):
            raise TypeError("Expected argument 'single_page_applications' to be a list")
        pulumi.set(__self__, "single_page_applications", single_page_applications)
        if support_url and not isinstance(support_url, str):
            raise TypeError("Expected argument 'support_url' to be a str")
        pulumi.set(__self__, "support_url", support_url)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if terms_of_service_url and not isinstance(terms_of_service_url, str):
            raise TypeError("Expected argument 'terms_of_service_url' to be a str")
        pulumi.set(__self__, "terms_of_service_url", terms_of_service_url)
        if webs and not isinstance(webs, list):
            raise TypeError("Expected argument 'webs' to be a list")
        pulumi.set(__self__, "webs", webs)

    @property
    @pulumi.getter
    def apis(self) -> Sequence['outputs.GetApplicationApiResult']:
        """
        An `api` block as documented below.
        """
        return pulumi.get(self, "apis")

    @property
    @pulumi.getter(name="appRoleIds")
    def app_role_ids(self) -> Mapping[str, str]:
        """
        A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
        """
        return pulumi.get(self, "app_role_ids")

    @property
    @pulumi.getter(name="appRoles")
    def app_roles(self) -> Sequence['outputs.GetApplicationAppRoleResult']:
        """
        A collection of `app_role` blocks as documented below. For more information see [official documentation on Application Roles](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
        """
        return pulumi.get(self, "app_roles")

    @property
    @pulumi.getter(name="applicationId")
    @_utilities.deprecated("""The `application_id` property has been replaced with the `client_id` property and will be removed in version 3.0 of the AzureAD provider""")
    def application_id(self) -> str:
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        The Client ID for the application.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceOnlyAuthEnabled")
    def device_only_auth_enabled(self) -> bool:
        """
        Specifies whether this application supports device authentication without a user.
        """
        return pulumi.get(self, "device_only_auth_enabled")

    @property
    @pulumi.getter(name="disabledByMicrosoft")
    def disabled_by_microsoft(self) -> str:
        """
        Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
        """
        return pulumi.get(self, "disabled_by_microsoft")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name for the app role that appears during app role assignment and in consent experiences.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="fallbackPublicClientEnabled")
    def fallback_public_client_enabled(self) -> bool:
        """
        The fallback application type as public client, such as an installed application running on a mobile device.
        """
        return pulumi.get(self, "fallback_public_client_enabled")

    @property
    @pulumi.getter(name="featureTags")
    def feature_tags(self) -> Sequence['outputs.GetApplicationFeatureTagResult']:
        """
        A `features` block as described below.
        """
        return pulumi.get(self, "feature_tags")

    @property
    @pulumi.getter(name="groupMembershipClaims")
    def group_membership_claims(self) -> Sequence[str]:
        """
        The `groups` claim issued in a user or OAuth 2.0 access token that the app expects.
        """
        return pulumi.get(self, "group_membership_claims")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identifierUri")
    def identifier_uri(self) -> str:
        return pulumi.get(self, "identifier_uri")

    @property
    @pulumi.getter(name="identifierUris")
    def identifier_uris(self) -> Sequence[str]:
        """
        A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
        """
        return pulumi.get(self, "identifier_uris")

    @property
    @pulumi.getter(name="logoUrl")
    def logo_url(self) -> str:
        """
        CDN URL to the application's logo.
        """
        return pulumi.get(self, "logo_url")

    @property
    @pulumi.getter(name="marketingUrl")
    def marketing_url(self) -> str:
        """
        URL of the application's marketing page.
        """
        return pulumi.get(self, "marketing_url")

    @property
    @pulumi.getter
    def notes(self) -> str:
        """
        User-specified notes relevant for the management of the application.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="oauth2PermissionScopeIds")
    def oauth2_permission_scope_ids(self) -> Mapping[str, str]:
        """
        A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
        """
        return pulumi.get(self, "oauth2_permission_scope_ids")

    @property
    @pulumi.getter(name="oauth2PostResponseRequired")
    def oauth2_post_response_required(self) -> bool:
        """
        Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. When `false`, only GET requests are allowed.
        """
        return pulumi.get(self, "oauth2_post_response_required")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> str:
        """
        The application's object ID.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter(name="optionalClaims")
    def optional_claims(self) -> Sequence['outputs.GetApplicationOptionalClaimResult']:
        """
        An `optional_claims` block as documented below.
        """
        return pulumi.get(self, "optional_claims")

    @property
    @pulumi.getter
    def owners(self) -> Sequence[str]:
        """
        A list of object IDs of principals that are assigned ownership of the application.
        """
        return pulumi.get(self, "owners")

    @property
    @pulumi.getter(name="privacyStatementUrl")
    def privacy_statement_url(self) -> str:
        """
        URL of the application's privacy statement.
        """
        return pulumi.get(self, "privacy_statement_url")

    @property
    @pulumi.getter(name="publicClients")
    def public_clients(self) -> Sequence['outputs.GetApplicationPublicClientResult']:
        """
        A `public_client` block as documented below.
        """
        return pulumi.get(self, "public_clients")

    @property
    @pulumi.getter(name="publisherDomain")
    def publisher_domain(self) -> str:
        """
        The verified publisher domain for the application.
        """
        return pulumi.get(self, "publisher_domain")

    @property
    @pulumi.getter(name="requiredResourceAccesses")
    def required_resource_accesses(self) -> Sequence['outputs.GetApplicationRequiredResourceAccessResult']:
        """
        A collection of `required_resource_access` blocks as documented below.
        """
        return pulumi.get(self, "required_resource_accesses")

    @property
    @pulumi.getter(name="serviceManagementReference")
    def service_management_reference(self) -> str:
        """
        References application context information from a Service or Asset Management database.
        """
        return pulumi.get(self, "service_management_reference")

    @property
    @pulumi.getter(name="signInAudience")
    def sign_in_audience(self) -> str:
        """
        The Microsoft account types that are supported for the current application. One of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
        """
        return pulumi.get(self, "sign_in_audience")

    @property
    @pulumi.getter(name="singlePageApplications")
    def single_page_applications(self) -> Sequence['outputs.GetApplicationSinglePageApplicationResult']:
        """
        A `single_page_application` block as documented below.
        """
        return pulumi.get(self, "single_page_applications")

    @property
    @pulumi.getter(name="supportUrl")
    def support_url(self) -> str:
        """
        URL of the application's support page.
        """
        return pulumi.get(self, "support_url")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        A list of tags applied to the application.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="termsOfServiceUrl")
    def terms_of_service_url(self) -> str:
        """
        URL of the application's terms of service statement.
        """
        return pulumi.get(self, "terms_of_service_url")

    @property
    @pulumi.getter
    def webs(self) -> Sequence['outputs.GetApplicationWebResult']:
        """
        A `web` block as documented below.
        """
        return pulumi.get(self, "webs")


class AwaitableGetApplicationResult(GetApplicationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationResult(
            apis=self.apis,
            app_role_ids=self.app_role_ids,
            app_roles=self.app_roles,
            application_id=self.application_id,
            client_id=self.client_id,
            description=self.description,
            device_only_auth_enabled=self.device_only_auth_enabled,
            disabled_by_microsoft=self.disabled_by_microsoft,
            display_name=self.display_name,
            fallback_public_client_enabled=self.fallback_public_client_enabled,
            feature_tags=self.feature_tags,
            group_membership_claims=self.group_membership_claims,
            id=self.id,
            identifier_uri=self.identifier_uri,
            identifier_uris=self.identifier_uris,
            logo_url=self.logo_url,
            marketing_url=self.marketing_url,
            notes=self.notes,
            oauth2_permission_scope_ids=self.oauth2_permission_scope_ids,
            oauth2_post_response_required=self.oauth2_post_response_required,
            object_id=self.object_id,
            optional_claims=self.optional_claims,
            owners=self.owners,
            privacy_statement_url=self.privacy_statement_url,
            public_clients=self.public_clients,
            publisher_domain=self.publisher_domain,
            required_resource_accesses=self.required_resource_accesses,
            service_management_reference=self.service_management_reference,
            sign_in_audience=self.sign_in_audience,
            single_page_applications=self.single_page_applications,
            support_url=self.support_url,
            tags=self.tags,
            terms_of_service_url=self.terms_of_service_url,
            webs=self.webs)


def get_application(application_id: Optional[str] = None,
                    client_id: Optional[str] = None,
                    display_name: Optional[str] = None,
                    identifier_uri: Optional[str] = None,
                    object_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationResult:
    """
    Use this data source to access information about an existing Application within Azure Active Directory.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_application(display_name="My First AzureAD Application")
    pulumi.export("applicationObjectId", example.object_id)
    ```


    :param str client_id: Specifies the Client ID of the application.
    :param str display_name: Specifies the display name of the application.
    :param str identifier_uri: Specifies any identifier URI of the application. See also the `identifier_uris` attribute which contains a list of all identifier URIs for the application.
           
           > One of `client_id`, `display_name`, `object_id`, or `identifier_uri` must be specified.
    :param str object_id: Specifies the Object ID of the application.
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    __args__['clientId'] = client_id
    __args__['displayName'] = display_name
    __args__['identifierUri'] = identifier_uri
    __args__['objectId'] = object_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azuread:index/getApplication:getApplication', __args__, opts=opts, typ=GetApplicationResult).value

    return AwaitableGetApplicationResult(
        apis=pulumi.get(__ret__, 'apis'),
        app_role_ids=pulumi.get(__ret__, 'app_role_ids'),
        app_roles=pulumi.get(__ret__, 'app_roles'),
        application_id=pulumi.get(__ret__, 'application_id'),
        client_id=pulumi.get(__ret__, 'client_id'),
        description=pulumi.get(__ret__, 'description'),
        device_only_auth_enabled=pulumi.get(__ret__, 'device_only_auth_enabled'),
        disabled_by_microsoft=pulumi.get(__ret__, 'disabled_by_microsoft'),
        display_name=pulumi.get(__ret__, 'display_name'),
        fallback_public_client_enabled=pulumi.get(__ret__, 'fallback_public_client_enabled'),
        feature_tags=pulumi.get(__ret__, 'feature_tags'),
        group_membership_claims=pulumi.get(__ret__, 'group_membership_claims'),
        id=pulumi.get(__ret__, 'id'),
        identifier_uri=pulumi.get(__ret__, 'identifier_uri'),
        identifier_uris=pulumi.get(__ret__, 'identifier_uris'),
        logo_url=pulumi.get(__ret__, 'logo_url'),
        marketing_url=pulumi.get(__ret__, 'marketing_url'),
        notes=pulumi.get(__ret__, 'notes'),
        oauth2_permission_scope_ids=pulumi.get(__ret__, 'oauth2_permission_scope_ids'),
        oauth2_post_response_required=pulumi.get(__ret__, 'oauth2_post_response_required'),
        object_id=pulumi.get(__ret__, 'object_id'),
        optional_claims=pulumi.get(__ret__, 'optional_claims'),
        owners=pulumi.get(__ret__, 'owners'),
        privacy_statement_url=pulumi.get(__ret__, 'privacy_statement_url'),
        public_clients=pulumi.get(__ret__, 'public_clients'),
        publisher_domain=pulumi.get(__ret__, 'publisher_domain'),
        required_resource_accesses=pulumi.get(__ret__, 'required_resource_accesses'),
        service_management_reference=pulumi.get(__ret__, 'service_management_reference'),
        sign_in_audience=pulumi.get(__ret__, 'sign_in_audience'),
        single_page_applications=pulumi.get(__ret__, 'single_page_applications'),
        support_url=pulumi.get(__ret__, 'support_url'),
        tags=pulumi.get(__ret__, 'tags'),
        terms_of_service_url=pulumi.get(__ret__, 'terms_of_service_url'),
        webs=pulumi.get(__ret__, 'webs'))


@_utilities.lift_output_func(get_application)
def get_application_output(application_id: Optional[pulumi.Input[Optional[str]]] = None,
                           client_id: Optional[pulumi.Input[Optional[str]]] = None,
                           display_name: Optional[pulumi.Input[Optional[str]]] = None,
                           identifier_uri: Optional[pulumi.Input[Optional[str]]] = None,
                           object_id: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationResult]:
    """
    Use this data source to access information about an existing Application within Azure Active Directory.

    ## API Permissions

    The following API permissions are required in order to use this data source.

    When authenticated with a service principal, this data source requires one of the following application roles: `Application.Read.All` or `Directory.Read.All`

    When authenticated with a user principal, this data source does not require any additional roles.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_azuread as azuread

    example = azuread.get_application(display_name="My First AzureAD Application")
    pulumi.export("applicationObjectId", example.object_id)
    ```


    :param str client_id: Specifies the Client ID of the application.
    :param str display_name: Specifies the display name of the application.
    :param str identifier_uri: Specifies any identifier URI of the application. See also the `identifier_uris` attribute which contains a list of all identifier URIs for the application.
           
           > One of `client_id`, `display_name`, `object_id`, or `identifier_uri` must be specified.
    :param str object_id: Specifies the Object ID of the application.
    """
    ...
