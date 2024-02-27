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
    /// Manages a single directory role assignment within Azure Active Directory.
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
    /// *Assignment for a built-in role*
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
    ///     var exampleDirectoryRole = new AzureAD.DirectoryRole("example", new()
    ///     {
    ///         DisplayName = "Security administrator",
    ///     });
    /// 
    ///     var exampleDirectoryRoleAssignment = new AzureAD.DirectoryRoleAssignment("example", new()
    ///     {
    ///         RoleId = exampleDirectoryRole.TemplateId,
    ///         PrincipalObjectId = example.Apply(getUserResult =&gt; getUserResult.ObjectId),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &gt; Note the use of the `template_id` attribute when referencing built-in roles.
    /// 
    /// *Assignment for a custom role*
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
    ///     var exampleCustomDirectoryRole = new AzureAD.CustomDirectoryRole("example", new()
    ///     {
    ///         DisplayName = "My Custom Role",
    ///         Enabled = true,
    ///         Version = "1.0",
    ///         Permissions = new[]
    ///         {
    ///             new AzureAD.Inputs.CustomDirectoryRolePermissionArgs
    ///             {
    ///                 AllowedResourceActions = new[]
    ///                 {
    ///                     "microsoft.directory/applications/basic/update",
    ///                     "microsoft.directory/applications/standard/read",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleDirectoryRoleAssignment = new AzureAD.DirectoryRoleAssignment("example", new()
    ///     {
    ///         RoleId = exampleCustomDirectoryRole.ObjectId,
    ///         PrincipalObjectId = example.Apply(getUserResult =&gt; getUserResult.ObjectId),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// *Scoped assignment for an application*
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// 	
    /// object NotImplemented(string errorMessage) 
    /// {
    ///     throw new System.NotImplementedException(errorMessage);
    /// }
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleDirectoryRole = new AzureAD.DirectoryRole("example", new()
    ///     {
    ///         DisplayName = "Cloud application administrator",
    ///     });
    /// 
    ///     var exampleApplication = new AzureAD.Application("example", new()
    ///     {
    ///         DisplayName = "My Application",
    ///     });
    /// 
    ///     var example = AzureAD.GetUser.Invoke(new()
    ///     {
    ///         UserPrincipalName = "jdoe@example.com",
    ///     });
    /// 
    ///     var exampleDirectoryRoleAssignment = new AzureAD.DirectoryRoleAssignment("example", new()
    ///     {
    ///         RoleId = exampleDirectoryRole.TemplateId,
    ///         PrincipalObjectId = example.Apply(getUserResult =&gt; getUserResult.ObjectId),
    ///         DirectoryScopeId = NotImplemented("format(\"/%s\",azuread_application.example.object_id)"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// &gt; Note the use of the `template_id` attribute when referencing built-in roles.
    /// 
    /// ## Import
    /// 
    /// Directory role assignments can be imported using the ID of the assignment, e.g.
    /// 
    /// ```sh
    /// $ pulumi import azuread:index/directoryRoleAssignment:DirectoryRoleAssignment example ePROZI_iKE653D_d6aoLHyr-lKgHI8ZGiIdz8CLVcng-1
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/directoryRoleAssignment:DirectoryRoleAssignment")]
    public partial class DirectoryRoleAssignment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&amp;tabs=http) for example usage. Changing this forces a new resource to be created.
        /// </summary>
        [Output("appScopeId")]
        public Output<string> AppScopeId { get; private set; } = null!;

        /// <summary>
        /// Identifier of the app-specific scope when the assignment scope is app-specific
        /// </summary>
        [Output("appScopeObjectId")]
        public Output<string> AppScopeObjectId { get; private set; } = null!;

        /// <summary>
        /// Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&amp;tabs=http) for example usage. Changing this forces a new resource to be created.
        /// </summary>
        [Output("directoryScopeId")]
        public Output<string> DirectoryScopeId { get; private set; } = null!;

        /// <summary>
        /// Identifier of the directory object representing the scope of the assignment
        /// </summary>
        [Output("directoryScopeObjectId")]
        public Output<string> DirectoryScopeObjectId { get; private set; } = null!;

        /// <summary>
        /// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Output("principalObjectId")]
        public Output<string> PrincipalObjectId { get; private set; } = null!;

        /// <summary>
        /// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;


        /// <summary>
        /// Create a DirectoryRoleAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DirectoryRoleAssignment(string name, DirectoryRoleAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/directoryRoleAssignment:DirectoryRoleAssignment", name, args ?? new DirectoryRoleAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DirectoryRoleAssignment(string name, Input<string> id, DirectoryRoleAssignmentState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/directoryRoleAssignment:DirectoryRoleAssignment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DirectoryRoleAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DirectoryRoleAssignment Get(string name, Input<string> id, DirectoryRoleAssignmentState? state = null, CustomResourceOptions? options = null)
        {
            return new DirectoryRoleAssignment(name, id, state, options);
        }
    }

    public sealed class DirectoryRoleAssignmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&amp;tabs=http) for example usage. Changing this forces a new resource to be created.
        /// </summary>
        [Input("appScopeId")]
        public Input<string>? AppScopeId { get; set; }

        /// <summary>
        /// Identifier of the app-specific scope when the assignment scope is app-specific
        /// </summary>
        [Input("appScopeObjectId")]
        public Input<string>? AppScopeObjectId { get; set; }

        /// <summary>
        /// Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&amp;tabs=http) for example usage. Changing this forces a new resource to be created.
        /// </summary>
        [Input("directoryScopeId")]
        public Input<string>? DirectoryScopeId { get; set; }

        /// <summary>
        /// Identifier of the directory object representing the scope of the assignment
        /// </summary>
        [Input("directoryScopeObjectId")]
        public Input<string>? DirectoryScopeObjectId { get; set; }

        /// <summary>
        /// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalObjectId", required: true)]
        public Input<string> PrincipalObjectId { get; set; } = null!;

        /// <summary>
        /// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        public DirectoryRoleAssignmentArgs()
        {
        }
        public static new DirectoryRoleAssignmentArgs Empty => new DirectoryRoleAssignmentArgs();
    }

    public sealed class DirectoryRoleAssignmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the app-specific scope when the assignment scope is app-specific. Cannot be used with `directory_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&amp;tabs=http) for example usage. Changing this forces a new resource to be created.
        /// </summary>
        [Input("appScopeId")]
        public Input<string>? AppScopeId { get; set; }

        /// <summary>
        /// Identifier of the app-specific scope when the assignment scope is app-specific
        /// </summary>
        [Input("appScopeObjectId")]
        public Input<string>? AppScopeObjectId { get; set; }

        /// <summary>
        /// Identifier of the directory object representing the scope of the assignment. Cannot be used with `app_scope_id`. See [official documentation](https://docs.microsoft.com/en-us/graph/api/rbacapplication-post-roleassignments?view=graph-rest-1.0&amp;tabs=http) for example usage. Changing this forces a new resource to be created.
        /// </summary>
        [Input("directoryScopeId")]
        public Input<string>? DirectoryScopeId { get; set; }

        /// <summary>
        /// Identifier of the directory object representing the scope of the assignment
        /// </summary>
        [Input("directoryScopeObjectId")]
        public Input<string>? DirectoryScopeObjectId { get; set; }

        /// <summary>
        /// The object ID of the principal for you want to create a role assignment. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
        /// </summary>
        [Input("principalObjectId")]
        public Input<string>? PrincipalObjectId { get; set; }

        /// <summary>
        /// The template ID (in the case of built-in roles) or object ID (in the case of custom roles) of the directory role you want to assign. Changing this forces a new resource to be created.
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        public DirectoryRoleAssignmentState()
        {
        }
        public static new DirectoryRoleAssignmentState Empty => new DirectoryRoleAssignmentState();
    }
}
