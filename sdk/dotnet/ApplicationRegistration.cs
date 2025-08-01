// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    /// <summary>
    /// Manages an application registration within Azure Active Directory.
    /// 
    /// For a more comprehensive alternative, please see the azuread.Application resource. Please note that this resource should not be used together with the `azuread.Application` resource when managing the same application.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`
    /// 
    /// When authenticated with a user principal, this resource may require one of the following directory roles: `Application Administrator` or `Global Administrator`
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.ApplicationRegistration("example", new()
    ///     {
    ///         DisplayName = "Example Application",
    ///         Description = "My example application",
    ///         SignInAudience = "AzureADMyOrg",
    ///         HomepageUrl = "https://app.example.com/",
    ///         LogoutUrl = "https://app.example.com/logout",
    ///         MarketingUrl = "https://example.com/",
    ///         PrivacyStatementUrl = "https://example.com/privacy",
    ///         SupportUrl = "https://support.example.com/",
    ///         TermsOfServiceUrl = "https://example.com/terms",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Application Registrations can be imported using the object ID of the application, in the following format.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/applicationRegistration:ApplicationRegistration example /applications/00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/applicationRegistration:ApplicationRegistration")]
    public partial class ApplicationRegistration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Client ID for the application, which is globally unique.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// A description of the application, as shown to end users.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
        /// </summary>
        [Output("disabledByMicrosoft")]
        public Output<string> DisabledByMicrosoft { get; private set; } = null!;

        /// <summary>
        /// The display name for the application.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
        /// </summary>
        [Output("groupMembershipClaims")]
        public Output<ImmutableArray<string>> GroupMembershipClaims { get; private set; } = null!;

        /// <summary>
        /// Home page or landing page of the application.
        /// </summary>
        [Output("homepageUrl")]
        public Output<string?> HomepageUrl { get; private set; } = null!;

        /// <summary>
        /// Whether this web application can request an access token using OAuth implicit flow.
        /// </summary>
        [Output("implicitAccessTokenIssuanceEnabled")]
        public Output<bool?> ImplicitAccessTokenIssuanceEnabled { get; private set; } = null!;

        /// <summary>
        /// Whether this web application can request an ID token using OAuth implicit flow.
        /// </summary>
        [Output("implicitIdTokenIssuanceEnabled")]
        public Output<bool?> ImplicitIdTokenIssuanceEnabled { get; private set; } = null!;

        /// <summary>
        /// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
        /// </summary>
        [Output("logoutUrl")]
        public Output<string?> LogoutUrl { get; private set; } = null!;

        /// <summary>
        /// URL of the marketing page for the application.
        /// </summary>
        [Output("marketingUrl")]
        public Output<string?> MarketingUrl { get; private set; } = null!;

        /// <summary>
        /// User-specified notes relevant for the management of the application.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// The object ID of the application within the tenant.
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// URL of the privacy statement for the application.
        /// </summary>
        [Output("privacyStatementUrl")]
        public Output<string?> PrivacyStatementUrl { get; private set; } = null!;

        /// <summary>
        /// The verified publisher domain for the application.
        /// </summary>
        [Output("publisherDomain")]
        public Output<string> PublisherDomain { get; private set; } = null!;

        /// <summary>
        /// The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `sign_in_audience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
        /// </summary>
        [Output("requestedAccessTokenVersion")]
        public Output<int?> RequestedAccessTokenVersion { get; private set; } = null!;

        /// <summary>
        /// References application context information from a Service or Asset Management database.
        /// </summary>
        [Output("serviceManagementReference")]
        public Output<string?> ServiceManagementReference { get; private set; } = null!;

        /// <summary>
        /// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
        /// </summary>
        [Output("signInAudience")]
        public Output<string?> SignInAudience { get; private set; } = null!;

        /// <summary>
        /// URL of the support page for the application.
        /// </summary>
        [Output("supportUrl")]
        public Output<string?> SupportUrl { get; private set; } = null!;

        /// <summary>
        /// URL of the terms of service statement for the application.
        /// </summary>
        [Output("termsOfServiceUrl")]
        public Output<string?> TermsOfServiceUrl { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationRegistration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationRegistration(string name, ApplicationRegistrationArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/applicationRegistration:ApplicationRegistration", name, args ?? new ApplicationRegistrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationRegistration(string name, Input<string> id, ApplicationRegistrationState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/applicationRegistration:ApplicationRegistration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApplicationRegistration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationRegistration Get(string name, Input<string> id, ApplicationRegistrationState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationRegistration(name, id, state, options);
        }
    }

    public sealed class ApplicationRegistrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the application, as shown to end users.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name for the application.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("groupMembershipClaims")]
        private InputList<string>? _groupMembershipClaims;

        /// <summary>
        /// Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
        /// </summary>
        public InputList<string> GroupMembershipClaims
        {
            get => _groupMembershipClaims ?? (_groupMembershipClaims = new InputList<string>());
            set => _groupMembershipClaims = value;
        }

        /// <summary>
        /// Home page or landing page of the application.
        /// </summary>
        [Input("homepageUrl")]
        public Input<string>? HomepageUrl { get; set; }

        /// <summary>
        /// Whether this web application can request an access token using OAuth implicit flow.
        /// </summary>
        [Input("implicitAccessTokenIssuanceEnabled")]
        public Input<bool>? ImplicitAccessTokenIssuanceEnabled { get; set; }

        /// <summary>
        /// Whether this web application can request an ID token using OAuth implicit flow.
        /// </summary>
        [Input("implicitIdTokenIssuanceEnabled")]
        public Input<bool>? ImplicitIdTokenIssuanceEnabled { get; set; }

        /// <summary>
        /// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
        /// </summary>
        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        /// <summary>
        /// URL of the marketing page for the application.
        /// </summary>
        [Input("marketingUrl")]
        public Input<string>? MarketingUrl { get; set; }

        /// <summary>
        /// User-specified notes relevant for the management of the application.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// URL of the privacy statement for the application.
        /// </summary>
        [Input("privacyStatementUrl")]
        public Input<string>? PrivacyStatementUrl { get; set; }

        /// <summary>
        /// The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `sign_in_audience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
        /// </summary>
        [Input("requestedAccessTokenVersion")]
        public Input<int>? RequestedAccessTokenVersion { get; set; }

        /// <summary>
        /// References application context information from a Service or Asset Management database.
        /// </summary>
        [Input("serviceManagementReference")]
        public Input<string>? ServiceManagementReference { get; set; }

        /// <summary>
        /// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
        /// </summary>
        [Input("signInAudience")]
        public Input<string>? SignInAudience { get; set; }

        /// <summary>
        /// URL of the support page for the application.
        /// </summary>
        [Input("supportUrl")]
        public Input<string>? SupportUrl { get; set; }

        /// <summary>
        /// URL of the terms of service statement for the application.
        /// </summary>
        [Input("termsOfServiceUrl")]
        public Input<string>? TermsOfServiceUrl { get; set; }

        public ApplicationRegistrationArgs()
        {
        }
        public static new ApplicationRegistrationArgs Empty => new ApplicationRegistrationArgs();
    }

    public sealed class ApplicationRegistrationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Client ID for the application, which is globally unique.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// A description of the application, as shown to end users.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. `DisabledDueToViolationOfServicesAgreement`
        /// </summary>
        [Input("disabledByMicrosoft")]
        public Input<string>? DisabledByMicrosoft { get; set; }

        /// <summary>
        /// The display name for the application.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("groupMembershipClaims")]
        private InputList<string>? _groupMembershipClaims;

        /// <summary>
        /// Configures the `groups` claim issued in a user or OAuth access token that the app expects. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
        /// </summary>
        public InputList<string> GroupMembershipClaims
        {
            get => _groupMembershipClaims ?? (_groupMembershipClaims = new InputList<string>());
            set => _groupMembershipClaims = value;
        }

        /// <summary>
        /// Home page or landing page of the application.
        /// </summary>
        [Input("homepageUrl")]
        public Input<string>? HomepageUrl { get; set; }

        /// <summary>
        /// Whether this web application can request an access token using OAuth implicit flow.
        /// </summary>
        [Input("implicitAccessTokenIssuanceEnabled")]
        public Input<bool>? ImplicitAccessTokenIssuanceEnabled { get; set; }

        /// <summary>
        /// Whether this web application can request an ID token using OAuth implicit flow.
        /// </summary>
        [Input("implicitIdTokenIssuanceEnabled")]
        public Input<bool>? ImplicitIdTokenIssuanceEnabled { get; set; }

        /// <summary>
        /// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
        /// </summary>
        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        /// <summary>
        /// URL of the marketing page for the application.
        /// </summary>
        [Input("marketingUrl")]
        public Input<string>? MarketingUrl { get; set; }

        /// <summary>
        /// User-specified notes relevant for the management of the application.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// The object ID of the application within the tenant.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// URL of the privacy statement for the application.
        /// </summary>
        [Input("privacyStatementUrl")]
        public Input<string>? PrivacyStatementUrl { get; set; }

        /// <summary>
        /// The verified publisher domain for the application.
        /// </summary>
        [Input("publisherDomain")]
        public Input<string>? PublisherDomain { get; set; }

        /// <summary>
        /// The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `sign_in_audience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `2`.
        /// </summary>
        [Input("requestedAccessTokenVersion")]
        public Input<int>? RequestedAccessTokenVersion { get; set; }

        /// <summary>
        /// References application context information from a Service or Asset Management database.
        /// </summary>
        [Input("serviceManagementReference")]
        public Input<string>? ServiceManagementReference { get; set; }

        /// <summary>
        /// The Microsoft account types that are supported for the current application. Must be one of `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`. Defaults to `AzureADMyOrg`.
        /// </summary>
        [Input("signInAudience")]
        public Input<string>? SignInAudience { get; set; }

        /// <summary>
        /// URL of the support page for the application.
        /// </summary>
        [Input("supportUrl")]
        public Input<string>? SupportUrl { get; set; }

        /// <summary>
        /// URL of the terms of service statement for the application.
        /// </summary>
        [Input("termsOfServiceUrl")]
        public Input<string>? TermsOfServiceUrl { get; set; }

        public ApplicationRegistrationState()
        {
        }
        public static new ApplicationRegistrationState Empty => new ApplicationRegistrationState();
    }
}
