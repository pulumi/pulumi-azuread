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
    /// Manages a single owner of an application registration.
    /// 
    /// &gt; This resource is incompatible with the `azuread.Application` resource, instead use this with the `azuread.ApplicationRegistration` resource.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.OwnedBy` or `Application.ReadWrite.All`
    /// 
    /// &gt; When using the `Application.ReadWrite.OwnedBy` application role, the principal being used to run Pulumi must be an owner of the application.
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
    ///         DisplayName = "example",
    ///     });
    /// 
    ///     var jane = new AzureAD.User("jane", new()
    ///     {
    ///         UserPrincipalName = "jane.fischer@example.com",
    ///         DisplayName = "Jane Fischer",
    ///         Password = "Ch@ngeMe",
    ///     });
    /// 
    ///     var exampleJane = new AzureAD.ApplicationOwner("example_jane", new()
    ///     {
    ///         ApplicationId = example.Id,
    ///         OwnerObjectId = jane.ObjectId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &gt; **Tip** For managing more application owners, create additional instances of this resource
    /// 
    /// ## Import
    /// 
    /// Application Owners can be imported using the object ID of the application and the object ID of the owner, in the following format.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/applicationOwner:ApplicationOwner example /applications/00000000-0000-0000-0000-000000000000/owners/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/applicationOwner:ApplicationOwner")]
    public partial class ApplicationOwner : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
        /// </summary>
        [Output("ownerObjectId")]
        public Output<string> OwnerObjectId { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationOwner resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationOwner(string name, ApplicationOwnerArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/applicationOwner:ApplicationOwner", name, args ?? new ApplicationOwnerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationOwner(string name, Input<string> id, ApplicationOwnerState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/applicationOwner:ApplicationOwner", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationOwner resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationOwner Get(string name, Input<string> id, ApplicationOwnerState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationOwner(name, id, state, options);
        }
    }

    public sealed class ApplicationOwnerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
        /// </summary>
        [Input("ownerObjectId", required: true)]
        public Input<string> OwnerObjectId { get; set; } = null!;

        public ApplicationOwnerArgs()
        {
        }
        public static new ApplicationOwnerArgs Empty => new ApplicationOwnerArgs();
    }

    public sealed class ApplicationOwnerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the application registration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The object ID of the owner to assign to the application, typically a user or service principal. Changing this forces a new resource to be created.
        /// </summary>
        [Input("ownerObjectId")]
        public Input<string>? OwnerObjectId { get; set; }

        public ApplicationOwnerState()
        {
        }
        public static new ApplicationOwnerState Empty => new ApplicationOwnerState();
    }
}
