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
    /// Manages a single directory role assignment scoped to an administrative unit within Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `AdministrativeUnit.ReadWrite.All` and `RoleManagement.ReadWrite.Directory`, or `Directory.ReadWrite.All`
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
    ///     var example = AzureAD.GetUser.Invoke(new()
    ///     {
    ///         UserPrincipalName = "jdoe@example.com",
    ///     });
    /// 
    ///     var exampleAdministrativeUnit = new AzureAD.AdministrativeUnit("example", new()
    ///     {
    ///         DisplayName = "Example-AU",
    ///     });
    /// 
    ///     var exampleDirectoryRole = new AzureAD.DirectoryRole("example", new()
    ///     {
    ///         DisplayName = "Security administrator",
    ///     });
    /// 
    ///     var exampleAdministrativeUnitRoleMember = new AzureAD.AdministrativeUnitRoleMember("example", new()
    ///     {
    ///         RoleObjectId = exampleDirectoryRole.ObjectId,
    ///         AdministrativeUnitObjectId = exampleAdministrativeUnit.Id,
    ///         MemberObjectId = example.Apply(getUserResult =&gt; getUserResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Administrative unit role members can be imported using the object ID of the administrative unit and the unique ID of the role assignment, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember example 00000000-0000-0000-0000-000000000000/roleMember/zX37MRLyF0uvE-xf2WH4B7x-6CPLfudNnxFGj800htpBXqkxW7bITqGb6Rj4kuTuS
    /// ```
    /// 
    ///  -&gt; This ID format is unique to Terraform and is composed of the Administrative Unit Object ID and the role assignment ID in the format `{AdministrativeUnitObjectID}/roleMember/{RoleAssignmentID}`.
    /// </summary>
    [AzureADResourceType("azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember")]
    public partial class AdministrativeUnitRoleMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("administrativeUnitObjectId")]
        public Output<string> AdministrativeUnitObjectId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
        /// </summary>
        [Output("memberObjectId")]
        public Output<string> MemberObjectId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Output("roleObjectId")]
        public Output<string> RoleObjectId { get; private set; } = null!;


        /// <summary>
        /// Create a AdministrativeUnitRoleMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AdministrativeUnitRoleMember(string name, AdministrativeUnitRoleMemberArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember", name, args ?? new AdministrativeUnitRoleMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AdministrativeUnitRoleMember(string name, Input<string> id, AdministrativeUnitRoleMemberState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/administrativeUnitRoleMember:AdministrativeUnitRoleMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AdministrativeUnitRoleMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AdministrativeUnitRoleMember Get(string name, Input<string> id, AdministrativeUnitRoleMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new AdministrativeUnitRoleMember(name, id, state, options);
        }
    }

    public sealed class AdministrativeUnitRoleMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("administrativeUnitObjectId", required: true)]
        public Input<string> AdministrativeUnitObjectId { get; set; } = null!;

        /// <summary>
        /// The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
        /// </summary>
        [Input("memberObjectId", required: true)]
        public Input<string> MemberObjectId { get; set; } = null!;

        /// <summary>
        /// The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleObjectId", required: true)]
        public Input<string> RoleObjectId { get; set; } = null!;

        public AdministrativeUnitRoleMemberArgs()
        {
        }
        public static new AdministrativeUnitRoleMemberArgs Empty => new AdministrativeUnitRoleMemberArgs();
    }

    public sealed class AdministrativeUnitRoleMemberState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The object ID of the administrative unit you want to add the member to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("administrativeUnitObjectId")]
        public Input<string>? AdministrativeUnitObjectId { get; set; }

        /// <summary>
        /// The object ID of the user, group or service principal you want to add as a member of the administrative unit. Changing this forces a new resource to be created.
        /// </summary>
        [Input("memberObjectId")]
        public Input<string>? MemberObjectId { get; set; }

        /// <summary>
        /// The object ID of the directory role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleObjectId")]
        public Input<string>? RoleObjectId { get; set; }

        public AdministrativeUnitRoleMemberState()
        {
        }
        public static new AdministrativeUnitRoleMemberState Empty => new AdministrativeUnitRoleMemberState();
    }
}
