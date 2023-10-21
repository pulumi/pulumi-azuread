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
    /// Manages a single directory role membership (assignment) within Azure Active Directory.
    /// 
    /// &gt; **Deprecation Warning:** This resource has been superseded by the azuread.DirectoryRoleAssignment resource and will be removed in version 3.0 of the AzureAD provider
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `RoleManagement.ReadWrite.Directory` or `Directory.ReadWrite.All`
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Privileged Role Administrator` or `Global Administrator`
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
    ///     var exampleUser = AzureAD.GetUser.Invoke(new()
    ///     {
    ///         UserPrincipalName = "jdoe@hashicorp.com",
    ///     });
    /// 
    ///     var exampleDirectoryRole = new AzureAD.DirectoryRole("exampleDirectoryRole", new()
    ///     {
    ///         DisplayName = "Security administrator",
    ///     });
    /// 
    ///     var exampleDirectoryRoleMember = new AzureAD.DirectoryRoleMember("exampleDirectoryRoleMember", new()
    ///     {
    ///         RoleObjectId = exampleDirectoryRole.ObjectId,
    ///         MemberObjectId = exampleUser.Apply(getUserResult =&gt; getUserResult.ObjectId),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Directory role members can be imported using the object ID of the role and the object ID of the member, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/directoryRoleMember:DirectoryRoleMember example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
    /// ```
    /// 
    ///  -&gt; This ID format is unique to Terraform and is composed of the Directory Role Object ID and the target Member Object ID in the format `{RoleObjectID}/member/{MemberObjectID}`.
    /// </summary>
    [AzureADResourceType("azuread:index/directoryRoleMember:DirectoryRoleMember")]
    public partial class DirectoryRoleMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Output("memberObjectId")]
        public Output<string?> MemberObjectId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("roleObjectId")]
        public Output<string?> RoleObjectId { get; private set; } = null!;


        /// <summary>
        /// Create a DirectoryRoleMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DirectoryRoleMember(string name, DirectoryRoleMemberArgs? args = null, CustomResourceOptions? options = null)
            : base("azuread:index/directoryRoleMember:DirectoryRoleMember", name, args ?? new DirectoryRoleMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DirectoryRoleMember(string name, Input<string> id, DirectoryRoleMemberState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/directoryRoleMember:DirectoryRoleMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DirectoryRoleMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DirectoryRoleMember Get(string name, Input<string> id, DirectoryRoleMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new DirectoryRoleMember(name, id, state, options);
        }
    }

    public sealed class DirectoryRoleMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Input("memberObjectId")]
        public Input<string>? MemberObjectId { get; set; }

        /// <summary>
        /// The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleObjectId")]
        public Input<string>? RoleObjectId { get; set; }

        public DirectoryRoleMemberArgs()
        {
        }
        public static new DirectoryRoleMemberArgs Empty => new DirectoryRoleMemberArgs();
    }

    public sealed class DirectoryRoleMemberState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of the principal you want to add as a member to the directory role. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Input("memberObjectId")]
        public Input<string>? MemberObjectId { get; set; }

        /// <summary>
        /// The object ID of the directory role you want to add the member to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleObjectId")]
        public Input<string>? RoleObjectId { get; set; }

        public DirectoryRoleMemberState()
        {
        }
        public static new DirectoryRoleMemberState Empty => new DirectoryRoleMemberState();
    }
}
