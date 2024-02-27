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
    /// Manages a single group membership within Azure Active Directory.
    /// 
    /// &gt; **Warning** Do not use this resource at the same time as the `members` property of the `azuread.Group` resource for the same group. Doing so will cause a conflict and group members will be removed.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `Group.ReadWrite.All` or `Directory.ReadWrite.All`.
    /// 
    /// However, if the authenticated service principal is an owner of the group being managed, an application role is not required.
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Groups Administrator`, `User Administrator` or `Global Administrator`
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
    ///     var example = AzureAD.GetUser.Invoke(new()
    ///     {
    ///         UserPrincipalName = "jdoe@example.com",
    ///     });
    /// 
    ///     var exampleGroup = new AzureAD.Group("example", new()
    ///     {
    ///         DisplayName = "my_group",
    ///         SecurityEnabled = true,
    ///     });
    /// 
    ///     var exampleGroupMember = new AzureAD.GroupMember("example", new()
    ///     {
    ///         GroupObjectId = exampleGroup.Id,
    ///         MemberObjectId = example.Apply(getUserResult =&gt; getUserResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Group members can be imported using the object ID of the group and the object ID of the member, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/groupMember:GroupMember example 00000000-0000-0000-0000-000000000000/member/11111111-1111-1111-1111-111111111111
    /// ```
    /// 
    ///  -&gt; This ID format is unique to Terraform and is composed of the Azure AD Group Object ID and the target Member Object ID in the format `{GroupObjectID}/member/{MemberObjectID}`.
    /// </summary>
    [AzureADResourceType("azuread:index/groupMember:GroupMember")]
    public partial class GroupMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("groupObjectId")]
        public Output<string> GroupObjectId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Output("memberObjectId")]
        public Output<string> MemberObjectId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMember(string name, GroupMemberArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/groupMember:GroupMember", name, args ?? new GroupMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupMember(string name, Input<string> id, GroupMemberState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/groupMember:GroupMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMember Get(string name, Input<string> id, GroupMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMember(name, id, state, options);
        }
    }

    public sealed class GroupMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("groupObjectId", required: true)]
        public Input<string> GroupObjectId { get; set; } = null!;

        /// <summary>
        /// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Input("memberObjectId", required: true)]
        public Input<string> MemberObjectId { get; set; } = null!;

        public GroupMemberArgs()
        {
        }
        public static new GroupMemberArgs Empty => new GroupMemberArgs();
    }

    public sealed class GroupMemberState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("groupObjectId")]
        public Input<string>? GroupObjectId { get; set; }

        /// <summary>
        /// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Input("memberObjectId")]
        public Input<string>? MemberObjectId { get; set; }

        public GroupMemberState()
        {
        }
        public static new GroupMemberState Empty => new GroupMemberState();
    }
}
