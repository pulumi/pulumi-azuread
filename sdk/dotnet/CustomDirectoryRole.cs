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
    /// Manages a Custom Directory Role within Azure Active Directory.
    /// 
    /// This resource is for managing custom directory roles. For management of built-in roles, see the azuread.DirectoryRole resource.
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
    ///     var example = new AzureAD.CustomDirectoryRole("example", new()
    ///     {
    ///         Description = "Allows reading applications and updating groups",
    ///         DisplayName = "My Custom Role",
    ///         Enabled = true,
    ///         Permissions = new[]
    ///         {
    ///             new AzureAD.Inputs.CustomDirectoryRolePermissionArgs
    ///             {
    ///                 AllowedResourceActions = new[]
    ///                 {
    ///                     "microsoft.directory/applications/basic/update",
    ///                     "microsoft.directory/applications/create",
    ///                     "microsoft.directory/applications/standard/read",
    ///                 },
    ///             },
    ///             new AzureAD.Inputs.CustomDirectoryRolePermissionArgs
    ///             {
    ///                 AllowedResourceActions = new[]
    ///                 {
    ///                     "microsoft.directory/groups/allProperties/read",
    ///                     "microsoft.directory/groups/allProperties/read",
    ///                     "microsoft.directory/groups/basic/update",
    ///                     "microsoft.directory/groups/create",
    ///                     "microsoft.directory/groups/delete",
    ///                 },
    ///             },
    ///         },
    ///         Version = "1.0",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support importing.
    /// </summary>
    [AzureADResourceType("azuread:index/customDirectoryRole:CustomDirectoryRole")]
    public partial class CustomDirectoryRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the custom directory role.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the custom directory role.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the role is enabled for assignment.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The object ID of the custom directory role.
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// A collection of `permissions` blocks as documented below.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.CustomDirectoryRolePermission>> Permissions { get; private set; } = null!;

        /// <summary>
        /// Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// The version of the role definition. This can be any arbitrary string between 1-128 characters.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a CustomDirectoryRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomDirectoryRole(string name, CustomDirectoryRoleArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/customDirectoryRole:CustomDirectoryRole", name, args ?? new CustomDirectoryRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomDirectoryRole(string name, Input<string> id, CustomDirectoryRoleState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/customDirectoryRole:CustomDirectoryRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomDirectoryRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomDirectoryRole Get(string name, Input<string> id, CustomDirectoryRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomDirectoryRole(name, id, state, options);
        }
    }

    public sealed class CustomDirectoryRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the custom directory role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the custom directory role.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Indicates whether the role is enabled for assignment.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("permissions", required: true)]
        private InputList<Inputs.CustomDirectoryRolePermissionArgs>? _permissions;

        /// <summary>
        /// A collection of `permissions` blocks as documented below.
        /// </summary>
        public InputList<Inputs.CustomDirectoryRolePermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.CustomDirectoryRolePermissionArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// The version of the role definition. This can be any arbitrary string between 1-128 characters.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public CustomDirectoryRoleArgs()
        {
        }
        public static new CustomDirectoryRoleArgs Empty => new CustomDirectoryRoleArgs();
    }

    public sealed class CustomDirectoryRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the custom directory role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the custom directory role.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Indicates whether the role is enabled for assignment.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The object ID of the custom directory role.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        [Input("permissions")]
        private InputList<Inputs.CustomDirectoryRolePermissionGetArgs>? _permissions;

        /// <summary>
        /// A collection of `permissions` blocks as documented below.
        /// </summary>
        public InputList<Inputs.CustomDirectoryRolePermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.CustomDirectoryRolePermissionGetArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// The version of the role definition. This can be any arbitrary string between 1-128 characters.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public CustomDirectoryRoleState()
        {
        }
        public static new CustomDirectoryRoleState Empty => new CustomDirectoryRoleState();
    }
}
