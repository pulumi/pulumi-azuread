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
    /// Manages a single catalog role assignment within Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `EntitlementManagement.ReadWrite.All` or `Directory.ReadWrite.All`
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Identity Governance administrator` or `Global Administrator`
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
    ///     var exampleAccessPackageCatalogRole = AzureAD.GetAccessPackageCatalogRole.Invoke(new()
    ///     {
    ///         DisplayName = "Catalog owner",
    ///     });
    /// 
    ///     var exampleAccessPackageCatalog = new AzureAD.AccessPackageCatalog("exampleAccessPackageCatalog", new()
    ///     {
    ///         DisplayName = "example-access-package-catalog",
    ///         Description = "Example access package catalog",
    ///     });
    /// 
    ///     var exampleAccessPackageCatalogRoleAssignment = new AzureAD.AccessPackageCatalogRoleAssignment("exampleAccessPackageCatalogRoleAssignment", new()
    ///     {
    ///         RoleId = exampleAccessPackageCatalogRole.Apply(getAccessPackageCatalogRoleResult =&gt; getAccessPackageCatalogRoleResult.ObjectId),
    ///         PrincipalObjectId = exampleUser.Apply(getUserResult =&gt; getUserResult.ObjectId),
    ///         CatalogId = exampleAccessPackageCatalog.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Catalog role assignments can be imported using the ID of the assignment, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment example 00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment")]
    public partial class AccessPackageCatalogRoleAssignment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("catalogId")]
        public Output<string> CatalogId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Output("principalObjectId")]
        public Output<string> PrincipalObjectId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPackageCatalogRoleAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPackageCatalogRoleAssignment(string name, AccessPackageCatalogRoleAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment", name, args ?? new AccessPackageCatalogRoleAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPackageCatalogRoleAssignment(string name, Input<string> id, AccessPackageCatalogRoleAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/accessPackageCatalogRoleAssignment:AccessPackageCatalogRoleAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccessPackageCatalogRoleAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPackageCatalogRoleAssignment Get(string name, Input<string> id, AccessPackageCatalogRoleAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessPackageCatalogRoleAssignment(name, id, state, options);
        }
    }

    public sealed class AccessPackageCatalogRoleAssignmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("catalogId", required: true)]
        public Input<string> CatalogId { get; set; } = null!;

        /// <summary>
        /// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalObjectId", required: true)]
        public Input<string> PrincipalObjectId { get; set; } = null!;

        /// <summary>
        /// The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        public AccessPackageCatalogRoleAssignmentArgs()
        {
        }
        public static new AccessPackageCatalogRoleAssignmentArgs Empty => new AccessPackageCatalogRoleAssignmentArgs();
    }

    public sealed class AccessPackageCatalogRoleAssignmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Catalog this role assignment will be scoped to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalObjectId")]
        public Input<string>? PrincipalObjectId { get; set; }

        /// <summary>
        /// The object ID of the catalog role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        public AccessPackageCatalogRoleAssignmentState()
        {
        }
        public static new AccessPackageCatalogRoleAssignmentState Empty => new AccessPackageCatalogRoleAssignmentState();
    }
}
