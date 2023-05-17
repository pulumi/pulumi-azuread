// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    /// <summary>
    /// Manages a delegated permission grant for a service principal, on behalf of a single user, or all users.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires the following application role: `Directory.ReadWrite.All`
    /// 
    /// When authenticated with a user principal, this resource requires one the following directory role: `Global Administrator`
    /// 
    /// ## Example Usage
    /// 
    /// *Delegated permission grant for all users*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var wellKnown = AzureAD.GetApplicationPublishedAppIds.Invoke();
    /// 
    ///     var msgraph = new AzureAD.ServicePrincipal("msgraph", new()
    ///     {
    ///         ApplicationId = wellKnown.Apply(getApplicationPublishedAppIdsResult =&gt; getApplicationPublishedAppIdsResult.Result?.MicrosoftGraph),
    ///         UseExisting = true,
    ///     });
    /// 
    ///     var exampleApplication = new AzureAD.Application("exampleApplication", new()
    ///     {
    ///         DisplayName = "example",
    ///         RequiredResourceAccesses = new[]
    ///         {
    ///             new AzureAD.Inputs.ApplicationRequiredResourceAccessArgs
    ///             {
    ///                 ResourceAppId = wellKnown.Apply(getApplicationPublishedAppIdsResult =&gt; getApplicationPublishedAppIdsResult.Result?.MicrosoftGraph),
    ///                 ResourceAccesses = new[]
    ///                 {
    ///                     new AzureAD.Inputs.ApplicationRequiredResourceAccessResourceAccessArgs
    ///                     {
    ///                         Id = msgraph.Oauth2PermissionScopeIds.Apply(oauth2PermissionScopeIds =&gt; oauth2PermissionScopeIds.Openid),
    ///                         Type = "Scope",
    ///                     },
    ///                     new AzureAD.Inputs.ApplicationRequiredResourceAccessResourceAccessArgs
    ///                     {
    ///                         Id = msgraph.Oauth2PermissionScopeIds.Apply(oauth2PermissionScopeIds =&gt; oauth2PermissionScopeIds.User_Read),
    ///                         Type = "Scope",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleServicePrincipal = new AzureAD.ServicePrincipal("exampleServicePrincipal", new()
    ///     {
    ///         ApplicationId = exampleApplication.ApplicationId,
    ///     });
    /// 
    ///     var exampleServicePrincipalDelegatedPermissionGrant = new AzureAD.ServicePrincipalDelegatedPermissionGrant("exampleServicePrincipalDelegatedPermissionGrant", new()
    ///     {
    ///         ServicePrincipalObjectId = exampleServicePrincipal.ObjectId,
    ///         ResourceServicePrincipalObjectId = msgraph.ObjectId,
    ///         ClaimValues = new[]
    ///         {
    ///             "openid",
    ///             "User.Read.All",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// *Delegated permission grant for a single user*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var wellKnown = AzureAD.GetApplicationPublishedAppIds.Invoke();
    /// 
    ///     var msgraph = new AzureAD.ServicePrincipal("msgraph", new()
    ///     {
    ///         ApplicationId = wellKnown.Apply(getApplicationPublishedAppIdsResult =&gt; getApplicationPublishedAppIdsResult.Result?.MicrosoftGraph),
    ///         UseExisting = true,
    ///     });
    /// 
    ///     var exampleApplication = new AzureAD.Application("exampleApplication", new()
    ///     {
    ///         DisplayName = "example",
    ///         RequiredResourceAccesses = new[]
    ///         {
    ///             new AzureAD.Inputs.ApplicationRequiredResourceAccessArgs
    ///             {
    ///                 ResourceAppId = wellKnown.Apply(getApplicationPublishedAppIdsResult =&gt; getApplicationPublishedAppIdsResult.Result?.MicrosoftGraph),
    ///                 ResourceAccesses = new[]
    ///                 {
    ///                     new AzureAD.Inputs.ApplicationRequiredResourceAccessResourceAccessArgs
    ///                     {
    ///                         Id = msgraph.Oauth2PermissionScopeIds.Apply(oauth2PermissionScopeIds =&gt; oauth2PermissionScopeIds.Openid),
    ///                         Type = "Scope",
    ///                     },
    ///                     new AzureAD.Inputs.ApplicationRequiredResourceAccessResourceAccessArgs
    ///                     {
    ///                         Id = msgraph.Oauth2PermissionScopeIds.Apply(oauth2PermissionScopeIds =&gt; oauth2PermissionScopeIds.User_Read),
    ///                         Type = "Scope",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleServicePrincipal = new AzureAD.ServicePrincipal("exampleServicePrincipal", new()
    ///     {
    ///         ApplicationId = exampleApplication.ApplicationId,
    ///     });
    /// 
    ///     var exampleUser = new AzureAD.User("exampleUser", new()
    ///     {
    ///         DisplayName = "J. Doe",
    ///         UserPrincipalName = "jdoe@hashicorp.com",
    ///         MailNickname = "jdoe",
    ///         Password = "SecretP@sswd99!",
    ///     });
    /// 
    ///     var exampleServicePrincipalDelegatedPermissionGrant = new AzureAD.ServicePrincipalDelegatedPermissionGrant("exampleServicePrincipalDelegatedPermissionGrant", new()
    ///     {
    ///         ServicePrincipalObjectId = exampleServicePrincipal.ObjectId,
    ///         ResourceServicePrincipalObjectId = msgraph.ObjectId,
    ///         ClaimValues = new[]
    ///         {
    ///             "openid",
    ///             "User.Read.All",
    ///         },
    ///         UserObjectId = exampleUser.ObjectId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Delegated permission grants can be imported using their ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant example aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant")]
    public partial class ServicePrincipalDelegatedPermissionGrant : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
        /// </summary>
        [Output("claimValues")]
        public Output<ImmutableArray<string>> ClaimValues { get; private set; } = null!;

        /// <summary>
        /// The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceServicePrincipalObjectId")]
        public Output<string> ResourceServicePrincipalObjectId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("servicePrincipalObjectId")]
        public Output<string> ServicePrincipalObjectId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
        /// 
        /// &gt; **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
        /// </summary>
        [Output("userObjectId")]
        public Output<string?> UserObjectId { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePrincipalDelegatedPermissionGrant resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePrincipalDelegatedPermissionGrant(string name, ServicePrincipalDelegatedPermissionGrantArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant", name, args ?? new ServicePrincipalDelegatedPermissionGrantArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServicePrincipalDelegatedPermissionGrant(string name, Input<string> id, ServicePrincipalDelegatedPermissionGrantState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServicePrincipalDelegatedPermissionGrant resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePrincipalDelegatedPermissionGrant Get(string name, Input<string> id, ServicePrincipalDelegatedPermissionGrantState? state = null, CustomResourceOptions? options = null)
        {
            return new ServicePrincipalDelegatedPermissionGrant(name, id, state, options);
        }
    }

    public sealed class ServicePrincipalDelegatedPermissionGrantArgs : global::Pulumi.ResourceArgs
    {
        [Input("claimValues", required: true)]
        private InputList<string>? _claimValues;

        /// <summary>
        /// A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
        /// </summary>
        public InputList<string> ClaimValues
        {
            get => _claimValues ?? (_claimValues = new InputList<string>());
            set => _claimValues = value;
        }

        /// <summary>
        /// The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceServicePrincipalObjectId", required: true)]
        public Input<string> ResourceServicePrincipalObjectId { get; set; } = null!;

        /// <summary>
        /// The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("servicePrincipalObjectId", required: true)]
        public Input<string> ServicePrincipalObjectId { get; set; } = null!;

        /// <summary>
        /// The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
        /// 
        /// &gt; **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
        /// </summary>
        [Input("userObjectId")]
        public Input<string>? UserObjectId { get; set; }

        public ServicePrincipalDelegatedPermissionGrantArgs()
        {
        }
        public static new ServicePrincipalDelegatedPermissionGrantArgs Empty => new ServicePrincipalDelegatedPermissionGrantArgs();
    }

    public sealed class ServicePrincipalDelegatedPermissionGrantState : global::Pulumi.ResourceArgs
    {
        [Input("claimValues")]
        private InputList<string>? _claimValues;

        /// <summary>
        /// A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
        /// </summary>
        public InputList<string> ClaimValues
        {
            get => _claimValues ?? (_claimValues = new InputList<string>());
            set => _claimValues = value;
        }

        /// <summary>
        /// The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceServicePrincipalObjectId")]
        public Input<string>? ResourceServicePrincipalObjectId { get; set; }

        /// <summary>
        /// The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("servicePrincipalObjectId")]
        public Input<string>? ServicePrincipalObjectId { get; set; }

        /// <summary>
        /// The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
        /// 
        /// &gt; **Granting Admin Consent** To grant admin consent for the service principal to impersonate all users, just omit the `user_object_id` property.
        /// </summary>
        [Input("userObjectId")]
        public Input<string>? UserObjectId { get; set; }

        public ServicePrincipalDelegatedPermissionGrantState()
        {
        }
        public static new ServicePrincipalDelegatedPermissionGrantState Empty => new ServicePrincipalDelegatedPermissionGrantState();
    }
}
