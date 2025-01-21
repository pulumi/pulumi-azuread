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
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.ApplicationRegistration("example", new()
    ///     {
    ///         DisplayName = "example",
    ///     });
    /// 
    ///     var examplePublic = new AzureAD.ApplicationRedirectUris("example_public", new()
    ///     {
    ///         ApplicationId = example.Id,
    ///         Type = "PublicClient",
    ///         RedirectUris = new[]
    ///         {
    ///             "myapp://auth",
    ///             "sample.mobile.app.bundie.id://auth",
    ///             "https://login.microsoftonline.com/common/oauth2/nativeclient",
    ///             "https://login.live.com/oauth20_desktop.srf",
    ///             "ms-appx-web://Microsoft.AAD.BrokerPlugin/00000000-1111-1111-1111-222222222222",
    ///             "urn:ietf:wg:oauth:2.0:foo",
    ///         },
    ///     });
    /// 
    ///     var exampleSpa = new AzureAD.ApplicationRedirectUris("example_spa", new()
    ///     {
    ///         ApplicationId = example.Id,
    ///         Type = "SPA",
    ///         RedirectUris = new[]
    ///         {
    ///             "https://mobile.example.com/",
    ///             "https://beta.example.com/",
    ///         },
    ///     });
    /// 
    ///     var exampleWeb = new AzureAD.ApplicationRedirectUris("example_web", new()
    ///     {
    ///         ApplicationId = example.Id,
    ///         Type = "Web",
    ///         RedirectUris = new[]
    ///         {
    ///             "https://app.example.com/",
    ///             "https://classic.example.com/",
    ///             "urn:ietf:wg:oauth:2.0:oob",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Application API Access can be imported using the object ID of the application and the URI type, in the following format.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/applicationRedirectUris:ApplicationRedirectUris example /applications/00000000-0000-0000-0000-000000000000/redirectUris/Web
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/applicationRedirectUris:ApplicationRedirectUris")]
    public partial class ApplicationRedirectUris : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// A set of redirect URIs to assign to the application.
        /// </summary>
        [Output("redirectUris")]
        public Output<ImmutableArray<string>> RedirectUris { get; private set; } = null!;

        /// <summary>
        /// The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationRedirectUris resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationRedirectUris(string name, ApplicationRedirectUrisArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/applicationRedirectUris:ApplicationRedirectUris", name, args ?? new ApplicationRedirectUrisArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationRedirectUris(string name, Input<string> id, ApplicationRedirectUrisState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/applicationRedirectUris:ApplicationRedirectUris", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationRedirectUris resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationRedirectUris Get(string name, Input<string> id, ApplicationRedirectUrisState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationRedirectUris(name, id, state, options);
        }
    }

    public sealed class ApplicationRedirectUrisArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        [Input("redirectUris", required: true)]
        private InputList<string>? _redirectUris;

        /// <summary>
        /// A set of redirect URIs to assign to the application.
        /// </summary>
        public InputList<string> RedirectUris
        {
            get => _redirectUris ?? (_redirectUris = new InputList<string>());
            set => _redirectUris = value;
        }

        /// <summary>
        /// The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ApplicationRedirectUrisArgs()
        {
        }
        public static new ApplicationRedirectUrisArgs Empty => new ApplicationRedirectUrisArgs();
    }

    public sealed class ApplicationRedirectUrisState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        [Input("redirectUris")]
        private InputList<string>? _redirectUris;

        /// <summary>
        /// A set of redirect URIs to assign to the application.
        /// </summary>
        public InputList<string> RedirectUris
        {
            get => _redirectUris ?? (_redirectUris = new InputList<string>());
            set => _redirectUris = value;
        }

        /// <summary>
        /// The type of redirect URIs to manage. Must be one of: `PublicClient`, `SPA`, or `Web`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ApplicationRedirectUrisState()
        {
        }
        public static new ApplicationRedirectUrisState Empty => new ApplicationRedirectUrisState();
    }
}
