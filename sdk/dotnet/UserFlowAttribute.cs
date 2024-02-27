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
    /// Manages user flow attributes in an Azure Active Directory (Azure AD) tenant.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires the following application role: `IdentityUserFlow.ReadWrite.All`
    /// 
    /// ## Example Usage
    /// 
    /// *Basic example*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AzureAD.UserFlowAttribute("example", new()
    ///     {
    ///         DisplayName = "Hobby",
    ///         Description = "Your hobby",
    ///         DataType = "string",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// User flow attributes can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/userFlowAttribute:UserFlowAttribute example extension_ecc9f88db2924942b8a96f44873616fe_Hobbyjkorv
    /// ```
    /// 
    ///  -&gt; This ID can be queried using the [User Flow Attributes API](https://learn.microsoft.com/en-us/graph/api/identityuserflowattribute-list?view=graph-rest-1.0&amp;tabs=http).
    /// </summary>
    [AzureADResourceType("azuread:index/userFlowAttribute:UserFlowAttribute")]
    public partial class UserFlowAttribute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The type of the user flow attribute. Values include `builtIn`, `custom` or `required`.
        /// </summary>
        [Output("attributeType")]
        public Output<string> AttributeType { get; private set; } = null!;

        /// <summary>
        /// The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("dataType")]
        public Output<string> DataType { get; private set; } = null!;

        /// <summary>
        /// The description of the user flow attribute that is shown to the user at the time of sign-up.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the user flow attribute. Changing this forces a new resource to be created.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;


        /// <summary>
        /// Create a UserFlowAttribute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserFlowAttribute(string name, UserFlowAttributeArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/userFlowAttribute:UserFlowAttribute", name, args ?? new UserFlowAttributeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserFlowAttribute(string name, Input<string> id, UserFlowAttributeState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/userFlowAttribute:UserFlowAttribute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserFlowAttribute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserFlowAttribute Get(string name, Input<string> id, UserFlowAttributeState? state = null, CustomResourceOptions? options = null)
        {
            return new UserFlowAttribute(name, id, state, options);
        }
    }

    public sealed class UserFlowAttributeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataType", required: true)]
        public Input<string> DataType { get; set; } = null!;

        /// <summary>
        /// The description of the user flow attribute that is shown to the user at the time of sign-up.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The display name of the user flow attribute. Changing this forces a new resource to be created.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        public UserFlowAttributeArgs()
        {
        }
        public static new UserFlowAttributeArgs Empty => new UserFlowAttributeArgs();
    }

    public sealed class UserFlowAttributeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the user flow attribute. Values include `builtIn`, `custom` or `required`.
        /// </summary>
        [Input("attributeType")]
        public Input<string>? AttributeType { get; set; }

        /// <summary>
        /// The data type of the user flow attribute. Possible values are `boolean`, `dateTime`, `int64`, `string` or `stringCollection`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataType")]
        public Input<string>? DataType { get; set; }

        /// <summary>
        /// The description of the user flow attribute that is shown to the user at the time of sign-up.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the user flow attribute. Changing this forces a new resource to be created.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        public UserFlowAttributeState()
        {
        }
        public static new UserFlowAttributeState Empty => new UserFlowAttributeState();
    }
}
