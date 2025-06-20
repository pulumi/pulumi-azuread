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
    ///     var exampleApplicationFallbackPublicClient = new AzureAD.ApplicationFallbackPublicClient("example", new()
    ///     {
    ///         ApplicationId = example.Id,
    ///         Enabled = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The Application Fallback Public Client setting can be imported using the object ID of the application, in the following format.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/applicationFallbackPublicClient:ApplicationFallbackPublicClient example /applications/00000000-0000-0000-0000-000000000000/fallbackPublicClient
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/applicationFallbackPublicClient:ApplicationFallbackPublicClient")]
    public partial class ApplicationFallbackPublicClient : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the application as a fallback public client.
        /// 
        /// &gt; Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationFallbackPublicClient resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationFallbackPublicClient(string name, ApplicationFallbackPublicClientArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/applicationFallbackPublicClient:ApplicationFallbackPublicClient", name, args ?? new ApplicationFallbackPublicClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationFallbackPublicClient(string name, Input<string> id, ApplicationFallbackPublicClientState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/applicationFallbackPublicClient:ApplicationFallbackPublicClient", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationFallbackPublicClient resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationFallbackPublicClient Get(string name, Input<string> id, ApplicationFallbackPublicClientState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationFallbackPublicClient(name, id, state, options);
        }
    }

    public sealed class ApplicationFallbackPublicClientArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// Whether to enable the application as a fallback public client.
        /// 
        /// &gt; Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ApplicationFallbackPublicClientArgs()
        {
        }
        public static new ApplicationFallbackPublicClientArgs Empty => new ApplicationFallbackPublicClientArgs();
    }

    public sealed class ApplicationFallbackPublicClientState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// Whether to enable the application as a fallback public client.
        /// 
        /// &gt; Some configurations may require the Fallback Public Client setting to be `null`, for this case simply destroy this resource (or don't use it)
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ApplicationFallbackPublicClientState()
        {
        }
        public static new ApplicationFallbackPublicClientState Empty => new ApplicationFallbackPublicClientState();
    }
}
