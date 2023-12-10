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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.%[1]s;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var wellKnown = AzureAD.GetApplicationPublishedAppIds.Invoke();
    /// 
    ///     var msgraph = AzureAD.GetServicePrincipal.Invoke(new()
    ///     {
    ///         ClientId = wellKnown.Apply(getApplicationPublishedAppIdsResult =&gt; getApplicationPublishedAppIdsResult.Result?.MicrosoftGraph),
    ///     });
    /// 
    ///     var example = new AzureAD.ApplicationRegistration("example", new()
    ///     {
    ///         DisplayName = "example",
    ///     });
    /// 
    ///     var exampleMsgraph = new AzureAD.ApplicationApiAccess("exampleMsgraph", new()
    ///     {
    ///         ApplicationId = example.Id,
    ///         ApiClientId = wellKnown.Apply(getApplicationPublishedAppIdsResult =&gt; getApplicationPublishedAppIdsResult.Result?.MicrosoftGraph),
    ///         RoleIds = new[]
    ///         {
    ///             msgraph.Apply(getServicePrincipalResult =&gt; getServicePrincipalResult.AppRoleIds?.Group_Read_All),
    ///             msgraph.Apply(getServicePrincipalResult =&gt; getServicePrincipalResult.AppRoleIds?.User_Read_All),
    ///         },
    ///         ScopeIds = new[]
    ///         {
    ///             msgraph.Apply(getServicePrincipalResult =&gt; getServicePrincipalResult.Oauth2PermissionScopeIds?.User_ReadWrite),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &gt; **Tip** For managing permissions for an additional API, create another instance of this resource
    /// 
    /// *Usage with azuread.Application resource*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleApplication = new AzureAD.Application("exampleApplication", new()
    ///     {
    ///         DisplayName = "example",
    ///     });
    /// 
    ///     var exampleApplicationApiAccess = new AzureAD.ApplicationApiAccess("exampleApplicationApiAccess", new()
    ///     {
    ///         ApplicationId = exampleApplication.Id,
    ///     });
    /// 
    ///     // ...
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Application API Access can be imported using the object ID of the application and the client ID of the API, in the following format.
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/applicationApiAccess:ApplicationApiAccess example /applications/00000000-0000-0000-0000-000000000000/apiAccess/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/applicationApiAccess:ApplicationApiAccess")]
    public partial class ApplicationApiAccess : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
        /// </summary>
        [Output("apiClientId")]
        public Output<string> ApiClientId { get; private set; } = null!;

        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// A set of role IDs to be granted to the application, as published by the API.
        /// </summary>
        [Output("roleIds")]
        public Output<ImmutableArray<string>> RoleIds { get; private set; } = null!;

        /// <summary>
        /// A set of scope IDs to be granted to the application, as published by the API.
        /// 
        /// &gt; At least one of `role_ids` or `scope_ids` must be specified.
        /// </summary>
        [Output("scopeIds")]
        public Output<ImmutableArray<string>> ScopeIds { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationApiAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationApiAccess(string name, ApplicationApiAccessArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/applicationApiAccess:ApplicationApiAccess", name, args ?? new ApplicationApiAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationApiAccess(string name, Input<string> id, ApplicationApiAccessState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/applicationApiAccess:ApplicationApiAccess", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationApiAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationApiAccess Get(string name, Input<string> id, ApplicationApiAccessState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationApiAccess(name, id, state, options);
        }
    }

    public sealed class ApplicationApiAccessArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiClientId", required: true)]
        public Input<string> ApiClientId { get; set; } = null!;

        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        [Input("roleIds")]
        private InputList<string>? _roleIds;

        /// <summary>
        /// A set of role IDs to be granted to the application, as published by the API.
        /// </summary>
        public InputList<string> RoleIds
        {
            get => _roleIds ?? (_roleIds = new InputList<string>());
            set => _roleIds = value;
        }

        [Input("scopeIds")]
        private InputList<string>? _scopeIds;

        /// <summary>
        /// A set of scope IDs to be granted to the application, as published by the API.
        /// 
        /// &gt; At least one of `role_ids` or `scope_ids` must be specified.
        /// </summary>
        public InputList<string> ScopeIds
        {
            get => _scopeIds ?? (_scopeIds = new InputList<string>());
            set => _scopeIds = value;
        }

        public ApplicationApiAccessArgs()
        {
        }
        public static new ApplicationApiAccessArgs Empty => new ApplicationApiAccessArgs();
    }

    public sealed class ApplicationApiAccessState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client ID of the API to which access is being granted. Changing this forces a new resource to be created.
        /// </summary>
        [Input("apiClientId")]
        public Input<string>? ApiClientId { get; set; }

        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        [Input("roleIds")]
        private InputList<string>? _roleIds;

        /// <summary>
        /// A set of role IDs to be granted to the application, as published by the API.
        /// </summary>
        public InputList<string> RoleIds
        {
            get => _roleIds ?? (_roleIds = new InputList<string>());
            set => _roleIds = value;
        }

        [Input("scopeIds")]
        private InputList<string>? _scopeIds;

        /// <summary>
        /// A set of scope IDs to be granted to the application, as published by the API.
        /// 
        /// &gt; At least one of `role_ids` or `scope_ids` must be specified.
        /// </summary>
        public InputList<string> ScopeIds
        {
            get => _scopeIds ?? (_scopeIds = new InputList<string>());
            set => _scopeIds = value;
        }

        public ApplicationApiAccessState()
        {
        }
        public static new ApplicationApiAccessState Empty => new ApplicationApiAccessState();
    }
}
